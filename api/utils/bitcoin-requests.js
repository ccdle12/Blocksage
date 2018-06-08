require('dotenv').config()
const config = require('config')

const request = require('request')
const base64Auth = require('../utils/base64-auth')

module.exports.send = (method_request, callback) => { 
    const options = {
        url: `http://${config.get('domain')}:${config.get('node-rpc')}`,
        headers: {
            Authorization: `Basic ${base64Auth.getEncodedAuth()}`
        },
        body: JSON.stringify({
            method: method_request
        })
    }
    
    request.post(options, callback)
}