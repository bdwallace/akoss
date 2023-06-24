<template>
  <div class="panel">
    <panel-title :title="$route.meta.title"></panel-title>
    <div class="panel-body"
         v-loading="load_data"
         element-loading-text="拼命加载中">
      <el-row>
        <el-col :span="24">
          <el-form ref="form" :model="form" :rules="rules">

            <el-form-item label="任务名称:" prop="Name" label-width="120px">
              <el-input v-model="form.Name" placeholder="请输入任务名称"
                        style="width: 600px;"></el-input>
            </el-form-item>

            <el-form-item label="说明:" prop="Direction" label-width="120px">
              <el-input v-model="form.Direction" placeholder="请输入任务说明"
                        style="width: 600px;"></el-input>
            </el-form-item>

            <el-form-item label="任务周期:" prop="Crontab" label-width="120px">
              <el-input v-model="form.Crontab" placeholder="请输入任务周期，如：*/30 * * * * *"
                        style="width: 600px;"></el-input>
            </el-form-item>

            <el-form-item label="是/否自动执行:" prop="Auto" label-width="120px" type="number">
              <el-radio-group v-model="form.Auto">
                <el-radio :key="0" :label=0>否</el-radio>
                <el-radio :key="1" :label=1>是</el-radio>
              </el-radio-group>
            </el-form-item>                  

            <el-form-item label="任务脚本:" prop="Script" label-width="120px">
              <el-input v-model="form.Script" placeholder="请输入任务脚本"
                        type="textarea" autosize style="width: 600px;"></el-input>
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
  import {port_crontab} from 'common/port_uri'
  import {tools_verify} from 'common/tools'
  import store from 'store'

  export default {
    data() {
      return {

        form: {
          	Id: 0,
	          Name: null,
	          Direction: null,
	          Crontab: null,
	          Auto: 0, 
            Script: null,
            Project: {
              Id: store.state.user_info.user.ProjectId
            },
        },
        route_id: this.$route.params.id,
        load_data: false,
        on_submit_loading: false,
        rules: {
          Name: [{required: true, message: '任务名称不能为空', trigger: 'blur'}],
          Crontab: [{required: true, message: '任务周期不能为空', trigger: 'blur'}],
          Script: [{required: true, message: '任务脚本不能为空', trigger: 'blur'}],
          Auto: [{required: true, message: '是否自动执行不能为空', trigger: 'blur'}],
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
            this.$http.get(port_crontab.id, {
              params: {
                  crontabId: this.route_id
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
            // this.$refs.form.validate((valid) => {
            // if (
            //     !valid
            // )
            //     return false

            this.$http.post(port_crontab.save, this.form)
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
            // })
        }
        },


        components: {
          panelTitle
        }
    }
</script>
