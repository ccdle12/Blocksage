// +build unit

package injector_test

import (
	"github.com/ccdle12/Blocksage/go-crawler/injector"
	"github.com/ccdle12/Blocksage/go-crawler/test-utils"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
	// "time"
)

// ===========================================================
// Testing Suite
// ===========================================================
type InjectorSuite struct {
	suite.Suite
	httpClient *http.Client
}

// This gets run automatically by `go test` so we call `suite.Run` inside it
func TestInjectorSuite(t *testing.T) {
	// This is what actually runs our suite
	suite.Run(t, new(InjectorSuite))
}

func (suite *InjectorSuite) SetupTest() {
	suite.httpClient = injector.DefaultHTTPClient()
}

// ===========================================================
// Unit Tests
// ===========================================================
// TestInitDBClient will test that a DBClient can be initialized.
// TestNewRequest will test the function to create a Node Request struct
// to pass to the usecase.
func (suite *InjectorSuite) TestNewRequest() {
	// TODO (ccdle12): Move http client to injector
	nodeReq := injector.NodeRequest(suite.httpClient, testutils.NodeAddress, testutils.Username, testutils.Password, testutils.GetBlock, testutils.GetBlockParams)

	// nodeReq should initialize
	suite.NotNil(nodeReq, "nodeReq should not be nil")

	// nodeReq should have the nodeAddress passed
	suite.EqualValues(nodeReq.Headers.Address, testutils.NodeAddress)

	// nodeReq should have the username passed
	suite.EqualValues(nodeReq.Headers.Username, testutils.Username)

	// nodeReq should have the password passed
	suite.EqualValues(nodeReq.Headers.Password, testutils.Password)

	// nodeReq should have the method passed, "getblock"
	suite.EqualValues(nodeReq.Body.Method, testutils.GetBlock)

	// nodeReq should have the params, "12345"
	suite.EqualValues(nodeReq.Body.Params[0], testutils.GetBlockParams[0])
}

// TestDefaultHttpClient will test that we can retrive a default Http Client.
func (suite *InjectorSuite) TestDefaultHttpClient() {
	suite.NotNil(suite.httpClient)
}

// TestGetBTCDomain will test that we can retrieve the env variable of the BTC Mainnet Node Address.
func (suite *InjectorSuite) TestGetBTCDomain() {
	btcDomain := injector.BTCDomain()

	suite.NotNil(btcDomain, "btcDomain should not be nil")
}

// TestGetBTCUsername will test that we can retrieve the env variable of the BTC Mainnet Node Username.
func (suite *InjectorSuite) TestGetBTCUsername() {
	btcUsername := injector.BTCUsername()

	suite.NotNil(btcUsername, "btcUser should not be nil")
}

// TestGetBTCPassword will test that we can retrieve the env variable of the BTC Mainnet Node Password.
func (suite *InjectorSuite) TestGetBTCPassword() {
	btcPassword := injector.BTCPassword()

	suite.NotNil(btcPassword, "btcPassword should not be nil")
}

// TestPostgresDBName will test that we can retrieve the env variable of the Postgres DB Name.
func (suite *InjectorSuite) TestPostgresDBName() {
	postgresDBName := injector.PostgresDBName()

	suite.NotNil(postgresDBName, "postgresDBName should not be nil")
}

// TestPostgresUserName will test that we can retrieve the env variable of the Postgres User Name.
func (suite *InjectorSuite) TestPostgresUserName() {
	postgresUsername := injector.PostgresUserName()

	suite.NotNil(postgresUsername, "postgresUsername should not be nil")
}

// TestPostgresPassword will test that we can retrieve the env variable of the Postgres Password.
func (suite *InjectorSuite) TestPostgresPassword() {
	postgresPassword := injector.PostgresPassword()

	suite.NotNil(postgresPassword, "postgresPassword should not be nil")
}

// TODO (ccdle12): These tests only work on the local version
// TestPostgresDomain will test that we can retrieve the env variable of the Postgres Domain.
func (suite *InjectorSuite) TestPostgresDomain() {
	postgresDomain := injector.PostgresDomain()

	suite.NotNil(postgresDomain, "postgresDomain should not be nil")
	suite.Equal("172.25.0.4", postgresDomain, "postgresDomain should equal 172.25.0.4")
}

// TestPostgresPort will test that we can retrieve the env variable of the Postgres Port.
func (suite *InjectorSuite) TestPostgresPort() {
	postgresPort := injector.PostgresPort()

	suite.NotNil(postgresPort, "postgresPort should not be nil")
	suite.Equal("5432", postgresPort, "postgresPort should equal 5432")
}
