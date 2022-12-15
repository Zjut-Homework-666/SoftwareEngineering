<!--suppress ALL -->
<template #default="scope">

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
                            v-model="searchForm.date"
                            format='YYYY-MM-DD'
                            value-format="YYYY-MM-DD"
                            type='date'
                            placeholder='选择日期'
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
    </div>

    <el-divider style="margin-top: 5px"/>

    <div id="listArea">
        <el-table ref="tableRef" row-key="date" :data="showPage" style="width: auto; height: 70%; margin: 0 auto" :key="Math.random()" @sort-change="sortChange">
            <el-table-column prop="flight" label="航班号" width="120" sortable='custom'/>
            <el-table-column prop="depPlace" label="出发" width="120" column-key="depPlace"/>
            <el-table-column prop="arrPlace" label="到达" width="120" column-key="depPlace"/>
            <el-table-column prop="depTime" label="起飞时间" width="180" column-key="depPlace" sortable='custom'/>
            <el-table-column prop="arrTime" label="到达时间" width="180" column-key="depPlace" sortable='custom'/>
            <el-table-column prop="price" label="最低价格" width="120" sortable='custom'/>
            <el-table-column prop="seats" label="剩余座位" width="120" sortable='custom'/>
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
                <template #default="scope">
                    <el-button type="primary" size="default" style="width: 40%" text @click="showChart(scope.row)" >
                        预订
                    </el-button>
                </template>
            </el-table-column>

        </el-table>
        <el-dialog
                v-model="dialogVisible"
                title="座位选择"
                width="30%"
                :open-delay="300"
        >
            <div ref="chartDom" style="width: 400px; height: 800px"></div>
            <template #footer>
              <span class="dialog-footer">
                <el-button @click="dialogVisible = false" style="width: 100px;height: 40px;">
                    取消
                </el-button>
                <el-button type="primary" @click="SchheduleButton()" style="width: 100px;height: 40px;">
                    确定
                </el-button>
              </span>
            </template>
        </el-dialog>
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
import {ref, reactive, getCurrentInstance, onMounted} from 'vue'
// import Qs from 'qs'
import axios from 'axios';
import {ElMessageBox, ElTable, type TableColumnCtx} from 'element-plus'
import * as echarts from 'echarts';
// import jutils from '../../../node_modules/jutils'
import router from '../../router'

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

interface seatDetailInfo {
    flight: string  // 机次
    seat: string  // 座位
    price: string  // 价格
    status: string  // 状态
}

// 查询表单
const searchForm = reactive({
    flight: '',         // 航班
    departure: '',      // 出发地
    destination: '',    // 目的地
    price: [0, 10000],   // 价格区间
    date:''             // 日期
})

const proxy :any = getCurrentInstance().appContext.config.globalProperties

let flightSeat = reactive<any>({});
const chartDom = ref();
const tableRef = ref<InstanceType<typeof ElTable>>()
const currentPage = ref(1)
const pageSize = ref(5)
const small = ref(false)
const background = ref(true)
const disabled = ref(false)
const tableData = ref([
    {
        flight: '',
        depPlace: '',
        arrPlace: '',
        depTime: '',
        arrTime: '',
        price: 0,
        seats: 0,
        status:  '',
    }])
const flag = ref(0);
const dialogVisible = ref(false)
let showPage = ref([])
let seatInfo = ref([{
    flight: '',  // 机次
    seat: '',  // 座位
    price: '',  // 价格
    status: ''  // 状态
}]);
const flightSelect = ref('')
let selectedNames = ref(String);  // 选择的座位

onMounted(() => {
    // 页面刷新时 默认请求空条件查询
    axios.get(proxy.$url+proxy.$BackendPort+"/flights", {
        params: {
            flight:'',        // 直接指定航班
            depPlace:'',    // 出发地（关键词匹配）
            arrPlace:'',   // 到达地（关键词匹配）
            minPrice:0,    // 选择的最低价格
            maxPrice:1000,    // 选择的最高价格   不能不填！！
            date:'',            // 选择的时间
        }
    }).then(function (ret) {
        // console.log(ret.data)
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
                status : ret.data.flights[i].status  // 状态
            }
            tableData.value.push(tempflight)
        }
        showPage.value = tableData.value.slice(0, pageSize.value);
    });
})

// 显示数据
// 每页条数控制
const handleSizeChange = () => {
    showPage.value = tableData.value.slice(0, pageSize.value);
}
// 换页控制
const handleCurrentChange = () => {
    showPage.value = tableData.value.slice((currentPage.value-1)*pageSize.value, currentPage.value*pageSize.value);
}

const collectSeatInfo : any = async (flight) => {
    // 座位信息
    await axios.get(proxy.$url+proxy.$BackendPort+"/seats", {
        params: {
            flight:flight,        // 直接指定航班
        }
    }).then(function (ret) {
        // console.log('ret.data.seats:'+JSON.stringify(ret.data.seats))
        seatInfo.value =  ret.data.seats;
        // console.log('seatInfo.value:'+JSON.stringify(seatInfo.value))
    });
}

