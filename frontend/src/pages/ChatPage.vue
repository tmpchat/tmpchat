<template>
  <v-container>
    <h1>Chat Page, Room Title</h1>
    <v-row>
      <v-col cols="12">
        <!-- TODO: change to list -->
        <v-card
          outlined
        >
          <v-list-item three-line style="text-align: left;">
            <v-list-item-content>
              <v-list-item-title >1. No name</v-list-item-title>
              <v-list-item-subtitle>this message</v-list-item-subtitle>
            </v-list-item-content>
          </v-list-item>
        </v-card>
        <v-card
          outlined
        >
          <v-list-item three-line style="text-align: left;">
            <v-list-item-content>
              <v-list-item-title >2. No name</v-list-item-title>
              <v-list-item-subtitle>ahaha</v-list-item-subtitle>
            </v-list-item-content>
          </v-list-item>
        </v-card>
      </v-col>
    </v-row>
    <v-row>
      <v-text-field label="Message" />
      <v-btn>
        <v-icon>mdi-pencil</v-icon>
      </v-btn>
    </v-row>
  </v-container>
</template>

<script>
  export default {
    name: 'ChatPage',

    data: () => ({}),
  }
</script>
