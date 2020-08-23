<template>
  <v-container>
    <h1>Top Page</h1>
    <li v-for="room in roomList" :key="room.uuid">
      <p>{{ room.title }}</p>
    </li>
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
