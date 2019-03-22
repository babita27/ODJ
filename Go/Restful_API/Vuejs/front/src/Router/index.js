import Vue from 'vue'
import VueRouter from 'vue-router'
import Hello from '@/components/Hello'
import Calculator from '@/components/Calculator'
import Home from '@/components/Home' 
Vue.use(VueRouter)
export default new Router({
    routes: [
        {
            path:"/",
            component: Home,
            name:"Home",
            children:[{
                path: "/hello",
                name: "Hello",
                component: Hello
            },
            {
                path: "/calculator",
                name: "Calculator",
                component: Calculator
            }]
        }]
})