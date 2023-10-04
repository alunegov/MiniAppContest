import { computed, ref } from 'vue'
import { defineStore } from 'pinia'

export const useCounterStore = defineStore('counter', () => {
  const APP_API = import.meta.env.VITE_APP_API;

  const items = ref<{item: Item; qty: number}[]>([]);

  const gotSelectedItems = computed(() => items.value.findIndex(it => it.qty > 0) !== -1);

  async function loadItems() {
    const resp = await fetch(`${APP_API}/goods`);
    console.log(resp);
    const goods: Item[] = await resp.json();
    items.value = goods.map(it => ({item: it, qty: 0}));
  }

  function addItem(item: Item) {
    const indx = items.value.findIndex(it => it.item.id === item.id);
    items.value[indx].qty++;
  }

  function removeItem(item: Item) {
    const indx = items.value.findIndex(it => it.item.id === item.id);
    items.value[indx].qty--;
  }

  async function makeOrder() {
    const resp = await fetch(`${APP_API}/order`, {
      method: 'POST',
      headers: {
        "Content-Type": "application/json",
      },
      body: "[]",
    });
    console.log(resp);
  }

  return {
    items,
    gotSelectedItems,
    
    loadItems,
    addItem,
    removeItem,
    makeOrder,
  };
})

export interface Item {
  id: number;
  name: string;
  price: number;
  pic: string;
}
