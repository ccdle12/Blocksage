require('dotenv').config()
const config = require('config')
const debug = require('debug')('app:transactions')

const bitcoin_request = require('../bitcoin-node/bitcoin-requests')
const validate_response = require('../bitcoin-node/validate-response')

const express = require('express')
const router = express.Router()

router.get('/:txhash', (req, res) =>
{
    (async () => {
        debug(`Tx Hash: ${req.params.txhash}`)

        let raw_tx_body = await bitcoin_request.send('getrawtransaction', [req.params.txhash])
        debug(`Raw Tx Body: ${raw_tx_body}`)

        let { status_code, message } = validate_response(raw_tx_body)
        debug(`Status Code from getrawtransaction: ${status_code}`)

        if (status_code !== 200) 
            return res.status(status_code).send(`${message} incorrect tx hash or is not hex`)

        const raw_tx = JSON.parse(raw_tx_body)['result']

        let decoded_tx_body = await bitcoin_request.send('decoderawtransaction', [raw_tx])
        debug(`decoded tx: ${decoded_tx_body}`)

        status_code, message = validate_response(decoded_tx_body)

        if (status_code !== 200) 
            return res.status(decoded_tx_status_code).send(decoded_tx_message)

        res.send(decoded_tx_body)
    })()
})

module.exports = router