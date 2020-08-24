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
  </v-container>
</template>

<script>
  export default {
    name: 'TopPage',
    data: () => ({
      roomList: []
    }),
    methods: {
      getRooms() {
        const axios = require('axios');

        axios.get('http://localhost:8888/rooms')
          .then(this.updateRoomList)
          .catch(function (error) {
            console.log(error);
          });
      },
      updateRoomList(response) {
        this.roomList = response.data;
      }
    },
    created: function() {
      this.getRooms();
    }
  }
</script>

<style scoped>
li {
  list-style: none;
}
</style>
