import client from './client'

var MCENTER_API = {
  Login: (data) => {
    return client.post('/api/devcloud/v1/token', data)
  },
  LabelList: (params) => {
    return client.get('/api/devcloud/v1/labels', { params })
  },
}

export default MCENTER_API
