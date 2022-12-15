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
            <el-dialog v-model="centerDialogVisible" title="核验成功" width="45%" align-center>
                <div class="dia">
                    <el-container>
                        <el-header class="el-header">AIR TICKET</el-header>
                        <el-main>
                            <el-descriptions title="BOARDING PASS" direction="vertical" >
                                <el-descriptions-item label="PASSANGER NAME">MURPHY</el-descriptions-item>
                                <el-descriptions-item label="DATE">2022.12.13</el-descriptions-item>
                                <el-descriptions-item label="BOARING TIME">10:30</el-descriptions-item>
                                <el-descriptions-item label="FLIGHT NO.">LJ6581</el-descriptions-item>
                                <el-descriptions-item label="FROM/TO">杭州/北京</el-descriptions-item>
                                <el-descriptions-item label="SEAT">21B</el-descriptions-item>
                            </el-descriptions>
                        </el-main>
                        <el-footer class="el-footer">GATE CLOSES 30 MINUTES BEFORE DEPARTURE</el-footer>
                    </el-container>
                </div>
                <template #footer>
                    <span class="dialog-footer">
                        <el-button type="primary" @click="centerDialogVisible = false">
                            确认
                        </el-button>
                    </span>
                </template>
            </el-dialog>
        </el-form>
    </div>
</template>

<style scoped>
.bt_ct {
    border-radius: 15px;
}
.el-header{
    width: 100%;
    height: 0%;
    padding-right: 300px;
    text-align: center;
    color:aliceblue;
    background-color: #1076a4;
    font-size: 30px;
    /* font-weight: bold */
}
.el-footer{
    width: 100%;
    height: 0%;
    padding-left: 300px;
    text-align: center;
    color:aliceblue;
    background-color: #1076a4;
    font-size: 10px;
    /* font-weight: bold */
}
.dia {
    height: 265px;
    overflow: auto;
}
.dialog-footer{
  margin-right: 10px;
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
    margin-top: 9%;
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

const centerDialogVisible = ref(false)
// const labelPosition = ref('Top')

const user_Info = ref({
    username: '',
    idnum: '',
    phonenum: '',
    email: '',
    ordernum:''
})

const InfoWrong = () => {// eslint-disable-line no-unused-vars
    ElMessageBox.alert('信息错误,请重新输入', '提示', {
    confirmButtonText: 'OK',
    // center:true
  })
}
const InfoEmpty = () => {// eslint-disable-line no-unused-vars
    ElMessageBox.alert('信息不可为空,请重新输入', '提示', {
    confirmButtonText: 'OK',
  })
}

const proxy :any = getCurrentInstance().appContext.config.globalProperties
const submitUserInfo = () => {// eslint-disable-line no-unused-vars
    if (user_Info.value.username.length > 0
        && user_Info.value.phonenum.length > 0
        && user_Info.value.idnum.length > 0
        && user_Info.value.email.length > 0
        && user_Info.value.ordernum.length > 0) {
        console.log("CheckInfo!")
        let config = {
            headers: { 'Content-Type': "multipart/json, charset=UTF-8" }
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
        centerDialogVisible.value=true
    } else
        InfoEmpty()
}

</script>
