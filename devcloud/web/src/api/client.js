import axios from 'axios'
import { Message } from '@arco-design/web-vue'

// 封装一个axios的实例，http cient实例
// https://axios-http.com/zh/docs/instance
const client = axios.create({
  timeout: 5000,
})

// 拦截API的返回结果, 如果是异常 提取异常信息，并展示
// 如果正常 则直接返回 API的响应结果
// https://axios-http.com/zh/docs/interceptors
client.interceptors.response.use(
  (data) => {
    return data.data
  },
  (error) => {
    // 提取错误
    var msg = error.message

    // 如果有更加详细的信息
    if (error.response.data.message) {
      msg = error.response.data.message
    }

    Message.error(msg)
    return error
  },
)

export default client
