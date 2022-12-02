<template>
    <!-- ct:  container -->
    <div id="ct">
        <div class="main_ct">
            <div class="logo_ct">
                <!-- 
                原型里的logo下不下来
                !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
                !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
                !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
                !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
                !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
                !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
                !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
                !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
                !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
                !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
                -->
                <span class=" logo_text"><p>机票预订系统</p></span>
            </div>
            <div class="bt_ct">
                <el-form class="bt_form">
                    <el-form-item class="bt_row" label-width="0">
                        <el-button class="bt" type="primary">
                            <span class="sp">预定机票</span>
                        </el-button>
                    </el-form-item>
                    <el-form-item class="bt_row" label-width="0">
                        <el-button class="bt" type="primary">
                            <span class="sp">核验机票</span>
                        </el-button>
                    </el-form-item>
                </el-form>

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

.main_ct {
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
}

.logo_text{
    font-size: 25px;
    font-weight: bold;
    text-align: left;
    color: #336B9F;
}

.bt_ct {
    background-color: #f8f6f6;
    border: 1px solid #cac8c8;
    box-shadow: 10px 10px 5px #8e8e8e;

    border-radius: 15px;
}

.bt_form {
    position: absolute;
    top: 0;
    bottom: 0;
    left: 0;
    right: 0;
    margin: auto;
    width: 70%;
    height: 50%;

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
    line-height: -15px;
    border-radius: 5px;
    background-color: #228EC0;
}
.bt:focus,
.bt:hover {
    background-color: #1076a4;
}
</style>

<script>

export default {
    name: 'HomePage'
}

</script>