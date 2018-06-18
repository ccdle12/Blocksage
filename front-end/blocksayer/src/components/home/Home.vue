<template>
  <div>
    {{ title }}

    <!-- Search Bar -->
    <form @submit.prevent="findSearchValue">
      <input v-model="searchValue" v-on:keyup.enter.space="findSearchValue" type="text" placeholder="Enter transaction hash/block hash">
    </form>

    <button v-on:click="findSearchValue"> Search </button>

    <!-- Error Message -->
    <span v-if="invalidSearch"> Invalid Search </span>

  </div>
</template>

<script>
export default {
  /** Member Variables */
  data: function() {
    return {
      title: 'Blocksayer',
      searchValue: '',
      invalidSearch: false
    }
  },

  computed: {
  },

  watch: {
  }, 

  // b1fea52486ce0c62bb442b530a3f0132b826c74e473d1f2c220bfa78111c5082
  methods: {
    findSearchValue: function() {
      let isValid = this.parseSearchValue(this.searchValue)
     
      if (isValid.valid && isValid.type === 'tx') this.navigateToTx()

      this.invalidSearch = true
    },

    parseSearchValue: function(search_value) {
      let acceptable = false
      let value = 'None'

      switch(search_value.length) {
        case 64:
          acceptable = true
          value = 'tx'
          break
      }

      return { valid: acceptable, type: value }
    },

    navigateToTx: function() {
      this.$router.push({ name: 'tx', params: { txHash: this.searchValue } })
    }
  },

} // export default
</script>

<style src='./Home.css' scoped></style>
