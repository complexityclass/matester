<template>
  <div>
    <div v-if="isLoaded">
      <h1 class="page-title">Users Page</h1>
      <table>
        <thead>
        <tr>
          <td v-for="thead in td_titles" :key="thead.id">
            {{ thead.title }}
          </td>
        </tr>
        </thead>
        <tbody>
        <tr v-for="(user, i) in users" :key="`user_${i}`">
          <td v-for="(thead, j) in td_titles" :key="`user_cell_${thead.id}_${j}`">
            <span>
              {{ thead.id === 'birth_date' ? formatDate(user[thead.id]['Time']) : user[thead.id] }}
            </span>
            <span class="add-icon" v-if="thead.id === 'add_friend'">
              <span v-if="!friends[user.login]" @click="addFriend(user)">&#10010;</span>
              <span v-if="friends[user.login]">&#10004;</span>
            </span>
          </td>
        </tr>
        </tbody>
      </table>
    </div>
    <div v-if="!isLoaded">
      Loaded all users...
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: "UsersPage",
  data() {
    return {
      isLoaded: false,
      users: [],
      friends: {},
      td_titles: [
        { id: "first_name", title: 'First name' },
        { id: "last_name", title: 'Last name' },
        { id: "login", title: 'Login' },
        { id: "birth_date", title: 'Birth date' },
        { id: "gender", title: 'Gender' },
        { id: "city", title: 'City' },
        { id: "add_friend", title: 'Add to friends' },
      ]
    }
  },
  computed: {
    basicAuth() {
      return localStorage.getItem('matesterBasicAuth');
    }
  },
  created() {
    this.getUsers();
  },
  methods: {
    getUsers() {
      axios.get(`https://matester23.herokuapp.com/users`, {
        headers: { 'Authorization': this.basicAuth }
      }).then(usersResponse => {
        console.log('usersResponse', usersResponse)
        if (usersResponse.data && usersResponse.data.length !== 0) {
          this.users = usersResponse.data;
          this.isLoaded = true;
        }
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
    addFriend(user) {
      console.log('add friend', user);
      this.friends[`${user.login}`] = true;
    }
  }
}
</script>

<style scoped>
.add-icon:hover {
  cursor: pointer;
}
</style>
