require('dotenv').config()
const config = require('config')

const base64Auth = require('../utils/base64-auth')

module.exports.createOptions = (method_request) => { 
    return options = {
        url: `http://${config.get('domain')}:${config.get('node-rpc')}`,
        headers: {
            Authorization: `Basic ${base64Auth.getEncodedAuth()}`
        },
        body: JSON.stringify({
            method: method_request
        })
    }
}