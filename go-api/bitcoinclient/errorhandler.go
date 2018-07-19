package bitcoinclient

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

var (
	// ErrUnresponsive is an error that catches an unresponsive Bitcoin Node
	ErrUnresponsive = errors.New("Timeout Error: Bitcoin Node is unresponsive")

	// ErrCreatingRequest is an error that catches internal error in creating a request
	ErrCreatingRequest = errors.New("Internal Error: Cannot create a request to Bitcoin Node")

	// ErrCreatingBody is an error that catches an internal error in creating a request
	ErrCreatingBody = errors.New("Internal Error: Cannot create body for request to Bitcoin Node")

	// ErrConvertingBodyToString is an error that catches an internal error in converting the body to a string
	ErrConvertingBodyToString = errors.New("Internal Error: Cannot convert body to string")

	// ErrFailedAuthentication is an error that catches an internal error in converting the body to a string
	ErrFailedAuthentication = errors.New("Authentication Error: Username and Password has failed to access Bitcoin Node")

	// ErrMethodRequestNotFound is an error that catches a request to the Bitcoin Node with a method that does not exist
	ErrMethodRequestNotFound = errors.New("Bitcoin Node Error: Method requested to the Bitcoin Node does not exist")

	// ErrIncorrectInput is an error that catches bad inputs in a request
	ErrIncorrectInput = errors.New("Input Error: Not found, does not exist or incorrect input format")

	// ErrInvalidAddr is an error that catches inputs that result in a malformed address, block, etc..
	ErrInvalidAddr = errors.New("Input Error: Invalid Address")

	// ErrInvalidParameter is an error that catches invalid params
	ErrInvalidParameter = errors.New("Input Error: Invalid Parameter")
)

// RPCBitcoinError is a struct that contains the format for errors when returned by the Bitcoin Node RPC
type RPCBitcoinError struct {
	Result *string  `json:"result"`
	Error  rpcError `json:"error"`
	ID     *string  `json:"id"`
}

type rpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// AuthenticateRPCResponse will take in a response string and determine whether there was an internal RPC error
// or whether there was a failure to authenticate
func (b *BitcoinClient) AuthenticateRPCResponse(res string) error {
	authErr := b.checkAuthentication(res)
	if authErr != nil {
		return authErr
	}

	err := b.checkRPCErrorCode(res)
	if err != nil {
		return err
	}

	return nil
}

// checkAuthentication checks if a response is empty, meaning RPC authentication to the Bitcoin Node has failed
func (b *BitcoinClient) checkAuthentication(res string) error {
	if len(res) == 0 {
		return ErrFailedAuthentication
	}

	return nil
}

func (b *BitcoinClient) checkRPCErrorCode(res string) error {
	var rpcError RPCBitcoinError
	json.Unmarshal([]byte(res), &rpcError)

	if rpcError.Error.Code != 0 {
		switch rpcError.Error.Code {
		case -32601:
			log.Println(ErrMethodRequestNotFound)
			return ErrMethodRequestNotFound

		case -3:
			log.Println(ErrIncorrectInput)
			return ErrIncorrectInput

		case -5:
			log.Println(ErrInvalidAddr)
			return ErrInvalidAddr

		case -8:
			log.Println(ErrInvalidParameter)
			return ErrInvalidParameter

		case 502:
			log.Println(ErrUnresponsive)
			return ErrUnresponsive
		}
	}

	return nil
}

// HandleStatusCodeError will check the error returned from the Bitcon Client and
// update ResponseWriter with the correct status code
func (b *BitcoinClient) HandleStatusCodeError(w http.ResponseWriter, err error) http.ResponseWriter {
	switch err {
	case ErrUnresponsive:
		w.WriteHeader(http.StatusServiceUnavailable)
		return w

	case ErrCreatingRequest:
		w.WriteHeader(http.StatusInternalServerError)
		return w

	case ErrCreatingBody:
		w.WriteHeader(http.StatusInternalServerError)
		return w

	case ErrConvertingBodyToString:
		w.WriteHeader(http.StatusInternalServerError)
		return w

	case ErrFailedAuthentication:
		w.WriteHeader(http.StatusInternalServerError)
		return w

	case ErrMethodRequestNotFound:
		w.WriteHeader(http.StatusNotFound)
		return w

	case ErrIncorrectInput:
		w.WriteHeader(http.StatusNotFound)
		return w

	case ErrInvalidAddr:
		w.WriteHeader(http.StatusNotFound)
		return w
	}

	return w
}
