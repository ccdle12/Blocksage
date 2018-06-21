<template>
  <div>
    <h1>TX PAGE</h1>
    {{ retrievedTx }}
  </div>
</template>

<script>
import btcApiService from '../bitcoin-mainnet/BTCAPIService'

export default {
  data: function() {
    return {
      retrievedTx: ''
    }
  },

  created: function() {
    this.getTx(this.$route.params.txHash)
  },

  methods: {
    getTx: function(txHash) {
      btcApiService
        .getTx(txHash)
        .then(response => this.retrievedTx = response.data)
        .catch(error => {
        if (error.response) {
          this.retrievedTx = error.response.data
        } else if (error.request) {
          console.log('Error reqest: ', error.request);
          this.retrievedTx = 'Something went wrong, there was no response'
        } else {
          return error.message
        }
      })
    }
  }
} // export default
</script>

<style src='./Tx.css' scoped></style>