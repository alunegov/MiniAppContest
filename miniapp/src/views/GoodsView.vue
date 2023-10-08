<script setup lang="ts">
  import { computed } from 'vue';
  import { useRouter } from 'vue-router';
  import { useBaseStore } from '../stores/base';
  import type { Item } from '../stores/base';
  import MainButton from '../components/MainButton.vue';
  import { useClosingConfirmation } from '../composables/useClosingConfirmation';

  const router = useRouter();
  const baseStore = useBaseStore();

  const isSmthSelected = computed(() => baseStore.isSmthSelected);
  useClosingConfirmation(isSmthSelected);

  function onBuyClicked(item: Item) {
    baseStore.addItem(item);
  }

  function onUnbuyClicked(item: Item) {
    baseStore.removeItem(item);
  }

  function onMainButtonClicked() {
    router.push('/order');
  }
</script>

<template>
  <div class="container mx-auto px-5 pt-2">
    <div class="grid grid-cols-2 gap-5 sm:grid-cols-3 lg:grid-cols-4">
      <div v-for="it in baseStore.items" :key="it.item.id" class="">
        <div class="relative w-full h-28">
          <span v-if="it.qty > 0" class="absolute end-0 rounded-full px-2 inline-flex justify-center items-center text-[--tg-theme-button-text-color] bg-[--tg-theme-button-color]">{{ it.qty }}</span>
          <img :src="it.item.pic" :alt="it.item.picAlt" class="h-full mx-auto">
        </div>

        <div class="mt-2 flex justify-between">
          <h3 class="">{{ it.item.name }}</h3>
          <p class="">${{ it.item.price }}</p>
        </div>

        <div class="mt-2 flex gap-2">
          <Transition name="ba">
            <button v-if="it.qty > 0" v-wave type="button" @click="onUnbuyClicked(it.item)" class="w-full h-11 rounded text-[--tg-theme-button-text-color] bg-red-400">-</button>
          </Transition>
          <button v-wave type="button" @click="onBuyClicked(it.item)" class="w-full h-11 rounded text-[--tg-theme-button-text-color] bg-[--tg-theme-button-color]">{{ it.qty === 0 ? 'ADD' : '+' }}</button>
        </div>
      </div>
    </div>

    <MainButton v-if="baseStore.isSmthSelected" :text="'VIEW ORDER'" @click="onMainButtonClicked" />
  </div>
</template>

<style>
.ba-enter-active,
.ba-leave-active {
  transition: width 200ms ease;
}

.ba-enter-from,
.ba-leave-to {
  @apply w-0;
}
</style>
