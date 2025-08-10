import token from '@/storage/token'

var witheList = ['ProductPage', 'LoginPage']

export var beforeEach = async (to) => {
  // 白名单
  for (var name of witheList) {
    if (name === to.name) {
      return true
    }
  }

  // 检查用户是否已登录
  if (token.value.access_token) {
    return true
  }

  // 将用户重定向到登录页面
  return { name: 'LoginPage' }
}
