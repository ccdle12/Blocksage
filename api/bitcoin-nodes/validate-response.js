const debug = require('debug')('app:validate-response')

module.exports = (body) =>
{
    debug(`Body before parse: ${body}`)
    const json_body = JSON.parse(body)
    debug(`JSON Body: ${json_body.error}`)

    if (json_body.error && json_body.error != null)
    {
        debug(`Error code: ${json_body.error.code}`)
        debug(`Error Code type: ${typeof(json_body.error.code)}`)

        // RPC Protocol - https://github.com/bitcoin/bitcoin/blob/v0.15.0.1/src/rpc/protocol.h
        switch(json_body.error.code)
        {
            case -32601:
                return { status_code: 404, message: 'Method request not found' }

            case -3:
                return { status_code: 404, message: 'Bad input type - not found, does not exist or incorrect input format' }

            case -5:
                return { status_code: 404, message: 'Invalid address or key - ' }

            case -8:
                return { status_code: 404, message: 'Invalid Parameter -' }

            case 502:
                return { status_code: 502, message: 'Bad Gateway - Node unresponsive' }
        }
    }

    return { status_code: 200, message: 'OK' }
}