<template>
  <a-modal v-model:visible="modelVisible" :title="editMode ? '编辑项目' : '新建项目'" @ok="handleOk" @cancel="handleCancel"
    :ok-text="editMode ? '保存' : '创建'" unmount-on-close>
    <a-form :model="form" layout="vertical">
      <a-form-item label="项目名称" field="name" :rules="[{ required: true, message: '请输入项目名称' }]"
        :validate-trigger="['change', 'blur']">
        <a-input v-model="form.name" placeholder="请输入项目名称" allow-clear />
      </a-form-item>

      <a-form-item label="项目描述" field="description">
        <a-textarea v-model="form.description" placeholder="请输入项目描述" :auto-size="{ minRows: 3, maxRows: 5 }"
          allow-clear />
      </a-form-item>

      <a-form-item label="项目状态" field="status" :rules="[{ required: true, message: '请选择项目状态' }]">
        <a-select v-model="form.status" placeholder="请选择状态" allow-clear>
          <a-option value="active">活跃</a-option>
          <a-option value="archived">已归档</a-option>
          <a-option value="planning">规划中</a-option>
        </a-select>
      </a-form-item>
    </a-form>
  </a-modal>
</template>

<script setup>
import { ref, watch, computed } from 'vue';
import { Message } from '@arco-design/web-vue';

const props = defineProps({
  visible: {
    type: Boolean,
    required: true
  },
  projectData: {
    type: Object,
    default: null
  }
});

const emit = defineEmits(['update:visible', 'submit']);

// 1. 先定义所有函数和变量
const resetForm = () => {
  form.value = {
    name: '',
    description: '',
    status: 'active'
  };
};

// 2. 再定义其他响应式数据
const modelVisible = computed({
  get: () => props.visible,
  set: (value) => emit('update:visible', value)
});

const editMode = computed(() => !!props.projectData);

const form = ref({
  name: '',
  description: '',
  status: 'active'
});

// 3. 最后定义watch（此时resetForm已经定义）
watch(() => props.projectData, (newVal) => {
  if (newVal) {
    form.value = { ...newVal };
  } else {
    resetForm();
  }
}, { immediate: true });

const handleOk = () => {
  if (!form.value.name.trim()) {
    Message.error('项目名称不能为空');
    return;
  }

  emit('submit', {
    ...form.value,
    name: form.value.name.trim()
  });
  modelVisible.value = false;
  resetForm(); // 提交后重置表单
};

const handleCancel = () => {
  modelVisible.value = false;
  resetForm(); // 取消时也重置表单
};
</script>

<style scoped>
/* 可以添加自定义样式 */
:deep(.arco-form-item-label) {
  font-weight: 500;
}
</style>
