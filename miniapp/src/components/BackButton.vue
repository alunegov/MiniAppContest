<!-- BackButton as Vue component -->

<script setup lang="ts">
  import WebApp from '@twa-dev/sdk';
  import { watchEffect } from 'vue';

  const emit = defineEmits<{
    // click event
    (event: 'click'): void;
  }>();

  // button click handler
  const onClick = () => emit('click');

  const button = WebApp.BackButton;

  // show on component create, hide on destroy
  watchEffect((onCleanup) => {
    button.show();
    onCleanup(() => {
      button.hide();
    });
  });

  // register click handler on component create, unregister on destroy
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
