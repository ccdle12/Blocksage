require('dotenv').config()
const config = require('config')
const debug = require('debug')('app:network-info')

const bitcoin_request = require('../utils/bitcoin-requests')

const express = require('express')
const router = express.Router()

router.get('/', (req, res) => {
    bitcoin_request.send('getnetworkinfo', (err, rpc_res, body) => {
        if (err) {
            res.status(500).send()
            debug(`Error: ${err}`)
        }
        
        const json_body = JSON.parse(body)
        debug(`Body: ${json_body}`)

        if (json_body.error && json_body.error.code === -32601)
            res.status(404).send('Method request not found')
        
        res.send(body)
    })
})

module.exports = router