<template>
    <div class="bt_ct">
        <el-form class="bt_form" label-width="70px">
            <p class="title">核验信息</p>
            <el-form-item label="姓名">
                <el-input v-model="user_Info.username" placeholder="请输入姓名" clearable></el-input>
            </el-form-item>
            <el-form-item label="身份证号">
                <el-input v-model="user_Info.idnum" placeholder="请输入身份证号" clearable></el-input>
            </el-form-item>
            <el-form-item label="手机号">
                <el-input v-model="user_Info.phonenum" placeholder="请输入手机号" clearable></el-input>
            </el-form-item>
            <el-form-item label="邮箱">
                <el-input v-model="user_Info.email" placeholder="请输入邮箱" clearable></el-input>
            </el-form-item>
            <el-form-item label="订单编号">
                <el-input v-model="user_Info.ordernum" placeholder="请输入订单编号" clearable></el-input>
            </el-form-item>
            <el-form-item class="bt_row" label-width="0">
                <el-button class="bt" type="primary" @click="submitUserInfo">
                    <span>提交</span>
                </el-button>
            </el-form-item>
        </el-form>
    </div>
</template>

<style scoped>
.bt_ct {
    border-radius: 15px;
}
.title {
    width: 100%;
    height: 7%;
    padding-left: 25px;
    text-align: center;
    color:#1076a4;
    font-size: 30px;
    font-weight: bold
}

.bt_form {
    position: absolute;
    top: 0;
    bottom: 0;
    left: 0;
    right: 0;
    margin: auto;
    width: 30%;
    height: 70%;

}

.bt_row {
    height: 10%;
    width: auto;
    margin-top: 10%;
    margin-left: 30%;
    justify-content: center;
    text-align: center;

}

.bt {
    width: 80%;
    height: 120%;
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
// import axios from 'axios';
// import router from '../../router'
import {getCurrentInstance, ref} from 'vue'
import { ElMessageBox } from 'element-plus'
import axios from "axios";
// import type { Action } from 'element-plus'

const user_Info = ref({
    username: '',
    idnum: '',
    phonenum: '',
    email: '', 
    ordernum:''
})

const InfoWrong = () => {// eslint-disable-line no-unused-vars
    ElMessageBox.alert('信息错误,请重新输入', '提示', {
    // if you want to disable its autofocus
    // autofocus: false,
    confirmButtonText: 'OK',
    // center:true
  })
}

const proxy :any = getCurrentInstance().appContext.config.globalProperties
const submitUserInfo = () =>{
    console.log("CheckInfo!")
    let config = {
        headers: {'Content-Type': "multipart/json, charset=UTF-8"}
    };
    let userInfo = {
        name: user_Info.value.username,
        phone:user_Info.value.phonenum,
        id:user_Info.value.idnum
    };
    axios.post(proxy.$url+proxy.$BackendPort+"/check",userInfo,config)
        .then(function (ret){
            console.log(ret.data)

    })
}

</script>