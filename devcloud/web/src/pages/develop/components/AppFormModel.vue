<template>
  <a-modal draggable v-model:visible="modelVisible" :ok-loading="createAppLoading" :title="editMode ? '编辑应用' : '新建应用'"
    @ok="handleOk" :on-before-ok="handleBeforeOk" @cancel="handleCancel" :ok-text="editMode ? '保存' : '创建'"
    unmount-on-close>
    <a-form ref="appFormRef" :model="form" layout="vertical">
      <a-form-item label="应用名称" field="name" :rules="[{ required: true, message: '请输入应用名称' }]"
        :validate-trigger="['change', 'blur']">
        <a-input v-model="form.name" placeholder="请输入应用名称" allow-clear />
      </a-form-item>

      <a-form-item label="应用描述" field="description">
        <a-textarea v-model="form.description" placeholder="请输入应用描述" :auto-size="{ minRows: 3, maxRows: 5 }"
          allow-clear />
      </a-form-item>

      <a-form-item label="应用类型" field="type" :rules="[{ required: true, message: '请选择应用类型' }]">
        <a-select v-model="form.type" placeholder="应用类型" allow-clear>
          <a-option value="SOURCE_CODE">源代码</a-option>
          <a-option value="CONTAINER_IMAGE">容器镜像</a-option>
          <a-option value="OTHER">其他</a-option>
        </a-select>
      </a-form-item>

      <div v-if="form.type === 'SOURCE_CODE'">
        <a-form-item label="代码仓库" field="code_repository.ssh_url" :rules="[{ required: true, message: '请输入代码仓库地址' }]">
          <a-input v-model="form.code_repository.ssh_url" placeholder="请输入代码仓库地址" allow-clear />
        </a-form-item>
      </div>

      <a-form-item label="团队" field="label.team" :rules="[{ required: true, message: '请输入团队名称' }]">
        <a-tree-select v-model="form.label.team" :field-names="{
          key: 'value',
          title: 'label'
        }" :data="teamOptions" :loading="fetchAppTeamsLoading" placeholder="请选择团队"></a-tree-select>
      </a-form-item>
    </a-form>
  </a-modal>
</template>

<script setup>
import { ref, watch, computed, onMounted } from 'vue';
import API from '@/api'
import { useTemplateRef } from 'vue'


const props = defineProps({
  visible: {
    type: Boolean,
    required: true
  },
  appData: {
    type: Object,
    default: null
  }
});

const emit = defineEmits(['update:visible', 'submit']);

const appFormRef = useTemplateRef('appFormRef')

// 1. 先定义所有函数和变量
const resetForm = () => {
  form.value = {
    name: '',
    description: '',
    code_repository: {
      ssh_url: ''
    },
    label: {
      team: ''
    }
  };
};

// 2. 再定义其他响应式数据
const modelVisible = computed({
  get: () => props.visible,
  set: (v) => emit('update:visible', v)
});

const editMode = computed(() => !!props.appData);

const form = ref({
  name: '',
  description: '',
  type: 'SOURCE_CODE',
  code_repository: {}
});

// 拉去团队的选型
const fetchAppTeamsLoading = ref(false);
const teamOptions = ref([])
const fetchAppTeams = async () => {
  try {
    fetchAppTeamsLoading.value = true
    const resp = await API.mcenter.LabelList({
      key: 'team',
    });
    console.log(resp.items)
    if (resp.items.length > 0) {
      teamOptions.value = resp.items[0].enum_options
    }

  } catch (error) {
    console.error('Error fetching app teams:', error);
    return [];
  } finally {
    fetchAppTeamsLoading.value = false
  }
};

onMounted(() => {
  if (props.appData) {
    form.value = { ...props.appData };
  } else {
    resetForm();
  }
  fetchAppTeams()
});

// 3. 最后定义watch（此时resetForm已经定义）
watch(() => props.visible, (newVal) => {
  if (props.appData) {
    if (newVal) {
      form.value = { ...props.appData };
    } else {
      resetForm();
    }
  }
}, { immediate: true, });


// 提交前校验, 阻止模态框关闭
const createAppLoading = ref(false)
const handleBeforeOk = async () => {
  // 检查
  const resp = await appFormRef.value.validate();
  if (resp) {
    return false;
  }

  // 创建
  createAppLoading.value = true
  try {
    createAppLoading.value = true

    if (editMode.value) {
      await API.mpaas.AppUpdate(props.appData.id, {
        ...form.value,
        name: form.value.name.trim()
      });
      return true
    } else {
      await API.mpaas.AppCreate({
        ...form.value,
        name: form.value.name.trim()
      });
    }
    return true
  } catch (error) {
    console.error('Error creating app:', error);
  } finally {
    createAppLoading.value = false
  }
  return false
}

const handleOk = async () => {
  emit('submit', {
    ...form.value,
    name: form.value.name.trim()
  });
  resetForm(); // 提交后重置表单
};

const handleCancel = () => {
  resetForm(); // 取消时也重置表单
};
</script>

<style scoped>
/* 可以添加自定义样式 */
:deep(.arco-form-item-label) {
  font-weight: 500;
}
</style>
