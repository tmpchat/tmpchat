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
      <v-text-field label="Message" v-model="postMessage"/>
      <v-btn v-on:click="sendMessage">
        <v-icon>mdi-pencil</v-icon>
      </v-btn>
    </v-row>
  </v-container>
</template>

<script>
  export default {
    name: 'ChatPage',

    data: () => ({
      messages: [],
      postMessage: "",
      socket: null
    }),
    methods: {
      sendMessage() {
        console.log(this.postMessage);
        this.socket.send(this.postMessage);
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
    }
  }
</script>

<style scoped>
li {
  list-style: none;
}
</style>
