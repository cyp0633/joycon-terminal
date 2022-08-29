<script setup>
import { reactive } from 'vue'
import { ConnectSerial, StartListen, StopListen, Disconnect, GetAvailablePorts } from '../../wailsjs/go/main/App'
import { NButton, NInput, NAlert, NSpace, NSelect, useLoadingBar, NTabs, NTabPane, NH1, NText } from 'naive-ui'

var data = reactive({
	name: "",
	resultText: "等待连接",
	status: false,
	alertBox: "info",
})

// var serialOptions = []

function getAvailablePorts() {
	GetAvailablePorts().then(result => {
		console.log(result)
		serialOptions = result
	})
}

const loadingBar = useLoadingBar()

function connect() {
	loadingBar.start()
	if (data.name == "") {
		data.resultText = "请输入串口名称或路径"
		data.alertBox = "error"
		loadingBar.error()
		return
	}
	if (data.status == true) {
		loadingBar.finish()
		data.resultText = "串口已连接"
		data.alertBox = "warning"
		return
	}
	ConnectSerial(data.name).then(result => {
		data.resultText = result
		if (result == "success") {
			data.status = true
			data.alertBox = "success"
			loadingBar.finish()
			StartListen()
		}
		else {
			data.alertBox = "error"
			loadingBar.error()
		}
	})
}

function disconnect() {
	StopListen()
	Disconnect()
	data.alertBox = "info"
	data.status = false
	data.resultText = "等待连接"
}


</script>

<template>
	<main>
		<n-tabs class="card-tabs mx-4 w-11/12" size="large" animated>
			<n-tab-pane name="connect" tab="连接">
				<n-h1 prefix="bar" class="text-left">
					<n-text type="primary">
						连接
					</n-text>
				</n-h1>
				<div id="tip" class="tip mx-auto m-1.5 text-base">
					<n-text>输入串口名/路径<br>对于 Windows，尝试 </n-text>
					<n-text code>COMx</n-text>
					<n-text>；否则，尝试 </n-text>
					<n-text code>/dev/ttyUSBx</n-text>
				</div>
				<n-space align="center" justify="center">
					<n-select children-field="children" label-field="label" value-field="value" filterable
						:options="serialOptions" @click="getAvailablePorts" placeholder="选择串口" />
					<n-input id="name" v-model:value="data.name" class="m-1.5 w-2/6" type="text" placeholder="串口名" />
				</n-space>
				<n-space align="center" justify="center">
					<n-button type="primary" @click="connect" class="m-1.5" text-color="#18a058">连接</n-button>
					<n-button type="error" @click="disconnect" class="m-1.5" secondary>断开</n-button>
				</n-space>
				<n-alert title="连接状态" v-bind:type="data.alertBox" class="w-2/6 mx-auto inset-x-0 m-1.5">{{
						data.resultText
				}}
				</n-alert>
			</n-tab-pane>
			<n-tab-pane name="settings" tab="设置">
				<n-h1 prefix="bar" class="text-left">
					<n-text type="primary">
						设置
					</n-text>
				</n-h1>
				<n-button secondary type="success">Test</n-button>
				<p>Nothing here now.</p>
			</n-tab-pane>
		</n-tabs>
	</main>
</template>

<style scoped>
</style>
