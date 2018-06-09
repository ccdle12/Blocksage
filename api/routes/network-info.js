require('dotenv').config()
const config = require('config')
const debug = require('debug')('app:network-info')

const bitcoin_request = require('../bitcoin-node/bitcoin-requests')
const validate_response = require('../bitcoin-node/validate-response')

const express = require('express')
const router = express.Router()

router.get('/', (req, res) =>
{
    bitcoin_request.send('getnetworkinfo', (err, rpc_res, body) =>
    {
        const { status_code, message } = validate_response(err, body)
        debug(`Status: ${status_code.toString()}\n Message: ${message}`)

        if (status_code !== 200)
        {
            res.status(status_code).send(message)
            return
        }

        res.send(body)
    })
})

module.exports = router