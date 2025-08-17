<template>
  <a-form layout="inline" :model="query">
    <div>
    </div>
    <div>
      <a-form-item label="Hover" field="hover">
        <a-switch v-model="query.hover" />
      </a-form-item>
    </div>
  </a-form>
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
