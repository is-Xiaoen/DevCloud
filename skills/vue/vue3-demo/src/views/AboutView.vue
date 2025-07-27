<template>
  <div class="about">
    <h1>This is an about page</h1>
    <!-- 响应式 -->
    <!-- 视图更新: Mode -> View -->
    <h2 id="name">{{ reverseName }}</h2>
    <!-- 数据更新 View -> Model -->
    <input v-model="hrefBindValue" type="text">

    <!-- 元素属性的动态绑定 -->
    <!-- href="{{ hrefBindValue }}" -->
    <!-- dom set set attr -->
    <!-- 编译器需要找到这个元素，为这个元素设置属性 set attr -->
    <!-- v-bind 语法   专门用于 元素属性的绑定 -->
    <!-- v-bind -> : -->
    <a :href="hrefBindValue" target="_blank"> {{ hrefBindValue }} </a>

    <!-- 原生的按钮点击 -->
    <!-- onclick="copyText()" -->
    <!-- 编译器需要找到这个元素，为这个元素添加事件 addEvent -->
    <!-- v-on 语法 专门用于 元素事件的绑定, onclick="copyText()" -->
    <!-- v-on -> @ -->
    <button @click="copyText">复制文本</button>

    <!-- class/style 都是元素的全局属性, 属性绑定的语法对他也生效 v-bind: class的值是有类型限制的 -->
    <!-- class="static active" -->
    <!-- { active: isActive } -->
    <span :class="dynamicClass">class样式绑定</span>
    <!-- style="color: red; font-size: 1em;" -->
    <span :style="dynamicStyle">style样式绑定</span>

    <!-- 表单的绑定
    https://cn.vuejs.org/guide/essentials/forms.html -->
    <!-- value属性绑定 Model -- View(用户输入框里面的值) -->
    <!-- input事件绑定 onInput View(用户在界面上的输入) -- Model -->
    <!-- <input :value="inputValue" @input="inputEventHandler"> -->

    <!-- 做了合并操作, 需要绑定元素的value, 并且监听input事件, 更新value属性的 这种操作
    给了个 上层封装: v-model -->
    <input v-model="inputValue">


    <!-- 列表渲染 -->
    <ul>
      <li v-for="item in items" :key="item.message">
        {{ item.message }}
      </li>
    </ul>

  </div>
</template>

<style>
@media (min-width: 1024px) {
  .about {
    min-height: 100vh;
    display: flex;
    align-items: center;
    flex-direction: column;
  }
}
</style>

<!-- vue组件里面 js逻辑部份 -->
<script setup>
import { computed, nextTick, onBeforeMount, onBeforeUnmount, onBeforeUpdate, onMounted, onUnmounted, onUpdated, ref } from 'vue';

// 声明响应式变量, reactive, ref
// {value: 'my vue3 demo'}
var name = ref('my vue3 demo')
// var name = 'my vue3 demo'

// name.value = 'xx'
// var name = reactive({ value: 'my vue3 demo' })
// name.value

// vue 组件生命周期验证
// mounted
onBeforeMount(() => {
  console.log('onBeforeMount')
})
onMounted(() => {
  console.log('onMounted')
})


// updated
onBeforeUpdate(() => {
  console.log('onBeforeUpdate')
})
onUpdated(() => {
  console.log('onUpdated')
})


// deleted
onBeforeUnmount(() => {
  console.log('onBeforeUnmount')
})
onUnmounted(() => {
  console.log('onUnmounted')
})


// 异步更新
// 只有等模版挂载好后，我门才能获取到对应的HTML元素
onMounted(() => {
  // 通过value来设置 基础类型的值(Setter方式)
  name.value = "张三";

  // 等待下一次更新周期的到来，才会修改dom(渲染)

  // UI上获取当前视图显示的值, 打印出来的值是什么
  // 初始化: my vue3 demo
  // 修改过后的值: 张三 ?
  //
  console.log(document.getElementById("name").innerText);


  // 怎么才能让他拿到新的喃?
  // 设置下载一次周期来的时候 在执行一个操作
  nextTick(() => {
    // 拿到就是界面更新后的数据
    console.log(document.getElementById("name").innerText);
  })
});


// 计算属性(缓存机制)
const reverseName = computed(() => name.value.split('').reverse().join(''))


// attr bind
const hrefBindValue = ref('https://www.baidu.com')



// 元素事件
const copyText = () => {
  alert('复制文本')
}


// 样式绑定, 使用map比使用字符串 更方便程序处理
// class="static active"
const dynamicClass = ref({ static: true, active: true })
dynamicClass.value.static = false

// color: red; font-size: 1em;
// 通过一个style的map
// css里面的属性 带有 - 需要被转换为驼峰
const dynamicStyle = ref({ color: 'red', fontSize: '2em' })

// 表单绑定
const inputValue = ref('')
// const inputEventHandler = (event) => {
//   // 用户输入的值，传递给inputValue, 输入框才会显示用户输入的值
//   inputValue.value = event.target.value
// }

// 响应式列表
const items = ref([{ message: 'msg1' }, { message: 'msg2' }, { message: 'msg3' }])

</script>
