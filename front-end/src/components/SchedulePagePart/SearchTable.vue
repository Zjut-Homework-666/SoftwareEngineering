<!--suppress ALL -->
<template>
    <div id="conditionArea">
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
    </div>
    <el-divider style="margin-top: 5px"/>
    <div id="listArea">
        <el-table ref="tableRef" row-key="date" :data="showPage" style="width: auto; height: 70%; margin: 0 auto" v-fit-columns>
            <el-table-column prop="flight" label="航班号" width="120"
            />
            <el-table-column prop="depPlace" label="出发" width="120" sortable column-key="depPlace"
                    :filters="[
        { text: '2016-05-01', value: '2016-05-01' },
        { text: '2016-05-02', value: '2016-05-02' },
        { text: '2016-05-03', value: '2016-05-03' },
        { text: '2016-05-04', value: '2016-05-04' },
      ]"
                    :filter-method="filterHandler"
            />
            <el-table-column prop="arrPlace" label="到达" width="120" sortable column-key="depPlace"
                    :filters="[
        { text: '2016-05-01', value: '2016-05-01' },
        { text: '2016-05-02', value: '2016-05-02' },
        { text: '2016-05-03', value: '2016-05-03' },
        { text: '2016-05-04', value: '2016-05-04' },
      ]"
                    :filter-method="filterHandler"
            />
            <el-table-column prop="depTime" label="起飞时间" width="180" sortable column-key="depPlace"
                    :filters="[
        { text: '2016-05-01', value: '2016-05-01' },
        { text: '2016-05-02', value: '2016-05-02' },
        { text: '2016-05-03', value: '2016-05-03' },
        { text: '2016-05-04', value: '2016-05-04' },
      ]"
                    :filter-method="filterHandler"
            />
            <el-table-column prop="arrTime" label="到达时间" width="180" sortable column-key="depPlace"
                    :filters="[
        { text: '2016-05-01', value: '2016-05-01' },
        { text: '2016-05-02', value: '2016-05-02' },
        { text: '2016-05-03', value: '2016-05-03' },
        { text: '2016-05-04', value: '2016-05-04' },
      ]"
                    :filter-method="filterHandler"
            />
            <el-table-column prop="price" label="最低价格" width="120" sortable/>
            <el-table-column prop="seats" label="剩余座位" width="120" />
            <el-table-column prop="status" label="航班状态" width="100"
                    :filters="[
        { text: '售票中', value: '售票中' },
        { text: '未开售', value: '未开售' },
        { text: '停飞', value: '停飞' },
        { text: '已满', value: '已满' },
      ]"
                    :filter-method="filterTag"
                    filter-placement="bottom-end"
            >
                <template #default="scope">
                    <el-tag
                            :type="tagCtrl(scope.row.status)"
                            disable-transitions
                    >{{ scope.row.status }}</el-tag
                    >
                </template>
            </el-table-column>
            <el-table-column label="操作" width="150">
                <el-button type="primary" size="small" @click="ScheduleButtonClicked">预订</el-button>
            </el-table-column>
        </el-table>
        <div id="pagination">
            <el-pagination
                    @size-change="handleSizeChange"
                    @current-change="handleCurrentChange"
                    style="margin: 0 auto"
                    v-model:current-page="currentPage"
                    v-model:page-size="pageSize"
                    :page-sizes="[5, 10, 15, 20]"
                    :small="small"
                    :disabled="disabled"
                    :background="background"
                    layout="total, sizes, prev, pager, next, jumper"
                    :total="tableData.length"
            />
        </div>
    </div>
</template>

<script lang="ts" setup>
import { ref,reactive,getCurrentInstance,onMounted } from 'vue'
// import Qs from 'qs'
import axios from 'axios';
import { ElTable, type TableColumnCtx } from 'element-plus'

interface flightDetailInfo {  // 机次信息
    flight: string  // 机次
    arrPlace: string  // 到达地
    depPlace: string  // 出发地
    depTime: string  // 起飞时间
    arrTime: string  // 到达时间
    price: number  // 最低价格
    seats: number  // 剩余座位
    status: string  // 状态
}

const searchForm = reactive({
    flight: '',         // 航班
    departure: '',      // 出发地
    destination: '',    // 目的地
    status: [],         // 状态
    price: [0, 1000],   // 价格区间
    date:''             // 日期
})

const proxy :any = getCurrentInstance().appContext.config.globalProperties

const tableRef = ref<InstanceType<typeof ElTable>>()
const currentPage = ref(4)
const pageSize = ref(5)
const small = ref(false)
const background = ref(true)
const disabled = ref(false)

const tagCtrl = (value: string) => {
    if (value == '售票中') return 'success';
    else if (value == '未开售') return 'warning';
    else if (value == '停飞') return 'info';
    else return 'danger';
}

const filterHandler = (
    value: string,
    row: flightDetailInfo,
    column: TableColumnCtx<flightDetailInfo>
) => {
    const property = column['property']
    return row[property] === value
}