// 处理返回的座位信息
const handleSeatInfo = (seatInfo: seatDetailInfo[]) => {
    var regions = [];
    for (var i = 0; i < seatInfo.length; i++) {
        if(seatInfo[i].status === '已预定') {
            regions.push({
                name: seatInfo[i].seat,
                silent: true,
                itemStyle: {
                    color: '#99CCFF'
                },
                emphasis: {
                    itemStyle: {
                        borderColor: '#aaa',
                        borderWidth: 1
                    }
                },
                select: {
                    itemStyle: {
                        color: '#99CCFF'
                    }
                }
            });
        } else if (seatInfo[i].status === '已付款') {
            regions.push({
                name: seatInfo[i].seat,
                silent: true,
                itemStyle: {
                    color: '#bf0e08'
                },
                emphasis: {
                    itemStyle: {
                        borderColor: '#aaa',
                        borderWidth: 1
                    }
                },
                select: {
                    itemStyle: {
                        color: '#bf0e08'
                    }
                }
            });
        }
    }
    // console.log('seats:'+ JSON.stringify(seatInfo))
    return regions;
}

// 显示预定飞机座位图形
const showChart = (rowData) => {
    flightSelect.value = rowData.flight
    dialogVisible.value = true;
    type EChartsOption = echarts.EChartsOption;
    var option: EChartsOption;

    if(flag.value === 1) {

        option = {
            tooltip: {},
            geo: {
                map: 'flight-seats',
                roam: false,
                selectedMode: 'single',
                layoutCenter: ['30%', '50%'],
                layoutSize: '200%',
                tooltip: {
                    show: true
                },
                itemStyle: {
                    color: '#fff'
                },
                emphasis: {
                    itemStyle: {
                        color: undefined,
                        borderColor: 'green',
                        borderWidth: 2
                    },
                    label: {
                        show: false
                    }
                },
                select: {
                    itemStyle: {
                        color: 'green'
                    },
                    label: {
                        show: false,
                        textBorderColor: '#fff',
                        textBorderWidth: 2
                    }
                },
                regions: []
            }
        };

        flightSeat.setOption(option);
    }

    setTimeout(async ()=>{

        await collectSeatInfo(rowData.flight)

        axios.get('./flight-seats.svg').then((ret) => {
            if(flag.value === 0) {
                flightSeat = echarts.init(chartDom.value)
                echarts.registerMap('flight-seats', { svg: ret.data });
            }

            option = {
                tooltip: {},
                geo: {
                    map: 'flight-seats',
                    roam: false,
                    selectedMode: 'single',
                    layoutCenter: ['30%', '50%'],
                    layoutSize: '200%',
                    tooltip: {
                        show: true
                    },
                    itemStyle: {
                        color: '#fff'
                    },
                    emphasis: {
                        itemStyle: {
                            color: undefined,
                            borderColor: 'green',
                            borderWidth: 2
                        },
                        label: {
                            show: false
                        }
                    },
                    select: {
                        itemStyle: {
                            color: 'green'
                        },
                        label: {
                            show: false,
                            textBorderColor: '#fff',
                            textBorderWidth: 2
                        }
                    },
                    regions: handleSeatInfo(seatInfo.value)
                }
            };

            flightSeat.setOption(option, true);

            // Get selected seats.
            if(flag.value === 0){
                flag.value = 1;
                flightSeat.on('geoselectchanged', function (params: any) {
                    selectedNames.value = params.allSelected[0].name[0];
                    console.log('select', selectedNames.value);
                });
            }
        });

        option && flightSeat.setOption(option);

    },100)
}

// 查询按钮
const submitSearchForm = () => {
    console.log(tableData.value)
    // var date = jutils.formatDate(new Date(1562672641*1000),"YYYY-MM-DD")
    console.log(searchForm.date)
    // var date2=date.Format("yyyy-MM-dd hh:mm:ss");
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
        // 清空 tableData
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
                status : ret.data.flights[i].status  // 状态
            }
            tableData.value.push(tempflight)

        }
        console.log(tableData)
        showPage.value = tableData.value.slice(0, pageSize.value);
    })
}

// 预定操作

// const handleClose = (done: () => void) => {
//     ElMessageBox.confirm('确认关闭?')
//             .then(() => {
//                 done()
//             })
//             .catch(() => {
//                 // catch error
//             })
// }

const filterTag = (value: string,
                   row: flightDetailInfo,
                   column: TableColumnCtx<flightDetailInfo>
) => {
    const property = column['property']
    return row[property] === value
}

// 表格排序、过滤、标签过滤

/**
 * @method
 * @param column
 * @subparam prop: 'xxxx', // el-table-column中的prop
 * @subparam order: 'xxxx', // 'ascending' or 'descending'
 */
