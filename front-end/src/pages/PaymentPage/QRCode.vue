
<template>
    <div class="bt_ct">
        <el-form class="bt_form" label-width="70px">
            <p class="title">请扫描二维码付款</p>
            <img id="QRCode" src="../../assets/QRCode.jpg" alt="" width=300 height=300>
            <div class="title">剩余支付时间：{{count}}</div>
            <el-form-item class="bt_row" label-width="0">
                <el-button class="bt" type="primary" @click="Cancel">
                    <span>取消</span>
                </el-button>
            </el-form-item>
        </el-form>
    </div>
</template>

<style scoped>
#QRCode {
  object-fit:contain;
  vertical-align: middle;
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

.bt_row {
    height: 10%;
    width: auto;
    margin-top: 5%;
    margin-left: 20%;
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

<script>
// import axios from 'axios';
import router from '../../router'
// import { ref } from 'vue'
import {  ElMessageBox } from 'element-plus'

export default {
    data() {
        return {
            count: '', //倒计时
            seconds: 900 // 10天的秒数
        }
    },
    mounted() {
        this.Time() //调用定时器
    },
    methods: {
        //分 秒 格式化函数
        countDown() {
            let m = parseInt(this.seconds / 60 % 60);
            m = m < 10 ? "0" + m : m
            let s = parseInt(this.seconds % 60);
            s = s < 10 ? "0" + s : s
            this.count = m + '分' + s + '秒'
        },
        //定时器没过1秒参数减1
        Time() {
            setInterval(() => {
                this.seconds -= 1
                this.countDown()
                if (this.seconds == 0) 
                    Return()
            }, 1000)
        },
    }
}
const Return = () => {
    console.log("HomePage")
    router.push('/HomePage')
    ElMessageBox.alert('支付超时，返回主界面', '提示', {
        confirmButtonText: 'OK',
    })
}
</script>