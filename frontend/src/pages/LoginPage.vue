<template>
  <div class="login page">
    <app-header/>
    <div class="form__wrapper d-flex-column">
      <form action="#" class="form d-flex-column">
        <div class="form-fields">
          <input-field v-for="field in fields" :key="field.id" :field-type="field.type" :id="field.id"
                       :field-name="field.name" @getValue="saveFieldValues"/>
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
        { id: 'login', type: 'text', name: 'Login' },
        { id: 'password', type: 'password', name: 'Password' },
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
    saveFieldValues(id, value) {
      this.values[id] = value;
    },
    submitForm() {
      let basicAuth = 'Basic ' + btoa(this.values.login + ':' + this.values.password);
      this.$store.dispatch('user/login', basicAuth).then(res => {
        console.log('dispatch user/login was successful', res);
        this.$router.push({ path: `/user/${this.values.login}` })
      })
    },
  }
}
</script>

<style scoped>

</style>
