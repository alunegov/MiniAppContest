import { toValue, watchEffect } from 'vue';
import WebApp from '@twa-dev/sdk';

export function useClosingConfirmation(v: any) {
  if (!WebApp.isVersionAtLeast('6.2')) return;
  watchEffect(() => {
    if (toValue(v)) {
      WebApp.enableClosingConfirmation();
    } else {
      WebApp.disableClosingConfirmation();
    }
  });
}
