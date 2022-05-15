import { writable } from "svelte/store";

import type { User } from "./types";

const defaultUser: User = {
  id: "",
  email: "",
  username: "",
};

const createStore = () => {
  const { subscribe, set, update } = writable<User>(defaultUser);

  return {
    subscribe,
    update: (user: User) =>
      update((cu) => ({
        ...cu,
        ...user,
      })),
    clear: () => set(defaultUser),
  };
};

export const userStore = createStore();
