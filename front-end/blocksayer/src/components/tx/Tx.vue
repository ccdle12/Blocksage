<template>
<h1>TX PAGE</h1>
</template>

<script>
export default {
    /** Member variables */
    data: function() {
        return {
            txHash: 0
        }
    },

    /** Life cycle hooks */
    created: function() {
        this.txHash = this.$route.params.txHash
    },

} // export default
</script>

<style src='./Tx.css' scoped></style>