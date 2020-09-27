<template>
  <v-container>
    <h1>Room list</h1>
    <v-simple-table>
      <template v-slot:default>
        <thead>
          <tr>
            <th class="text-left">Title</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="room in roomList" :key="room.uuid">
            <td><a v-bind:href="'/chat/' + room.uuid">{{ room.title }}</a></td>
          </tr>
        </tbody>
      </template>
    </v-simple-table>

    <template>
      <v-text-field label="Title" v-model="title"/>
      <v-btn v-on:click="createRoom">
        <v-icon>mdi-plus</v-icon>
      </v-btn>
    </template>

    <v-footer
      app
      color="primary"
      dark
    >
      <v-col
        class="text-center"
        cols="12"
      >
        {{ currentYear }} - <strong>tmpchat</strong>
      </v-col>
    </v-footer>
  </v-container>
</template>

<script>
  const axios = require('axios')
  export default {
    name: 'TopPage',
    data: () => ({
      roomList: [],
      title: '',
      currentYear: ''
    }),
    methods: {
      updateRooms() {
        axios.get('http://localhost:8888/rooms')
          .then(this.setRoomList)
          .catch(function (error) {
            console.log(error);
          });
      },
      setRoomList(response) {
        this.roomList = response.data;
      },
      createRoom() {
        axios.post('http://localhost:8888/rooms', {
            title: this.title
          })
          .then(function (response) {
            console.log(response);
          })
          .catch(function (error) {
            console.log(error);
          });
      },
      getCurrentYear() {
        this.currentYear = new Date().getFullYear();
      }
    },
    created: function() {
      this.updateRooms();
      this.getCurrentYear();
    }
  }
</script>

<style scoped>
li {
  list-style: none;
}
</style>
