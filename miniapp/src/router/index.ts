import { createRouter, createWebHistory } from 'vue-router';
import P1 from '../views/P1.vue';

const routes = [
  {
    name: 'p1',
    path: '/',
    component: P1,
  },
  {
    name: 'p2',
    path: '/order',
    component: () => import('../views/P2.vue'),
  },
];

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
});

export default router;
