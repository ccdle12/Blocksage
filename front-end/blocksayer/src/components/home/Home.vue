<template>
  <div>
    {{ name }}

    <form @submit.prevent="findSearchValue">
      <input v-model="searchValue" v-on:keyup.enter.space="findSearchValue" type="text" placeholder="Enter transaction hash/block hash">
    </form>

    <button v-on:click="findSearchValue"> Search </button>

    <span v-if="invalidSearch"> Invalid Search </span>

  </div>
</template>

<script>
export default {
  data: function() {
    return {
      /** Member Variables */
      name: 'Blocksayer',
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
     
      if (isValid.valid && isValid.type === 'Tx') 
        this.navigateToTx()

      this.invalidSearch = true
    },

    parseSearchValue: function(search_value) {
      let len_search_value = search_value.length

      let acceptable = false
      let value = 'None'

      switch(len_search_value) {
        case 64:
          acceptable = true
          value = 'Tx'
          break
      }

      return { valid: acceptable, type: value }
    },

    navigateToTx: function() {
      this.$router.push('/tx')
    }


  } // methods
} // export default
</script>

<style src='./Home.css' scoped></style>
