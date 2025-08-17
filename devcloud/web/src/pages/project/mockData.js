// mockData.js

/**
 * 生成随机日期（最近365天内）
 */
const generateRandomDate = () => {
  const now = new Date()
  const past = new Date(now)
  past.setDate(now.getDate() - 365)
  return new Date(past.getTime() + Math.random() * (now.getTime() - past.getTime()))
}

/**
 * 生成随机项目描述
 */
const generateDescription = (index) => {
  const descriptors = [
    '企业级应用开发项目',
    '微服务架构重构项目',
    '前端技术升级专项',
    'DevOps平台建设项目',
    'AI能力集成项目',
    '大数据分析平台',
    '移动端应用重设计',
    '云原生迁移项目',
  ]
  const types = ['核心业务系统', '技术中台', '数据平台', '用户增长项目', '基础设施升级']
  return `${types[index % types.length]} - ${descriptors[index % descriptors.length]}`
}

/**
 * 模拟用户数据
 */
export const mockUsers = [
  { id: 'user-1', name: '张管理员', role: 'admin' },
  { id: 'user-2', name: '李开发', role: 'dev' },
  { id: 'user-3', name: '王测试', role: 'qa' },
  { id: 'user-4', name: '赵产品', role: 'product' },
  { id: 'user-5', name: '刘架构', role: 'arch' },
]

/**
 * 模拟项目数据
 */
export const mockProjects = Array.from({ length: 36 }, (_, i) => {
  const statuses = ['active', 'archived', 'planning']
  const status = statuses[i % 3]

  return {
    id: `project-${i + 1}`,
    name: `研发项目 ${String(i + 1).padStart(2, '0')}`,
    description: generateDescription(i),
    status,
    owner: mockUsers[i % mockUsers.length].id,
    members: mockUsers.slice(0, (i % 4) + 1),
    createTime: generateRandomDate().getTime(),
    updateTime: generateRandomDate().getTime(),
    stats: {
      tasks: Math.floor(Math.random() * 50) + 10,
      completed: Math.floor(Math.random() * 45),
      progress: Math.floor(Math.random() * 100),
    },
  }
})

/**
 * 项目状态映射
 */
export const projectStatusMap = {
  active: { text: '活跃', color: '#00B42A' },
  archived: { text: '已归档', color: '#86909C' },
  planning: { text: '规划中', color: '#FF7D00' },
}

/**
 * 获取项目状态信息
 */
export const getProjectStatusInfo = (status) => {
  return projectStatusMap[status] || { text: '未知', color: '#86909C' }
}

/**
 * 模拟API响应格式
 */
export const mockApiResponse = (data, options = {}) => {
  return {
    code: 200,
    data,
    message: 'success',
    ...options,
  }
}

/**
 * 模拟获取项目列表API
 */
export const fetchProjects = (params = {}) => {
  const { page = 1, pageSize = 10, search = '', status } = params

  let filtered = [...mockProjects]

  // 模拟筛选
  if (search) {
    const keyword = search.toLowerCase()
    filtered = filtered.filter(
      (p) =>
        p.name.toLowerCase().includes(keyword) || p.description.toLowerCase().includes(keyword),
    )
  }

  if (status) {
    filtered = filtered.filter((p) => p.status === status)
  }

  // 模拟分页
  const start = (page - 1) * pageSize
  const end = start + pageSize
  const pageData = filtered.slice(start, end)

  return new Promise((resolve) => {
    setTimeout(() => {
      resolve(
        mockApiResponse({
          list: pageData,
          total: filtered.length,
          page,
          pageSize,
        }),
      )
    }, 500) // 模拟网络延迟
  })
}

/**
 * 模拟获取项目详情API
 */
export const fetchProjectDetail = (id) => {
  const project = mockProjects.find((p) => p.id === id)
  return new Promise((resolve) => {
    setTimeout(() => {
      if (project) {
        resolve(mockApiResponse(project))
      } else {
        resolve(mockApiResponse(null, { code: 404, message: '项目不存在' }))
      }
    }, 300)
  })
}

/**
 * 模拟创建项目API
 */
export const createProject = (data) => {
  const newProject = {
    id: `project-${mockProjects.length + 1}`,
    ...data,
    createTime: Date.now(),
    updateTime: Date.now(),
    members: [mockUsers[0]],
    stats: {
      tasks: 0,
      completed: 0,
      progress: 0,
    },
  }

  mockProjects.unshift(newProject)

  return new Promise((resolve) => {
    setTimeout(() => {
      resolve(mockApiResponse(newProject))
    }, 400)
  })
}

/**
 * 模拟更新项目API
 */
export const updateProject = (id, data) => {
  const index = mockProjects.findIndex((p) => p.id === id)

  return new Promise((resolve) => {
    setTimeout(() => {
      if (index >= 0) {
        const updated = { ...mockProjects[index], ...data, updateTime: Date.now() }
        mockProjects.splice(index, 1, updated)
        resolve(mockApiResponse(updated))
      } else {
        resolve(mockApiResponse(null, { code: 404, message: '项目不存在' }))
      }
    }, 400)
  })
}
