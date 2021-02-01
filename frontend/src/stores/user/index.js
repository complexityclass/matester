import axios from 'axios';
import { router } from "@/main";

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
      // console.log('setData'. data);
      localStorage.setItem('matesterLogin', data.login);
      state.userData.login = data.login;
      state.userData.first_name = data.first_name;
      state.userData.last_name = data.last_name;
      state.userData.birth_date = data.birth_date;
      state.userData.gender = data.gender;
      state.userData.city = data.city;
      state.userData.hobbies = data.hobbies;
    },
    setAuth(state, basicAuth) {
      // console.log('setAuth', basicAuth);
      state.basicAuth = basicAuth;
      localStorage.setItem('matesterBasicAuth', basicAuth);
    },
    openProfile() {
      router.push('/user');
    },
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
      // console.log('login in store', basicAuth)
      axios.post('https://matester23.herokuapp.com', {}, {
        headers: { 'Authorization': basicAuth }
      }).then(authResponse => {
        console.log('Authenticated true', authResponse);
        context.commit('setAuth', basicAuth);
        context.commit('setData', authResponse.data);
        context.commit('openProfile');
      }).catch(function(error) {
        console.log('Error on Authentication', error);
      });
    },
  }
}
