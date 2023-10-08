import { createRouter, createWebHistory } from 'vue-router';
import GoodsView from '../views/GoodsView.vue';

const routes = [
  {
    name: 'GoodsView',
    path: '/',
    component: GoodsView,
  },
  {
    name: 'OrderView',
    path: '/order',
    component: () => import('../views/OrderView.vue'),
  },
];

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
});

export default router;
