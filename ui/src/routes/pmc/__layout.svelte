<script lang="ts">
  import { goto } from "$app/navigation";
  import { beforeUpdate } from 'svelte';

  const logout = () => {
    sessionStorage.setItem("token", "");
    goto("/sign-in");
  }

  beforeUpdate(() => {
    const token = sessionStorage.getItem("token");

    if (!token) {
      goto("/sign-in")
    }
  });

</script>

<div>
  <button on:click={logout}>logout</button>
</div>
<slot />