require('dotenv').config()
const config = require('config')

module.exports.getRPCAuth = () => { return Buffer.from(`${config.get('username')}:${config.get('password')}`).toString('base64') }