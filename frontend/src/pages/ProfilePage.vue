<template>
  <div class="page profile-wrapper">
    <div v-if="!isLoaded"> Loaded ...</div>
    <div v-else class="user-profile">
<!--      <div v-for=""></div>-->
    </div>
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
      isLoaded: false,
      profileData: {}
    }
  },
  computed: {
    basicAuth() {
      return localStorage.getItem('matesterBasicAuth');
    },
    user() {
      return this.$store.getters['user/userData'].login;
    }
  },
  created() {
    console.log('loaded profile page')
    this.getProfile();
  },
  methods: {
    getProfile() {
      axios.get(`https://matester23.herokuapp.com/user?user=${this.user}`, {
        headers: { 'Authorization': this.basicAuth }
      }).then(userResponse => {
        console.log('userResponse', userResponse)
        this.profileData = userResponse.data;
        this.isLoaded = true;
      })
    }
  }
}
</script>

<style>
.page {
  background-color: #fff;
  padding: 10px;
  //min-height: 400px;
}
</style>
