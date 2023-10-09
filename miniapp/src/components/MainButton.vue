<!-- MainButton as Vue component -->

<script setup lang="ts">
  import WebApp from '@twa-dev/sdk';
  import { watchEffect } from 'vue';

  const props = defineProps<{
    // text property
    text: string;
  }>();

  const emit = defineEmits<{
    // click event
    (event: 'click'): void;
  }>();

  // button click handler
  const onClick = () => emit('click');

  const button = WebApp.MainButton;

  // show on component create, hide on destroy
  watchEffect((onCleanup) => {
    button.show();
    onCleanup(() => {
      button.hide();
    });
  });

  // update text
  watchEffect(() => button.setText(props.text));

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
