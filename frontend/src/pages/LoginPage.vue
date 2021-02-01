<template>
  <div class="login page">
    <app-header/>
    <div class="form__wrapper d-flex-column">
      <form action="#" class="form d-flex-column">
        <div class="form-fields">
          <input-field v-for="field in fields" :key="field.id"
                       :data="field" @getValue="saveFieldValues"/>
          <router-link class="router-link" to="/register">No account?</router-link>
        </div>
        <input class="btn btn-submit" type="button" value="Login" @click="submitForm">
      </form>
    </div>
    <app-footer/>
  </div>
</template>

<script>
import AppHeader from "../components/AppHeader";
import AppFooter from "../components/AppFooter";
import InputField from "../components/InputField";

export default {
  name: "LoginPage",
  components: {
    AppHeader,
    AppFooter,
    InputField
  },
  data() {
    return {
      values: {},
      fields: [
        { id: 'login', type: 'text', data: {name: 'Login'} },
        { id: 'password', type: 'password', data: {name: 'Password'} },
      ]
    }
  },
  mounted() {
    this.initValues();
  },
  methods: {
    initValues() {
      this.fields.forEach(field => {
        this.values[field.id] = "";
      })
    },
    saveFieldValues(data) {
      // console.log('saveFieldValues', data)
      this.values[data.id] = data.label;
    },
    submitForm() {
      // console.log('click login')
      let basicAuth = 'Basic ' + btoa(this.values.login + ':' + this.values.password);
      this.$store.dispatch('user/login', basicAuth).then(res => {
        console.log('dispatch user/login was successful', res);
      })
    },
  }
}
</script>

<style scoped>

</style>
