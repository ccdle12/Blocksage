'use strict'
const merge = require('webpack-merge')
const prodEnv = require('./prod.env')

module.exports = merge(prodEnv, {
  NODE_ENV: '"development"',
  DEBUG_MODE: true,
  API_BASE_URL: '"http://127.0.0.1:8548/api"'
})
