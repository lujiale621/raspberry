import { createRouter, createWebHistory } from "vue-router";
import Home from "../views/Home_Page.vue";
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [{ path: "/", name: "Home", component: Home }],
});

export default router;
