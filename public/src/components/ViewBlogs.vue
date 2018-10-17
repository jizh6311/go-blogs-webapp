<template>
  <div id="viewBlogs">
    <v-container>
      <v-layout row wrap justify-space-around>
        <v-flex xs12 md8>
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
                <v-btn flat class="blue--text">Delete</v-btn>
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
    imageSrc (blogID) {
      return 'http://localhost:8081/blog/' + blogID
    }
  },
  data () {
    return {
      blogs: null
    }
  },
  mounted () {
    axios.get('/blogs/jianzhang')
      .then((response) => {
        this.blogs = response.data
      })
      .catch(() => {
        console.error('Load blog data unsuccessfully!')
      })
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
