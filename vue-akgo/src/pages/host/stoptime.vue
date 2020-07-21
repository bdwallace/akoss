<template>
  <div class="panel">
    <panel-title :title="$route.meta.title"></panel-title>
    <div class="panel-body"
         v-loading="load_data"
         element-loading-text="拼命加载中">
      <el-row>
        <el-col :span="24">
          <div class="panel-body" style="clear: both;">
            <el-table
                    ref="multipleTable"
                    :data="table_data"
                    v-loading="load_data"
                    element-loading-text="拼命加载中"
                    border
                    style="width: 100%;"
                    @selection-change="handleSelectionChange">

                <el-table-column
                type="selection"
                width="50">
                </el-table-column>

                <el-table-column
                        prop="Id"
                        label="id"
                        width="80">
                </el-table-column>

                <el-table-column
                        prop="Name"
                        label="主机名称"
                        width="250">
                </el-table-column>

                <el-table-column
                    label="定时关机"
                    width="115">
                    <template scope="props">
                        {{props.row.StartTime}}-{{props.row.EndTime}}
                    </template>
                </el-table-column>

                <el-table-column
                        prop="Ip"
                        label="内网ip"
                        width="140">
                </el-table-column>

                <el-table-column
                        prop="IpPub"
                        label="外网ip"
                        width="140">
                </el-table-column>

                <el-table-column
                        prop="InstanceId"
                        label="实例Id"
                        width="200">
                </el-table-column>

                <el-table-column
                        prop="Region"
                        label="所在地区"
                        width="140">
                </el-table-column>

                <el-table-column
                        prop="project_name"
                        label="关联项目">
                </el-table-column>
            </el-table>

          <el-form style="margin-top: 30px">
            <el-form-item label="选择选择关闭时间:">
                <el-time-select
                  placeholder="起始时间"
                  v-model="start_time"
                  clearable
                  :picker-options="{
                    start: '00:00',
                    step: '00:30',
                    end: '24:00'
                  }">
                </el-time-select>

                <el-time-select
                  placeholder="结束时间"
                  v-model="end_time"
                  clearable
                  :picker-options="{
                    start: '00:00',
                    step: '00:30',
                    end: '24:00',
                    minTime: start_time
                  }">
                </el-time-select> 
            </el-form-item>

            <el-form-item>
              <el-button type="primary" @click="on_submit_form" :loading="on_submit_loading">立即提交
              </el-button>
              <el-button @click="$router.back()">取消</el-button>
            </el-form-item>
          </el-form>

          </div>

        </el-col>
      </el-row>
    </div>
  </div>
</template>
<script type="text/javascript">
  import {panelTitle} from 'components'
  import {port_host} from 'common/port_uri'
  import {tools_verify} from 'common/tools'
  import store from 'store'

  export default {
    data() {
      return {
        classData: [],
        // level_data: [],
        StartTime: "",
        EndTime: "",
        ProjectId: store.state.user_info.user.ProjectId,

        table_data: [],
        form: [],
        start_time: "",
        end_time: "",
        itemAwsHost: [],
        awsHost: {},
        // awsHost: {
        //   Region: this.form.Region,
        //   EC2InstanceInfo: {InstanceID: this.form.InstanceId}
        // },
        load_data: false,
        on_submit_loading: false,
      }
    },
    created() {

      this.get_tables_data()
    },
    methods: {
      //获取数据
      get_tables_data() {
        this.load_data = true
        this.$http.get(port_host.list, {
          params: {
            project_id: this.ProjectId,
          }
        })
          .then(({data: {data}}) => {
            this.table_data = data
            console.log("---------", this.table_data)
            this.load_data = false
          })
          .catch(() => {
            this.load_data = false
          })
      },

      handleSelectionChange(val) {
          this.form = val;
      }, 
    
      //提交
      on_submit_form() {
            console.log("------1---")
        this.on_submit_loading = true
        for(var i in this.form) {
          this.form[i].StartTime = this.start_time
          this.form[i].EndTime = this.end_time
          delete this.form[i].project_name
        }

        this.$http.post(port_host.hostsavestoptime, this.form)
        .then(({data: {msg}}) => {
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
