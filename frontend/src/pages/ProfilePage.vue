<template>
  <div class="page profile-wrapper">
    <div v-if="!isLoaded"> Loaded ...</div>
    <h1 v-else>HELLO</h1>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: "ProfilePage",
  components: {},
  props: {
    user_id: { type: String }
  },
  data() {
    return {
      isLoaded: false
    }
  },
  computed: {
    basicAuth() {
      return this.$store.getters['user/basicAuth'];
    }
  },
  created() {
    console.log('loaded profile page')
    axios.get(`https://matester23.herokuapp.com/user?user=${this.user_id}`, {
      headers: { 'Authorization': this.basicAuth }
    }).then(userResponse => {
      console.log('userResponse', userResponse)
      this.isLoaded = true;
    })
  },
  methods: {

  }
}
</script>

<style scoped>

</style>
