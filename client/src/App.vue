<template>
  <v-app>
    <Navbar />
    <v-spacer></v-spacer>
    <router-view/>
  </v-app>
</template>

<script>
import Navbar from './components/Navbar';

export default {
  name: "App",
  components: {
    Navbar
  },
};
</script>
