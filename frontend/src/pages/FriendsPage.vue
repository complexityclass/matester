<template>
  <div>
    <div v-if="isLoaded">
      <h1 class="page-title">Friends Page</h1>
      <div v-if="title">
        <h4>{{ title }}</h4>
      </div>
      <table v-else class="table">
        <thead>
        <tr>
          <td v-for="thead in td_titles" :key="thead.id">
            {{ thead.title }}
          </td>
        </tr>
        </thead>
        <tbody>
        <tr v-for="(friend, i) in friends" :key="`user_${i}`">
          <td v-for="(thead, j) in td_titles" :key="`friend_cell_${thead.id}_${j}`">
            <span>
              {{ thead.id === 'birth_date' ? formatDate(friend[thead.id]['Time']) : friend[thead.id] }}
            </span>
            <span class="remove-icon" v-if="thead.id === 'delete_friend'">
              <span @click="deleteFriend(friend)">&#10008;</span>
            </span>
          </td>
        </tr>
        </tbody>
      </table>
    </div>
    <div v-if="!isLoaded">
      Loaded friends...
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: "FriendsPage",
  data() {
    return {
      title: '',
      isLoaded: false,
      friends: [],
      td_titles: [
        { id: "first_name", title: 'First name' },
        { id: "last_name", title: 'Last name' },
        { id: "login", title: 'Login' },
        { id: "birth_date", title: 'Birth date' },
        { id: "gender", title: 'Gender' },
        { id: "city", title: 'City' },
        { id: "delete_friend", title: 'Delete from friends' },
      ]
    }
  },
  computed: {
    basicAuth() {
      return localStorage.getItem('matesterBasicAuth');
    }
  },
  created() {
    this.getFriends();
  },
  methods: {
    getFriends() {
      axios.get(`https://matester23.herokuapp.com/friends`, {
        headers: { 'Authorization': this.basicAuth }
      }).then(friendsResponse => {
        // console.log('friendsResponse', friendsResponse)
        if (friendsResponse.data && friendsResponse.data.length !== 0) {
          this.friends = friendsResponse.data;
          this.isLoaded = true;
        } else {
          this.title = 'There are no friends in your list... Choose someone in all users'
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
    deleteFriend(user) {
      axios.patch(`https://matester23.herokuapp.com/unfriend?user=${user.login}`, {}, {
        headers: { 'Authorization': this.basicAuth }
      }).then(() => {
        let ind = this.friends.findIndex(_user => _user.login === user.login);
        this.friends.splice(ind, 1);
      })
    }
  }
}
</script>

<style>
.remove-icon:hover {
  cursor: pointer;
}
table {
  border-collapse: collapse;
  border-spacing: 0;
}
table thead td:first-child {
  border-top-left-radius: 5px;
}
table thead td:last-child {
  border-top-right-radius: 5px;
}
table thead td {
  font-weight: bold;
  width: 150px;
  text-align: left;
  padding: 12px 5px;
  background-color: #6F8BE9;
  color: whitesmoke;
}
table tbody td {
  text-align: left;
  border: 1px solid #ccc;
  padding: 5px 5px;
}
table tbody tr {
  transition: background-color 0.08s ease-in;
}
table tbody tr:hover {
  cursor: text;
  background-color: #d5e0fc !important;
}
table tbody tr:nth-child(2n) {
  background-color: aliceblue;
}
.page-title {
  font-weight: 400;
  text-align: left;
  margin-top: 0;
}
</style>
