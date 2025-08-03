<template>
  <div>
    <p> {{ modelValue }} </p>
    <input type="text" @input="handleInput" >
  </div>
</template>

<script setup>

// 组件属性
const { modelValue } =  defineProps({
  modelValue: {
    type: String,
    required: true,
    default: '-'
  }
})

// 组件事件
const emit = defineEmits(['update:modelValue'])

// 组件内不能修改属性, 因此我们需要生命另外一个变量去获取他的值
const handleInput = (e) => {
  emit('update:modelValue', e.target.value)
}
</script>

<style lang="css" scoped>

</style>
