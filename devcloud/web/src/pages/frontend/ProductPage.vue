<template>
  <div class="product-page">
    <!-- 核心价值部分 -->
    <section class="section">
      <a-divider class="driver" orientation="center">
        <a-typography-title :heading="4">平台核心价值</a-typography-title>
      </a-divider>

      <a-row :gutter="24" justify="center" class="value-cards">
        <a-col :xs="24" :sm="12" :md="8" v-for="(value, index) in values" :key="index">
          <a-card class="value-card" hoverable>
            <div class="card-icon">
              <component :is="value.icon" size="48" />
            </div>
            <a-card-meta :title="value.title" :description="value.description" class="card-meta" />
          </a-card>
        </a-col>
      </a-row>
    </section>

    <!-- 功能特性部分 -->
    <section class="section">
      <a-divider class="driver" orientation="center">
        <a-typography-title :heading="4">功能特性</a-typography-title>
      </a-divider>

      <a-list :grid="{ gutter: 16, xs: 1, sm: 2, lg: 2, xl: 2 }" :data="features" class="feature-list">
        <template #item="{ item }">
          <a-list-item>
            <a-card hoverable class="feature-card">
              <a-list-item-meta>
                <template #avatar>
                  <a-avatar :size="40" :style="{ backgroundColor: 'rgb(var(--primary-6))' }">
                    <icon-check />
                  </a-avatar>
                </template>
                <template #title>
                  <a-typography-text strong>{{ item.title }}</a-typography-text>
                </template>
                <template #description>
                  <a-typography-paragraph :ellipsis="{ rows: 2 }">{{ item.description }}</a-typography-paragraph>
                </template>
              </a-list-item-meta>
            </a-card>
          </a-list-item>
        </template>
      </a-list>
    </section>

    <!-- 客户案例部分 -->
    <!-- <section class="section">
      <a-divider orientation="center">
        <a-typography-title :heading="4">客户案例</a-typography-title>
      </a-divider>

      <div class="case-container">
        <a-carousel :auto-play="true" :interval="4000" animation-name="fade" show-arrow="hover" indicator-type="line"
          class="case-carousel">
          <a-carousel-item v-for="caseItem in cases" :key="caseItem.id">
            <a-card :title="caseItem.company" class="case-card" hoverable>
              <a-typography-paragraph>
                {{ caseItem.content }}
              </a-typography-paragraph>
              <template #extra>
                <a-link>查看详情</a-link>
              </template>
            </a-card>
          </a-carousel-item>
        </a-carousel>
      </div>
    </section> -->

    <!-- 行动号召部分 -->
    <section class="section cta-section">
      <a-result status="success" title="准备好提升您的研发效能了吗?" sub-title="我们的研发交付平台将帮助您的团队实现高效、稳定的软件交付">
        <template #extra>
          <a-space :size="24">
            <a-button type="primary" size="large" shape="round">免费试用</a-button>
            <a-button size="large" shape="round">预约演示</a-button>
          </a-space>
        </template>
      </a-result>
    </section>
  </div>
</template>

<script setup>
import {
  IconSync,
  IconBranch,
  IconUserGroup,
  IconCheck
} from '@arco-design/web-vue/es/icon';

const values = [
  {
    icon: IconSync,
    title: "持续集成",
    description: "自动化构建、测试流程，快速反馈代码质量，提升开发效率"
  },
  {
    icon: IconBranch,
    title: "持续交付",
    description: "一键部署，快速迭代，确保软件随时可发布"
  },
  {
    icon: IconUserGroup,
    title: "高效协作",
    description: "跨团队无缝协作，可视化流程，透明化管理"
  }
];

const features = [
  {
    title: '自动化构建与测试',
    description: '支持多种语言和框架的自动化构建，集成单元测试、集成测试等质量门禁'
  },
  {
    title: '多环境部署',
    description: '支持开发、测试、预发、生产等多环境一键部署，配置灵活可定制'
  },
  {
    title: '可视化流水线',
    description: '直观展示构建、测试、部署全流程，实时监控各环节状态'
  },
  {
    title: '质量门禁',
    description: '代码规范检查、测试覆盖率要求、安全扫描等多维度质量管控'
  }
];

// const cases = [
//   {
//     id: 1,
//     company: "某大型互联网公司",
//     content: "使用研发交付平台后，我们的部署频率提升了300%，故障率降低了60%，研发团队可以更专注于业务创新。",
//     image: 'https://p1-arco.byteimg.com/tos-cn-i-uwbnlip3yd/cd7a1aaea8e1c5e3d26fe2591e561798.png~tplv-uwbnlip3yd-webp.webp',
//   },
//   {
//     id: 2,
//     company: "某金融机构",
//     content: "平台的多环境管理和质量门禁功能帮助我们实现了合规要求，同时大幅提升了交付效率。",

//   },
//   {
//     id: 3,
//     company: "某电商企业",
//     content: "可视化流水线和团队协作功能让我们的跨部门协作更加顺畅，版本发布时间从原来的2周缩短到2天。"
//   }
// ];

</script>

<style scoped>
.product-page {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 16px;
}

.section {
  margin-bottom: 48px;
}

.driver {
  margin: 80px 0px;
}

/* 核心价值卡片 */
.value-cards {
  margin: 32px 0;
}

.value-card {
  height: 200px;
  text-align: center;
  transition: transform 0.2s;
}

.value-card:hover {
  transform: translateY(-5px);
}

.card-icon {
  display: flex;
  justify-content: center;
  padding: 24px 0 16px;
  color: rgb(var(--primary-6));
}

.card-meta {
  padding: 0 16px 24px;
}

/* 功能特性列表 */
.feature-list {
  margin: 32px 0;
}

.feature-card {
  height: 100%;
  transition: transform 0.2s;
}

.feature-card:hover {
  transform: translateY(-3px);
}

/* 客户案例 */
.case-container {
  padding: 0 16px;
}

.case-carousel {
  margin: 32px auto;
  max-width: 800px;
}

.case-card {
  padding: 24px;
  margin: 0 8px;
}

/* 行动号召部分 */
.cta-section {
  background: var(--color-fill-2);
  padding: 48px 24px;
  border-radius: 8px;
  margin-top: 64px;
}

/* 响应式调整 */
@media (max-width: 768px) {
  .section {
    margin-bottom: 32px;
  }

  .value-card,
  .feature-card {
    margin-bottom: 16px;
  }

  .cta-section {
    margin-top: 48px;
    padding: 32px 16px;
  }
}
</style>
