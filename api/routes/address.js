require('dotenv').config()
const config = require('config')
const debug = require('debug')('app:address')

const bitcoin_request = require('../bitcoin-node/bitcoin-requests')
const validate_response = require('../bitcoin-node/validate-response')

const express = require('express')
const router = express.Router()

router.get('/:address/balance', (req, res) =>
{
    debug(`Address: ${req.params.address}`) 

    bitcoin_request.send('getreceivedbyaddress', [req.params.address, 1], (err, rpc_res, body) =>
    {
        const { status_code, message } = validate_response(err, body)

        if (status_code !== 200)
            return res.status(status_code).send(message)

        res.send(body)
    })
})

module.exports = router