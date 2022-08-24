<script setup>
import HelloWorld from './components/HelloWorld.vue'
</script>

<template>
  <img id="logo" alt="a logo to be replaced later" src="./assets/images/logo-universal.png"/>
  <HelloWorld/>
</template>

<style>
#logo {
  display: block;
  width: 50%;
  height: auto;
  margin: auto;
  padding: 10% 0 0;
  background-position: center;
  background-repeat: no-repeat;
  background-size: 100% 100%;
  background-origin: content-box;
}
</style>
