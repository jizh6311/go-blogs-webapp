<template>
  <div id="postBlogs">
    <h1>{{ msg }}</h1>
  </div>
</template>

<script>
export default {
  name: 'PostBlogs',
  data () {
    return {
      msg: 'Your blogs can be posted here'
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  margin-top: 10%;
  font-size: 20px;
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
