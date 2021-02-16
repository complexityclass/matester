<template>
  <footer class="main-footer">
    &#169;Matester, {{ currentYear }}
  </footer>
</template>

<script>
export default {
  name: 'AppFooter',
  data() {
    return {
      currentYear: ""
    }
  },
  mounted() {
    this.setCurrentYear();
  },
  methods: {
    setCurrentYear() {
      let date = new Date();
      this.currentYear = date.getFullYear();
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.main-footer {
  height: 50px;
}
</style>
