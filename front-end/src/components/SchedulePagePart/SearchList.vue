<!--suppress ALL -->
<template>
    <el-table ref="tableRef" row-key="date" :data="showPage" style="width: auto; height: 70%; margin: 0 auto" v-fit-columns>
        <el-table-column
                prop="flight"
                label="航班号"
                width="120"
        />
        <el-table-column
                prop="depPlace"
                label="出发"
                width="120"
                sortable
                column-key="depPlace"
                :filters="[
        { text: '2016-05-01', value: '2016-05-01' },
        { text: '2016-05-02', value: '2016-05-02' },
        { text: '2016-05-03', value: '2016-05-03' },
        { text: '2016-05-04', value: '2016-05-04' },
      ]"
                :filter-method="filterHandler"
        />
        <el-table-column
                prop="arrPlace"
                label="到达"
                width="120"
                sortable
                column-key="depPlace"
                :filters="[
        { text: '2016-05-01', value: '2016-05-01' },
        { text: '2016-05-02', value: '2016-05-02' },
        { text: '2016-05-03', value: '2016-05-03' },
        { text: '2016-05-04', value: '2016-05-04' },
      ]"
                :filter-method="filterHandler"
        />
        <el-table-column
                prop="depTime"
                label="起飞时间"
                width="180"
                sortable
                column-key="depPlace"
                :filters="[
        { text: '2016-05-01', value: '2016-05-01' },
        { text: '2016-05-02', value: '2016-05-02' },
        { text: '2016-05-03', value: '2016-05-03' },
        { text: '2016-05-04', value: '2016-05-04' },
      ]"
                :filter-method="filterHandler"
        />
        <el-table-column
                prop="arrTime"
                label="到达时间"
                width="180"
                sortable
                column-key="depPlace"
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

        <el-table-column
                prop="status"
                label="航班状态"
                width="100"
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


        <el-table-column label="操作" width="100">
                <el-button type="primary" size="small">预订</el-button>
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
</template>
<script lang="ts" setup>
import { ref } from 'vue'
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

const tableRef = ref<InstanceType<typeof ElTable>>()
const currentPage = ref(4)
const pageSize = ref(5)
const small = ref(false)
const background = ref(true)
const disabled = ref(false)

// const flightStatus : {[key:string]:number} = {
//     '售票中': 0,
//     '未开售': 1,
//     '停飞': 2,
//     '已满': 3
// };

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

// 总数据
const tableData: flightDetailInfo[] = [
    {
        flight: 'B-1234',
        depPlace: '北京',
        arrPlace: '杭州',
        depTime: '2022-10-12 10:00',
        arrTime: '2022-10-12 12:00',
        price: 500,
        seats: 50,
        status:  '售票中',
    },
    {
        flight: 'B-1234',
        depPlace: '北京',
        arrPlace: '杭州',
        depTime: '2022-10-12 10:00',
        arrTime: '2022-10-12 12:00',
        price: 500,
        seats: 50,
        status:  '未开售',
    },
    {
        flight: 'B-1234',
        depPlace: '北京',
        arrPlace: '杭州',
        depTime: '2022-10-12 10:00',
        arrTime: '2022-10-12 12:00',
        price: 500,
        seats: 50,
        status:  '停飞',
    },
    {
        flight: 'B-1234',
        depPlace: '北京',
        arrPlace: '杭州',
        depTime: '2022-10-12 10:00',
        arrTime: '2022-10-12 12:00',
        price: 500,
        seats: 50,
        status:  '已满',
    },
    {
        flight: 'B-1234',
        depPlace: '北京',
        arrPlace: '杭州',
        depTime: '2022-10-12 10:00',
        arrTime: '2022-10-12 12:00',
        price: 500,
        seats: 50,
        status:  '已满',
    },
    {
        flight: 'B-1234',
        depPlace: '北京',
        arrPlace: '杭州',
        depTime: '2022-10-13 10:00',
        arrTime: '2022-10-13 12:00',
        price: 500,
        seats: 50,
        status:  '已满',
    },
    {
        flight: 'B-1234',
        depPlace: '北京',
        arrPlace: '杭州',
        depTime: '2022-10-14 10:00',
        arrTime: '2022-10-14 12:00',
        price: 500,
        seats: 50,
        status:  '已满',
    },
    {
        flight: 'B-1234',
        depPlace: '北京',
        arrPlace: '杭州',
        depTime: '2022-10-15 10:00',
        arrTime: '2022-10-15 12:00',
        price: 500,
        seats: 50,
        status:  '已满',
    },
]
// 显示数据
let showPage: flightDetailInfo[] = tableData.slice(0, pageSize.value);

// 每页条数控制
const handleSizeChange = () => {
    showPage = tableData.slice(0, pageSize.value);
}

// 换页控制
const handleCurrentChange = () => {
    showPage = tableData.slice((currentPage.value-1)*pageSize.value, currentPage.value*pageSize.value);
}
</script>

<style scoped>
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

