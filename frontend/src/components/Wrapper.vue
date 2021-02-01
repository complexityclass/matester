<template>
  <div class="container">
    <header class="header">
      <img class="header-logo" src="../assets/logo_transparent.png" alt="header_logo">
    </header>
    <div class="common-content">
      <main-menu/>
      <div class="workarea">
        <div class="content">
          <router-view></router-view>
        </div>
      </div>
    </div>
    <footer>
      &#169;Matester, 2021
    </footer>
  </div>
</template>

<script>
import MainMenu from "./MainMenu.vue";

export default {
  name: "Wrapper",
  components: { MainMenu }
}
</script>

<style scoped>
.header-logo {
  width: 80px;
}
.container {
  display: flex;
  flex-direction: column;
  width: auto;
  margin: 0 60px;
  flex: 1 0 auto;
  margin-bottom: 40px;
}
header.header {
  display: flex;
  padding-top: 20px;
  padding-bottom: 20px;
}
.common-content {
  display: flex;
  height: 100%;
}
.workarea {
  height: 100%;
  margin-left: 20px;
  width: 100%;
  min-width: 78%;
}
footer {
  margin-top: auto;
}
</style>
