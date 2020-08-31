<template>
  <v-app>
    <v-app-bar
      app
      color="primary"
      dark
    >
      <div class="d-flex align-center">
        <v-toolbar-title>
          <router-link to="/" class="toolbar-title">tmpchat</router-link>
        </v-toolbar-title>
      </div>

      <v-spacer></v-spacer>

      <v-btn
        href="https://github.com/tmpchat/tmpchat"
        target="_blank"
        text
        fab
      >
        <v-icon large>mdi-github</v-icon>
      </v-btn>
    </v-app-bar>

    <v-main>
      <router-view/>
    </v-main>
  </v-app>
</template>

<script>
export default {
  name: 'App',

  components: {},

  data: () => ({
    //
  }),
};
</script>

<style scoped>
.toolbar-title {
  color: white;
  text-decoration: none;
}
</style>
