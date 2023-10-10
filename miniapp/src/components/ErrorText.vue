<script setup lang="ts">
  /*const props = */defineProps<{
    // text property
    text: string;
  }>();
</script>

<template>
  <Transition name="fade">
    <div v-if="text !== ''" class="mt-6 text-red-600">{{ text }}</div>
  </Transition>
</template>

<style>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 200ms ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
