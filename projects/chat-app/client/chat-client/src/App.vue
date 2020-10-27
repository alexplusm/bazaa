<template>
  <img alt="Vue logo" src="./assets/logo.png">
  <HelloWorld msg="Welcome to Your Vue.js App"/>
  <MyFile />
</template>

<script>
import HelloWorld from './components/HelloWorld.vue'
import MyFile from './components/MyFile.vue'

export default {
  name: 'App',
  components: {
    HelloWorld,
    MyFile
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
