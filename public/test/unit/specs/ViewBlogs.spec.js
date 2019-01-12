import Vue from 'vue'
import ViewBlogs from '@/components/ViewBlogs'

describe('ViewBlogs.vue', () => {
  it('should render correct contents on view blog page', () => {
    const Constructor = Vue.extend(ViewBlogs)
    const vm = new Constructor().$mount()

    // Mock out Axios for testing data here
  })
})
