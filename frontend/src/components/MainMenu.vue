<template>
  <div class="main-menu">
    <div class="title-wrapper" v-for="title in titles" :key="title.id">
      <span class="title">{{ title.name }}</span>
    </div>
  </div>
</template>

<script>
export default {
  name: "MainMenu",
  data() {
    return {
      titles: [
        { id: 'users', name: 'All users' },
        { id: 'friends', name: 'Friends' },
        { id: 'profile', name: 'Profile' },
        { id: 'logout', name: 'Logout' }
      ]
    }
  }
}
</script>

<style scoped>
.main-menu {
  display: flex;
  flex-direction: column;
}
</style>
