<template>
  <div id="app">
    <div id="appHeader">
      <h1>Welcome to your blogs!</h1>
    </div>
    <v-tabs left color='grey lighten-4'>
      <v-tab :to="{path:'/'}">View Blogs</v-tab>
      <v-tab :to="{path:'/postBlogs'}">Post Blogs</v-tab>
    </v-tabs>
    <router-view />
  </div>
</template>

<script>
export default {
  name: 'App'
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
