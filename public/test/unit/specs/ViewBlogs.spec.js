import Vue from 'vue'
import ViewBlogs from '@/components/ViewBlogs'
import axios from 'axios'
import moxios from 'moxios'
import sinon from 'sinon'

describe('ViewBlogs.vue', () => {
  beforeEach(function () {
    // import and pass your custom axios instance to this method
    moxios.install()
  })

  afterEach(function () {
    // import and pass your custom axios instance to this method
    moxios.uninstall()
  })

  it('should render correct contents on view blog page', () => {
    var mockData = [
      {
        'id': '5be895136cface1b0010a6d7',
        'Username': 'jianzhang',
        'Description': 'This is the first sample',
        'PostDate': '2018-11-11T15:46:11.171-05:00'
      },
      {
        'id': '5be895136cface1b0010a6d8',
        'Username': 'jianzhang',
        'Description': 'This is the second sample',
        'PostDate': '2018-11-11T15:46:11.171-05:00'
      }
    ]

    moxios.stubRequest('/blogs/jianzhang', {
      status: 200,
      response: mockData
    })

    let onFulfilled = sinon.spy()
    axios.get('/blogs/jianzhang').then(onFulfilled)

    moxios.wait(function () {
      equal(onFulfilled.getCall(0).args[0].data, mockData)
      done()
    })

    const Constructor = Vue.extend(ViewBlogs)
    const vm = new Constructor().$mount()
  })
})
