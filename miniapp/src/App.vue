<script setup lang="ts">
  import { ref } from 'vue'
  import WebApp from '@twa-dev/sdk';
  import { useCounterStore } from './stores/counter';
  import type { Item } from './stores/counter';
  import BackButton from './components/BackButton.vue';
  import MainButton from './components/MainButton.vue';

  WebApp.ready();

  WebApp.expand();

  //console.log("sdafasdf");

  const counterStore = useCounterStore();

  const bb = ref(false);

  function onAddItemClicked() {
    counterStore.loadItems();
  }

  function onBuyClicked(item: Item) {
    counterStore.addItem(item);
  }

  function onUnbuyClicked(item: Item) {
    counterStore.removeItem(item);
  }
</script>

<template>
  <BackButton v-if="bb" @click="console.log('back')" />
  <div>test</div>
  <button @click="bb = !bb">Toggle TG buttons</button>
  <button @click="onAddItemClicked">Add item</button>
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
  <MainButton v-if="counterStore.gotSelectedItems" :text="'VIEW ORDER'" @click="console.log('main')" />
</template>
