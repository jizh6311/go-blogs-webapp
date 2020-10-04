<template>
  <div id="viewBlogs">
    <input class="username-input" v-model="username" placeholder="input username" />
    <button type="submit" @click="submitUsername">Submit</button>
    <v-container>
      <v-layout row wrap justify-space-around>
        <v-flex xs12 md8>
          <!-- eslint-disable-next-line -->
          <div v-for="oneBlog in blogs">
            <v-card class="my-3" hover>
              <v-card-media
                class="white--text"
                height="170px"
                :src="imageSrc(oneBlog.id)"
              >
                <v-container fill-height fluid>
                  <v-layout>
                    <v-card-text>
                      {{ oneBlog.Username }}
                    </v-card-text>
                  </v-layout>
                </v-container>
              </v-card-media>
              <v-card-text>
                {{ oneBlog.Description }}
              </v-card-text>
              <v-card-text>
                {{ oneBlog.PostDate }}
              </v-card-text>
              <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn flat class="blue--text" v-on:click.native="updateBlog(oneBlog.id)">Update</v-btn>
                <v-btn flat class="blue--text" v-on:click.native="deleteBlog(oneBlog.id)">Delete</v-btn>
              </v-card-actions>
            </v-card>
          </div>
        </v-flex>
      </v-layout>
    </v-container>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'ViewBlogs',
  methods: {
    submitUsername () {
      axios.get(`/blogs/${this.username}`)
        .then((response) => {
          this.blogs = response.data
        })
        .catch(() => {
          console.error('Load blog data unsuccessfully!')
        })
    },

    imageSrc (blogID) {
      return 'http://localhost:8081/blog/' + blogID
    },

    updateBlog (blogID) {
      // TODO: leverage upcoming update blog API
    },

    deleteBlog (blogID) {
      axios.delete(`http://localhost:8081/blog/${blogID}`)
        .then((response) => {
          console.info(`Succeeded to delete blog by ID: ${blogID}`)
          location.reload()
        })
        .catch(() => {
          console.error(`Failed to delete blog by ID: ${blogID}`)
        })
    }
  },
  data () {
    return {
      blogs: null
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.username-input {
  margin: 15px 10px 0px 0px;
  padding: 0 5px 0 8px;
}
</style>
