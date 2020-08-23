<template>
  <v-container>
    <h1>Top Page</h1>
  </v-container>
</template>

<script>
  export default {
    name: 'TopPage',
    data: () => ({}),
    methods: {
      getRooms() {
        const axios = require('axios');

        axios.get('http://localhost:8888/rooms')
          .then(function (response) {
            console.log(response);
          })
          .catch(function (error) {
            console.log(error);
          });
      }
    },
    created: function() {
      this.getRooms();
    }
  }
</script>