const sortChange = (column) => {
    currentPage.value = 1
    if (column.prop === 'flight') {
        if (column.order === 'descending') {
            tableData.value = tableData.value.sort(flightDescSort)
        } else if (column.order === 'ascending') {
            tableData.value = tableData.value.sort(flightAscSort)
        }
    } else if (column.prop === 'depTime') {
        if (column.order === 'descending') {
            tableData.value = tableData.value.sort(depTimeDescSort)
        } else if (column.order === 'ascending') {
            tableData.value = tableData.value.sort(depTimeAscSort)
        }
    } else if (column.prop === 'arrTime') {
        if (column.order === 'descending') {
            tableData.value = tableData.value.sort(arrTimeDescSort)
        } else if (column.order === 'ascending') {
            tableData.value = tableData.value.sort(arrTimeAscSort)
        }
    } else if (column.prop === 'price') {
        if (column.order === 'descending') {
            tableData.value = tableData.value.sort(priceDescSort)
        } else if (column.order === 'ascending') {
            tableData.value = tableData.value.sort(priceAscSort)
        }
    } else if (column.prop === 'seats') {
        if (column.order === 'descending') {
            tableData.value = tableData.value.sort(seatsDescSort)
        } else if (column.order === 'ascending') {
            tableData.value = tableData.value.sort(seatsAscSort)
        }
    }
    showPage.value = tableData.value.slice(0, pageSize.value)
}

// 排序函数
const flightDescSort = (a, b) => {
    if (a.flight > b.flight) {
        return -1
    } else if (a.flight < b.flight) {
        return 1
    } else {
        return 0
    }
}
const flightAscSort = (a, b) => {
    if (a.flight < b.flight) {
        return -1
    } else if (a.flight > b.flight) {
        return 1
    } else {
        return 0
    }
}
const depTimeDescSort = (a, b) => {
    console.log(a)
    if (a.arrTime > b.arrTime) {
        return -1
    } else if (a.arrTime < b.arrTime) {
        return 1
    } else {
        return 0
    }
}
const depTimeAscSort = (a, b) => {
    if (a.arrTime < b.arrTime) {
        return -1
    } else if (a.arrTime > b.arrTime) {
        return 1
    } else {
        return 0
    }
}
const arrTimeDescSort = (a, b) => {
    if (a.arrTime > b.arrTime) {
        return -1
    } else if (a.arrTime < b.arrTime) {
        return 1
    } else {
        return 0
    }
}
const arrTimeAscSort = (a, b) => {
    if (a.arrTime < b.arrTime) {
        return -1
    } else if (a.arrTime > b.arrTime) {
        return 1
    } else {
        return 0
    }
}
const seatsDescSort = (a, b) => {
    if (a.seats > b.seats) {
        return -1
    } else if (a.seats < b.seats) {
        return 1
    } else {
        return 0
    }
}
const seatsAscSort = (a, b) => {
    if (a.seats < b.seats) {
        return -1
    } else if (a.seats > b.seats) {
        return 1
    } else {
        return 0
    }
}
const priceDescSort = (a, b) => {
    if (a.price > b.price) {
        return -1
    } else if (a.price < b.price) {
        return 1
    } else {
        return 0
    }
}
const priceAscSort = (a, b) => {
    if (a.price < b.price) {
        return -1
    } else if (a.price > b.price) {
        return 1
    } else {
        return 0
    }
}

// 标签
const tagCtrl = (value: string) => {
    if (value == '售票中') return 'success';
    else if (value == '未开售') return 'warning';
    else if (value == '停飞') return 'info';
    else return 'danger';
}

// const store = useStore()

// const store = useStore()/
// import bus from '../../utils'

const SchheduleButton = () =>{
    console.log(flightSelect.value)
    let userInfo = {
        name : 'zz五',
        sex : '女',
        phone : 11231231231,
        mail : '1470603076@qq.com',
        id : '12312312312312312'
    }
    let flightSeat = {
        flight : flightSelect.value,
        seat : selectedNames.value
    }
    let data = {
        userInfo,
        flightSeat
    }

    let Res = {
        OrderId:0,
        payUrl:'',
        cancelUrl:''
    }
    console.log(data)
    axios.post(proxy.$url+proxy.$BackendPort+'/reserve',data)
            .then(function (ret) {
                console.log(ret.data)
                Res.payUrl = ret.data.payUrl;
                Res.cancelUrl = ret.data.cancelUrl;
                Res.OrderId = ret.data.orderId;
                // 如果成功
                if(ret.data.responeInfo.code == 0){
                    router.push({
                        path: '/PaymentPage',
                        query: {
                            id : Res.OrderId,
                            pay : Res.payUrl,
                            cancel : Res.cancelUrl
                        }
                    })
                    dialogVisible.value = false;
                }
                else{
                    console.log('这个座位已被预定')
                    ElMessageBox.alert('这个座位已被预定', '提示', {
                        confirmButtonText: '确定',
                    })
                }
            })


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
    margin-top: 6px;
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

.dialog-footer button:first-child {
    margin-right: 10px;
}
</style>
