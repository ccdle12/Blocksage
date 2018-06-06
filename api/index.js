require('dotenv').config()
const config = require('config')
const debug = require('debug')('app:startup')

const express = require('express')
const networkInfo = require('./routes/network-info')

const app = express()

app.use(express.json())
app.use('/api/network-info', networkInfo)

app.listen(config.get('port'), () => console.log(`Server running on port: ${config.get('port')}`))