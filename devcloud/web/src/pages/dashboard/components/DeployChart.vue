<template>
  <a-card title="部署统计" :bordered="false" class="chart-card">
    <v-chart class="chart" :option="deployOption" autoresize />
  </a-card>
</template>

<script setup>
import { ref, computed } from 'vue';
import { use } from 'echarts/core';
import { CanvasRenderer } from 'echarts/renderers';
import { BarChart, LineChart } from 'echarts/charts';
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
  LineChart,
  GridComponent,
  TooltipComponent,
  LegendComponent,
  TitleComponent
]);

// Arco Design 颜色配置（使用您调整后的颜色）
const colors = {
  primary: 'rgb(45,140,240)',   // --blue-6
  success: 'rgb(76,210,99)',    // --green-4
  danger: 'rgb(249,137,129)',   // --red-4
  warning: 'rgb(255,182,93)',   // --orange-4
  text1: '#1D2129',             // --color-text-1
  text2: '#4E5969',             // --color-text-2
  border: '#E5E6EB'             // --color-border
};

// 部署数据
const deployData = ref({
  environments: ['开发', '测试', '预发', '生产'],
  counts: [120, 85, 45, 32],
  successRate: [98.5, 95.2, 99.1, 99.8]
});

// 部署图表配置
const deployOption = computed(() => ({
  tooltip: {
    trigger: 'axis',
    axisPointer: {
      type: 'cross',
      crossStyle: {
        color: colors.border
      }
    },
    formatter: params => {
      const count = params[0].value;
      const rate = params[1].value;
      return `
        <div style="font-weight:bold;margin-bottom:5px">${params[0].axisValue}</div>
        <div>部署次数: <span style="color:${colors.primary}">${count}</span></div>
        <div>成功率: <span style="color:${colors.success}">${rate}%</span></div>
      `;
    }
  },
  legend: {
    data: ['部署次数', '成功率'],
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
    data: deployData.value.environments,
    axisLine: {
      lineStyle: {
        color: colors.border
      }
    },
    axisLabel: {
      color: colors.text2
    }
  },
  yAxis: [
    {
      type: 'value',
      name: '部署次数',
      min: 0,
      max: Math.max(...deployData.value.counts) * 1.2,
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
    {
      type: 'value',
      name: '成功率',
      min: 90,
      max: 100,
      axisLine: {
        lineStyle: {
          color: colors.border
        }
      },
      axisLabel: {
        color: colors.text2,
        formatter: '{value}%'
      },
      splitLine: {
        show: false
      }
    }
  ],
  series: [
    {
      name: '部署次数',
      type: 'bar',
      barWidth: 20,
      data: deployData.value.counts,
      itemStyle: {
        color: colors.primary
      },
      label: {
        show: true,
        position: 'top',
        color: colors.text1
      }
    },
    {
      name: '成功率',
      type: 'line',
      yAxisIndex: 1,
      data: deployData.value.successRate,
      symbol: 'circle',
      symbolSize: 8,
      itemStyle: {
        color: colors.success
      },
      lineStyle: {
        width: 3
      },
      label: {
        show: true,
        formatter: '{c}%',
        color: colors.text1,
        position: 'top'
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

.chart {
  width: 100%;
  height: 100%;
  min-height: 300px;
}
</style>
