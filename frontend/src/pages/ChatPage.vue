<template>
  <v-container>
    <h1>{{ roomInfo.title }}</h1>
    <div class="chatList">
      <v-row>
        <v-col cols="12">
          <li v-for="message in messages" :key="message.uuid" class="pb-1">
            <v-card
              outlined
            >
              <v-list-item single-line style="text-align: left;">
                <v-list-item-content>
                  {{ message.value }}
                </v-list-item-content>
              </v-list-item>
            </v-card>
          </li>
        </v-col>
      </v-row>
    </div>
    <v-footer fixed color="white">
      <v-container class="ma-0 pa-0">
        <v-row no-gutters>
          <v-col>
            <div class="d-flex flex-row align-center">
              <template v-if="state === 'idle'">
                <v-text-field autofocus label="Message" v-model="postMessage" @keypress.enter="sendMessage"/>
                <v-btn v-on:click="sendMessage" icon class="ml-4">
                  <v-icon>mdi-send</v-icon>
                </v-btn>
              </template>
              <template v-else-if="state === 'sendError'">
                <v-text-field label="Message" v-model="postMessage" error/>
                <v-btn v-on:click="retrySendMessage" class="ml-4">
                  Retry
                </v-btn>
                <v-btn v-on:click="clearMessage" class="ml-4" color="error">
                  Cancel
                </v-btn>
              </template>
            </div>
          </v-col>
        </v-row>
      </v-container>
    </v-footer>
  </v-container>
</template>

<script>
  const axios = require('axios')
  export default {
    name: 'ChatPage',

    data: () => ({
      messages: [],
      postMessage: "",
      socket: null,
      state: "idle",
      roomInfo: {}
    }),
    methods: {
      openWebSocket(openedCallback) {
        console.log("open websocket", this.socket)
        this.socket = new WebSocket('ws://localhost:8081/broker/' + this.$route.params.id);
        if (openedCallback != null) {
          this.socket.onopen = openedCallback;
        }
        this.socket.onmessage = this.receiveMessage;
      },
      setIdle() {
        this.state = "idle";
      },
      sendMessage() {
        const readyState = this.socket.readyState;
        console.log(`readyState value is: ${readyState}`);

        switch (readyState) {
          case 0:
            console.log('Socket has been created. Please waiting for a moment.');
            this.state = "sendError";
            break;
          case 1:
            console.log('The connection is ready!!');
            console.log(this.postMessage);
            this.socket.send(this.postMessage);
            this.clearMessage();
            this.state = "idle";
            break;
          default:
            console.log('WebSocket is already in CLOSING state.');
            this.state = "sendError";
            break;
        }
      },
      retrySendMessage() {
        this.openWebSocket(this.reopend);
      },
      reopend() {
        this.setIdle();
        this.sendMessage();
      },
      clearMessage() {
        this.postMessage = '';
        this.state = "idle"
      },
      receiveMessage(event) {
        console.log('Message from server ', event.data);
        this.messages.push({
          value: event.data
        });
      },
      updateRoom(roomId) {
        axios.get('http://localhost:8888/rooms/' + roomId)
          .then(this.setRoom)
          .catch(function (error) {
            console.log(error);
          });
      },
      setRoom(response) {
        console.log(response);
        this.roomInfo = response.data;
      }
    },
    created: function() {
      this.openWebSocket(this.setIdle);
      this.updateRoom(this.$route.params.id);
    },
    beforeDestroy: function() {
      this.socket.close();
    }
  }
</script>

<style scoped>
li {
  list-style: none;
}
.chatList {
  height: auto;
  overflow: scroll;
  -ms-overflow-style: none;    /* IE, Edge 対応 */
  scrollbar-width: none;       /* Firefox 対応 */
}
.chatList::-webkit-scrollbar {  /* Chrome, Safari 対応 */
  display:none;
}
</style>
