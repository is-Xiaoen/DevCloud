<template>
  <div>
    <div class="action-bar">
      <a-space>
        <a-button type="primary">
          <template #icon><icon-plus /></template>
          新建项目空间
        </a-button>
      </a-space>
      <a-space>
        <a-select v-model="query.ready" :style="{ width: '220px', backgroundColor: '#fff' }" @change="fetchAppList"
          placeholder="选择状态" allow-clear>
          <a-option :value="true">就绪</a-option>
          <a-option :value="false">未就绪</a-option>
        </a-select>
        <a-input-search v-model="searchKey" placeholder="搜索项目名称/描述" allow-clear
          style="width: 300px;background-color: #fff;" />
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
              <a-button @click="editApp(record)">编辑</a-button>
              <a-button @click="deleteApp(record)" type="primary">删除</a-button>
            </a-space>
          </template>
        </a-table-column>
      </template>
    </a-table>
  </div>
</template>

<script setup>
import { onMounted, reactive, ref } from 'vue';
import API from '@/api'

const query = reactive({
  page_size: 20,
  page_number: 1,

});

// 通过API获取应用列表
const data = ref({
  items: [],
  total: 0,
});
const fetchAppLoading = ref(false);
const fetchAppList = async () => {
  try {
    fetchAppLoading.value = true;
    const resp = await API.mpaas.AppList(query);
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
</script>
