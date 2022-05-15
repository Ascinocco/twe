import { getWsUrl } from "$lib/utils/network";

export const wsHandler = () => {
  const socket = new WebSocket(getWsUrl());

  socket.onopen = () => {
    console.log("connected");
    // @TODO: fetch user
    // @TODO: fetch pmc
    socket.send(
      JSON.stringify({
        token: sessionStorage.getItem("token"),
        type: "hello",
      }),
    );
  };

  socket.onclose = (e) => {
    console.log("socket closed", e);
  };

  socket.onerror = (err) => {
    console.log("socket error", err);
  };

  socket.onmessage = (msg) => {
    // @TODO: Figure out how to structure message handlers
    // @TODO: Somewhere you'll need to leverage decoders after
    // parsing the message data
    console.log("message recieved", JSON.parse(msg.data));
  };

  return socket;
};
