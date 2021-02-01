import RegisterPage from './RegisterPage.vue'
import LoginPage from './LoginPage.vue'
import FriendsPage from './FriendsPage.vue'
import ProfilePage from "./ProfilePage.vue";
import UsersPage from "./UsersPage.vue";
import Wrapper from "../components/Wrapper.vue";

export default [
  {
    path: '/',
    redirect: { name: 'registerPage' },
  },
  {
    path: '/register',
    props: true,
    name: 'registerPage',
    component: RegisterPage,
  },
  {
    path: '/login',
    props: true,
    name: 'loginPage',
    component: LoginPage,
  },
  {
    path: '',
    props: true,
    name: 'wrapper',
    component: Wrapper,
    children: [
      {
        path: '/friends',
        props: true,
        name: 'friendsPage',
        component: FriendsPage
      },
      {
        path: '/user',
        props: true,
        name: 'profilePage',
        component: ProfilePage
      },
      {
        path: '/users',
        props: true,
        name: 'userPage',
        component: UsersPage
      }
    ]
  },
]
