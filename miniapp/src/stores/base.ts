import { computed, ref } from 'vue'
import { defineStore } from 'pinia'

export const useBaseStore = defineStore('base', () => {
  const APP_API = import.meta.env.VITE_APP_API;

  const items = ref<{item: Item; qty: number}[]>([]);

  const selectedItems = computed(() => items.value.filter(it => it.qty > 0));

  const isSmthSelected = computed(() => selectedItems.value.length !== 0);

  async function loadItems() {
    const resp = await fetch(`${APP_API}/goods`, {
      headers: {
        'Ngrok-Skip-Browser-Warning': 'da',
      },
    });
    const goods: Item[] = await resp.json();
    items.value = goods.map(it => ({item: it, qty: 0}));
  }

  function buyItem(item: Item) {
    const indx = items.value.findIndex(it => it.item.id === item.id);
    if (indx === -1) {
      // TODO: error
      return
    }
    items.value[indx].qty++;
  }

  function unbuyItem(item: Item) {
    const indx = items.value.findIndex(it => it.item.id === item.id);
    if (indx === -1) {
      // TODO: error
      return
    }
    console.assert(items.value[indx].qty > 0);
    items.value[indx].qty--;
  }

  async function makeOrder() {
    const payload = selectedItems.value.map(it => ({id: it.item.id, qty: it.qty}));

    /*const resp = */await fetch(`${APP_API}/order`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Ngrok-Skip-Browser-Warning': 'da',
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

export interface Item {
  id: number;
  name: string;
  price: number;
  pic: string;
  picAlt: string;
}
