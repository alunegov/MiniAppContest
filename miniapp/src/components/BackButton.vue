<script setup lang="ts">
  import WebApp from '@twa-dev/sdk';
  import { watchEffect } from 'vue';

  const emit = defineEmits<{
    (event: 'click'): void;
  }>();

  const onClick = () => emit('click');

  const button = WebApp.BackButton;

  watchEffect((onCleanup) => {
    button.show();
    onCleanup(() => {
      button.hide();
    });
  });

  watchEffect((onCleanup) => {
    button.onClick(onClick);
    onCleanup(() => {
      button.offClick(onClick);
    });
  });
</script>

<template>
  <div />
</template>
