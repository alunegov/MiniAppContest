<script setup lang="ts">
  //console.log('P2');

  import { useRouter } from 'vue-router';
  import WebApp from '@twa-dev/sdk';
  import { useCounterStore } from '../stores/counter';
  import BackButton from '../components/BackButton.vue';
  import MainButton from '../components/MainButton.vue';
  import { useClosingConfirmation } from '../composables/useClosingConfirmation';

  const router = useRouter();
  const counterStore = useCounterStore();

  if (!counterStore.isSmthSelected) {
    //console.log('empty, going to P1');
    router.replace('/');
  }

  useClosingConfirmation(true);

  function onBackButtonClicked() {
    router.back();
  }

  async function onMainButtonClicked() {
    await counterStore.makeOrder();
    WebApp.showPopup({message: 'done'}, _ => WebApp.close());
    //WebApp.close();
  }
</script>

<template>
  <div>
    <BackButton @click="onBackButtonClicked" />
    <ul>
      <li v-for="it in counterStore.selectedItems" :key="it.item.id">
        {{ it.item.name }}
        {{ it.qty }}x
        {{ it.qty * it.item.price }}
      </li>
    </ul>
    <MainButton :text="'MAKE ORDER'" @click="onMainButtonClicked" />
  </div>
</template>
