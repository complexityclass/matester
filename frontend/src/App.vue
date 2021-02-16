<template>
  <div id="app">
    <div class="main">
      <router-view></router-view>
    </div>
  </div>
</template>

<script>

export default {
  name: 'App',
}
</script>

<style scoped>
.main {
  display: flex;
  width: 100%;
  height: 100%;
  position: relative;
  flex-direction: column;
  background: transparent;
}
</style>
