<template>
  <a-card class="project-card" :style="{ borderTop: `4px solid ${statusColor}` }" hoverable>
    <a-card-meta :title="project.name">
      <template #description>
        <div class="project-meta">
          <div class="project-desc">{{ project.description }}</div>

          <div class="project-footer">
            <a-space>
              <a-avatar-group :max-count="3">
                <a-avatar v-for="member in project.members" :key="member.id" :size="24">
                  {{ member.name.charAt(0) }}
                </a-avatar>
              </a-avatar-group>
              <span>创建于 {{ formatDate(project.createTime) }}</span>
            </a-space>
          </div>
        </div>
      </template>
    </a-card-meta>

    <template #actions>
      <a-tooltip content="编辑">
        <a-button type="text" @click.stop="$emit('edit', project)">
          <icon-edit />
        </a-button>
      </a-tooltip>

      <a-tooltip content="删除">
        <a-popconfirm content="确认删除该项目？" @ok="$emit('delete', project.id)">
          <a-button type="text">
            <icon-delete />
          </a-button>
        </a-popconfirm>
      </a-tooltip>

      <a-dropdown @click.stop>
        <a-button type="text">
          <icon-more />
        </a-button>
        <template #content>
          <a-doption @click="$emit('archive', project.id)">归档</a-doption>
        </template>
      </a-dropdown>
    </template>
  </a-card>
</template>

<script setup>
import { computed } from 'vue';
import {
  IconEdit,
  IconDelete,
  IconMore
} from '@arco-design/web-vue/es/icon';

const props = defineProps({
  project: {
    type: Object,
    required: true,
    default: () => ({
      id: '',
      name: '',
      description: '',
      status: 'active',
      members: [],
      createTime: Date.now()
    })
  }
});

defineEmits(['edit', 'delete', 'archive']);

const statusColor = computed(() => {
  const map = {
    active: '#00B42A',
    archived: '#86909C',
    planning: '#FF7D00'
  };
  return map[props.project.status] || '#86909C';
});

const formatDate = (timestamp) => {
  return new Date(timestamp).toLocaleDateString('zh-CN');
};
</script>

<style scoped>
.project-card {
  margin-bottom: 16px;
  border-radius: 4px;
  transition: all 0.2s;
  height: 100%;
}

.project-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.project-meta {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.project-desc {
  color: var(--color-text-2);
  font-size: 12px;
  display: -webkit-box;
  -webkit-box-orient: vertical;
  overflow: hidden;
  min-height: 40px;
}

.project-footer {
  margin-top: 12px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

:deep(.arco-card-actions) {
  border-top: none;
  padding: 0 8px 8px;
}

:deep(.arco-card-meta-description) {
  margin-top: 8px;
}
</style>
