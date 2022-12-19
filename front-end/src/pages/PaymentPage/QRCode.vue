<template>
    <div class="bt_ct">
        <el-form class="bt_form" label-width="70px">
            <p class="title">请扫描二维码付款</p>
            <vue-qr background-color="#f8f6f6" :text='QRCodeURL' :size="300"></vue-qr>
            <div class="title">剩余支付时间：{{count}}</div>

            <el-form-item class="bt_row" label-width="0">
                <el-button class="bt" type="primary" @click="Cancel">
                    <span>取消</span>
                </el-button>
            </el-form-item>
            <el-dialog v-model="centerDialogVisible" title="付款成功" width="50%" align-center>
                <div class="dia">
                    <el-container>
                        <el-header>
                            <el-descriptions
                            title="航班信息"
                            direction="vertical"
                            :column="3"
                            :size="size"
                            >
                                <el-descriptions-item label="航班编号">{{ flightInfo.flight }}</el-descriptions-item>
                                <el-descriptions-item label="出发/到达">{{ flightInfo.depPlace }}/{{ flightInfo.arrPlace }}</el-descriptions-item>
                                <el-descriptions-item label="起飞时间">{{ flightInfo.depTime }}</el-descriptions-item>
                                <el-descriptions-item label="到达时间">{{ flightInfo.arrTime }}</el-descriptions-item>
                                <el-descriptions-item label="座位号">{{  flightInfo.seat  }}</el-descriptions-item>
                            </el-descriptions>
                            <el-descriptions
                            title="用户信息"
                            :column="4"
                            :size="size"
                            direction="vertical"
                            :style="blockMargin"
                            border
                            >
                            <el-descriptions-item label="姓名">{{  userInfo.name  }}</el-descriptions-item>
                            <el-descriptions-item label="电话号码">{{ userInfo.phone }}</el-descriptions-item>
                            <el-descriptions-item label="身份证号">{{ userInfo.id }}</el-descriptions-item>
                            <el-descriptions-item label="邮箱">{{ userInfo.mail }}</el-descriptions-item>
                            </el-descriptions>
                            <el-descriptions
                            title="订单信息"
                            :column="4"
                            :size="size"
                            direction="vertical"
                            :style="blockMargin"
                            border
                            >
                            <el-descriptions-item label="生成时间">{{  orderInfo.orderTime  }}</el-descriptions-item>
                            <el-descriptions-item label="价格">{{  orderInfo.price  }}</el-descriptions-item>
                            <el-descriptions-item label="订单编号">{{ orderInfo.orderId }}</el-descriptions-item>
                            <el-descriptions-item label="订单状态">{{  orderInfo.orderStatus  }}</el-descriptions-item>
                            </el-descriptions>
                        </el-header>
                    </el-container>
                </div>
                <template #footer>
                    <span class="dialog-footer">
                        <el-button type="primary" @click="Cancel">
                            确定
                        </el-button>
                    </span>
                </template>
            </el-dialog>
        </el-form>
    </div>
</template>

<style scoped>
#QRCode {
  object-fit:contain;
  vertical-align: middle;
}
.dia {
    height: 420px;
    overflow: auto;
}
.dialog-footer{
  margin-right: 10px;
}
.bt_ct {
    border-radius: 15px;
}
.title {
    width: 91%;
    height: 7%;
    padding-left: 25px;
    text-align: center;
    color:#1076a4;
    font-size: 25px;
    font-weight: bold
}

.bt_form {
    position: absolute;
    top: 0;
    bottom: 0;
    left: 0;
    right: 0;
    margin: auto;
    width: 40%;
    height: 90%;

}
.dia_form {
    position: absolute;
    top: 0;
    bottom: 0;
    left: 0;
    right: 0;
    margin: auto;
    width: 40%;
    height: 90%;

}

.bt_row {
    height: 10%;
    width: auto;
    margin-top: 5%;
    margin-left: 20%;
    justify-content: center;
    text-align: center;

}
.dia_row {
    height: 3%;
    width: auto;
    margin-top: 20%;
    margin-left: 0%;
    justify-content: center;
    text-align: center;
}
.bt {
    width: 75%;
    height: 100%;
    font-size: 20px;
    vertical-align: middle;
    text-align: center;
    line-height: 30px;
    border-radius: 5px;
    background-color: #228EC0;
}

