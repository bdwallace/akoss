<template>
  <div class="panel">
    <panel-title :title="$route.meta.title"></panel-title>
    <div class="panel-body"
         v-loading="load_data"
         element-loading-text="拼命加载中">
      <el-row>
        <el-col :span="8">
          <el-form ref="form" :model="form" :rules="rules">

            <el-form-item label="名称:" prop="Name" label-width="150px">
              <el-input v-model="form.Name" placeholder="名称"
                        style="width: 300px;"></el-input>
            </el-form-item>

            <el-form-item label="系统别名:" prop="Alias" label-width="150px">
              <el-input v-model="form.Alias" placeholder="系统别名"
                        style="width: 300px;"></el-input>
            </el-form-item>

<!--            ----------                  -->
            <el-form-item label="Nacos-1:" prop="Nacos1" label-width="150px">
              <el-input v-model="form.Nacos1" placeholder="Nacos-1"
                        style="width: 300px;"></el-input>
            </el-form-item>

            <el-form-item label="UserNameNacos1:" prop="UserNameNacos1" label-width="150px">
              <el-input v-model="form.Nacos1UserName" placeholder="Nacos-1 UserName"
                        style="width: 300px;"></el-input>
            </el-form-item>

            <el-form-item label="PasswordNacos1:" prop="PasswordNacos1" label-width="150px">
              <el-input v-model="form.Nacos1Pwd" placeholder="Nacos-1 Password"
                        style="width: 300px;"></el-input>
            </el-form-item>

<!--            ----------                  -->
            <el-form-item label="Nacos-2:" prop="Nacos2" label-width="150px">
              <el-input v-model="form.Nacos2" placeholder="Nacos-2"
                        style="width: 300px;"></el-input>
            </el-form-item>

            <el-form-item label="UserNameNacos2:" prop="UserNameNacos2" label-width="150px">
              <el-input v-model="form.Nacos2UserName" placeholder="Nacos-2 UserName"
                        style="width: 300px;"></el-input>
            </el-form-item>

            <el-form-item label="PasswordNacos2:" prop="PasswordNacos2" label-width="150px">
              <el-input v-model="form.Nacos2Pwd" placeholder="Nacos2 Password"
                        style="width: 300px;"></el-input>
            </el-form-item>

<!--            ----------                  -->
            <el-form-item label="Nacos-3:" prop="Nacos3" label-width="150px">
              <el-input v-model="form.Nacos3" placeholder="Nacos-3"
                        style="width: 300px;"></el-input>
            </el-form-item>

            <el-form-item label="UserNameNacos3:" prop="UserNameNacos3" label-width="150px">
              <el-input v-model="form.Nacos3UserName" placeholder="Nacos-3 UserName"
                        style="width: 300px;"></el-input>
            </el-form-item>

            <el-form-item label="PasswordNacos3:" prop="PasswordNacos3" label-width="150px">
              <el-input v-model="form.Nacos3Pwd" placeholder="Nacos-3 Password"
                        style="width: 300px;"></el-input>
            </el-form-item>

            <!--            ----------                  -->
            <el-form-item label="Prometheus:" prop="Prometheus" label-width="150px">
              <el-input v-model="form.Prometheus" placeholder="Prometheus"
                        style="width: 300px;"></el-input>
            </el-form-item>

            <el-form-item label="Aws默认区域:" prop="AwsRegion" label-width="150px">
              <el-input v-model="form.AwsRegion" placeholder="AwsRegion"
                        style="width: 300px;"></el-input>
            </el-form-item>

            <div v-show="userId == 1" >
              <el-form-item label="AwsKeyId:" prop="AwsKeyId" label-width="150px">
                <el-input v-model="form.AwsKeyId" placeholder="AwsKeyId"
                          style="width: 300px;"></el-input>
              </el-form-item>

              <el-form-item label="AwsKeySecret:" prop="AwsKeySecret" label-width="150px">
                <el-input v-model="form.AwsKeySecret" placeholder="AwsKeySecret"
                          style="width: 300px;"></el-input>
              </el-form-item>

              <!-- <el-form-item label="Aws默认安全组:" prop="AwsSgroup" label-width="120px">
                <el-input v-model="form.AwsSgroup" placeholder="请类型名称"
                          style="width: 300px;"></el-input>
              </el-form-item> -->

              <el-form-item>
                <el-button type="primary" @click="on_submit_form" :loading="on_submit_loading">立即提交
                </el-button>
                <el-button @click="$router.back()">取消</el-button>
              </el-form-item>
            </div>

          </el-form>
        </el-col>

        <!-- <el-col :span="8">
          <el-form>

            <el-form-item label="导入host数据，初始化项目主机关联:">
                <el-upload
                    class="upload-demo"
                    action="http://127.0.0.1:8192"
                    :http-request = "handleGetFile">
                    <el-button size="small" type="primary">点击上传</el-button>
                </el-upload>
            </el-form-item>

          </el-form>
        </el-col> -->
      </el-row>
    </div>
  </div>
</template>
<script type="text/javascript">
  import {panelTitle} from 'components'
  import {port_project} from 'common/port_uri'
  import {tools_verify} from 'common/tools'
  import store from "store"

  export default {
    data() {
      return {
        userId: store.state.user_info.user.Id,
        form: {Name: null},
        route_id: this.$route.params.id,
        load_data: false,
        on_submit_loading: false,
        rules: {
          Name: [{required: true, message: '名称不能为空', trigger: 'blur'}],
          // Nacos: [{required: true, message: 'Nacos不能为空', trigger: 'blur'}],
          Alias: [{required: true, message: '系统别名不能为空', trigger: 'blur'}],
        }
      }
    },
    created() {
      if (this.route_id) {
        this.get_form_data()
      }
    },
    methods: {
        //获取数据
        get_form_data() {
            this.load_data = true
            this.$http.get(port_project.id, {
              params: {
                  id: this.route_id
              }
            })
            .then(({data: {data}}) => {
                this.form = data
                this.load_data = false
            })
            .catch(() => {
                this.load_data = false
            })
        },


        //提交
        on_submit_form() {
          this.$refs.form.validate((valid) => {
            if ( ! valid ) {
              return false
            }

            this.on_submit_loading = true
            this.$http.post(port_project.save, this.form)
            .then(({data: {msg}}) => {
              this.on_submit_loading = false
              this.$message({
                  message: msg,
                  type: 'success'
              })
              setTimeout(() => {
                  this.$router.back()
                  },
                  500
              )
            })
            .catch(() => {
              this.on_submit_loading = false
            })


          })

        }
    },

    components: {
      panelTitle
    }
  }
</script>
