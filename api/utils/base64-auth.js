require('dotenv').config()
const config = require('config')
const debug = require('debug')('app:base64-auth')

module.exports.getEncodedAuth = () =>
{
    debug(Buffer.from(`${config.get('username')}:${config.get('password')}`).toString('base64'))
    return Buffer.from(`${config.get('username')}:${config.get('password')}`).toString('base64')
}