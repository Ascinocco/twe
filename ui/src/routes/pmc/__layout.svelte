<script lang="ts">
  import { goto } from "$app/navigation";
  import { beforeUpdate, onMount } from "svelte";
  import { wsHandler } from "$lib/messaging";

  const logout = () => {
    sessionStorage.setItem("token", "");
    goto("/sign-in");
  };

  beforeUpdate(() => {
    const token = sessionStorage.getItem("token");

    if (!token) {
      goto("/sign-in");
    }
  });

  onMount(() => {
    const socket = wsHandler();
    return () => socket.close();
  });
</script>

<div class="navbar bg-base-100">
  <div class="flex-1">
    <!-- @TODO: replace "the war economy" with pmc name -->
    <!-- @TODO: Add drop down to change pmc's -->
    <a href="/pmc" class="btn btn-ghost normal-case text-xl">The war economy</a>
  </div>
  <div class="flex-none">
    <ul class="menu menu-horizontal p-0">
      <li>
        <span>user name</span>
      </li>
      <li>
        <button on:click={logout}>logout</button>
      </li>
    </ul>
  </div>
</div>

<slot />
