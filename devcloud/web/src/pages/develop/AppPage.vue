<template>
  <div>
    <a-card>
      <div class="action-bar">
        <a-space>
          <a-button type="outline" @click="addApp()">
            <template #icon><icon-plus /></template>
            新建应用
          </a-button>
        </a-space>
        <a-space>
          <a-select v-model="query.ready" :style="{ width: '220px' }" @change="fetchAppList" @clear="query.ready = null"
            placeholder="选择状态" allow-clear>
            <a-option :value="true">就绪</a-option>
            <a-option :value="false">未就绪</a-option>
          </a-select>
          <a-input-search v-model="query.keywords" placeholder="搜索项目名称/描述" allow-clear @press-enter="fetchAppList"
            @search="fetchAppList" style="width: 300px" />
        </a-space>
      </div>
      <a-table :data="data.items" :loading="fetchAppLoading">
        <template #columns>
          <a-table-column title="名称" data-index="name"></a-table-column>
          <a-table-column title="描述" data-index="description"></a-table-column>
          <a-table-column title="团队" data-index="label.team"></a-table-column>
          <a-table-column title="创建时间" data-index="create_at"></a-table-column>
          <a-table-column title="类型" data-index="type"></a-table-column>
          <a-table-column title="仓库地址" data-index="code_repository.ssh_url"></a-table-column>
          <a-table-column title="是否就绪" data-index="ready">
            <template #cell="{ record }">
              <a-switch type="round" disabled v-model="record.ready">
                <template #checked>
                  就绪
                </template>
                <template #unchecked>
                  未就绪
                </template>
              </a-switch>
            </template>
          </a-table-column>
          <a-table-column title="操作">
            <template #cell="{ record }">
              <a-space>
                <a-button @click="editApp(record)" type="primary">编辑</a-button>
                <a-popconfirm :content="`确定要删除${record.name}吗？`" :on-before-ok="() => doDeleteApp(record)"
                  :ok-loading="deleteAppLoading" @ok="deleteApp(record)" type="warning">
                  <a-button status="danger">删除</a-button>
                </a-popconfirm>
              </a-space>
            </template>
          </a-table-column>
        </template>
      </a-table>
      <AppFormModel v-model:visible="formVisible" @update:visible="(v) => console.log(v)" :appData="currentApp"
        @submit="fetchAppList" />
    </a-card>
  </div>
</template>

<script setup>
import { onMounted, ref } from 'vue';
import API from '@/api'
import AppFormModel from './components/AppFormModel.vue';

const query = ref({
  page_size: 20,
  page_number: 1,
  keywords: '',
});

// 通过API获取应用列表
const data = ref({
  items: [],
  total: 0,
});
const fetchAppLoading = ref(false);
const fetchAppList = async () => {
  console.log(query.value.ready)
  if (query.value.ready === null || query.value.ready === undefined || query.value.ready === '') {
    delete query.value.ready;
  }
  try {
    fetchAppLoading.value = true;
    const resp = await API.mpaas.AppList(query.value);
    data.value = resp;
    console.log('Fetched app list:', data.value);
  } catch (error) {
    console.error('Error fetching app list:', error);
  } finally {
    fetchAppLoading.value = false;
  }
};

onMounted(() => {
  fetchAppList();
});


// 编辑新增
const formVisible = ref(false);
const currentApp = ref(null);

const addApp = () => {
  formVisible.value = true;
  currentApp.value = null
}

const editApp = (app) => {
  currentApp.value = app;
  formVisible.value = true;
};

const deleteAppLoading = ref(false);
const doDeleteApp = async (app) => {
  try {
    deleteAppLoading.value = true;
    await API.mpaas.AppDelete(app.id);

  } catch (error) {
    console.error('Error deleting app:', error);
  } finally {
    deleteAppLoading.value = false;
  }
}
const deleteApp = () => {
  fetchAppList();
};
</script>
