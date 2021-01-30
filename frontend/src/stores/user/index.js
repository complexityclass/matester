import axios from 'axios';

export default {
  namespaced: true,
  state() {
    return {
      userData: {
        login: '',
        first_name: '',
        last_name: '',
        birth_date: '',
        gender: '',
        city: '',
        hobbies: '',
      },
      basicAuth: '',
    }
  },
  mutations: {
    setData(state, data) {
      state.userData.login = data.login;
      state.userData.first_name = data.first_name;
      state.userData.last_name = data.last_name;
      state.userData.birth_date = data.birth_date;
      state.userData.gender = data.gender;
      state.userData.city = data.city;
      state.userData.hobbies = data.hobbies;
    },
    setAuth(state, basicAuth) {
      state.basicAuth = basicAuth;
    }
  },
  getters: {
    userData(state) {
      return state.userData;
    },
    basicAuth(state) {
      return state.basicAuth;
    }
  },
  actions: {
    login(context, basicAuth) {
      return axios.post('https://matester23.herokuapp.com', {}, {
        headers: { 'Authorization': basicAuth }
      }).then(authResponse => {
        console.log('Authenticated true', authResponse);
        context.commit('setAuth', basicAuth);
      }).catch(function(error) {
        console.log('Error on Authentication', error);
      });
    },
  }
}
