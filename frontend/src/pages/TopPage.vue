<template>
  <v-container>
    <h1>Top Page</h1>
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
  </v-container>
</template>

<script>
  export default {
    name: 'TopPage'
  }
</script>
