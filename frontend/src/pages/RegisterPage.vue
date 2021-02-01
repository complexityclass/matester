<template>
  <div class="register page">
    <app-header/>
    <div class="form__wrapper d-flex-column">
      <form action="#" class="form d-flex-column">
        <div class="form-fields">
          <div v-for="field in fields" :key="field.id">
            <component :is="field.type + 'Field'" :data="field" @getValue="saveFieldValues"/>
          </div>
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
import SelectField from "../components/SelectField";
import DatepickerField from "../components/DatepickerField";
import axios from "axios";

export default {
  name: 'RegisterPage',
  components: {
    AppHeader,
    AppFooter,
    InputField,
    SelectField,
    DatepickerField
  },
  data() {
    return {
      values: {},
      fields: [
        { id: 'login', type: 'Input', data: { name: 'Login' } },
        { id: 'pass', type: 'Input', data: { name: 'Password' } },
        { id: 'first_name', type: 'Input', data: { name: 'First name' } },
        { id: 'last_name', type: 'Input', data: { name: 'Last name' } },
        { id: 'birth_date', type: 'datepicker', data: { name: 'Date of Birth' } },
        {
          id: 'gender',
          type: 'Select',
          name: 'Gender',
          data: {
            options: [{ id: 'male', label: 'Male' }, { id: 'female', label: 'Female' }, {
              id: 'other',
              label: 'Other'
            }],
            label: 'Gender'
          }
        },
        { id: 'city', type: 'Input', data: { name: 'City' } },
        { id: 'hobbies', type: 'Input', data: { name: 'Hobbies' }},
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
    saveFieldValues(data) {
      console.log('saveFieldValues', data)
      this.values[data.id] = data.label;
    },
    submitForm() {
      console.log('submitForm', this.values)
      this.values.hobbies = this.values.hobbies.split(',');
      let configData = {
        pass: this.values['pass'],
        profile: {
          login: this.values.login,
          first_name: this.values.first_name,
          last_name: this.values.last_name,
          birth_date: { Time: this.values.birth_date, Valid: true },
          gender: this.values.gender,
          city: this.values.city,
          hobbies: this.values.hobbies,
        }
      }
      console.log('configData', configData)
      axios.post(`https://matester23.herokuapp.com/register`, configData, {
        // headers: { 'X-Requested-With': 'XMLHttpRequest' },
      }).then(res => {
        console.log('register', res)
        this.$router.push({ path: `/login` });
        //23542345, asdf
      })
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
