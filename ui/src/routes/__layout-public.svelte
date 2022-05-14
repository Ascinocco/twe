<script lang="ts">
  import { onMount } from "svelte";
  import { goto } from "$app/navigation";

  onMount(() => {
    const token = sessionStorage.getItem("token");

    if (token) {
      goto("/pmc")
    }
  })
</script>

<div class="navbar bg-base-100">
  <div class="flex-1">
    <a href="/" class="btn btn-ghost normal-case text-xl">The war economy</a>
  </div>
  <div class="flex-none">
    <ul class="menu menu-horizontal p-0">
      <li>
        <a href="/sign-in">Sign In</a>
        <a href="/sign-up">Sign Up</a>
      </li>
    </ul>
  </div>
</div>

<slot />