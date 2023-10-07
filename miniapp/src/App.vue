<script setup lang="ts">
  import WebApp from '@twa-dev/sdk';
  import { useBaseStore } from './stores/base';

  WebApp.ready();
  //WebApp.expand();

  const baseStore = useBaseStore();
  baseStore.loadItems();
</script>

<template>
  <router-view />
</template>
