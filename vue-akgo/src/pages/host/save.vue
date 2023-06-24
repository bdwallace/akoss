<template>
  <div class="panel">
    <panel-title :title="$route.meta.title"></panel-title>
    <div class="panel-body"
         v-loading="load_data"
         element-loading-text="拼命加载中">
      <el-row>
        <el-form ref="form" :model="form" :rules="rules">
          <el-col :span="7">
            <el-form-item label="主机名:" prop="Name" label-width="130px">
              <el-input v-model="form.Name" placeholder="请输入主机名"
                        style="width: 300px;"></el-input>
            </el-form-item>

            <el-form-item label="内网IP:" prop="PrivateIp" label-width="130px">
              <el-input v-model="form.PrivateIp" placeholder="请输入IP"
                        style="width: 300px;"></el-input>
            </el-form-item>

            <!-- <el-form-item label="使用IP:" prop="PublicIp" label-width="130px">
              <el-input v-model="form.UseIp" placeholder="请输入公网IP"
                        style="width: 600px;"></el-input>
            </el-form-item> -->

            <el-form-item  label="是否使用公网地址：" prop="UsePublic" label-width="140px">
              <el-switch
                v-model="UseIp"
                on-text="是"
                off-text="否">
              </el-switch>
            </el-form-item>

            <el-form-item label="选择选择关闭时间:">
                <el-time-select
                  placeholder="起始时间"
                  v-model="form.StartTime"
                  clearable
                  style="width: 140px;"
                  :picker-options="{
                    start: '00:00',
                    step: '00:30',
                    end: '24:00'
                  }">
                </el-time-select>

                <el-time-select
                  placeholder="结束时间"
                  v-model="form.EndTime"
                  clearable
                  style="width: 140px;"
                  :picker-options="{
                    start: '00:00',
                    step: '00:30',
                    end: '24:00',
                    minTime: form.StartTime
                  }">
                </el-time-select>
            </el-form-item>

          </el-col>

          <el-col :span="10">
            <el-form-item label="查询aws实例:" prop="PublicIp" label-width="130px">
              <!-- <el-select v-model="awsHost"
                          filterable
                          value-key
                          clearable
                          :placeholder="`${form.Region} . ${form.InstanceId}`"
                          style="width: 600px;"
                          @change="check_itemAwsHost(awsHost)">
                <el-option
                  v-for="item in itemAwsHost"
                  :key="item"
                  :label="`${item.Region} . ${item.EC2InstanceInfo.InstanceID} . ${item.EC2InstanceInfo.InstanceName} . ${item.EC2InstanceInfo.PrivateIP} . ${item.EC2InstanceInfo.PublicIP}`"
                  :value="item">
                </el-option>
              </el-select> -->

              <el-button @click.stop="get_itemAwsHost" size="mini" :loading="on_submit_loading">
                <i class="fa fa-refresh"></i>
              </el-button>
            </el-form-item>

            <el-form-item label="公网IP:" prop="PublicIp" label-width="130px">
              <el-input v-model="form.PublicIp" placeholder="请输入公网IP"
                        style="width: 300px;">
              </el-input>
              <el-button
                v-if="itemAwsHost[0] && form.PublicIp != itemAwsHost[0].EC2InstanceInfo.PublicIP"
                :type="awsColor"
                @click="form.PublicIp = itemAwsHost[0].EC2InstanceInfo.PublicIP"
                :loading="on_submit_loading" >
                  更新为：{{itemAwsHost[0].EC2InstanceInfo.PublicIP}}
              </el-button>
            </el-form-item>

            <el-form-item label="aws实例ID:" prop="InstanceId" label-width="130px">
              <el-input v-model="form.InstanceId" placeholder="请 查询aws实例"
                        style="width: 300px;">
              </el-input>
              <el-button
                v-if="itemAwsHost[0] && form.InstanceId != itemAwsHost[0].EC2InstanceInfo.InstanceID"
                :type="awsColor"
                @click="form.InstanceId = itemAwsHost[0].EC2InstanceInfo.InstanceID"
                :loading="on_submit_loading" >
                  更新为：{{itemAwsHost[0].EC2InstanceInfo.InstanceID}}
              </el-button>
            </el-form-item>

            <el-form-item label="aws实例类型:" prop="InstanceType" label-width="130px">
              <el-input v-model="form.InstanceType" placeholder="请 查询aws实例类型"
                        style="width: 300px;">
              </el-input>
              <el-button
                v-if="itemAwsHost[0] && form.InstanceType != itemAwsHost[0].EC2InstanceInfo.InstanceType"
                :type="awsColor"
                @click="form.InstanceType = itemAwsHost[0].EC2InstanceInfo.InstanceType"
                :loading="on_submit_loading" >
                更新为：{{itemAwsHost[0].EC2InstanceInfo.InstanceType}}
              </el-button>
            </el-form-item>


            <el-form-item label="aws区域:" prop="Region" label-width="130px">
              <el-input v-model="form.Region" placeholder="请 查询aws实例"
                        style="width: 300px;">
              </el-input>
              <el-button
                v-if="itemAwsHost[0] && form.Region != itemAwsHost[0].Region"
                :type="awsColor"
                @click="form.Region = itemAwsHost[0].Region"
                :loading="on_submit_loading" >
                  更新为：{{itemAwsHost[0].Region}}
              </el-button>
            </el-form-item>


          </el-col>

          <el-col>
            <el-form-item>
              <el-button type="primary" @click="on_submit_form" :loading="on_submit_loading">立即提交
              </el-button>
              <el-button @click="$router.back()">取消</el-button>
            </el-form-item>
          </el-col>

        </el-form>
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
        UseIp: false,

        form: {
          Name: null,
          PrivateIp: null,
          PublicIp: null,
          UseIp: null,
          InstanceId: null,
          Region: null,
          StartTime: null,
          EndTime: null,
          Project: {
            Id: store.state.user_info.user.ProjectId
          },
        },
        itemAwsHost: [],
        awsHost: {},
        // awsHost: {
        //   Region: this.form.Region,
        //   EC2InstanceInfo: {InstanceID: this.form.InstanceId}
        // },
        route_id: this.$route.params.id,
        load_data: false,
        on_submit_loading: false,
        awsColor: "",
        intervalid1: null,
        rules: {
          Name: [{required: true, message: '参数名称不能为空', trigger: 'blur'}],
          PrivateIp: [{required: true, message: 'IP不能为空', trigger: 'blur'}],
        }
      }
    },

    beforeDestroy () {
        clearInterval(this.intervalid1)
    },

    created() {
      if (this.route_id) {
        this.get_form_data()
      }
    },

    methods: {
      set_interval() {
          var _this = this;
          this.intervalid1 = setInterval(function () {
              if (_this.awsColor == "") {
                _this.awsColor = "warning"
              } else {
                _this.awsColor = ""
              }
          }, 1000)
      },

      clear_interval() {
          clearInterval(this.intervalid1)
      },

      //获取数据
      get_form_data() {
        this.load_data = true
        this.$http.get(port_host.id, {
          params: {
            id: this.route_id
          }
        })
          .then(({data: {data}}) => {
            this.form = data
            if (this.form.UseIp === this.form.PublicIp) {
              this.UseIp = true
            }
            // this.get_aws_host()
            this.awsHost = {
              Region: this.form.Region,
              EC2InstanceInfo: {
                InstanceID: this.form.InstanceId,
                InstanceName: "",
                PrivateIP: "",
                PublicIP: ""
                }
            },
            this.load_data = false
          })
          .catch(() => {
            this.load_data = false
          })
      },

      //获取数据
      check_itemAwsHost(aws) {
        this.on_submit_loading = true
        this.form.Region = aws.Region
        this.form.InstanceId = aws.EC2InstanceInfo.InstanceID
        this.on_submit_loading = false
      },


      //获取数据
      get_itemAwsHost() {
        this.on_submit_loading = true
        if (this.form.PrivateIp == "") {
              this.$message({
                message: '请输入要查询的内网IP',
                type: "warning"
              })
              this.on_submit_loading = false
              return
        }

        this.$http.post(port_host.hostrsyncaws, this.form)
          .then(({data: {data}}) => {
            this.itemAwsHost= data
            this.set_interval()
            // if (this.form.InstanceId == "" || this.form.InstanceId == null) {
            //   this.check_itemAwsHost(this.itemAwsHost[0])
            // }
            this.on_submit_loading = false
          })
          .catch(() => {
            this.on_submit_loading = false
          })
      },

      //提交
      on_submit_form() {
        this.$refs.form.validate((valid) => {
          if (
            !valid
          )
            return false

          this.on_submit_loading = true
          if (this.UseIp) {
            this.form.UseIp = this.form.PublicIp
          } else {
            this.form.UseIp = this.form.PrivateIp
          }
          this.$http.post(port_host.save, this.form)
            .then(({data: {msg}}) => {
              this.$message({
                message: msg,
                type: 'success'
              })
              this.clear_interval()
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
