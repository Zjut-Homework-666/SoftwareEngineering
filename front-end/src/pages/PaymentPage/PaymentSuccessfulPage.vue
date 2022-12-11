<template>
    <!-- ct:  container -->
    <div id="ct">
        <div class="Main_ct">
            <div class="logo_ct">
                <div style="width:45px; height:45px">
                  <img id="logo" src="../../assets/logo.png" alt="" width=45 height=45>
                </div>
                <span class=" logo_text">
                    <p>机票预订系统 支付界面</p>
                </span>
            </div>
            <div class="mainBox_ct">
                <QRCode></QRCode>
            </div>
        </div>
    </div>
    
</template>

<style scoped>
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
  height: 60px;
  line-height: 45px;
  text-align: left;
  display: flex;
  flex-direction: row;
}

#logo {
  object-fit:contain;
  vertical-align: middle;
}

.logo_text {
  width: 100%;
  height: 100%;
  font-size: 25px;
  font-weight: bold;
  text-align: left;
  color: #336B9F;
  margin-left: 10px;
  vertical-align: middle;
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

<script setup>
// import HomePage_bts from "./HomePage_bts.vue"
import QRCode from "./QRCode.vue";
// export default {
//     name: 'HomePage',
//     methods: {}
// }

</script>
