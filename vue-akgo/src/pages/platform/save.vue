<template>
  <div class="panel">
    <panel-title :title="$route.meta.title"></panel-title>
    <div class="panel-body"
         v-loading="load_data"
         element-loading-text="拼命加载中">
      <el-row>
        <el-col :span="24">
          <el-form ref="form" :model="form" :rules="rules">

            <el-form-item label="名称:" prop="Name" label-width="120px">
              <el-input v-model="form.Name" placeholder="请类型名称"
                        style="width: 300px;"></el-input>
            </el-form-item>

            <el-form-item label="平台Id:" prop="ChannelId" label-width="120px">
              <el-input v-model="form.ChannelId" placeholder="请输入平台Id"
                        style="width: 300px;"></el-input>
            </el-form-item>

            <el-form-item label="平台参数：" label-width="120px">
            <table class="table">
              <tr>
                <th style="text-align:right;width:300px;padding-right:10px">key值</th>
                <th style="text-align:left;width:500px;padding-left:10px">value值</th>
                <th style="padding-left: 20px">删除</th>
              </tr>

              <tr v-for="(item, index) in jsons" :key="item" style="margin-bottom:50px">
                <td style="text-align:right;width:300px;padding-right:10px">
                    <el-input type="textarea" autosize placeholder="输入内容" v-model="item.Key" style="direction: rtl;"></el-input>
                </td>

                <td style="text-align:left;width:500px;padding-left:10px">
                    <el-input type="textarea" autosize placeholder="输入内容" v-model="item.Value"></el-input>
                </td>

                <td style="padding-left: 20px">
                  <el-button type="danger" size="mini" @click="del_value(index)">
                      <i class="el-icon-delete"></i>
                  </el-button>
                </td>
              </tr>

              <tr><td> &nbsp; </td></tr>
              <tr><td> &nbsp; </td></tr>

              <tr>
                <td style="text-align:right;width:300px;padding-right:10px">
                    <el-input type="textarea" autosize placeholder="输入内容" v-model="input.Key"></el-input>
                </td>
                <td style="text-align:left;width:500px;padding-left:10px">
                    <el-input type="textarea" autosize placeholder="输入内容" v-model="input.Value"></el-input>
                </td>
                <td style="padding-left: 20px;">
                  <el-button type="primary" icon="plus" size="mini"   @click="add_value()"></el-button>
                </td>
              </tr>
            </table>
            </el-form-item>

            <el-form-item>
              <el-button type="primary" @click="on_submit_form" :loading="on_submit_loading">立即提交
              </el-button>
              <el-button @click="$router.back()">取消</el-button>
            </el-form-item>

          </el-form>
        </el-col>

      </el-row>
    </div>
  </div>
</template>
<script type="text/javascript">
  import {panelTitle} from 'components'
  import {port_platform} from 'common/port_uri'
  import {tools_verify} from 'common/tools'
  import store from 'store'

  export default {
    data() {
      return {
        form: {
          Name: null,
          ChannelId: null,
          Value: null, 
          Project: {
            Id: store.state.user_info.user.ProjectId
          } 
        },
        jsons: [{
            Key: "CHANNEL_ID",
            Value: ""
          }, {
            Key: "CHANNEL_NAME",
            Value: ""
          }, {
            Key: "ENCRYPTED_KEY",
            Value: ""
          }, {
            Key: "KEYWORDS",
            Value: ""
          }, {
            Key: "META_DESC",
            Value: ""
          }, {
            Key: "PLATFORM",
            Value: ""
          }, {
            Key: "SKIN_NAME",
            Value: ""
          // }, {
          //   Key: "TITLE",
          //   Value: ""
          }],
        input: {
          Key: "",
          Value: ""
        },
        route_id: this.$route.params.id,
        load_data: false,
        on_submit_loading: false,
        rules: {
          Name: [{required: true, message: '名称不能为空', trigger: 'blur'}],
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
            this.$http.get(port_platform.id, {
              params: {
                  id: this.route_id
              }
            })
            .then(({data: {data}}) => {
                this.form = data
                if (this.form.Value == "") {
                  this.jsons = []
                } else {
                  this.jsons = JSON.parse(this.form.Value)
                }
                this.load_data = false
            })
            .catch(() => {
                this.load_data = false
            })
        },

        del_value(index) {
            this.load_data = true
            delete this.jsons.splice(index, 1)
            this.load_data = false
        },

        add_value() {
          if(this.input.Key == ''){
            this.$message({
              message: "请输入docker参数名",
              type: 'warning'
            })
          }else{
            if(JSON.stringify(this.jsons) !== '{}') {
              for(var i in this.josns) {
                if( this.jsons[i].Key == this.input.Key) {
                  this.$message({
                    message: "输入了相同的docker参数名",
                    type: 'warning'
                  })

                  this.input.Key=''
                  return
                }
              }
            }

              this.jsons.push({Key:this.input.Key, Value:this.input.Value})
              this.input.Key=''
              this.input.Value=''
              //可以强制重新渲染页面
              // this.$forceUpdate();
          }
        },

        //提交
        on_submit_form() {
          this.$refs.form.validate((valid) => {
            if ( ! valid ) {
              return false
            }

            if ( this.input.Key != "" ) {
              this.$message({
                message: "请保存正在编辑的参数!",
                type: 'warning'
              })
              return false
            }
            this.on_submit_loading = true
            this.form.value = JSON.stringify(this.jsons)
            this.$http.post(port_platform.save, this.form)
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
