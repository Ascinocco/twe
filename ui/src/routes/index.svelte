<script lang="ts">
  import { onMount } from "svelte";

  onMount(() => {
    let socket = new WebSocket("ws://localhost:8080/ws")

    socket.onopen = () => {
      console.log("connected");
      socket.send("hi from client");
    }
    
    socket.onclose = (e) => {
      console.log("socket closed", e)
    }
    
    socket.onerror = (err) => {
      console.log("socket error", err);
    }
    
    socket.onmessage = (msg) => {
      console.log("message recieved", msg)
    }
  })
</script>

<h1 class="text-3xl font-bold underline">
  Hello world!
</h1>