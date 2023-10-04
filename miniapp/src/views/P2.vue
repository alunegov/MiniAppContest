<script setup lang="ts">
  import { useRouter } from 'vue-router';
  import WebApp from '@twa-dev/sdk';
  import { useCounterStore } from '../stores/counter';
  import BackButton from '../components/BackButton.vue';
  import MainButton from '../components/MainButton.vue';

  const router = useRouter();
  const counterStore = useCounterStore();

  function onBackButtonClicked() {
    console.log('back');
    router.back();
  }

  async function onMainButtonClicked() {
    console.log('main');
    await counterStore.makeOrder();
    //WebApp.close();
    WebApp.showPopup({message: 'done'});
  }
</script>

<template>
  <BackButton @click="onBackButtonClicked" />
  <ul>
    <li v-for="it in counterStore.items" :key="it.item.id">
      <div v-if="it.qty > 0">
        {{ it.item.name }}
        {{ it.qty }}x
        {{ it.qty * it.item.price }}
      </div>
    </li>
  </ul>
  <MainButton :text="'MAKE ORDER'" @click="onMainButtonClicked" />
</template>
