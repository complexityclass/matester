<template>
  <div class="page profile-wrapper">
    <div v-if="!isLoaded"> Loaded ...</div>
    <div v-else class="user-profile">
      <h1 class="page-title">Profile Page</h1>
      <div class="profile-info">
      <div class="info-item" v-for="(key, i) in infoKeys" :key="`${key}_${i}`">
        <span class="thead">{{ dic[key] }}</span>
        <span class="cell"> {{ profileData[key] }}</span>
      </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: "ProfilePage",
  components: {},
  props: {
    user_id: { type: String }
  },
  data() {
    return {
      isLoaded: false,
      dic: {
        birth_date: 'Birth date',
        city: 'City',
        first_name: 'First name',
        gender: 'Gender',
        hobbies: 'Hobbies',
        last_name: 'Last name',
        login: 'Login'

      },
      profileData: {},
      infoKeys: [],
    }
  },
  computed: {
    basicAuth() {
      return localStorage.getItem('matesterBasicAuth');
    },
    user() {
      return this.$store.getters['user/userData'].login || localStorage.getItem('matesterLogin');
    }
  },
  created() {
    console.log('loaded profile page')
    this.getProfile();
  },
  methods: {
    getProfile() {
      axios.get(`https://matester23.herokuapp.com/user?user=${this.user}`, {
        headers: { 'Authorization': this.basicAuth }
      }).then(userResponse => {
        console.log('userResponse', userResponse)
        Object.keys(userResponse.data).forEach(key => {
          this.infoKeys.push(key);
          switch(key) {
            case 'birth_date':
              this.profileData[key] = this.formatDate(userResponse.data[key].Time);
              break;
            case 'hobbies':
              this.profileData[key] = userResponse.data[key].join(', ');
              break;
            default:
              this.profileData[key] = userResponse.data[key]
          }
        })
        this.isLoaded = true;
      })
    },
    formatDate(date) {
      let currentDate = new Date(date);
      let year = currentDate.getFullYear();
      let month = currentDate.getMonth() + 1;
      let dt = currentDate.getDate();

      if (dt < 10) {
        dt = '0' + dt;
      }
      if (month < 10) {
        month = '0' + month;
      }

      return year + '-' + month + '-' + dt;
    },
  }
}
</script>

<style>
.page {
  background-color: #fff;
  padding: 10px;
}
.info-item {
  display: flex;
}
span.thead {
  min-width: 150px;
  text-align: left;
  padding: 10px 5px;
  border: 1px solid #cccccc;
}
span.cell {
  min-width: 120px;
  text-align: left;
  padding: 10px 5px;
  border-top: 1px solid #cccccc;
  border-right: 1px solid #cccccc;
  border-bottom: 1px solid #cccccc;


}
</style>
