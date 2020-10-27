<template>
    <div>123</div>
</template>

<script>
export default {
    name: 'MyFile',
    data() {
        return {
            ws: null,
            serverUrl: "ws://localhost:8080/ws"
        }
    },
    mounted: function() {
      this.connectToWebsocket()
    },
    methods: {
      connectToWebsocket() {
        this.ws = new WebSocket( this.serverUrl );
        this.ws.addEventListener('open', (event) => { this.onWebsocketOpen(event) });
      },
      onWebsocketOpen() {
        console.log("connected to WS!");
      }
    }
}
</script>
