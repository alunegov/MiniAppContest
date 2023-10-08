<script setup lang="ts">
  import { computed } from 'vue';
  import { useRouter } from 'vue-router';
  import WebApp from '@twa-dev/sdk';
  import { useBaseStore } from '../stores/base';
  import BackButton from '../components/BackButton.vue';
  import MainButton from '../components/MainButton.vue';
  import { useClosingConfirmation } from '../composables/useClosingConfirmation';

  const router = useRouter();
  const baseStore = useBaseStore();

  if (!baseStore.isSmthSelected) {
    // empty cart (reload or direct link), going to GoodsView
    router.replace('/');
  }

  useClosingConfirmation(true);

  const total = computed(() => baseStore.selectedItems.map(it => it.qty * it.item.price).reduce((prev, it) => prev + it, 0));

  function onBackButtonClicked() {
    router.back();
  }

  async function onMainButtonClicked() {
    await baseStore.makeOrder();
    WebApp.showPopup({message: 'Done'}, () => WebApp.close());
  }
</script>

<template>
  <div class="container mx-auto px-5 pt-2">
    <BackButton @click="onBackButtonClicked" />

    <div class="flex mb-3">
      <div class="">YOUR ORDER</div>
      <button v-wave type="button" @click="onBackButtonClicked" class="ms-4 px-6 rounded text-[--tg-theme-button-color] hover:bg-[--tg-theme-secondary-bg-color]">Edit</button>
    </div>

    <div v-for="it in baseStore.selectedItems" :key="it.item.id" class="relative flex items-center mb-2">
      <img :src="it.item.pic" :alt="it.item.picAlt" class="w-10 h-10">
      <div class="ms-2">{{ it.item.name }}</div>
      <div class="ms-2 text-[--tg-theme-button-color]">{{ it.qty }}x</div>
      <div class="absolute end-0">${{ it.qty * it.item.price }}</div>
    </div>

    <div class="h-8 flex items-center place-content-end">
      <div class="font-bold">TOTAL</div>
      <div class="ms-5 font-bold">${{ total }}</div>
    </div>

    <MainButton :text="'PLACE ORDER'" @click="onMainButtonClicked" />
  </div>
</template>
