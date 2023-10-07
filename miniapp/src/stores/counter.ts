import { computed, ref } from 'vue'
import { defineStore } from 'pinia'

export const useCounterStore = defineStore('counter', () => {
  const APP_API = import.meta.env.VITE_APP_API;

  const items = ref<{item: Item; qty: number}[]>([]);

  const selectedItems = computed(() => items.value.filter(it => it.qty > 0));

  const isSmthSelected = computed(() => selectedItems.value.length !== 0);

  async function loadItems() {
    const resp = await fetch(`${APP_API}/goods`);
    //console.log(resp);
    const goods: Item[] = await resp.json();
    items.value = goods.map(it => ({item: it, qty: 0}));
  }

  function addItem(item: Item) {
    const indx = items.value.findIndex(it => it.item.id === item.id);
    if (indx === -1) {
      // TODO: error
      return
    }
    items.value[indx].qty++;
  }

  function removeItem(item: Item) {
    const indx = items.value.findIndex(it => it.item.id === item.id);
    if (indx === -1) {
      // TODO: error
      return
    }
    console.assert(items.value[indx].qty > 0);
    items.value[indx].qty--;
  }

  async function makeOrder() {
    const resp = await fetch(`${APP_API}/order`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(selectedItems.value),
    });
    //console.log(resp);
  }

  return {
    items,
    selectedItems,
    isSmthSelected,
    
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
