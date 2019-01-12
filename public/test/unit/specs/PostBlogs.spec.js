import Vue from 'vue'
import PostBlogs from '@/components/PostBlogs'

describe('PostBlogs.vue', () => {
  it('should render correct contents on post blog page', () => {
    const Constructor = Vue.extend(PostBlogs)
    const vm = new Constructor().$mount()
    expect(vm.$el.querySelector('div h1').textContent)
      .toEqual('Post your blogs here!')
    expect(vm.$el.querySelector('label[for=name]').textContent)
      .toEqual('Your Name:')
    expect(vm.$el.querySelector('div #postButton').textContent)
      .toEqual('Post Blogs')
  })
})
