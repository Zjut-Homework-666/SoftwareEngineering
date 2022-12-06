<template>
<div class="video-container">
  <video  :style="fixStyle" autoplay loop muted class="fillWidth" :canplay="canplay" id="login-bg-vedio">
    <source autoplay src="@/assets/login-video-bg.mp4" />
  </video>
  <div class="container mx-auto my-60 w-1/3 h-2/5 login-form">
    <el-main>
      <div class="container mx-auto form-header">
        <label class="">欢迎登陆提瓦特学生管理系统</label>
      </div>
      <label>工号/学号:</label>
      <el-input
        ref="username"
        v-model="loginForm.username"
        placeholder="请输入用户名"
        name="username"
        type="text"
        tabindex="1"
        autocomplete="on"
        class="max-h-10"
      />
      <label>密码:</label>
      <el-input
          ref="password"
          v-model="loginForm.password"
          type="password"
          placeholder="请输入密码"
          name="password"
          tabindex="2"
          autocomplete="on"
          @keyup.enter="handleLogin"
          class="max-h-10"
      />
      <el-button class="sub-btn" :loading="loading" type="primary" style="width:100%;margin-bottom:30px;" @click="handleLogin">Login</el-button>
    </el-main>
  </div>

</div>
  
</template>

<style>
*{
  margin: 0;
  padding: 0; 
  width: 100%;
  height: 100%;
/* 清除浏览器边距用 */
}
.homepage-hero-module,
.video-container {
    position: relative;
    heviight: 100vh;
    overflow: hidden;
}

.video-container {
  z-index: 0;
  position: absolute;
}
.fixStyle{
  object-fit: cover 
}
.fillWidth {
  width: 100%;
}
video{
  position: fixed;

  right:0;

  bottom: 0;

  min-width: 100%;

  min-height: 100%;

  width: auto;

  height: auto;

  z-index: -9999;

  /*灰色调*/

  /*-webkit-filter:grayscale(100%)*/
}
.login-form {
  background-color: rgba(129, 126, 126, 0.502);
  border-radius: 25px;
  justify-content: center;
  
}
.form-header{
  max-height: 10%;
  justify-content: center;
  text-align: center;
}
.login-formitem{
  padding:0 0 0 0;
}
.sub-btn{
  margin-left: 10px;
  margin-top: 20px;
  justify-content: center;
  max-width: 30%;
  height: auto;
}
</style>

<script>
import Qs from 'qs'
// import LoginForm from "@/pages/Login/LoginForm.vue";

export default {
  name: 'Login',
  data() {
    const validateUsername = (rule, value, callback) => {
        callback()
    }
    const validatePassword = (rule, value, callback) => {
      if (value.length < 6) {
        callback(new Error('The password can not be less than 6 digits'))
      } else {
        callback()
      }
    }
    return {
      loginForm: {
        username: 'S01',
        password: '123456789'
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
    // window.onresize = () => {
    //   const windowWidth = document.body.clientWidth
    //   const windowHeight = document.body.clientHeight
    //   const windowAspectRatio = windowHeight / windowWidth
    //   let videoWidth
    //   let videoHeight
    //   if (windowAspectRatio < 0.5625) {
    //     videoWidth = windowWidth
    //     videoHeight = videoWidth * 0.5625
    //     this.fixStyle = {
    //       height: windowWidth * 0.5625 + 'px',
    //       width: windowWidth + 'px',
    //       'margin-bottom': (windowHeight - videoHeight) / 2 + 'px',
    //       'margin-left': 'initial'
    //     }
    //   } else {
    //     videoHeight = windowHeight
    //     videoWidth = videoHeight / 0.5625
    //     this.fixStyle = {
    //       height: windowHeight + 'px',
    //       width: windowHeight / 0.5625 + 'px',
    //       'margin-left': (windowWidth - videoWidth) / 2 + 'px',
    //       'margin-bottom': 'initial'
    //     }
    //   }
    // }
    // window.onresize();
  },
  destroyed() {
    // window.removeEventListener('storage', this.afterQRScan)
  },
  methods: {
    handleLogin() {
      console.log(this.loginForm.password)
      this.axios.post('http://127.0.0.1:8099/login',
      Qs.stringify({
        No:this.loginForm.username,
        Password:this.loginForm.password
      }),{headers: {'Content-Type': 'application/x-www-form-urlencoded'}
      }).then(function (ret) {
            console.log(ret.data)
      }).catch(function (ret) {
            console.log('error')
      });
    },
    canplay() {
        this.vedioCanPlay = true
    }
  }
}

</script>