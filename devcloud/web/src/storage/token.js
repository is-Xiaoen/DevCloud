import { useStorage } from '@vueuse/core'

export default useStorage('token', {}, localStorage, { mergeDefaults: true })
