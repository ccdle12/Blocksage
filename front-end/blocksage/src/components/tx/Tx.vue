<template>
  <div>
    <h1>TX PAGE</h1>
    {{ TX }}
  </div>
</template>

<script>
import btcAPIService from '../bitcoin-mainnet/BTCAPIService'

export default {
  data: function() {
    return {
      TX: ''
    }
  },

  created: function() {
    this.requestTX(this.$route.params.txHash)
  },

  methods: {
    requestTX: function(txHash) {
      btcAPIService
        .getTX(txHash)
        .then(response => this.TX = response.data)
        .catch(error => {
          if (error.response) {
            this.TX = error.response.data
          } else if (error.request) {
            console.log('Error reqest: ', error.request);
            this.TX = 'Something went wrong, there was no response'
          } else {
            return error.message
          }
      })
    }
  }
} // export default
</script>

<style src='./Tx.css' scoped></style>