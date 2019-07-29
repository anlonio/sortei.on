<template>
  <v-app id="inspire" dark>
    <v-toolbar app fixed clipped-left>
      <v-toolbar-title>SorTei.oN</v-toolbar-title>
    </v-toolbar>
    <v-content>
      <v-container fluid class="px-0">
        <v-layout justify-center align-center class="px-0">
          <draft></draft>
        </v-layout>
      </v-container>
    </v-content>
    <v-footer app fixed>
      <span style="margin-left:1em">&copy; Antonio Machado - 2019</span>
    </v-footer>
  </v-app>
</template>

<script>
import Draft from "./components/Draft.vue";

export default {
  components: {
    Draft
  },
  props: {
    source: String
  }
};
</script>