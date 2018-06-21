<template>

  <div>
    <h1>TX PAGE</h1>
    {{ retrievedTx }}
  </div>

</template>

<script>
export default {
    /** Member variables */
    data: function() {
        return {
            txHash: 0,
            retrievedTx: 'No Tx Data'
        }
    },

    /** Life cycle hooks */
    created: function() {
        this.txHash = this.$route.params.txHash
        this.getTx()
    },

    methods: {
        getTx: function() {
            this.$http.get(`http://localhost:8548/api/txs/${this.txHash}`)
                .then((res) => {
                    this.retrievedTx = res.body.result
                })
        }
    }

} // export default
</script>

<style src='./Tx.css' scoped></style>