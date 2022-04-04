<template>
  <div class="panel" style="padding:15px">
    <span> {{$route.meta.title}} </span>
    <!-- <panel-title :title="$route.meta.title"> -->
    <panel-title>
    </panel-title>
    <div class="panel-body"
         v-loading="load_data"
         element-loading-text="拼命加载中">
      <el-row>
        <el-col :span="24">
          <el-form ref="form" :model="form" :rules="rules">
            <el-form-item label="" label-width="120px">
              <el-radio-group style="float: left;margin-top: 15px" v-model="ProjectId" @change="get_table_data()">
              <el-radio v-for="item in project_list" :key="item.id" :label="item.id">{{item.name}}</el-radio>
              </el-radio-group>
            </el-form-item>

            <el-form-item label="安全组:" prop="GroupIds" label-width="120px">
              <el-select v-model="form.GroupIds[0]" placeholder="选择安全组"
                          filterable
                          style="width: 400px;">
                          <!-- @change="itemGroups_change(form.GroupIds[0])"> -->
                <el-option
                  v-for="item in itemGroups"
                  :key="item.Id"
                  :label="item.Name"
                  :value="item.Id">
                </el-option>
              </el-select>
            </el-form-item>

            <el-form-item label="访问端口:" prop="ToPort" label-width="120px">
              <el-select v-model="form.ToPort" placeholder="选择端口"
                          filterable
                          style="width: 100px;">
                <el-option
                  v-for="item in itemPort"
                  :key="item"
                  :label="item"
                  :value="item">
                </el-option>
              </el-select>
              <!-- <el-input v-model="form.ToPort" placeholder="80"
                        style="width: 100px;"></el-input> -->
            </el-form-item>

            <el-form-item label="来源IP:" prop="CidrIp" label-width="120px">
              <el-input v-model="form.CidrIp" placeholder="1.2.3.4/32"
                        style="width: 600px;"></el-input>
            </el-form-item>

            <el-form-item label="说明:" prop="Description" label-width="120px">
              <el-input v-model="form.Description" placeholder="只能输入全英文"
                        style="width: 600px;"></el-input>
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
  import {port_sgroup, port_project} from 'common/port_uri'
  import store from "store"

  export default {
    data() {
      return {
        project_list: [],
        // ProjectId: store.state.user_info.user.ProjectId,
        ProjectId: null,
        form: {
          GroupIds: [],
          ToPort: 8081,
          Project: null,
          CidrIp: null,
          Description: null
        },
        table_data: [],
        itemGroups: [],
        itemPort: [],
        mask: 32,
        port: this.$route.params.port,
        load_data: false,
        on_submit_loading: false,
        rules: {
          GroupIds: [{required: true, message: '安全组不能为空', trigger: 'blur'}],
          ToPort: [{required: true, message: '访问端口不能为空', trigger: 'blur'}],
          CidrIp: [{required: true, message: '来源IP不能为空', trigger: 'blur'}],
        }
      }
    },
    created() {
      this.get_project_list()
    },
    methods: {
      //刷新
      on_refresh(){
          this.get_table_data()
      },

      get_project_list() {
          this.load_data = true
          this.$http.get(port_project.list)
                  .then(({data: {data}}) => {
              this.project_list = data.project_list
                    console.log("project_list: ", this.project_list)
              this.load_data = false
          })
      },

      itemGroups_change(groupid) {
        let itemPort = []
        for(var i in this.table_data) {
          if ( this.table_data[i].GroupId == groupid) {
            for(var j in this.table_data[i].IpPermissions) {
              itemPort.push(this.table_data[i].IpPermissions[j].ToPort)
            }
            this.itemPort = itemPort
            // this.form.ToPort = this.table_data[i].IpPermissions[0].ToPort
            this.form.ToPort = 8081
          }
        }
      },

      //获取数据
      get_table_data(){
        this.load_data = true

        this.form = {
          GroupIds: [],
          ToPort: 8081,
          Level: 0,
          CidrIp: null,
          Description: null
        },
        this.itemGroups = []
        this.$http.get(port_sgroup.list, {
                    params: {
                      projectId: this.ProjectId,
                    }
                })
        .then(({data: {data}}) => {
          this.table_data = data.SecurityGroups
          var title = ""
          for(var i in this.table_data) {
              for(var j in this.table_data[i].IpPermissions) {
                  title = this.table_data[i].IpPermissions[j].ToPort + " 端口"
                  this.$set(this.table_data[i].IpPermissions[j], 'title', title)
              }
              this.itemGroups.push({Id: this.table_data[i].GroupId, Name: this.table_data[i].GroupName})
          }
          this.load_data = false
        })
        .catch(() => {
          this.$forceUpdate()
          this.load_data = false
        })
      },

      //提交
      on_submit_form() {
          this.on_submit_loading = true
          if ( this.form.CidrIp.split("/").length == 1) {
            this.form.CidrIp = this.form.CidrIp + "/32"
          }
          this.form.Project = this.ProjectId
          this.$http.post(port_sgroup.save, this.form)
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
      }
    },
    components: {
      panelTitle
    }
  }
</script>