.bt:focus,
.bt:hover {
    background-color: #1076a4;
}
</style>

<script setup>
import axios from 'axios';
import router from '../../router/index.js'
// import { useRoute } from "vue-router";
import { ElMessageBox } from 'element-plus'

import {getCurrentInstance, onMounted, ref,beforeDestroy} from "vue";
import VueQr from 'vue-qr/src/packages/vue-qr.vue'


let count = ref('');  //倒计时
let seconds= 60; // 1分钟的秒数

const proxy = getCurrentInstance().appContext.config.globalProperties


let id = router.currentRoute.value.query.id
let pay = router.currentRoute.value.query.pay
let cancel = router.currentRoute.value.query.cancel
const QRCodeURL = pay;
// const status = ref (1)
const centerDialogVisible = ref(false)

onMounted(()=>{
    Time() //调用定时器
    ReserveStatus()
})

const flag = ref(false)
beforeDestroy(()=>{
    flag.value = true;
})

let flightInfo = ref({
    flight:'',
    arrPlace:'',
    depPlace:'',
    depTime:'',
    arrTime:'',
    seat:''
})
let orderInfo = ref({
    orderTime:'',
    price:0,
    orderId:0,
    orderStatus:'',
})
let userInfo = ref({
    name:'',
    sex:'',
    phone:0,
    mail:'',
    id:'',
})

const ReserveStatus = () =>{
    axios.get(proxy.$url+proxy.$BackendPort+'/reserveStatus',{
        params:{
            orderId:id
        }
    }).then(function (ret){
        if(ret.data.responeInfo.code == 1){
            ReserveStatus()
        }
        if(ret.data.responeInfo.code == 0){
            //付款成功
            centerDialogVisible.value = true;
            console.log(ret.data)
            flightInfo.value.flight = ret.data.flightInfo.flight
            flightInfo.value.arrPlace = ret.data.flightInfo.arrPlace
            flightInfo.value.depPlace = ret.data.flightInfo.depPlace
            flightInfo.value.arrTime = ret.data.flightInfo.arrTime
            flightInfo.value.depTime = ret.data.flightInfo.depTime
            flightInfo.value.seat = ret.data.orderInfo.flightSeat.seat

            orderInfo.value.orderTime = ret.data.orderInfo.orderTime
            orderInfo.value.price = ret.data.orderInfo.price
            orderInfo.value.orderId = ret.data.orderInfo.orderId
            orderInfo.value.orderStatus = ret.data.orderInfo.orderStatus

            userInfo.value.name = ret.data.orderInfo.userInfo.name
            userInfo.value.sex = ret.data.orderInfo.userInfo.sex
            userInfo.value.phone = ret.data.orderInfo.userInfo.phone
            userInfo.value.mail = ret.data.orderInfo.userInfo.mail
            userInfo.value.id = ret.data.orderInfo.userInfo.id
        }
        if(ret.data.responeInfo.code == 2){
            //超时未付款
            Return()
        }
    })
}

// eslint-disable-next-line no-unused-vars
// const GetPaymentInfo = (pay,cancel,orderID) =>{
//     payURL.value = pay;
//     cancelURL.value = cancel;
//     orderId.value = orderID;
// }

const Time = () => {
    setInterval(() => {
        seconds -= 1
        countDown()
        if (seconds == 0)
            Return()
    }, 1000)
}

const countDown = () => {
    let m = parseInt((seconds/60%60).toString());
    m = m < 10 ? "0" + m : m
    let s = seconds % 60;
    s = s < 10 ? "0" + s : s
    count.value = m.toString() + '分' + s.toString() + '秒'
    // console.log(count.value)
    // count = '14:59'
}

const Cancel = () => {// eslint-disable-line no-unused-vars
    axios.get(cancel)
    router.push('/')
    centerDialogVisible.value = false;
}



const Return = () => {
    if(flag.value==false){
        axios.get(cancel)
        ElMessageBox.alert('支付超时，返回主界面', '提示', {
            confirmButtonText: '确定',
        })
        router.push('/')
    }
}
</script>
