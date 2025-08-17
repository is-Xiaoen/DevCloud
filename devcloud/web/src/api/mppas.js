import client from './client'

var MPAAS_API = {
  AppList: (params) => {
    return client.get('/api/devcloud/v1/applications', { params })
  },
  AppCreate: (data) => {
    return client.post('/api/devcloud/v1/applications', data)
  },
  AppUpdate: (id, data) => {
    return client.put(`/api/devcloud/v1/applications/${id}`, data)
  },
  AppDelete: (id) => {
    return client.delete(`/api/devcloud/v1/applications/${id}`)
  },
}

export default MPAAS_API
