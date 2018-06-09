const debug = require('debug')('app:validate-response')

module.exports = (err, body) => 
{
    if (err) 
    {
        debug(`Error: ${err}`)
        return { status_code: 502, message: 'Bad Gateway - Node unresponsive' }
    }
    
    const json_body = JSON.parse(body)
    debug(`Body: ${JSON.stringify(json_body)}`)

    if (json_body.error && json_body.error.code === -32601)
        return { status_code: 404, message: 'Method request not found' }

    return { status_code: 200, message: 'OK' }
}