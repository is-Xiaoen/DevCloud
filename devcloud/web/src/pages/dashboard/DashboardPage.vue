<template>
  <div class="dashboard-container">
    <!-- 顶部数据概览 -->
    <section class="metrics-section">
      <a-row :gutter="20">
        <a-col :xs="24" :sm="12" :md="6" v-for="metric in metrics" :key="metric.title">
          <a-card class="metric-card" hoverable>
            <div class="metric-content">
              <div class="metric-icon" :style="{ backgroundColor: metric.color }">
                <component :is="metric.icon" size="20" />
              </div>
              <div class="metric-info">
                <a-typography-text class="metric-title">{{ metric.title }}</a-typography-text>
                <a-typography-title :heading="4" class="metric-value">{{ metric.value }}</a-typography-title>
                <a-typography-text class="metric-trend" :type="metric.trend.type">
                  {{ metric.trend.value }}
                  <icon-caret-up v-if="metric.trend.direction === 'up'" />
                  <icon-caret-down v-else />
                </a-typography-text>
              </div>
            </div>
          </a-card>
        </a-col>
      </a-row>
    </section>

    <!-- 构建状态和活动流水线 -->
    <section class="main-section">
      <a-row :gutter="20">
        <!-- 构建状态图表 -->
        <a-col :xs="24" :md="16">
          <BuildStatusChar />
        </a-col>

        <!-- 活动流水线 -->
        <a-col :xs="24" :md="8">
          <a-card title="活动流水线" :bordered="false" class="pipeline-card">
            <template #extra>
              <a-link>查看全部</a-link>
            </template>
            <a-list :bordered="false" :split="false">
              <a-list-item v-for="pipeline in activePipelines" :key="pipeline.id">
                <a-list-item-meta>
                  <template #avatar>
                    <a-badge :status="pipeline.status" />
                  </template>
                  <template #title>
                    <a-link>{{ pipeline.name }}</a-link>
                  </template>
                  <template #description>
                    <a-space>
                      <a-typography-text type="secondary">#{{ pipeline.id }}</a-typography-text>
                      <a-typography-text type="secondary">{{ pipeline.duration }}</a-typography-text>
                    </a-space>
                  </template>
                </a-list-item-meta>
                <a-tag :color="pipeline.tagColor">{{ pipeline.stage }}</a-tag>
              </a-list-item>
            </a-list>
          </a-card>
        </a-col>
      </a-row>
    </section>

    <!-- 部署统计和最近活动 -->
    <section class="secondary-section">
      <a-row :gutter="20">
        <!-- 部署统计 -->
        <a-col :xs="24" :md="12">
          <DeployChart />
        </a-col>

        <!-- 最近活动 -->
        <a-col :xs="24" :md="12">
          <a-card title="最近活动" :bordered="false" class="activity-card">
            <a-timeline>
              <a-timeline-item v-for="activity in recentActivities" :key="activity.id" :color="activity.color">
                <a-space direction="vertical" size="2">
                  <div class="activity-content">
                    <a-typography-text strong>{{ activity.user }}</a-typography-text>
                    <a-typography-text type="secondary">{{ activity.action }}</a-typography-text>
                    <a-link>{{ activity.target }}</a-link>
                  </div>
                  <a-typography-text type="secondary" class="activity-time">{{ activity.time }}</a-typography-text>
                </a-space>
              </a-timeline-item>
            </a-timeline>
          </a-card>
        </a-col>
      </a-row>
    </section>
  </div>
</template>

<script setup>
import { ref, shallowRef } from 'vue';
import {
  IconCheckCircle,
  IconClockCircle,
  IconCloseCircle,
  IconCode,
  IconCaretUp,
  IconCaretDown
} from '@arco-design/web-vue/es/icon';

import DeployChart from './components/DeployChart.vue';
import BuildStatusChar from './components/BuildStatusChar.vue';

const metrics = shallowRef([
  {
    title: '成功构建',
    value: '1,248',
    icon: IconCheckCircle,
    color: 'var(--color-success-light-1)',
    trend: {
      value: '12.5%',
      direction: 'up',
      type: 'success'
    }
  },
  {
    title: '失败构建',
    value: '56',
    icon: IconCloseCircle,
    color: 'var(--color-danger-light-1)',
    trend: {
      value: '3.2%',
      direction: 'down',
      type: 'success'
    }
  },
  {
    title: '平均构建时间',
    value: '2m 45s',
    icon: IconClockCircle,
    color: 'var(--color-warning-light-1)',
    trend: {
      value: '5.1%',
      direction: 'down',
      type: 'success'
    }
  },
  {
    title: '今日部署',
    value: '42',
    icon: IconCode,
    color: 'var(--color-primary-light-1)',
    trend: {
      value: '8.7%',
      direction: 'up',
      type: 'success'
    }
  }
]);

const activePipelines = ref([
  {
    id: '2356',
    name: 'frontend-web',
    status: 'success',
    duration: '2m 12s',
    stage: '部署生产',
    tagColor: 'green'
  },
  {
    id: '2355',
    name: 'backend-service',
    status: 'processing',
    duration: '1m 45s',
    stage: '运行测试',
    tagColor: 'orange'
  },
  {
    id: '2354',
    name: 'mobile-app',
    status: 'danger',
    duration: '3m 28s',
    stage: '构建失败',
    tagColor: 'red'
  },
  {
    id: '2353',
    name: 'data-pipeline',
    status: 'warning',
    duration: '4m 15s',
    stage: '等待审批',
    tagColor: 'gold'
  }
]);

const recentActivities = ref([
  {
    id: 1,
    user: '张开发',
    action: '触发了构建',
    target: 'frontend-web #2356',
    time: '10分钟前',
    color: 'green'
  },
  {
    id: 2,
    user: '李测试',
    action: '部署了版本',
    target: 'backend-service v1.2.3',
    time: '25分钟前',
    color: 'blue'
  },
  {
    id: 3,
    user: '王运维',
    action: '创建了新环境',
    target: 'staging-environment',
    time: '1小时前',
    color: 'purple'
  },
  {
    id: 4,
    user: '系统',
    action: '完成了扫描',
    target: 'security-scan #142',
    time: '2小时前',
    color: 'gray'
  }
]);
</script>

<style scoped>
.dashboard-container {
  padding: 16px;
  max-width: 1600px;
  margin: 0 auto;
}

/* 指标卡片 */
.metrics-section {
  margin-bottom: 20px;
}

.metric-card {
  margin-bottom: 20px;
}

.metric-content {
  display: flex;
  align-items: center;
}

.metric-icon {
  width: 40px;
  height: 40px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-right: 16px;
  color: var(--color-white);
}

.metric-info {
  flex: 1;
}

.metric-title {
  display: block;
  color: var(--color-text-2);
  font-size: 12px;
}

.metric-value {
  margin: 4px 0;
}

.metric-trend {
  display: flex;
  align-items: center;
  font-size: 12px;
}

/* 主内容区 */
.main-section {
  margin-bottom: 20px;
}

.chart-card {
  margin-bottom: 20px;
  height: 500px;
}

.pipeline-card {
  height: 500px;
}

/* 活动时间线 */
.activity-card {
  height: 500px;
}

.activity-content {
  display: flex;
  align-items: center;
  gap: 8px;
}

.activity-time {
  font-size: 12px;
}

/* 响应式调整 */
@media (max-width: 768px) {
  .dashboard-container {
    padding: 8px;
  }

  .metric-card {
    margin-bottom: 12px;
  }
}
</style>
