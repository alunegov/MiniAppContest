<script setup lang="ts">
  //console.log('P1');

  import { computed } from 'vue';
  import { useRouter } from 'vue-router';
  import { useCounterStore } from '../stores/counter';
  import type { Item } from '../stores/counter';
  import MainButton from '../components/MainButton.vue';
  import { useClosingConfirmation } from '../composables/useClosingConfirmation';

  const router = useRouter();
  const counterStore = useCounterStore();

  const isSmthSelected = computed(() => counterStore.isSmthSelected);
  useClosingConfirmation(isSmthSelected);

  function onBuyClicked(item: Item) {
    counterStore.addItem(item);
  }

  function onUnbuyClicked(item: Item) {
    counterStore.removeItem(item);
  }

  function onMainButtonClicked() {
    router.push('/order');
  }
</script>

<template>
  <div>
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
    <MainButton v-if="counterStore.isSmthSelected" :text="'VIEW ORDER'" @click="onMainButtonClicked" />
  </div>
</template>
