<template>
  <div class="main-menu">
    <div class="title-wrapper" v-for="title in titles" :key="title.id">
      <router-link :key="title.id"
                   :to="title.url"
      > {{ title.name }}
      </router-link>
    </div>
  </div>
</template>

<script>
export default {
  name: "MainMenu",
  data() {
    return {
      titles: [
        { id: 'users', name: 'All users', url: '/users' },
        { id: 'friends', name: 'Friends', url: '/friends' },
        { id: 'profile', name: 'Profile', url: '/user' },
        { id: 'logout', name: 'Logout', url: '/login' }
      ]
    }
  }
}
</script>

<style scoped>
.main-menu {
  display: flex;
  flex-direction: column;
  width: 250px;
}
.title-wrapper {
  padding: 20px 0;
  text-transform: uppercase;
  text-align: left;
}
.title-wrapper a {
  text-decoration: none;
  color: #373434;
}
.title-wrapper a.router-link-active {
  color: #6F8BE9;
  font-weight: bold;
}
</style>
