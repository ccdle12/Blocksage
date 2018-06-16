require('dotenv').config()
const config = require('config')
const debug = require('debug')('app:bitcoin-requests')

const request = require('request')
const base64_auth = require('../utils/base64-auth')

module.exports.send = (method_request, callback) => 
{ 
    debug(`http://${config.get('domain')}:${config.get('node-rpc')}`)
    debug(`Authorization Basic ${base64_auth.getEncodedAuth()}`)

    const options =
    {
        url: `http://${config.get('domain')}:${config.get('node-rpc')}`,
        headers: {
            Authorization: `Basic ${base64_auth.getEncodedAuth()}`
        },
        body: JSON.stringify({
            method: method_request
        })
    }
    
    request.post(options, callback)
}