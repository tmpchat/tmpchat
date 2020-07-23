<template>
  <v-container>
    <h1>Chat Page, Room Title</h1>
    <v-row>
      <v-col cols="12">
        <li v-for="message in messages" :key="message.title">
          <v-card
            outlined
          >
            <v-list-item three-line style="text-align: left;">
              <v-list-item-content>
                <v-list-item-title >{{ message.title }}</v-list-item-title>
                <v-list-item-subtitle>{{ message.value }}</v-list-item-subtitle>
              </v-list-item-content>
            </v-list-item>
          </v-card>
        </li>
      </v-col>
    </v-row>
    <v-row>
      
    <template v-if="state === 'idle'">
      <v-text-field label="Message" v-model="postMessage"/>
      <v-btn v-on:click="sendMessage">
        <v-icon>mdi-send</v-icon>
      </v-btn>
    </template>
    <template v-else-if="state === 'sendError'">
      <v-text-field label="Message" v-model="postMessage" error/>
      <v-btn v-on:click="sendMessage">
        Retry
      </v-btn>
      <v-btn v-on:click="sendMessage" color="error">
        Cancel
      </v-btn>
    </template>
    </v-row>
  </v-container>
</template>

<script>
  export default {
    name: 'ChatPage',

    data: () => ({
      messages: [],
      postMessage: "",
      socket: null,
      state: "idle"
    }),
    methods: {
      sendMessage() {
        const readyState = this.socket.readyState;
        console.log(`readyState value is: ${readyState}`);

        switch ( readyState ) {
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
          case 2:
            console.log('WebSocket is already in CLOSING state.');
            this.state = "sendError";
            break;
          case 3:
            console.log('WebSocket is already in CLOSED state.');
            this.state = "sendError";
            break;
        }
      },
      clearMessage() {
        this.postMessage = '';
      },
      receiveMessage(event) {
        console.log('Message from server ', event.data);
        this.messages.push({
          title: "Example",
          value: event.data
        });
      }
    },
    created: function() {
      this.socket = new WebSocket('ws://localhost:8081/broker');
      this.socket.onmessage = this.receiveMessage;
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
</style>
