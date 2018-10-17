<template>
  <div>
    <h1>Post your blogs here!</h1>
    <form
      class="postForm"
      @submit="checkForm"
    >
      <div class="row">
        <label for="name">Your Name:</label>
        <input type="text" v-model="username" name="name" placeholder="Show me your name" />
      </div>
      <div class="row">
        <label for="description">Description:</label>
        <textarea id="description" v-model="description" type="text" name="description" placeholder="Add some notes"></textarea>
      </div>
      <div class="row">
        <input type="file" id="image" ref="image" v-on:change="handleImageUpload()"/>
      </div>
      <div>
        <button id="postButton" v-on:click="postBlog" type="button">Post Blogs</button>
      </div>
    </form>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'PostBlogs',
  methods: {
    handleImageUpload () {
      this.file = this.$refs.image.files[0]
    },

    postBlog () {
      let formData = new FormData()
      formData.append('file', this.file)
      formData.append('Username', this.username)
      formData.append('Description', this.description)
      axios.post('/blogs',
        formData,
        {
          headers: {
            'Content-Type': 'application/json, multipart/form-data'
          }
        }
      ).then(function () {
        console.log('Image Loaded successfully!')
      }).catch(function () {
        console.error('Image Loaded unsuccessfully!')
      })
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  margin-top: 5%;
  font-size: 20px;
  font-style: oblique;
  font-family: "Times New Roman", Times, serif;
}
.postForm {
  margin: 5%;
  font-size: 16px;
  font-style: bold;
  font-family: "Times New Roman", Times, serif;
  border-radius: 20px;
  background-color: #66CDAA;
  padding: 20px;
  align-items: center;
}
input[type=text], textarea{
  width: 50%;
  padding: 12px 20px;
  margin: 2%;
  box-sizing: border-box;
  background-color: white;
  background-position: 10px 10px;
  background-repeat: no-repeat;
  padding-left: 40px;
}
#description {
  height: 150px;
}
#postButton {
  background-color: #ccff33;
  font-size: 15px;
  border-radius: 30%;
  padding: 1%;
  margin-top: 5%;
}
label {
  display : block;
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
