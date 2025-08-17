import { useStorage } from '@vueuse/core'

export default useStorage(
  'app',
  {
    current_system: 'DashBoard',
    current_menu: {
      DashBoard: {
        sub_menu: '',
        menu_item: '1',
      },
      ProjectSystem: {
        sub_menu: '',
        menu_item: '1',
      },
      DevelopSystem: {
        sub_menu: '',
        menu_item: '1',
      },
      ArtifactSystem: {
        sub_menu: '',
        menu_item: 'RegistryPage',
      },
    },
  },
  localStorage,
  { mergeDefaults: true },
)
