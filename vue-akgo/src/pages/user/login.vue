<template>
    <div class="login-body">
        <div class="loginWarp"
             v-loading="load_data"
             element-loading-text="正在登陆中..."
             @keyup.enter="submit_form">
            <div class="login-title">
                <div> 自动发布系统</div>
            </div>
            <div class="login-form">
                <el-form ref="form" :model="form" :rules="rules" label-width="0">
                    <el-form-item prop="user_name" class="login-item">
                        <el-input v-model="form.user_name" placeholder="请输入账户名：" class="form-input"
                                  :autofocus="true"></el-input>
                    </el-form-item>
                    <el-form-item prop="user_password" class="login-item">
                        <el-input type="password" v-model="form.user_password" placeholder="请输入账户密码："
                                  class="form-input"></el-input>
                    </el-form-item>

<!--
                    <el-radio-group v-model="project">
                    <el-radio v-for="item in project_list" :key="item" :label="item">{{item.Name}}</el-radio>
                    </el-radio-group>
-->

                  <el-select v-model="project" clearable placeholder="请选择">
                    <el-option
                      v-for="item in project_name_list"
                      :key="item"
                      :label="item"
                      :value="item">
                    </el-option>
                  </el-select>

<!--                  <el-select v-model="value" clearable placeholder="请选择">
                    <el-option
                      v-for="item in options"
                      :key="item.value"
                      :label="item.label"
                      :value="item.value">
                    </el-option>
                  </el-select>
                  -->

                    <el-form-item></el-form-item>

                    <el-form-item class="login-item">
                        <el-button size="large"  class="form-submit" @click="submit_form">登录</el-button>
                    </el-form-item>
                </el-form>
            </div>
        </div>
    </div>
</template>
<script type="text/javascript">
    import {port_user, port_project} from 'common/port_uri'
    import {mapActions} from 'vuex'
    import {SET_USER_INFO} from 'store/actions/type'

    export default{
        data(){
            return {
              project_resp: null,
              project:"",
              project_list:[],
              project_name_list:[],
              form: {
                    user_name: "",
                    user_password: "",
                    project_name: "",
                },
                rules: {
                    user_name: [{required: true, message: '请输入账户名！', trigger: 'blur'}],
                    user_password: [{required: true, message: '请输入账户密码！', trigger: 'blur'}]
                },
                //请求时的loading效果
                load_data: false
            }
        },

        created() {
          this.get_project_list()
        },

        methods: {
                ...mapActions({
                    set_user_info: SET_USER_INFO
                }),

            get_project_list() {
                this.load_data = true
                this.$http.get(port_project.list)
                        .then(({data: {data}}) => {
                    this.project_resp = data
                    this.project_list = data.project_list
                    this.project_name_list = data.project_name_list
                    this.load_data = false
                })
            },

        //提交
        submit_form() {
            this.$refs.form.validate((valid) => {
                if (valid) {
                    this.load_data = true

                  if (this.project_list.length === 0) {
                      this.$message({
                        message: "获取系统环境失败",
                        type: 'warning'
                      });
                    } else {
                      this.form.project_name = this.project
                      if (this.project === "") {
                            this.$message({
                            message: "请选择环境",
                            type: 'warning'
                            });
                            this.load_data = false
                            return
                        }
                    }
                    //登录提交
                    this.$http.post(port_user.login, this.form)
                            .then(({data: {data, msg}}) => {
                        let isNull = data !== null
                        this.set_user_info({
                            user: data,
                            login: true
                        })
                        this.$message({
                            message: msg,
                            type: 'success'
                        })
                        setTimeout(() => {
                            this.$router.push({path: '/'})
                        }, 500)
                    })
                    .catch(() => {
                        this.load_data = false
                    })
                } else {
                    return false
                }
            }
        )
    },
    to_tasklist(){
          this.$router.push({path: '/task/searchlist'})
    }
    }
    }
</script>
<style lang="scss" type="text/css" rel="stylesheet/scss">
    .login-body {
        position: absolute;
        left: 0;
        top: 0;
        width: 100%;
        height: 100%;
        background-image: url(./images/login_bg.jpg);
        background-repeat: no-repeat;
        background-position: center;
        background-size: cover;

        .loginWarp {
            width: 300px;
            padding: 25px 15px;
            margin: 100px auto;
            background-color: #fff;
            border-radius: 5px;

            .login-title {
                margin-bottom: 25px;
                text-align: center;
            }

            .login-item {

                .el-input__inner {
                    margin: 0 !important;
                }

            }
            .form-input {

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
            .form-submit {
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
