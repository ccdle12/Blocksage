require('dotenv').config()
const debug = require('debug')('app:address')

const bitcoin_request = require('../bitcoin-node/bitcoin-requests')
const validate_response = require('../bitcoin-node/validate-response')

const express = require('express')
const router = express.Router()

router.get('/:block_hash', (req, res) =>
{
    (async () => {
        debug(`Block Hash: ${req.params.block_hash}`)

        let body = await bitcoin_request.send('getblock', [req.params.block_hash])

        const { status_code, message } = validate_response(body)
        debug(`Status Code: ${status_code}`)

        if (status_code !== 200) 
            return res.status(status_code).send(`${message} Block not found or block hash is incorrect`)

        res.send(body)
    })()
})

module.exports = router