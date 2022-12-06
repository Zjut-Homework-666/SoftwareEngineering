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
                <p class="formLabel">航班状态</p>
            </el-col>
            <el-col :span="5">
                <el-form-item>
                    <el-checkbox-group v-model="searchForm.status">
                        <el-checkbox label="售票中" name="status" />
                        <el-checkbox label="未开售" name="status" />
                        <el-checkbox label="停飞" name="status" />
                        <el-checkbox label="已满" name="status" />
                    </el-checkbox-group>
                </el-form-item>
            </el-col>
        </el-row>

        <el-row>
            <el-col :span="2">
                <p class="formLabel">航班始末</p>
            </el-col>
            <el-col :span="20">
                <el-form-item id="destination">
                    <el-select style="width: 200px" v-model="searchForm.departure" placeholder="出发地">
                        <el-option label="Zone one" value="shanghai" />
                        <el-option label="Zone two" value="beijing" />
                    </el-select>
                    <p style="margin: auto 20px">飞往</p>
                    <el-select style="width: 200px" v-model="searchForm.destination" placeholder="目的地">
                        <el-option label="Zone one" value="shanghai" />
                        <el-option label="Zone two" value="beijing" />
                    </el-select>
                </el-form-item>
            </el-col>
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
                    <el-button id="searchSubmitBtn" type="primary" @click="submitSearchForm">查调</el-button>
                </el-form-item>
            </el-col>
        </el-row>

    </el-form>
</template>

<script lang="ts" setup>
import { reactive,getCurrentInstance } from 'vue'
// import Qs from 'qs'
import axios from 'axios';

const searchForm = reactive({
    flight: '',
    departure: '',
    destination: '',
    status: [],
    price: [0, 1000],
})

const proxy :any = getCurrentInstance().appContext.config.globalProperties

const submitSearchForm = () => {
    console.log('submit!')

    axios.get('http://localhost:'+proxy.$BackendPort+"/flights").then(function (ret){
        console.log("successuful!")
        console.log(ret)
    })

    // axios.post('http://127.0.0.1:"+BackendPort+"/api/Pub/Course' + '?Limit=10&Pages=' + Number(1).toString(), {
    //     "No": SearchNo.value,
    //     "Name": SearchName.value,
    //     "Credit": SearchCredit.value,
    //     "Hour": SearchHour.value,
    //     "Method": SearchMethod.value,
    //     "Sem": SearchSem.value,
    //     "Tno": SearchTno.value,
    //     "Tname": SearchTname.value
    // })
    //     .then(function (ret) {
    //         console.log(ret.data)
    //         for (let i = 0; i < ret.data.len; i++) {
    //             // totalPrice += (this.books[i].price) * (this.books[i].count);
    //             tableData.value.push(ret.data.data[i])
    //         }
    //         TeaTotal.value = ret.data.count
    //     })
}
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
