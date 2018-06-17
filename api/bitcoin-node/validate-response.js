const debug = require('debug')('app:validate-response')

module.exports = (body) =>
{
    const json_body = JSON.parse(body)

    if (json_body.error)
    {
        switch(json_body.error.code)
        {
            case 502:
                return { status_code: 502, message: "Bad Gateway - Node unresponsive" }

            case -32601:
                return { status_code: 404, message: 'Method request not found' }

            case -3:
                return { status_code: 404, message: 'Bad input type - not found, does not exist or incorrect input format'}

            case -5:
                return { status_code: 404, message: 'Block not found'}
        }
    }

    return { status_code: 200, message: 'OK' }
}