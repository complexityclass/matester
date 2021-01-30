<template>
  <div class="register page">
    <app-header/>
    <div class="form__wrapper d-flex-column">
      <form action="#" class="form d-flex-column">
        <div class="form-fields">
          <input-field v-for="field in fields" :key="field.id" :field-type="field.type" :id="field.id"
                       :field-name="field.name" @getValue="saveFieldValues"/>
          <router-link class="router-link" to="/login">Already have account?</router-link>
        </div>
        <input class="btn btn-submit" type="button" value="Register" @click="submitForm">
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
  name: 'RegisterPage',
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
        { id: 'pass', type: 'password', name: 'Login' },
        { id: 'first_name', type: 'text', name: 'First name' },
        { id: 'last_name', type: 'text', name: 'Last name' },
        { id: 'birth_date', type: 'datepicker', name: 'Date of Birth' },
        { id: 'gender', type: 'select', name: 'Gender' },
        { id: 'city', type: 'text', name: 'City' },
        { id: 'hobbies', type: 'text', name: 'Hobbies' },
      ]
    }
  },
  mounted() {
    this.initValues()
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
      console.log('submitForm', this.values)
      // this.$http.post(`https://matester23.herokuapp.com/register`, this.values, {
      //   headers: { 'X-Requested-With': 'XMLHttpRequest' },
      // }).then(res => {
      //   console.log('result', res)
      // })
    },
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style>
.form__wrapper {
  margin-top: 30px;

}

.form {
  width: 50%;
}

.form-fields {
  display: flex;
  flex-direction: column;
  width: 60%;
  margin-bottom: 20px;
}

.router-link {
  align-self: flex-end;
  margin-top: 1rem;
}

.btn-submit {
  color: #ffffff;
  border-style: none;
  background: #6F8BE9;
  padding: 5px;
  border-radius: 4px;
  width: 60%;
  font-size: 16px;
  font-family: inherit;
  transition: background-color 0.3s ease-in;
}

.btn-submit:hover {
  cursor: pointer;
  background-color: #6b84d7;
}

.btn-submit:focus {
  cursor: pointer;
  background-color: #526ab7;
}
</style>
