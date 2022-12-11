<template>
    <div class="bt_ct">
        <el-form class="bt_form">
            <el-form-item class="bt_row" label-width="0">
                <el-button class="bt" type="primary" @click="handleSchedule">
                    <span>预定机票</span>
                </el-button>
            </el-form-item>
            <el-form-item class="bt_row" label-width="0">
                <el-button class="bt" type="primary" @click="handleCheck">
                    <span>核验机票</span>
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

<script setup>
// import axios from 'axios';
import router from '../../router'

const handleSchedule = () => {
    console.log("Schedule")
    router.push('/Schedule')

}
const handleCheck = () => {
    console.log("PaymentSuccessful")
    router.push('/PaymentSuccessful')
}

// const handleCheck = () => {
//     console.log("Check")
//     router.push('/Check')
// }

</script>