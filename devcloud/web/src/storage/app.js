import { useStorage } from '@vueuse/core'

export default useStorage(
  'app',
  {
    current_system: 'DashBoard',
    system_menu: [
      {
        key: 'DashBoard',
        current_sub_menu: '',
        current_menu_item: 'DashboardPage',
        title: '工作台',
      },
      {
        key: 'ProjectSystem',
        current_sub_menu: '',
        current_menu_item: 'ProjectList',
        title: '项目管理',
        menus: [
          {
            key: 'ProjectList',
            icon: 'IconLock',
            title: '项目空间',
          },
        ],
      },
      {
        key: 'ResourceSystem',
        current_sub_menu: '',
        current_menu_item: 'ResourceSearch',
        title: '资源管理',
        menus: [
          {
            key: 'ResourceSearch',
            icon: 'IconSearch',
            title: '资源检索',
          },
          {
            key: 'EnvManage',
            icon: 'IconLocation',
            title: '环境管理',
          },
          {
            key: 'AssetSync',
            icon: 'IconCloud',
            title: '资产同步',
          },
          {
            key: 'K8sCluster',
            icon: 'IconCommon',
            title: 'k8s集群',
          },
        ],
      },
      {
        key: 'DevelopSystem',
        current_sub_menu: '',
        current_menu_item: 'AppPage',
        title: '研发交付',
        menus: [
          {
            key: 'AppPage',
            icon: 'IconApps',
            title: '应用管理',
          },
          {
            key: 'VersionIteration',
            icon: 'IconTags',
            title: '版本迭代',
          },
          {
            key: 'PipelineTemplate',
            icon: 'IconSettings',
            title: '流水线模板',
          },
        ],
      },
      {
        key: 'ArtifactSystem',
        current_sub_menu: '',
        current_menu_item: 'RegistryPage',
        title: '制品库',
        menus: [
          {
            key: 'RegistryPage',
            icon: 'IconTags',
            title: '制品管理',
          },
          {
            key: 'AssetPage',
            icon: 'IconBranch',
            title: '制品管理',
          },
        ],
      },
      {
        key: 'TestSystem',
        current_sub_menu: '',
        current_menu_item: 'RegistryPage',
        title: '测试中心',
        menus: [
          {
            key: 'TestApply',
            icon: 'IconStamp',
            title: '提测申请',
          },
          {
            key: 'TestCase',
            icon: 'IconExperiment',
            title: '测试用例',
          },
          {
            key: 'TestExecution',
            icon: 'IconSend',
            title: '用例执行',
          },
          {
            key: 'DefectManage',
            icon: 'IconBug',
            title: '缺陷管理',
          },
          {
            key: 'TestReport',
            icon: 'IconDashboard',
            title: '测试报告',
          },
        ],
      },
      {
        key: 'SecuritySystem',
        current_sub_menu: '',
        current_menu_item: 'RegistryPage',
        title: '安全合规',
        menus: [
          {
            key: 'VulnerabilityScan',
            icon: 'IconFindReplace',
            title: '漏洞扫描',
          },
        ],
      },
      {
        key: 'OpsSystem',
        current_sub_menu: '',
        current_menu_item: 'RegistryPage',
        title: '运维中心',
        menus: [
          {
            key: 'OnlineApply',
            icon: 'IconStamp',
            title: '上线申请',
          },
          {
            key: 'DeployManage',
            icon: 'IconLayers',
            title: '应用部署',
          },

          {
            key: 'MonitorAlarm',
            icon: 'IconNotification',
            title: '监控告警',
          },
        ],
      },
    ],
  },
  localStorage,
  { mergeDefaults: true },
)
