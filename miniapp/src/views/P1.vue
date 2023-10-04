<script setup lang="ts">
  import { useRouter } from 'vue-router';
  import { useCounterStore } from '../stores/counter';
  import type { Item } from '../stores/counter';
  import MainButton from '../components/MainButton.vue';

  const router = useRouter();
  const counterStore = useCounterStore();

  counterStore.loadItems();

  function onBuyClicked(item: Item) {
    counterStore.addItem(item);
  }

  function onUnbuyClicked(item: Item) {
    counterStore.removeItem(item);
  }

  function onMainButtonClicked() {
    console.log('main');
    router.push('/order');
  }
</script>

<template>
  <ul>
    <li v-for="it in counterStore.items" :key="it.item.id">
      {{ it.item.id }}
      {{ it.item.name }}
      {{ it.item.price }}
      {{ it.item.pic }}
      {{ it.qty }}
      <button v-if="it.qty>0" @click="onUnbuyClicked(it.item)">Unbuy</button>
      <button @click="onBuyClicked(it.item)">Buy</button>
    </li>
  </ul>
  <MainButton v-if="counterStore.gotSelectedItems" :text="'VIEW ORDER'" @click="onMainButtonClicked" />
</template>
