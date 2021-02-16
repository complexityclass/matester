<template>
  <div class="page-wrapper">
    <div class="logo">
      <img class="page-logo" src="../assets/logo_transparent.png" alt="logo">
    </div>
    <div class="menu">
      <MainMenu />
    </div>
    <router-view></router-view>
  </div>
</template>

<script>
import MainMenu from "./MainMenu.vue";
export default {
  name: "MainWrapper",
  components: { MainMenu },
  data() {
    return {}
  },
  created() {
  },
  methods: {

  }
}
</script>

<style scoped>
.page-logo {
  width: 80px;
}
</style>
