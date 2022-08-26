<script setup>
import { reactive } from 'vue'
import { Greet, ConnectSerial } from '../../wailsjs/go/main/App'
import { NButton, NInput, NAlert, NSpace } from 'naive-ui'

const data = reactive({
	name: "",
	resultText: "等待连接",
	status: false,
	alertBox: "info",
})

function connect() {
	if (data.name == "") {
		data.resultText = "请输入串口名称或路径"
		data.alertBox = "error"
		return
	}
	if (data.status == true) {
		data.resultText = "串口已连接"
		data.alertBox = "warning"
		return
	}
	ConnectSerial(data.name).then(result => {
		data.resultText = result
	})
	if (result == "success") {
		data.status = true
		data.alertBox = "success"
	}
	else {
		data.alertBox = "error"
	}
}

</script>

<template>
	<main>
		<div id="tip" class="tip h-16 mx-auto">
			<h3>输入串口名/路径</h3>对于 Windows，尝试 <code class="code">COMx</code>；否则，尝试 <code class="code">/dev/ttyUSBx</code>
		</div>
		<n-space align="center" justify="center">
			<n-input id="name" v-model:value="data.name" class="m-1.5 w-2/6" type="text" placeholder="串口名" />
			<n-button type="primary" @click="connect" class="m-1.5">连接</n-button>
		</n-space>
		<n-alert title="连接状态" v-bind:type="data.alertBox" class="w-2/6 mx-auto inset-x-0">{{ data.resultText }}
		</n-alert>
	</main>
</template>

<style scoped>
.result {
	height: 20px;
	line-height: 20px;
	margin: 1.5rem auto;
}

.input-box .btn {
	width: 60px;
	height: 30px;
	line-height: 30px;
	border-radius: 3px;
	border: none;
	margin: 0 0 0 20px;
	padding: 0 8px;
	cursor: pointer;
}

.input-box .btn:hover {
	background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
	color: #333333;
}

.input-box .input {
	border: none;
	border-radius: 3px;
	outline: none;
	height: 30px;
	line-height: 30px;
	padding: 0 10px;
	background-color: rgba(240, 240, 240, 1);
	-webkit-font-smoothing: antialiased;
}

.input-box .input:hover {
	border: none;
	background-color: rgba(255, 255, 255, 1);
}

.input-box .input:focus {
	border: none;
	background-color: rgba(255, 255, 255, 1);
}

.code {
	font-family: Consolas, "Liberation Mono", Menlo, Courier, monospace;
	font-size: 14px;
	color: rgba(255, 255, 255, 1);
	border-radius: 3px;
	padding: 0 10px;
}
</style>
