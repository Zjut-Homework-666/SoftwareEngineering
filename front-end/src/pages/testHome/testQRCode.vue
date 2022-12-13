<template>
    <div>二维码生成</div>
    <qrcode-vue :value="data.payUrl" :size="data.size" level="H" />
</template>

<script lang="ts" setup>
import { reactive } from 'vue'
import QrcodeVue from 'qrcode.vue'

const data = reactive<{
    payUrl: string
    size: number
}>({
    payUrl: '千万别搭理s---b',
    size: 256,
})
</script>