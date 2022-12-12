<!--suppress ALL -->
<template>
    <el-form id="searchForm" :inline="true" :model="searchForm" class="demo-form-inline">
        <el-row>
            <el-col :span="2">
                <p class="formLabel">航班号</p>
            </el-col>
            <el-col :span="3">
                <el-form-item id="flightNoInput">
                    <el-input style="width: 200px" v-model="searchForm.flight" placeholder="航班号" />
                </el-form-item>
            </el-col>
        </el-row>

        <el-row>
            <el-col :span="2">
                <p class="formLabel">航班始末</p>
            </el-col>
            <el-col :span="20">
                <el-form-item id="destination">
                    <el-input style="width: 200px" v-model="searchForm.departure" placeholder="出发地" />
                    <p style="margin: auto 20px">飞往</p>
                    <el-input style="width: 200px" v-model="searchForm.destination" placeholder="目的地" />
                </el-form-item>
            </el-col>
        </el-row>

        <el-row>
            <el-col :span="2">
                <p class="formLabel">日期</p>
            </el-col>
            <div class="block">
                <el-date-picker
                        v-model="value1"
                        type="date"
                        placeholder="选择日期"
                        style="width: 200px"
                />
            </div>
        </el-row>

        <el-row>
            <el-col :span="2">
                <p class="formLabel">价格区间</p>
            </el-col>
            <el-col :span="21" >
                <el-form-item id="priceSection">
                    <div class="priceSlider">
                        <el-slider v-model="searchForm.price" range show-stops :max="1000" show-input/>
                    </div>
                    <el-input v-model="searchForm.price[0]" style="margin-left: 40px; width: 100px; " class="w-50 m-2"/>
                    <p style="margin-left: 20px;">~</p>
                    <el-input v-model="searchForm.price[1]" style="margin-left: 20px; width: 100px; " class="w-50 m-2"/>
                    <p style="margin-left: 10px;">元</p>
                    <el-button id="searchSubmitBtn" type="primary" @click="submitSearchForm">查询</el-button>
                </el-form-item>
            </el-col>
        </el-row>

    </el-form>
</template>

<script lang="ts" setup>
import { reactive, ref } from 'vue'

const value1 = ref('')

const searchForm = reactive({
    flight: '',
    departure: '',
    destination: '',
    price: [0, 1000],
})

const submitSearchForm = () => {
    console.log('submit!')
}

// const disabledDate = (time: Date) => {
//     return time.getTime() > Date.now()
// }
</script>

<style scoped>
* {
    margin: 0;
    padding: 0;
}

#searchForm {
    display: flex;
    justify-content: left;
    flex-flow: column;
}

.priceSlider {
    display: flex;
    align-items: center;
    width: 400px;
    margin-top: 0;
    margin-left: 12px;
}
#priceSection {
    display: flex;
    flex-flow: row;
}

#destination {
    display: flex;
    flex-flow: column;
    margin-top: 4px;
}

#searchSubmitBtn {
    width: 100px;
    height: 40px;
    margin-left: auto;
}

.formLabel {
    width: 100px;
    height: 40px;
    padding-left: 20px;
    text-align: left;
}

.el-checkbox-group {
    display: flex;
    flex-flow: row;
    justify-content:space-around;
    gap: 7px;
}

.el-col {
    text-align: center;
    float: left;
}

.el-row {
    line-height: 40px;
}

</style>
