import { getWsUrl } from "$lib/utils/network";

export const wsHandler = () => {
  const socket = new WebSocket(getWsUrl());

  socket.onopen = () => {
    console.log("connected");
    // @TODO: fetch user
    // @TODO: fetch pmc
    socket.send("hi from client");
  };

  socket.onclose = (e) => {
    console.log("socket closed", e);
  };

  socket.onerror = (err) => {
    console.log("socket error", err);
  };

  socket.onmessage = (msg) => {
    console.log("message recieved", msg);
  };

  return socket;
};
