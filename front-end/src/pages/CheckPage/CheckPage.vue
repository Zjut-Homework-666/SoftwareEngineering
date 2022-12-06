<template>
    <!-- ct:  container -->
    <div id="ct">
        <div class="main_ct">
            <div class="logo_ct">
                <span class=" logo_text">
                    <p>机票预订系统 核验界面</p>
                </span>
            </div>
            <div class="mainBox_ct">
                核验信息
                <el-input
                v-model="input1"
                class="w-50 m-2"
                size="large"
                placeholder="姓名"
                />
            </div>
        </div>
    </div>
    
</template>

<style>
* {
    margin: 0;
    padding: 0;
    width: 100%;
    height: 100%;

}

#ct {
    background-color: #e9e9e9;

}

.Main_ct {
    width: 70%;
    height: 70%;
    text-align: center;
    position: absolute;
    top: 0;
    bottom: 0;
    left: 0;
    right: 0;
    margin: auto;

}

.logo_ct {
    width: 100%;
    height: 10%;
    text-align: left;
    margin-top: -5%;
}

.logo_text {
    font-size: 25px;
    font-weight: bold;
    text-align: left;
    color: #336B9F;
}
.mainBox_ct {
    position: absolute;
    background-color: #f8f6f6;
    box-shadow: rgba(50, 50, 93, 0.25) 0px 50px 100px -20px,
        rgba(0, 0, 0, 0.3) 0px 30px 60px -30px,
        rgba(10, 37, 64, 0.35) 0px -2px 6px 0px inset;
    border-radius: 15px;
}

</style>

<script>
// import HomePage_bts from "./HomePage_bts.vue"

export default {
    name: 'HomePage',
    methods: {}
}

</script>