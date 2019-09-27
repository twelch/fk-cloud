import Vue from "vue";
import Router from "vue-router";
import Login from "./views/LoginView.vue";
import DataView from "./views/DataView.vue";
import ProjectsView from "./views/ProjectsView.vue";
import StationsView from "./views/StationsView.vue";

Vue.use(Router);

export default new Router({
    mode: "history",
    base: process.env.BASE_URL,
    routes: [
        {
            path: "/",
            alias: "/login",
            name: "login",
            component: Login
        },
        {
            path: "/dashboard",
            name: "projects",
            component: ProjectsView
        },
        {
            path: "/dashboard/stations",
            name: "stations",
            component: StationsView
        },
        {
            path: "/dashboard/data",
            name: "data",
            component: DataView
        }
    ]
});
