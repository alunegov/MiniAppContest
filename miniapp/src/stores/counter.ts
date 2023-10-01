import { computed, ref } from 'vue'
import { defineStore } from 'pinia'

export const useCounterStore = defineStore('counter', () => {
  let id: number = 0;

  const items = ref<{item: Item; qty: number}[]>([
    {item: {id: id++, name: "Pen", price: 33, pic: ""}, qty: 0},
    {item: {id: id++, name: "Pineapple", price: 33, pic: ""}, qty: 0},
    {item: {id: id++, name: "Apple", price: 33, pic: ""}, qty: 0},
  ]);

  const gotSelectedItems = computed(() => items.value.findIndex(it => it.qty > 0) !== -1);

  async function loadItems() {
    items.value.push({item: {id: id++, name: "smth", price: 99, pic: ""}, qty: 0});
    /*const resp = await fetch("");
    const body = await resp.json();
    return body;*/
  }

  function addItem(item: Item) {
    const indx = items.value.findIndex(it => it.item.id === item.id);
    items.value[indx].qty++;
  }

  function removeItem(item: Item) {
    const indx = items.value.findIndex(it => it.item.id === item.id);
    items.value[indx].qty--;
  }

  return { items, gotSelectedItems, loadItems, addItem, removeItem };
})

export interface Item {
  id: number;
  name: string;
  price: number;
  pic: string;
}
