<script setup>
import HelloWorld from './components/HelloWorld.vue'
import { NLoadingBarProvider} from 'naive-ui'
</script>

<template>
	<n-loading-bar-provider>
		<HelloWorld />
	</n-loading-bar-provider>

</template>