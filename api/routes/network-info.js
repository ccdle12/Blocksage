require('dotenv').config()
const debug = require('debug')('app:network-info')

const bitcoin_request = require('../bitcoin-node/bitcoin-requests')
const validate_response = require('../bitcoin-node/validate-response')

const express = require('express')
const router = express.Router()

router.get('/', (req, res) =>
{
    (async () => {
        let body = await bitcoin_request.send('getnetworkinfo', [])

        const { status_code, message } = validate_response(body)
        debug(`Status Code: ${status_code}`)

        if (status_code !== 200) 
            return res.status(status_code).send(message)

        res.send(body)
    })()
})

module.exports = router