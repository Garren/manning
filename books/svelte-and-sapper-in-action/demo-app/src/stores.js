import { derived, writable } from 'svelte/store';

export const dogStore = writable({});

const items = [
  { name: 'pencil', cost: 0.5 },
  { name: 'backpack', cost: 40 }
];
export const itemsStore = writable(items);

export const taxStore = writable(0.08);
export const itemsWithTaxStore = derived(
  [itemsStore, taxStore],
  ([$itemsStore, $taxStore]) => {
    const tax = 1 + $taxStore;
    return $itemsStore.map(
      item => ({ ...item, total: item.cost * tax })
    );
  }
);

const { subscribe, set, update } = writable(0);
export const count = {
  subscribe,
  increment: () => update(n => n + 1),
  decrement: () => update(n => n - 1),
  reset: () => set(0)
};

import Line from './line';
import Point from './point';
export const pointStore = writable(new Point(0, 0));
export const lineStore = writable(new Line(new Point(0, 0), new Point(0, 0)));
