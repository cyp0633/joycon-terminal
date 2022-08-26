<script setup>
import { reactive } from 'vue'
import { Greet, ConnectSerial } from '../../wailsjs/go/main/App'

const data = reactive({
	name: "",
	resultText: "等待连接",
	status: false,
})

function connect() {
	if (data.name == "") {
		data.resultText = "请输入串口名称或路径"
		return
	}
	if (data.status == true) {
		data.resultText = "串口已连接"
		return
	}
	ConnectSerial(data.name).then(result => {
		data.resultText = result
	})
	if (result == "success") {
		data.status = true
	}
}

</script>

<template>
	<main>
		<div id="tip" class="tip h-16 mx-auto">
			<h3>输入串口名/路径</h3>对于 Windows，尝试 <code class="code">COMx</code>；否则，尝试 <code class="code">/dev/ttyUSBx</code>
		</div>
		<div id="input" class="input-box">
			<input id="name" v-model="data.name" autocomplete="off" class="input" type="text" />
			<button class="btn" @click="connect">连接</button>
		</div>
		<div id="result" class="result">{{ data.resultText }}</div>
	</main>
</template>

<style scoped>
.result {
	height: 20px;
	line-height: 20px;
	margin: 1.5rem auto;
}

/* .tip {
	height: 50px;
	line-height: 20px;
	margin: 1.5rem auto;
} */

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
