<script setup>
import HelloWorld from './components/HelloWorld.vue'
import { NLoadingBarProvider, NConfigProvider, NGlobalStyle, useOsTheme, darkTheme } from 'naive-ui'

// dark theme
const osThemeRef = useOsTheme();
function theme() {
	if(osThemeRef.value === 'dark') {
		return  {
			theme: "dark",
			osTheme: osThemeRef
		};
	}
	else return {
		theme: null,
		osTheme: osThemeRef
	};
}
</script>

<template>
	<n-config-provider :theme="theme()">
		<n-global-style />
		<n-loading-bar-provider>
			<HelloWorld />
		</n-loading-bar-provider>
	</n-config-provider>
</template>