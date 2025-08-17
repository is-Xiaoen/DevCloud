import client from './client'

var MCENTER_API = {
  Login: (data) => {
    return client.post('/api/devcloud/v1/token', data)
  },
}

export default MCENTER_API
