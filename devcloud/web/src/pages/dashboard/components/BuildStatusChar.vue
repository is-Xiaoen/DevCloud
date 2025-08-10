<template>
  <a-card class="chart-card" title="构建状态统计" :bordered="false">
    <template #extra>
      <a-select v-model="buildTimeRange" size="small" style="width: 120px">
        <a-option value="week">最近一周</a-option>
        <a-option value="month">最近一月</a-option>
        <a-option value="quarter">最近三月</a-option>
      </a-select>
    </template>
    <div class="chart-wrapper">
      <v-chart class="chart" :option="buildOption" autoresize />
    </div>
  </a-card>
</template>

<script setup>
import { computed, ref } from 'vue';
import { use } from 'echarts/core';
import { CanvasRenderer } from 'echarts/renderers';
import { BarChart } from 'echarts/charts';
import {
  GridComponent,
  TooltipComponent,
  LegendComponent,
  TitleComponent
} from 'echarts/components';
import VChart from 'vue-echarts';

// 注册必要的组件
use([
  CanvasRenderer,
  BarChart,
  GridComponent,
  TooltipComponent,
  LegendComponent,
  TitleComponent
]);

// Arco Design 标准颜色值（亮色主题）
const colors = {
  success: 'rgb(76,210,99)', // --green-4
  danger: 'rgb(249,137,129)',  // --red-4
  warning: 'rgb(255,182,93)', // --orange-4
  text1: '#1D2129',   // --color-text-1
  text2: '#4E5969',   // --color-text-2
  border: '#E5E6EB'   // --color-border
};

const buildTimeRange = ref('week');

// 构建数据（根据时间范围动态生成）
const buildData = computed(() => {
  const dataMap = {
    week: {
      dates: ['周一', '周二', '周三', '周四', '周五', '周六', '周日'],
      success: [32, 45, 38, 52, 47, 60, 42],
      fail: [2, 3, 1, 0, 2, 1, 3],
      running: [5, 3, 6, 4, 7, 2, 4]
    },
    month: {
      dates: Array.from({ length: 30 }, (_, i) => `${i + 1}日`),
      success: Array.from({ length: 30 }, () => Math.floor(Math.random() * 50) + 30),
      fail: Array.from({ length: 30 }, () => Math.floor(Math.random() * 5)),
      running: Array.from({ length: 30 }, () => Math.floor(Math.random() * 8) + 2)
    },
    quarter: {
      dates: ['1月', '2月', '3月'],
      success: [1420, 1560, 1890],
      fail: [45, 32, 28],
      running: [120, 95, 110]
    }
  };

  return dataMap[buildTimeRange.value];
});

// 构建图表配置
const buildOption = computed(() => ({
  tooltip: {
    trigger: 'axis',
    axisPointer: {
      type: 'shadow'
    }
  },
  legend: {
    data: ['成功', '失败', '进行中'],
    right: 10,
    top: 0,
    textStyle: {
      color: colors.text1
    }
  },
  grid: {
    left: '3%',
    right: '4%',
    bottom: '3%',
    containLabel: true
  },
  xAxis: {
    type: 'category',
    data: buildData.value.dates,
    axisLine: {
      lineStyle: {
        color: colors.border
      }
    },
    axisLabel: {
      color: colors.text2,
      interval: buildTimeRange.value === 'month' ? 4 : 0
    }
  },
  yAxis: {
    type: 'value',
    name: '构建次数',
    nameTextStyle: {
      color: colors.text2
    },
    axisLine: {
      lineStyle: {
        color: colors.border
      }
    },
    axisLabel: {
      color: colors.text2
    },
    splitLine: {
      lineStyle: {
        color: colors.border
      }
    }
  },
  series: [
    {
      name: '成功',
      type: 'bar',
      stack: 'total',
      emphasis: {
        focus: 'series'
      },
      data: buildData.value.success,
      itemStyle: {
        color: colors.success
      }
    },
    {
      name: '失败',
      type: 'bar',
      stack: 'total',
      emphasis: {
        focus: 'series'
      },
      data: buildData.value.fail,
      itemStyle: {
        color: colors.danger
      }
    },
    {
      name: '进行中',
      type: 'bar',
      stack: 'total',
      emphasis: {
        focus: 'series'
      },
      data: buildData.value.running,
      itemStyle: {
        color: colors.warning
      }
    }
  ]
}));
</script>

<style scoped>
.chart-card {
  height: 100%;
  border-radius: 4px;
}

.chart-card :deep(.arco-card-body) {
  padding: 0;
  height: calc(100% - 56px);
}

.chart-wrapper {
  padding: 16px;
  height: 100%;
}

.chart {
  width: 100%;
  height: 100%;
}
</style>
