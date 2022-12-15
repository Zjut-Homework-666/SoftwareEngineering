<template>
    <div class="bt_ct">
        <el-form class="bt_form" label-width="70px">
            <p class="title">请扫描二维码付款</p>
           <!-- <img id="QRCode" src="../../assets/QRCode.jpg" alt="" width=300 height=300> -->
<!--            <vue-qr text="Hello world!" :callback="test" qid="testid"></vue-qr>-->
            <vue-qr background-color="#f8f6f6" :text='QRCodeURL' :size="300"></vue-qr>
            <div class="title">剩余支付时间：{{count}}</div>
            <el-form-item class="bt_row" label-width="0">
                <el-button class="bt" type="primary" @click="Cancel">
                    <span>取消</span>
                </el-button>
            </el-form-item>
<!--            <el-dialog v-model="centerDialogVisible" title="付款成功" width="30%" align-center>-->
<!--                <div class="dia">-->
<!--                    <el-form class="dia_form">-->
<!--                    <el-form-item class="dia_row" label-width="75" label="姓名:">-->
<!--                        XXXXXX-->
<!--                    </el-form-item>-->
<!--                    <el-form-item class="dia_row" label-width="75" label="性别:">-->
<!--                        XXXXXX-->
<!--                    </el-form-item>-->
<!--                    <el-form-item class="dia_row" label-width="75" label="身份证号:">-->
<!--                        XXXXXX-->
<!--                    </el-form-item>-->
<!--                    <el-form-item class="dia_row" label-width="75" label="电话号码:">-->
<!--                        XXXXXX-->
<!--                    </el-form-item>-->
<!--                    <el-form-item class="dia_row" label-width="75" label="订单编号:">-->
<!--                        XXXXXX-->
<!--                    </el-form-item>-->
<!--                    </el-form>-->
<!--                </div>-->
<!--                <template #footer>-->
<!--                    <span class="dialog-footer">-->
<!--                        <el-button type="primary" @click="centerDialogVisible = false">-->
<!--                            Confirm-->
<!--                        </el-button>-->
<!--                    </span>-->
<!--                </template>-->
<!--            </el-dialog>-->
        </el-form>
    </div>
</template>

<style scoped>
#QRCode {
  object-fit:contain;
  vertical-align: middle;
}
.dia {
    height: 230px;
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

<script lang="ts" setup>
import axios from 'axios';
import router from '../../router/index.js'
// import { useRoute } from "vue-router";
import { ElMessageBox } from 'element-plus'

import {onMounted, ref} from "vue";
// import VueQr from 'vue-qr/src/packages/vue-qr.vue'


let count = ref('');  //倒计时
let seconds= 900; // 15分钟的秒数

// const proxy :any = getCurrentInstance().appContext.config.globalProperties



// let id = router.currentRoute.value.query.id
let pay = router.currentRoute.value.query.pay
let cancel = router.currentRoute.value.query.cancel
const QRCodeURL = pay;

onMounted(()=>{
    Time() //调用定时器
})

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
    let m = parseInt(seconds/60%60);
     m = m < 10 ? 0 + m : m
    let s = seconds % 60;
    s = s < 10 ? "0" + s : s
    count.value = m.toString() + '分' + s.toString() + '秒'
    // count = '14:59'
}

const Cancel = () => {
    axios.get(cancel)
    router.push('/')
}

const Return = () => {
    axios.get(cancel)
    ElMessageBox.alert('支付超时，返回主界面', '提示', {
        confirmButtonText: 'OK',
    })
    router.push('/')
}
</script>