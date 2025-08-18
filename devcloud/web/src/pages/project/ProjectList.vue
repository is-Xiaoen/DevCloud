<template>
  <div class="project-space-container">
    <!-- 顶部操作栏 -->
    <div class="action-bar">
      <a-space>
        <a-button type="primary" @click="showCreateModal = true">
          <template #icon><icon-plus /></template>
          新建项目空间
        </a-button>
      </a-space>

      <a-space :size="18">
        <a-input-search v-model="searchKey" placeholder="搜索项目名称/描述" @search="handleSearch" allow-clear
          style="width: 300px" />
        <a-select v-model="filterParams.status" placeholder="项目状态" style="width: 120px" allow-clear
          @change="handleSearch">
          <a-option value="active">活跃</a-option>
          <a-option value="archived">已归档</a-option>
          <a-option value="planning">规划中</a-option>
        </a-select>

        <a-button type="text" @click="resetFilters">
          <template #icon><icon-refresh /></template>
          重置
        </a-button>
      </a-space>
    </div>

    <!-- 项目卡片列表 -->
    <div class="project-list">
      <a-row :gutter="16">
        <a-col v-for="project in projectList" :key="project.id" :xs="24" :sm="12" :md="8" :lg="6">
          <ProjectCard :project="project" @edit="handleEdit" @delete="handleDelete" @archive="handleArchive" />
        </a-col>
      </a-row>

      <!-- 空状态 -->
      <a-empty v-if="!loading && projectList.length === 0" description="暂无项目数据" />
    </div>

    <!-- 分页 -->
    <div class="pagination-wrapper">
      <a-pagination v-model="pagination.current" :total="pagination.total" :page-size="pagination.pageSize" show-total
        show-jumper @change="handlePageChange" />
    </div>

    <!-- 新建/编辑模态框 -->
    <ProjectFormModal v-model:visible="showCreateModal" :project-data="currentProject" @submit="handleSubmit" />
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue';
import { Message } from '@arco-design/web-vue';
import ProjectCard from './components/ProjectCard.vue';
import ProjectFormModal from './components/ProjectFormModal.vue';
import { mockProjects } from './mockData';

// 数据状态
const loading = ref(false);
const projectList = ref([]);
const searchKey = ref('');
const filterParams = reactive({
  status: '',
});
const pagination = reactive({
  current: 1,
  pageSize: 12,
  total: 0
});

// 表单相关
const showCreateModal = ref(false);
const currentProject = ref(null);

// 初始化加载数据
const loadProjects = async () => {
  try {
    loading.value = true;
    // 模拟API调用
    await new Promise(resolve => setTimeout(resolve, 500));

    let filtered = [...mockProjects];

    if (searchKey.value) {
      const key = searchKey.value.toLowerCase();
      filtered = filtered.filter(p =>
        p.name.toLowerCase().includes(key) ||
        p.description.toLowerCase().includes(key)
      );
    }

    if (filterParams.status) {
      filtered = filtered.filter(p => p.status === filterParams.status);
    }

    // 分页处理
    const start = (pagination.current - 1) * pagination.pageSize;
    const end = start + pagination.pageSize;
    projectList.value = filtered.slice(start, end);
    pagination.total = filtered.length;
  } finally {
    loading.value = false;
  }
};

// 事件处理
const handleSearch = () => {
  pagination.current = 1;
  loadProjects();
};

const resetFilters = () => {
  searchKey.value = '';
  filterParams.status = '';
  handleSearch();
};

const handlePageChange = (page) => {
  pagination.current = page;
  loadProjects();
};

const handleEdit = (project) => {
  currentProject.value = project;
  showCreateModal.value = true;
};

const handleDelete = (id) => {
  Message.success(`已删除项目 ${id}`);
  loadProjects();
};

const handleArchive = (id) => {
  Message.success(`已归档项目 ${id}`);
  loadProjects();
};

const handleSubmit = () => {
  Message.success(currentProject.value ? '项目更新成功' : '项目创建成功');
  showCreateModal.value = false;
  loadProjects();
  currentProject.value = null;
};

// 初始化
onMounted(() => {
  loadProjects();
});
</script>

<style scoped>
.project-space-container {
  padding: 16px;
  display: flex;
  flex-direction: column;
  height: 100%;
}

.action-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
  flex-wrap: wrap;
  gap: 12px;
}

.project-list {
  flex: 1;
  margin-bottom: 16px;
  min-height: 300px;
}

.pagination-wrapper {
  display: flex;
  justify-content: center;
  margin-top: 16px;
}

@media (max-width: 768px) {
  .action-bar {
    flex-direction: column;
    align-items: flex-start;
  }

  .action-bar>* {
    width: 100%;
    margin-bottom: 8px;
  }
}
</style>
