import RegisterPage from './RegisterPage.vue'
import LoginPage from './LoginPage.vue'
import FriendsPage from './FriendsPage'
import ProfilePage from "./ProfilePage";

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
    path: '/friends',
    props: true,
    name: 'friendsPage',
    component: FriendsPage
  },
  {
    path: '/user/:user_id',
    props: true,
    name: 'profilePage',
    component: ProfilePage
  }
]
