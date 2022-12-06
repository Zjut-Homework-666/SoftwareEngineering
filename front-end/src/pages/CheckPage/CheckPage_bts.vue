<template>
    <div class="bt_ct">
        <el-form class="bt_form" label-width="100px" :model="formLabelAlign">
            <el-form-item label="信息核验" label-position="Top"></el-form-item>
            <el-form-item label="姓名">
                <el-input v-model="userInfo.username.username" placeholder="请输入姓名"></el-input>
            </el-form-item>
            <el-form-item label="身份证号">
                <el-input v-model="userInfo.idnum" placeholder="请输入身份证号"></el-input>
            </el-form-item>
            <el-form-item label="手机号">
                <el-input v-model="userInfo.phonenum" placeholder="请输入手机号"></el-input>
            </el-form-item>
            <el-form-item label="邮箱">
                <el-input v-model="userInfo.email" placeholder="请输入邮箱"></el-input>
            </el-form-item>
            <el-form-item label="订单编号">
                <el-input v-model="userInfo.ordernum" placeholder="请输入订单编号"></el-input>
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

.bt_form {
    position: absolute;
    top: 0;
    bottom: 0;
    left: 0;
    right: 0;
    margin: auto;
    width: 80%;
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
    width: 70%;
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
import { ref } from 'vue'

const userInfo = ref({
    username: '',
    idnum: '',
    phonenum: '',
    email: '', 
    ordernum:''
})

</script>