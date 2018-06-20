require('dotenv').config()
const config = require('config')
const debug = require('debug')('app:startup')
const morgan = require('morgan')

const express = require('express')

const address = require('./routes/address')
const blocks = require('./routes/blocks')
const transactions = require('./routes/transactions')
const network_info = require('./routes/network-info')
const cors = require('cors')

const app = express()

if (app.get('env') === 'development-docker') {
    app.use(morgan('tiny'))
    app.use(cors({origin: `http://${config.get('domain')}:${config.get('front-end-port')}`}))
    debug("Morgan Enabled...")
}

app.use(express.json())
app.use('/api/address', address)
app.use('/api/blocks', blocks)
app.use('/api/txs', transactions)
app.use('/api/network-info', network_info)

app.listen(config.get('port'), () => console.log(`Server running on port: ${config.get('port')}`))