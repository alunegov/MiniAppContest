import { computed, ref } from 'vue'
import { defineStore } from 'pinia'

// Base store
export const useBaseStore = defineStore('base', () => {
  // server address
  const APP_API = import.meta.env.VITE_APP_API;

  // shop items
  const items = ref<{
    item: Item;
    qty: number;
  }[]>([]);

  // selected items
  const selectedItems = computed(() => items.value.filter(it => it.qty > 0));

  // flag: something has been selected
  const isSmthSelected = computed(() => selectedItems.value.length !== 0);

  // load shop items from server
  async function loadItems() {
    const resp = await fetch(`${APP_API}/goods`, {
      headers: {
        'Ngrok-Skip-Browser-Warning': 'da',  // ngrok shows warn in free accounts
      },
    });
    const goods: Item[] = await resp.json();
    items.value = goods.map(it => ({item: it, qty: 0}));
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

  // place order to server
  async function makeOrder() {
    const payload = selectedItems.value.map<{
      id: number;
      qty: number;
    }>(it => ({id: it.item.id, qty: it.qty}));

    /*const resp = */await fetch(`${APP_API}/order`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Ngrok-Skip-Browser-Warning': 'da',  // ngrok shows warn in free accounts
      },
      body: JSON.stringify(payload),
    });
  }

  return {
    items,
    selectedItems,
    isSmthSelected,
    
    loadItems,
    buyItem,
    unbuyItem,
    makeOrder,
  };
})

// Shop item
export interface Item {
  id: number;
  name: string;
  price: number;
  pic: string;
  picAlt: string;
}
