import { createRouter, createWebHistory } from 'vue-router'
import Home from './pages/Home.vue'
import Post from './pages/Post.vue'
import About from './pages/About.vue'
import Projects from './pages/Projects.vue'
import Resume from './pages/Resume.vue'
import Subscribe from './pages/Subscribe.vue'

const routerHistory = createWebHistory()

const router = createRouter({
  scrollBehavior(to) {
    if (to.hash) {
      window.scroll({ top: 0 })
    } else {
      document.querySelector('html').style.scrollBehavior = 'auto'
      window.scroll({ top: 0 })
      document.querySelector('html').style.scrollBehavior = ''
    }
  },  
  history: routerHistory,
  routes: [
    {
      path: '/',
      component: Home
    },
    {
      path: '/post',
      component: Post
    },
    {
      path: '/about',
      component: About
    },
    {
      path: '/projects',
      component: Projects
    },
    {
      path: '/resume',
      component: Resume
    },        
    {
      path: '/subscribe',
      component: Subscribe
    },
  ]
})

export default router
