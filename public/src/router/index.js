import Vue from 'vue'
import Router from 'vue-router'
import Vuetify from 'vuetify'
import 'vuetify/dist/vuetify.min.css'
import ViewBlogs from '@/components/ViewBlogs'
import PostBlogs from '@/components/PostBlogs'

Vue.use(Router)
Vue.use(Vuetify)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'ViewBlogs',
      component: ViewBlogs
    },

    {
      path: '/postBlogs',
      name: 'PostBlogs',
      component: PostBlogs
    }
  ]
})
