<template>
  <div class="login-bodya">
    <div class="loginWarpa">
      <div class="login-titlea">
        <div> 用户注册</div>
      </div>
      <div class="login-forma">
        <el-form ref="form" :model="form" :rules="rules" label-width="0">
          <el-form-item prop="register_username" class="login-itema">
            <label class="labela">账户名 ：</label>
            <el-input v-bind:readonly="isReadonly" v-model="form.register_username" placeholder="请输入账户名：" class="form-inputa"
                      :autofocus="true"></el-input>
          </el-form-item>

          <el-form-item prop="register_realname" class="login-itema">
            <label class="labela">花名.实名 ：</label>
            <el-input v-model="form.register_realname" placeholder="输入规范如：春哥.李宇春" class="form-inputa"
                      :autofocus="true"></el-input>
          </el-form-item>

          <el-form-item prop="register_email" class="login-itema">
            <label class="labela">邮箱 ：</label>
            <el-input v-model="form.register_email" placeholder="请输入联系邮箱：" class="form-inputa"
                      :autofocus="true"></el-input>
          </el-form-item>
          <el-form-item label="用户类型:" label-width="100px">
            <el-radio-group v-model="form.Role">
              <el-radio :label="1">管理员</el-radio>
              <el-radio :label="2">运维</el-radio>
              <el-radio :label="3">运营</el-radio>
            </el-radio-group>
          </el-form-item>
          <el-form-item class="login-item">
            <el-button size="large" class="form-submita" @click="submit_forma">确认注册</el-button>
          </el-form-item>
        </el-form>
      </div>
    </div>
  </div>
</template>
<script type="text/javascript">
  import {port_user, port_conf, port_code} from 'common/port_uri'
  import {mapActions} from 'vuex'
  import {SET_USER_INFO} from 'store/actions/type'

  export default {
    data() {
      return {
        projects: null,
        userId:this.$route.query.id|0,
        form: {
          Role: 3,
          pro_ids: ""
        },
        isReadonly:false,
        rules: {
          register_username: [{required: true, message: '请输入账户名！', trigger: 'blur'}],
          register_realname: [{required: true, message: '请输入花名.实名！', trigger: 'blur'}],
          register_email: [{required: true, message: '请输入邮箱！', trigger: 'blur'}]

        },
        //请求时的loading效果
        load_data: false
      }
    },
    created() {
      if(this.userId>0){
        this.get_data()
      }
      this.get_project_data()
    },
    methods: {
      ...mapActions({
        set_user_info: SET_USER_INFO
      }),
      //提交
      get_data() {

        this.load_data = true
        this.$http.get(port_user.users,{
          params: {
            id: this.userId
          }
        })
          .then(({data: {data}}) => {
            console.log(data.username)
            this.form.register_username=data.username
            this.form.register_realname=data.realname
            this.form.register_email=data.email
            this.form.Role=data.role|0
            this.isReadonly=true
            this.load_data = false
          }).catch(() => {
          this.load_data = false
        })
      },

      submit_forma() {
        this.$http.post(port_user.register+"?id="+this.userId, this.form)
          .then(({data: {msg}}) => {
            this.$message({
              message: msg,
              type: 'success'
            })
            setTimeout(() => {
                this.$router.push({path: '/'})
              },
              500
            )
          })
      }
    }
  }
</script>
<style lang="scss" type="text/css" rel="stylesheet/scss">
  .login-bodya {
    position: relative;
    left: 0;
    top: 0;
    width: auto;
    height: auto;
    margin: 0 auto;

  .loginWarpa {
    width: 500px;
    padding: 25px 15px;
    margin: 0 auto;
    background-color: #fff;
    border-radius: 5px;

  .login-titlea {
    margin-bottom: 25px;
    text-align: center;
  }

  .login-itema {

  .el-input__inner {
    margin: 0 !important;
  }

  }
  .form-inputa {

  input {
    margin-bottom: 15px;
    font-size: 12px;
    height: 40px;
    border: 1px solid #eaeaec;
    background: #eaeaec;
    border-radius: 5px;
    color: #555;
  }

  }
  .form-submita {
    width: 100%;
    color: #fff;
    border-color: #6bc5a4;
    background: #6bc5a4;

  &
  :active,

  &
  :hover {
    border-color: #6bc5a4;
    background: #6bc5a4;
  }

  }
  }
  }
</style>
