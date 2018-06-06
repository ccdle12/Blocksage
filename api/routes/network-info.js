require('dotenv').config()
const config = require('config')
const debug = require('debug')('app:network-info')
const request = require('request')

const express = require('express')
const router = express.Router()

router.get('/', (req, res) => {
    //TODO - Move into own middleware or service
    console.log(Buffer.from("bitcoin:password").toString('base64'))
    const options = {
        url: `http://${config.get('domain')}:${config.get('node-rpc')}`,
        headers: {
            Authorization: 'Basic Yml0Y29pbjpwYXNzd29yZA=='
        },
        body: JSON.stringify({
            method: 'getnetworkinfo'
        })
    }

    function callback(err, res, body) {
        if (err)
            debug(`Error: ${err}`)
        
        debug(`Body: ${body}`)
    }

    request.post(options, callback)
})

module.exports = router