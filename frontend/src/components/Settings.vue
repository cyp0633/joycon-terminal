<script setup>
import { reactive } from 'vue'
import { NButton, NInput, NAlert, NSpace, NSelect, useLoadingBar, NTabs, NTabPane, NH1, NText, NH2, NA } from 'naive-ui'
import { GetDevices,SetKey } from '../../wailsjs/go/main/App'

var data = reactive({
    manualSelectDevice: null,
    manualSelectKey: null,
    computerKey:null,
    availableDevices: [],
    availableKeys: [
        {
            type: "group",
            label: "按键",
            key: "keys",
            children: [
                {
                    label: "Key 1",
                    value: 1,
                },
                {
                    label: "Key 2",
                    value: 2,
                },
                {
                    label: "Key 3",
                    value: 3,
                },
            ]
        },
        {
            type: "group",
            label: "方向",
            key: "nav",
            children: [
                {
                    label: "上",
                    value: 4,
                },
                {
                    label: "下",
                    value: 5,
                },
                {
                    label: "左",
                    value: 6,
                },
                {
                    label: "右",
                    value: 7,
                },
                {
                    label: "确认",
                    value: 8,
                },
            ],
        },
    ],
})

function getDevices() {
    GetDevices().then(result => {
        console.log(result)
        data.availableDevices = result
    })
}

function setKey(){
    SetKey(data.manualSelectDevice,data.manualSelectKey,data.computerKey)
}

</script>

<template>
    <main>
        <n-h1 prefix="bar" class="text-left">
            <n-text type="primary">
                设置
            </n-text>
        </n-h1>
        <n-h2 class="text-left" prefix="bar">
            预设
        </n-h2>
        <div class="m-2">
            <n-text>在下面寻找你玩的游戏，快速进行键位设置。</n-text>
        </div>
        <n-space align="center" justify="center">
            <n-button>《极限竞速：地平线5》</n-button>
            <n-button>《原神》</n-button>
            <n-button>《森林冰火人》</n-button>
        </n-space>
        <n-h2 class="text-left" prefix="bar">自定映射</n-h2>
        <div class="m-2">
            <n-text>对于更广泛的用途，您可以在这里自定义按键对应关系。</n-text>
        </div>
        <div class="justify-center items-center flex flex-col">
            <div class="grid grid-cols-2 content-center justify-center items-center gap-4 w-3/6">
                <n-select children-field="children" label-field="label" value-field="value" filterable
                    :options="data.availableDevices" placeholder="选择设备" v-model:value="data.manualSelectDevice"
                    @click="getDevices" />
                <n-select children-field="children" label-field="label" value-field="value" filterable
                    :options="data.availableKeys" placeholder="选择按键" v-model:value="data.manualSelectKey" />
                <n-input v-model:value="data.computerKey" type="text" placeholder="映射的 PC 按键"></n-input>
                <n-button secondary type="primary" @click="setKey">设定</n-button>
            </div>
            <div class="m-2">
                <n-text>可用的按键名称请参见 </n-text>
                <n-a href="https://github.com/go-vgo/robotgo/blob/master/docs/keys.md" target="_blank">RobotGo 文档</n-a>
            </div>
        </div>
    </main>
</template>