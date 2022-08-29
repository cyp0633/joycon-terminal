<script setup>
import HelloWorld from './components/HelloWorld.vue'
import { NLoadingBarProvider, NConfigProvider, NGlobalStyle, useOsTheme, darkTheme } from 'naive-ui'
</script>

<template>
	<n-config-provider :theme="theme">
		<n-global-style />
		<n-loading-bar-provider>
			<HelloWorld />
		</n-loading-bar-provider>
	</n-config-provider>
</template>