require('dotenv').config()
const config = require('config')
const debug = require('debug')('app:startup')
const morgan = require('morgan')

const express = require('express')
const networkInfo = require('./routes/network-info')

const app = express()

if (app.get('env') === 'development-docker') {
    app.use(morgan('tiny'))
    debug("Morgan Enabled...")
}

app.use(express.json())
app.use('/api/network-info', networkInfo)

app.listen(config.get('port'), () => console.log(`Server running on port: ${config.get('port')}`))