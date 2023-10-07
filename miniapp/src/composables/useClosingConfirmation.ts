import { toValue, watchEffect } from 'vue';
import WebApp from '@twa-dev/sdk';

export function useClosingConfirmation(v: any) {
  if (!WebApp.isVersionAtLeast('6.2')) return;
  watchEffect(() => {
    if (toValue(v)) {
      //console.log('enableClosingConfirmation');
      WebApp.enableClosingConfirmation();
    } else {
      //console.log('disableClosingConfirmation');
      WebApp.disableClosingConfirmation();
    }
  });
}
