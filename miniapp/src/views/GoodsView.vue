<!-- Page shows list of shop items with prices and ability to select -->

<script setup lang="ts">
  import { computed, watchEffect } from 'vue';
  import { useRouter } from 'vue-router';
  import WebApp from '@twa-dev/sdk';
  import { useBaseStore } from '../stores/base';
  import type { Item } from '../stores/base';
  import ErrorText from '../components/ErrorText.vue';
  import MainButton from '../components/MainButton.vue';
  import { useClosingConfirmation } from '../composables/useClosingConfirmation';

  const router = useRouter();
  const baseStore = useBaseStore();

  // exit confirmation if something has been selected
  // baseStore.isSmthSelected looses reactivity being passed directly
  const isSmthSelected = computed(() => baseStore.isSmthSelected);
  useClosingConfirmation(isSmthSelected);

  watchEffect(() => {
    if (baseStore.errorText !== '') {
      WebApp.HapticFeedback.notificationOccurred('error');
    }
  });

  // add to cart handler
  function onBuyClicked(item: Item) {
    WebApp.HapticFeedback.selectionChanged();
    baseStore.buyItem(item);
  }

  // remove from cart handler
  function onUnbuyClicked(item: Item) {
    WebApp.HapticFeedback.selectionChanged();
    baseStore.unbuyItem(item);
  }

  // view order handler, go to OrderView
  function onMainButtonClicked() {
    router.push('/order');
  }
</script>

<template>
  <div class="container mx-auto px-5 pt-2">
    <div class="grid grid-cols-2 gap-5 sm:grid-cols-3 lg:grid-cols-4">
      <!-- Items list -->
      <div v-for="it in baseStore.items" :key="it.item.id" class="">
        <!-- Item image and selected counter -->
        <div class="relative w-full h-28">
          <span v-if="it.qty > 0" class="absolute end-0 rounded-full px-2 inline-flex justify-center items-center text-[--tg-theme-button-text-color] bg-[--tg-theme-button-color]">{{ it.qty }}</span>
          <img :src="it.item.pic" :alt="it.item.picAlt" class="h-full mx-auto">
        </div>

        <!-- Item name and price -->
        <div class="mt-2 flex justify-between">
          <h3 class="">{{ it.item.name }}</h3>
          <p class="">${{ it.item.price }}</p>
        </div>

        <!-- Add/remove to cart buttons -->
        <div class="mt-2 flex gap-2">
          <Transition name="resize-x">
            <button v-if="it.qty > 0" v-wave type="button" @click="onUnbuyClicked(it.item)" class="w-full h-11 rounded text-[--tg-theme-button-text-color] bg-red-400">-</button>
          </Transition>
          <button v-wave type="button" @click="onBuyClicked(it.item)" class="w-full h-11 rounded text-[--tg-theme-button-text-color] bg-[--tg-theme-button-color]">{{ it.qty === 0 ? 'ADD' : '+' }}</button>
        </div>
      </div>
    </div>

    <ErrorText :text="baseStore.errorText" />

    <MainButton v-if="baseStore.isSmthSelected" :text="'VIEW ORDER'" @click="onMainButtonClicked" />
  </div>
</template>

<style>
.resize-x-enter-active,
.resize-x-leave-active {
  transition: width 200ms ease;
}

.resize-x-enter-from,
.resize-x-leave-to {
  @apply w-0;
}
</style>
