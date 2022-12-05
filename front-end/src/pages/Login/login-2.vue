<template>
  <div class="video-container">
    <video :style="fixStyle" autoplay loop muted class="fillWidth" :canplay="canplay" id="login-bg-vedio">
      <source autoplay src="@/assets/login-video-bg.mp4" />
    </video>
  </div>
  <div class="mask">
  </div>
  <!-- <el-main color="#fbbf24" class="container mx-auto my-35 login-container max-h-650px max-w-600px login-form"> -->
  <el-main color="#fbbf24" class="container mx-auto my-40 login-container max-h-11/20 max-w-1/3 login-form">
    <el-header>
      <span
        class="text-gray-500 text-center tracking-wide font-medium subpixel-antialiased text-3xl form-header">欢迎登录提瓦特学生管理系统</span>
    </el-header>
    <el-form class="max-h-3/5" ref="loginForm" :model="loginForm" :rules="loginRules" autocomplete="on"
      label-position="top">
      <el-form-item label="" prop="username" class="max-h-10 m-top-10 max-w-9/10">
        <el-input ref="username" v-model="loginForm.username" placeholder="请输入学号或工号" name="username" type="text"
          tabindex="1" autocomplete="on" class="mx-2 input" />
      </el-form-item>
      <el-form-item label="" prop="password" class="max-h-10 m-top-10 max-w-9/10">
        <el-input ref="password" v-model="loginForm.password" type="password" placeholder="请输入密码" name="password"
          tabindex="2" autocomplete="on" @keyup.enter="handleLogin" class="mx-2 input" />
      </el-form-item>

      <!-- <div class="max-h-40px left-2 xieyi text-left">
        <el-radio class="max-w-300px text-2xl text-left" border label="2" size="large">同意《用户协议》和隐私政策</el-radio>
      </div> -->
      <el-button class="sub-btn :hover:bg-yellow-700" type="primary" :center="true" color=#393b40 :loading="loading"
        style="width:100%;
        margin-bottom:30px;" @click="handleLogin">
        <span class="text-color">登 录</span>
      </el-button>
    </el-form>
  </el-main>
</template>

<style scoped>
.homepage-hero-module,
.video-container {
  position: relative;
  height: 100vh;
  overflow: hidden;
}

.xieyi {
  text-align: left;
  padding-left: 50px;
  font-size: 20px;
}

.video-container {
  z-index: -1000;
  position: absolute;
}

.fixStyle {
  object-fit: cover
}

.fillWidth {
  width: 100%;
  height: auto;
}

video {
  position: fixed;

  right: 0;

  bottom: 0;

  min-width: 100%;

  min-height: 100%;

  width: auto;

  height: auto;

  z-index: -9999;

  /*灰色调*/

  /*-webkit-filter:grayscale(100%)*/
}

::-webkit-scrollbar {
  width: 0 !important;
}

::-webkit-scrollbar {
  width: 0 !important;
  height: 0;
}

.login-container {
  padding: 15px !important;
  text-align: center;
}

.login-form {
  background-color: rgb(255, 255, 255);
  border-radius: 10px;
  justify-content: center;
  text-align: center;
  opacity: 1;
  z-index: 100;
}

.form-header {
  max-height: 10%;
  justify-content: center;
  text-align: center;
  margin-left: auto;
  margin-right: auto;
  margin-top: 1rem;
  max-width: 50%;
  font-family: "Microsoft YaHei UI";
}

.login-formitem {
  padding: 0 0 0 0;
  opacity: 1;
}

.sub-btn {
  padding-top: 3%;
  margin-top: 20px;
  margin-left: 30px;
  margin-right: auto;
  margin-top: 1rem;
  max-width: 50%;
  font-family: "Microsoft YaHei UI";
  font-weight: normal;
  font-size: 25px;
  text-align: center;
  opacity: 1;
  z-index: 100;
  min-width: 400px;
  min-height: 60px;
  border-radius: 10px;
}

.mask {
  background-color: rgb(0, 0, 0);
  opacity: 0.4;
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: -5
}

.input {
  min-height: 80px;
  padding-bottom: 20px;
  font-size: 25px;
  border-radius: 20px;
  height: 300% !important;
}

.text-color {
  --tw-text-opacity: 1;
  color: #f4d8a8;
  font-size: 30px;
}
</style>

<script>
import Qs from 'qs'
import cookies from "vue-cookies";
import router from '../../router';

export default {
  name: 'Login',
  data() {
    const validateUsername = (rule, value, callback) => {
      if (value.length < 1) {
        callback(new Error('请输入正确的学号或工号'))
      } else {
        callback()
      }
    }
    const validatePassword = (rule, value, callback) => {
      if (value.length < 6) {
        callback(new Error('请输入正确格式的密码'))
      } else {
        callback()
      }
    }
    // const formref = ref(null)
    return {
      No: "S02",
      loginForm: {
        username: '',
        password: ''
      },
      loginRules: {
        username: [{ required: true, trigger: 'blur', validator: validateUsername }],
        password: [{ required: true, trigger: 'blur', validator: validatePassword }]
      },
      loading: false,
      showDialog: false,
      vedioCanPlay: false,
      fixStyle: '',
    }
  },
  created() {
    // window.addEventListener('storage', this.afterQRScan)
  },
  mounted() {

    if (this.loginForm.username === '') {
      this.$refs.username.focus()
    } else if (this.loginForm.password === '') {
      this.$refs.password.focus()
    }

    let m = function (e) { e.preventDefault(); };
    document.body.style.overflow = 'hidden';
    document.addEventListener("touchmove", m, { passive: false });//禁止页面滑动
  },
  destroyed() {
    // window.removeEventListener('storage', this.afterQRScan)
  },
  changed() {
    this.$forceUpdate()
  },
  methods: {
    handleLogin() {
      // this.axios.post('api/Stu/Login',
      if (this.loginForm.username[0] == 'S') {
        this.axios.post('http://47.96.8.176:12345/api/Stu/Login',
          Qs.stringify({
            No: this.loginForm.username,
            Password: this.loginForm.password
          }), {
          headers: { 'Content-Type': 'application/x-www-form-urlencoded' }
        }).then(function (ret) {
          console.log(ret.data)
          if (ret.data.code == 0) {
            // cookies.set('No', ret.data.no)
            router.push('/Student')
          }
        }).catch(function (ret) {
          console.log('error')
        });
        // console.log(cookies.get('No'))
      }
      else if (this.loginForm.username[0] == 'A'){
        router.push('/Admin')
      }
      else {
        this.axios.post('http://47.96.8.176:12345/api/Tea/Login',
          Qs.stringify({
            No: this.loginForm.username,
            Password: this.loginForm.password
          }), {
          headers: { 'Content-Type': 'application/x-www-form-urlencoded' }
        }).then(function (ret) {
          console.log(ret.data)
          if (ret.data.code == 0) {
            // cookies.set('No', ret.data.no)
            router.push('/Teacher')
          }
        }).catch(function (ret) {
          console.log('error')
        });
      }
    },
    canplay() {
      this.vedioCanPlay = true
    },

  }
}

</script>