<script setup>
import { reactive } from 'vue'
import { NText, NH1, NH2, NA } from 'naive-ui'
</script>

<template>
    <main>
        <n-h1 prefix="bar" class="text-left">
            <n-text type="primary">
                关于
            </n-text>
        </n-h1>
        <n-text>Powered by Golang and Wails.</n-text><br />
        <n-text>Frontend: Vue, NaiveUI, TailwindCSS</n-text><br />
        <n-text>Copyright©cyp0633</n-text><br />
        <n-text>Open source at </n-text>
        <n-a href="https://github.com/cyp0633/joycon-terminal">GitHub</n-a>
        <n-text>.</n-text>
    </main>
</template>