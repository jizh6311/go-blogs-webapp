<template>
  <div id="viewBlogs">
    <v-container>
      <v-layout row wrap justify-space-around>
        <v-flex xs12 md8>
          <div v-for="mockBlog in mockBlogs">
            <v-card class="my-3" hover>
              <v-card-media
                class="white--text"
                height="170px"
                :src="mockBlog.imgUrl"
              >
                <v-container fill-height fluid>
                  <v-layout>
                    <v-card-text>
                      {{ mockBlog.Username }}
                    </v-card-text>
                  </v-layout>
                </v-container>
              </v-card-media>
              <v-card-text>
                {{ mockBlog.Description }}
              </v-card-text>
              <v-card-text>
                {{ mockBlog.PostDate }}
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
export default {
  name: 'ViewBlogs',
  data () {
    return {
      mockBlogs: [
        {
          'id': '5ba80c897a884e1df38ecfe1',
          'Username': 'Mock_User_1',
          'Description': 'Mock_Description_1: Before I sign off for the night, I am curious, ' +
            'does anyone have any upsets they are picking from this weekend games?',
          'PostDate': '2018-09-23T17:58:33.347-04:00'
        },
        {
          'id': '5ba80c897a884e1df38ecfe1',
          'Username': 'Mock_User_2',
          'Description': 'Mock_Description_2: I really wanted to go with the Broncos over the Patriots,' +
            'but punked out at the end and just picked Miami to cover. So the Browns over the Raiders, ' +
            'on the road, is my upset of the week. Anyone feeling bold?',
          'PostDate': '2018-09-24T17:58:33.347-04:00'
        },
        {
          'id': '5ba80c897a884e1df38ecfe1',
          'Username': 'Mock_User_3',
          'Description': 'Mock_Description_3: What do I have to say here',
          'PostDate': '2018-09-24T17:58:33.347-04:00'
        }
      ]
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
