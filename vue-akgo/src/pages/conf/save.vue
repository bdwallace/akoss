<template>
  <el-form ref="form" :model="form" :rules="rules">
    <div class="panel">
      <panel-title :title="$route.meta.title">

      </panel-title>
      <div class="panel-body"
          v-loading="load_data"
          element-loading-text="拼命加载中">
        <el-row>
          <el-col :span="24">
            <el-form-item label="公共参数名称:" prop="Name" label-width="120px">
              <el-input v-model="form.Name" placeholder="请输入公共参数名称"
                        style="width: 600px;"></el-input>
            </el-form-item>

            <el-form-item label="公共参数：" label-width="120px">
            <table class="table">
              <tr>
                <th style="text-align:right;width:300px;padding-right:10px">key值</th>
                <th style="text-align:left;width:500px;padding-left:10px">value值</th>
                <!-- <th style="padding-left: 20px">编辑</th> -->
                <th style="padding-left: 20px">删除</th>
              </tr>

              <tr v-for="(item, index) in jsons" :key="item" style="margin-bottom:50px">
                <!-- <td style="text-align:right;width:300px;padding-right:10px">{{key}}</td>
                <td style="text-align:left;width:500px;padding-left:10px">{{jsons[key]}}</td> -->
                <!-- <td style="text-align:right;width:300px;padding-right:10px">
                    <el-input type="textarea" autosize placeholder="输入内容" v-model="key" style="direction: rtl;" v-if="inpub[key]"></el-input>
                    <el-input type="textarea" autosize placeholder="输入内容" v-model="key" style="direction: rtl;" v-else disabled></el-input>
                </td>

                <td style="text-align:left;width:500px;padding-left:10px">
                    <el-input type="textarea" autosize placeholder="输入内容" v-model="jsons[key]" v-if="inpub[key]"></el-input>
                    <el-input type="textarea" autosize placeholder="输入内容" v-model="jsons[key]" v-else disabled></el-input>
                </td> -->
                <td style="text-align:right;width:300px;padding-right:10px">
                    <el-input type="textarea" autosize placeholder="输入内容" v-model="item.Key" style="direction: rtl;"></el-input>
                </td>

                <td style="text-align:left;width:500px;padding-left:10px">
                    <el-input type="textarea" autosize placeholder="输入内容" v-model="item.Value"></el-input>
                </td>

                <!-- <td style="padding-left: 20px">
                  <el-button type="warning" size="mini" @click="etc_value(index, key)">
                      <i class="el-icon-edit"></i>
                  </el-button>
                </td> -->

                <td style="padding-left: 20px">
                  <el-button type="danger" size="mini" @click="del_value(index)">
                      <i class="el-icon-delete"></i>
                  </el-button>
                </td>
              </tr>

              <tr><td> &nbsp; </td></tr>

              <tr>
                <!-- <td> &nbsp; </td>
                <td> &nbsp; </td> -->
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

              <tr><td> &nbsp; </td></tr>
            </table>
            </el-form-item>

            <el-form-item label="关联服务：" prop="Services" label-width="120px">
                <!-- :filter-method="filterMethod" -->
              <el-transfer
                filterable
                filter-placeholder="请输入关键字"
                :titles="titles"
                v-model="keyService"
                :data="itemService">
              </el-transfer>
            </el-form-item>

            <el-form-item>
              <el-button type="primary" @click="on_submit_form" :loading="on_submit_loading">立即提交
              </el-button>
              <el-button @click="$router.back()">取消</el-button>
            </el-form-item>

          </el-col>
        </el-row>
      </div>
    </div>
  </el-form>
</template>
<script type="text/javascript">
  import {panelTitle} from 'components'
  import {port_conf, port_service} from 'common/port_uri'
  import {tools_verify} from 'common/tools'
  import store from 'store'

  export default {
    data() {
      return {
        form: {
          Name: null,
          Value: null,
          Services: [],
          Project: {
            Id: store.state.user_info.user.ProjectId
          } 
        },
        jsons: [],
        input: {
          Key: "",
          Value: ""
        },
        itemService: [],
        keyService: [],
        titles: ["未关联", "已关联"],
        route_id: this.$route.params.id,
        load_data: false,
        on_submit_loading: false,
        rules: {
          Name: [{required: true, message: '参数名称不能为空', trigger: 'blur'}],
          Value: [{required: true, message: '参数值不能为空', trigger: 'blur'}],
        }
      }
    },
    created() {
      if (this.route_id) {
        this.get_form_data()
      }
      this.get_itemService()
    },
    methods: {
      // 穿梭框过滤选项
      // filterMethod(query, item) {
      //   return item..indexOf(query) > -1;
      // },
        
      // 获取services数据
      get_itemService() {
        this.load_data = true
        this.$http.get(port_service.list, {
              params: {
                project_id: this.form.Project.Id
              }
            })
          .then(({data: {data}}) => {
          for(let i in data) {
            this.itemService.push({
              key: data[i].Id,
              label: data[i].Name
            });
          }

          this.load_data = false
        })
        .catch(() => {
          this.load_data = false
        })
      },
        
      // 获取数据
      get_form_data() {
        this.load_data = true
        this.$http.get(port_conf.id, {
          params: {
            id: this.route_id
          }
        })
          .then(({data: {data}}) => {
          this.form = data
          for(let i in this.form.Services) {
            this.keyService.push(data.Services[i].Id)
          }
          this.jsons = JSON.parse(this.form.Value)

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
        this.form.Value = JSON.stringify(this.jsons)
        this.form.Services = []
        for(let i in this.keyService) {
          this.form.Services.push({Id: this.keyService[i]})
        }

        this.$http.post(port_conf.save, this.form)
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
