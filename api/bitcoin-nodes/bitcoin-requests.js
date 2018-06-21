require('dotenv').config()
const config = require('config')
const debug = require('debug')('app:bitcoin-requests')

const request = require('request')
const base64_auth = require('../utils/base64-auth')

module.exports.send = (method_request, args) =>
{ 
    debug(`http://${config.get('btc-mainnet-domain')}:${config.get('node-rpc')}`)
    debug(`Authorization Basic ${base64_auth.getEncodedAuth()}`)

    const options =
    {
        url: `http://${config.get('btc-mainnet-domain')}:${config.get('node-rpc')}`,
        headers: {
            Authorization: `Basic ${base64_auth.getEncodedAuth()}`
        },
        body: JSON.stringify({
            method: method_request,
            params: args
        })
    }

    return new Promise((resolve, reject) => {
        request.post(options, (err, res, body) => {
            if (err) 
            {
                debug(`Request Err: ${err}`)
                return resolve(JSON.stringify({
                    "result": null,
                    "error": {
                        "code": 502,
                        "message": "Bad Gateway - Node Unresponsive"
                    },
                    "id": null
                }))
            }

            debug(`Request Body: ${body}`)
            return resolve(body)
        })
    })
}