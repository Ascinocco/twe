import { writable } from "svelte/store";
import type { Pmc } from "./types";

const defaultPmc: Pmc = {
  id: "",
  userId: "",
  name: "",
};

const createStore = () => {
  const { subscribe, set, update } = writable<Pmc>(defaultPmc);

  return {
    subscribe,
    update: (pmc: Pmc) =>
      update((cPmc) => ({
        ...cPmc,
        ...pmc,
      })),
    clear: () => set(defaultPmc),
  };
};

export const pmcStore = createStore();
