require('dotenv').config()
const config = require('config')
const debug = require('debug')('app:transactions')

const bitcoin_request = require('../bitcoin-node/bitcoin-requests')
const validate_response = require('../bitcoin-node/validate-response')

const express = require('express')
const router = express.Router()

router.get('/:txhash', (req, res) =>
{
    debug(`Tx Hash: ${req.params.txhash}`) 

    bitcoin_request.send('getrawtransaction', [req.params.txhash], (err, rpc_res, body) =>
    {
        const { status_code, message } = validate_response(err, body)

        if (status_code !== 200)
            return res.status(status_code).send(message)

        const raw_tx = JSON.parse(body)['result']
        debug(`Raw Tx: ${raw_tx}`)

        bitcoin_request.send('decoderawtransaction', [raw_tx], (err, rpc_res, body) => 
        {
            const { status_code, message } = validate_response(err, body)

            if (status_code !== 200)
                return res.status(status_code).send(message)

            res.send(body)
        })
    })
})

module.exports = router