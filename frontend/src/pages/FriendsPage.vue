<template>
  <div class="users-page">
    <div v-for="user in users" :key="user.id">
      {{ user }}
    </div>
  </div>
</template>

<script>
export default {
  name: "FriendsPage",
  data() {
    return {
      users: [],
    }
  },
  created() {
    this.getUsers();
  },
  mounted() {
  },
  methods: {
    getUsers() {
      this.$http.post(`https://matester23.herokuapp.com/friends`,{
        headers: { 'X-Requested-With': 'XMLHttpRequest' },
      }).then(result => {
        console.log('friends', result);
        if (result.data.users && result.data.users.length !== 0) {
          this.users = result.data.users;
        }
      })
    }
  }
}
</script>

<style scoped>

</style>