const tableData = ref([
    {
        flight: 'B-1234',
        depPlace: '北京',
        arrPlace: '杭州',
        depTime: '2022-10-12 10:00',
        arrTime: '2022-10-12 12:00',
        price: 500,
        seats: 50,
        status:  '售票中',
    },{
        flight: 'B-1234',
        depPlace: '北京',
        arrPlace: '杭州',
        depTime: '2022-10-12 10:00',
        arrTime: '2022-10-12 12:00',
        price: 500,
        seats: 50,
        status:  '未开售',
    },{
        flight: 'B-1234',
        depPlace: '北京',
        arrPlace: '杭州',
        depTime: '2022-10-12 10:00',
        arrTime: '2022-10-12 12:00',
        price: 500,
        seats: 50,
        status:  '停飞',
    },{
        flight: 'B-1234',
        depPlace: '北京',
        arrPlace: '杭州',
        depTime: '2022-10-12 10:00',
        arrTime: '2022-10-12 12:00',
        price: 500,
        seats: 50,
        status:  '已满',
    },{
        flight: 'B-1234',
        depPlace: '北京',
        arrPlace: '杭州',
        depTime: '2022-10-12 10:00',
        arrTime: '2022-10-12 12:00',
        price: 500,
        seats: 50,
        status:  '已满',
    },{
        flight: 'B-1234',
        depPlace: '北京',
        arrPlace: '杭州',
        depTime: '2022-10-13 10:00',
        arrTime: '2022-10-13 12:00',
        price: 500,
        seats: 50,
        status:  '已满',
    },{
        flight: 'B-1234',
        depPlace: '北京',
        arrPlace: '杭州',
        depTime: '2022-10-14 10:00',
        arrTime: '2022-10-14 12:00',
        price: 500,
        seats: 50,
        status:  '已满',
    },{
        flight: 'B-1234',
        depPlace: '北京',
        arrPlace: '杭州',
        depTime: '2022-10-15 10:00',
        arrTime: '2022-10-15 12:00',
        price: 500,
        seats: 50,
        status:  '已满',
    }])

// 显示数据
let showPage: flightDetailInfo[] = tableData.value.slice(0, pageSize.value);

// 每页条数控制
const handleSizeChange = () => {
    showPage = tableData.value.slice(0, pageSize.value);
}

// 换页控制
const handleCurrentChange = () => {
    showPage = tableData.value.slice((currentPage.value-1)*pageSize.value, currentPage.value*pageSize.value);
}
// 页面刷新时 默认请求空条件查询
onMounted(() => {
    axios.get(proxy.$url+proxy.$BackendPort+"/flights", {
        params: {
            flight:'',        // 直接指定航班
            depPlace:'',    // 出发地（关键词匹配）
            arrPlace:'',   // 到达地（关键词匹配）
            minPrice:0,    // 选择的最低价格
            maxPrice:1000,    // 选择的最高价格   不能不填！！
            date:'',            // 选择的时间
        }
    })
        .then(function (ret) {
            console.log(ret.data)
            // tableData.value = []
            for(var i=0;i<ret.data.flights.length;i++){
                let tempflight : flightDetailInfo = {
                    flight : ret.data.flights[i].flight,  // 机次
                    arrPlace : ret.data.flights[i].arrPlace,  // 到达地
                    depPlace : ret.data.flights[i].depPlace,  // 出发地
                    depTime : ret.data.flights[i].depTime,  // 起飞时间
                    arrTime : ret.data.flights[i].arrTime,  // 到达时间
                    price : ret.data.flights[i].lowestPrice,  // 最低价格
                    seats : ret.data.flights[i].seatLeft,  // 剩余座位
                    status : ret.data.flights[i].seatLeft  // 状态
                }
                tableData.value.push(tempflight)
            }
        })
})
// 查询按钮
const submitSearchForm = () => {
    console.log('submit!')
    console.log(searchForm.flight)
    console.log(searchForm.departure)
    console.log(searchForm.destination)
    console.log(searchForm.price[0])
    console.log(searchForm.price[1])
    console.log(searchForm.date)

    axios.get(proxy.$url+proxy.$BackendPort+"/flights",{
        params:{
            flight:searchForm.flight,        // 直接指定航班
            depPlace:searchForm.departure,    // 出发地（关键词匹配）
            arrPlace:searchForm.destination,   // 到达地（关键词匹配）
            minPrice:searchForm.price[0],    // 选择的最低价格
            maxPrice:searchForm.price[1],    // 选择的最高价格   不能不填！！
            date:searchForm.date,            // 选择的时间
        }
    }).then(function (ret){
        console.log(ret.data)
        tableData.value = []
        for(var i=0;i<ret.data.flights.length;i++){
            let tempflight : flightDetailInfo = {
                flight : ret.data.flights[i].flight,  // 机次
                arrPlace : ret.data.flights[i].arrPlace,  // 到达地
                depPlace : ret.data.flights[i].depPlace,  // 出发地
                depTime : ret.data.flights[i].depTime,  // 起飞时间
                arrTime : ret.data.flights[i].arrTime,  // 到达时间
                price : ret.data.flights[i].lowestPrice,  // 最低价格
                seats : ret.data.flights[i].seatLeft,  // 剩余座位
                status : ret.data.flights[i].seatLeft  // 状态
            }
            tableData.value.push(tempflight)
        }
    })
}

const ScheduleButtonClicked = () =>{

}
</script>

<style scoped>
#conditionArea {
    width: 100%;
    height: 40%;
    margin-top: 1%;
    margin-left: 5%;
}

#listArea {
    padding-top: 10px;
    width: 100%;
    height: 100%;
    display: flex;
    flex-flow: column;
    overflow-x: hidden;
}

.el-divider--horizontal {
    margin: 0 auto;
    display: block;
    height: 1px;
    width: 96%;
    border-top: 1px var(--el-border-color) var(--el-border-style);
}
/*SearchOption*/
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
    margin-right: auto;
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

/*SearchList*/
.demo-pagination-block + #pagination {
    margin-top: 10px;
}
.demo-pagination-block .demonstration {
    margin-bottom: 16px;
}
#pagination {
    margin: 20px auto 0 auto;
    width: 100%;
    height: 40px;
    display: flex;
    flex-flow: row;
}
</style>