const validate_response = require('../../bitcoin-node/validate-response')
const debug = require('debug')('app:validate-response-test')

function create_error_body(err_message, err_code) {
    return {
        "result": null,
        "error": {
            "code": err_code,
            "message": err_message
        },
        "id": null
    }
}

describe('validate-response-errors', () =>
{
    it('should return 502 - Bad Gateway Node unresponsive', () =>
    {
        const body = JSON.stringify(create_error_body("ERRConnect node unresponsive", 502))

        const { status_code, message } = validate_response(body)
        expect(status_code).toEqual(502)
    })

    it('should return 404 - Method Request Not Found', () =>
    {
        const body = JSON.stringify(create_error_body("Method Request Not Found", -32601))
        
        const { status_code, message } = validate_response(body)
        expect(status_code).toEqual(404)
    })

    it('should return 404 - Method Request Not Found', () =>
    {
        const body = JSON.stringify(create_error_body("Bad input type - not found, does not exist or incorrect input format", -3))
        
        const { status_code, message } = validate_response(body)
        expect(status_code).toEqual(404)
    })

    it('should return 404 - Block not found', () =>
    {
        const body = JSON.stringify(create_error_body("Block Not Found", -5))
        
        const { status_code, message } = validate_response(body)
        expect(status_code).toEqual(404)
    })
})