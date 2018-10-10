<template>

  <div>
    {{ title }}

    <!-- Search Bar -->
    <form @submit.prevent="submitSearch">
      <input v-model="userSearch" v-on:keyup.enter.space="submitSearch" type="text" placeholder="Enter transaction hash/block hash">
    </form>

    <button v-on:click="submitSearch"> Search </button>

    <!-- Error Message -->
    <span v-if="invalidSearch"> Invalid Search </span>
  </div>

</template>

<script>
import searchValidator from './SearchValidation'
import navigationService from '../navigation/NavigationService'

export default {
  data: function() {
    return {
      title: 'Blocksage',
      userSearch: '',
      invalidSearch: false
    }
  },

  computed: {
  },

  watch: {
  }, 

  // b1fea52486ce0c62bb442b530a3f0132b826c74e473d1f2c220bfa78111c5082
  methods: {
    submitSearch: function() {
      let searchRequest = searchValidator.validate(this.userSearch)
      navigationService.navigate(this.$router, searchRequest, this.userSearch)

      // Only gets executed if navigation has failed due to invalid search
      this.invalidSearch = true
    },
  },

} // export default
</script>

<style src='./Home.css' scoped></style>
