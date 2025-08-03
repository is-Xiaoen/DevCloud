import client from './client'

var API = {
  Login: (data) => {
    return client.post('/api/devcloud/v1/token', data)
  },
}

export default API
