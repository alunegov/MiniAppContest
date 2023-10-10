import { computed, ref } from 'vue';
import { defineStore } from 'pinia';
import WebApp from '@twa-dev/sdk';

// Base store
export const useBaseStore = defineStore('base', () => {
  // server address
  const APP_API = import.meta.env.VITE_APP_API;

  // ops error text, mostly for loadItems and makeOrder
  const errorText = ref('');

  // shop items
  const items = ref<StoreItem[]>([]);

  // selected items
  const selectedItems = computed(() => items.value.filter(it => it.qty > 0));

  // flag: something has been selected
  const isSmthSelected = computed(() => selectedItems.value.length !== 0);

  // set error text, after 2.5 sec clears it
  function applyError(text: string) {
    errorText.value = text;
    if (errorText.value !== '') {
      setTimeout(() => clearError(), 2500);
    }
  }

  // clear error text
  function clearError() {
    errorText.value = '';
  }

  // load shop items from server
  async function loadItems() {
    try {
      clearError();

      const resp = await fetch(`${APP_API}/goods`, {
        headers: {
          'Ngrok-Skip-Browser-Warning': 'da',  // ngrok shows warn in free accounts
        },
      });
      const goods: Item[] = await resp.json();

      items.value = goods.map(it => ({item: it, qty: 0}));
    } catch (err) {
      applyError((err as Error).message);
    }
  }

  // add item to cart
  function buyItem(item: Item) {
    const indx = items.value.findIndex(it => it.item.id === item.id);
    if (indx === -1) {
      console.assert(false);
      return
    }
    items.value[indx].qty++;
  }

  // remove item from cart
  function unbuyItem(item: Item) {
    const indx = items.value.findIndex(it => it.item.id === item.id);
    if (indx === -1) {
      console.assert(false);
      return
    }
    console.assert(items.value[indx].qty > 0);
    items.value[indx].qty--;
  }

  // place order to server, it returns invoice link
  async function makeOrder(): Promise<boolean> {
    try {
      clearError();

      const payload = selectedItems.value.map<OrderItem>(it => ({id: it.item.id, qty: it.qty}));

      const resp = await fetch(`${APP_API}/order`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Init-Data': WebApp.initData,
          'Ngrok-Skip-Browser-Warning': 'da',  // ngrok shows warn in free accounts
        },
        body: JSON.stringify(payload),
      });
      const invoiceLink = await resp.json();

      if (!invoiceLink.ok) {
        applyError(invoiceLink.description);
        return false;
      }

      WebApp.openInvoice(invoiceLink.result, status => {
        if (status === 'paid') {
          WebApp.close();
        } else if (status === 'failed') {
          applyError('Payment failed');
        } else {
          applyError('Payment canceled');
        }
      });

      return true;
    } catch (err) {
      applyError((err as Error).message);
      return false;
    }
  }

  return {
    errorText,
    items,
    selectedItems,
    isSmthSelected,
    
    clearError,
    loadItems,
    buyItem,
    unbuyItem,
    makeOrder,
  };
})

// Shop item
// refs back.Item
export interface Item {
  id: number;
  name: string;
  price: number;
  pic: string;
  picAlt: string;
};

interface StoreItem {
  item: Item;
  // amount of selected items
  qty: number;
};

// refs back.OrderItem
interface OrderItem {
  // item id
  id: number;
  // amount of selected items
  qty: number;
};
