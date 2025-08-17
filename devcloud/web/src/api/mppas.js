import client from './client'

var MPAAS_API = {
  AppList: (params) => {
    return client.get('/api/devcloud/v1/applications', { params })
  },
}

export default MPAAS_API
