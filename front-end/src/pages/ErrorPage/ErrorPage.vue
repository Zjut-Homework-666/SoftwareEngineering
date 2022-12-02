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
                <span class=" logo_text">
                    <p>机票预订系统</p>
                </span>
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

.logo_text {
    font-size: 25px;
    font-weight: bold;
    text-align: left;
    color: #336B9F;
}

</style>

<script>
// import HomePage_bts from "./HomePage_bts.vue"

export default {
    name: 'HomePage',
    methods: {}
}

</script>