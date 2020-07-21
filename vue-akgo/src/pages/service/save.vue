<template>
  <el-form ref="form" :model="form" :rules="rules">
    <div class="panel"> 
      <panel-title :title="$route.meta.title">
      </panel-title>
                        
    </div>
    <div class="panel-body" v-loading="load_data" element-loading-text="拼命加载中">
      <el-row>
        <el-col :span="11">
            <div class="panel" style="margin-top:5px;margin-bottom:15px;padding:15px;weight:500px"> 
              <el-form-item label="服务名称:" prop="Name" label-width="100px">
                <el-input v-model="form.Name" placeholder="请输入项目名称" style="width: 400px">
                          </el-input>
              </el-form-item>

              <el-form-item label="仓库地址:" prop="ImagePath" label-width="100px">
                <el-input v-model="form.ImagePath" style="width: 400px;"></el-input>
              </el-form-item>

              <el-form-item label="端口:" prop="Port" label-width="100px">
                <el-tooltip class="item" effect="dark"
                            content="192.168.57.0:8080:80,8443:443" placement="top">
                  <el-input v-model="form.Port" placeholder="请输入端口"
                            style="width: 400px;"></el-input>
                </el-tooltip>
              </el-form-item>

              <el-form-item label="容器名:" prop="DockerName" label-width="100px">
                <el-tooltip class="item" effect="dark"
                            content="指定容器名，为兼容OSS容器名而设置，默认容器名==项目名" placement="top">
                  <el-input v-model="form.DockerName" placeholder="项目名_系统别名"
                            style="width: 400px;"></el-input>
                </el-tooltip>
              </el-form-item>

              <el-form-item label="日志关键字:" prop="LogKeyword" label-width="100px">
                <el-tooltip class="item" effect="dark"
                            content="如果有关键字，则在第一台目标机docker启动时一直到跟踪检测到有这个关键字为止" placement="top">
                  <el-input v-model="form.LogKeyword" placeholder="Started .* seconds"
                            style="width: 400px;"></el-input>
                </el-tooltip>
              </el-form-item>

              <!-- <el-form-item label="强制关闭时间:" prop="KillTime" label-width="100px">
                <el-tooltip class="item" effect="dark"
                            content="指定容器的强制关闭时间，docker默认是10秒" placement="top">
                  <el-input v-model="form.KillTime" placeholder="默认不指定，10秒"
                            style="width: 400px;"></el-input>
                </el-tooltip>
              </el-form-item> -->

              <el-form-item label="健康监测:" prop="Health" label-width="100px">
                <el-tooltip class="item" effect="dark"
                            content="探测应用是否正常的URL" placement="top">
                  <el-input v-model="form.Health" placeholder="${service_name}/actuator/health"
                            style="width: 400px;"></el-input>
                </el-tooltip>
              </el-form-item>
            </div> 

            <div class="panel" style="margin-top:5px;margin-bottom:15px;padding:15px;weight:500px"> 
                  <el-form-item label="上传地址:" prop="ReleaseTo" label-width="100px">
                    <el-input v-model="form.ReleaseTo" style="width: 400px;" placeholder="发布目录"></el-input>
                  </el-form-item>
            </div>

            <div class="panel" style="margin-top:5px;margin-bottom:15px;padding:15px;weight:500px"> 
              <el-form-item label="机器列表:" prop="Hosts" label-width="100px">
                <el-form-item v-for="(item, hostindex) in form.Hosts" :key="hostindex" style="margin-top: 5px">
                  <el-button @click.stop="get_itemHost()" size="mini">
                    <i class="fa fa-refresh"></i>
                  </el-button>
                  <el-select v-model="form.Hosts[hostindex]" placeholder="选择主机" 
                              filterable
                              clearable
                              value-key="Id"
                              style="width: 400px;">
                      <!-- :label="`${item.Name} / ${item.PrivateIp} / ${item.PublicIp} / ${item.UseIp}`" -->
                    <el-option   
                      v-for="item in itemHost"
                      :key="item"
                      :label="`${item.Name} / ${item.UseIp}`"
                      :value="item">
                    </el-option>
                  </el-select>

                  <el-button type="warning" icon="delete" size="mini" @click="del_Host(hostindex)">删除</el-button>
                </el-form-item>
              </el-form-item>
              <el-button type="primary" icon="plus" size="mini" @click="add_Host()">添加</el-button>
            </div>

            <div class="panel" style="margin-top:5px;margin-bottom:15px;padding:15px;weight:500px"> 
              <el-form-item label="自定义参数:" v-for="(list,index) in valueJson" :key="index" label-width="100px">
                <el-input v-model="valueJson[index].Value" placeholder="请输入自定义参数" 
                      type="textarea" autosize 
                      style="width: 400px;">
                </el-input>

                <el-select v-model="valueJson[index].HostId" placeholder="不选为全部主机" 
                            filterable
                            clearable
                            style="width: 150px;">
                  <el-option   
                    v-for="item in form.Hosts"
                    :key="item.Id"
                    :label="`${item.Name} / ${item.UseIp}`"
                    :value="item.Id">
                  </el-option>
                </el-select>

                <el-button type="warning" icon="delete" size="mini"   @click="del_valueJson(index)">删除</el-button>

              </el-form-item>

              <el-button type="primary" icon="plus" size="mini"   @click="add_valueJson()">添加</el-button>
            </div>

        </el-col>

        <el-col :span="12" style="margin-left: 10px">
          <!-- <div class="panel" style="margin-top:5px;margin-bottom:15px;padding:15px">
            <el-form-item label="Dockerfile:">
              <el-input
                type="textarea"
                :autosize="{ minRows: 2, maxRows: 100}"
                placeholder="请输入内容"
                v-model="form.Dockerfile">
              </el-input>
            </el-form-item>
          </div> -->

          <div class="panel" style="margin-top:5px;margin-bottom:15px;padding:15px">
            <el-form-item label="关联公有参数:" v-for="(conf,confindex) in form.Confs" :key="confindex">
              <el-select v-model="form.Confs[confindex]" placeholder="选择已有公有类" 
                          filterable
                          value-key="Id"
                          style="width: 400px;">
                <el-option   
                  v-for="item in itemConf"
                  :key="item"
                  :label="item.Name"
                  :value="item">
                </el-option>
              </el-select>

              <el-button @click.stop="get_itemConf" size="mini">
                <i class="fa fa-refresh"></i>
              </el-button>

              <el-button type="warning" icon="delete" size="mini" @click="del_Conf(confindex)">删除</el-button>

              <el-form-item>
                <table v-if="form.Confs[confindex].Value">
                  <tr v-for="(item, index) in JSON.parse(form.Confs[confindex].Value)" :key="index" style="line-height:20px">
                      <td align="right" style="width:300px;padding:5px">
                          {{ item.Key }}
                      </td>

                      <td align="left" style="padding:5px">
                          {{ item.Value }}
                      </td>
                  </tr>
                </table>
              </el-form-item>

            </el-form-item>

            <el-button type="primary" icon="plus" size="mini"   @click="add_Conf()">添加</el-button>
          </div>

          <div class="panel" style="margin-top:5px;margin-bottom:15px;padding:15px">
            <el-form-item label="服务类型:" label-width="100px">
              <el-select v-model="form.Class" placeholder="选择服务" 
                          filterable
                          style="width: 400px;"
                          @change="change_class()">
                <el-option   
                  v-for="item in itemClass"
                  :key="item"
                  :label="item"
                  :value="item">
                </el-option>
              </el-select>
            </el-form-item>                  

            <el-form-item v-if="form.Class != 'java'">
            <el-form-item label="所属平台:" label-width="100px">
              <el-select v-model="form.Platforms[0]" placeholder="选择平台" 
                          value-key="Id"
                          filterable
                          clearable
                          style="width: 400px;"
                          @change="change_platform()">
                <el-option   
                  v-for="item in itemPlatform"
                  :key="item.Name"
                  :label="item.Name"
                  :value="item">
                </el-option>
              </el-select>
              <el-button @click.stop="get_itemPlatform()" size="mini">
                <i class="fa fa-refresh"></i>
              </el-button>
            </el-form-item>

            <td v-if="form.Platforms[0]&&!form.Platforms[0].Value" align="right" style="width:300px;padding:5px">
              无平台参数
            </td>
            <table v-if="form.Platforms[0]&&form.Platforms[0].Value" >
              <tr v-for="(item, index) in JSON.parse(form.Platforms[0].Value)" :key="index"
                style="line-height:20px">
                  <td align="right" style="width:300px;padding:5px">
                      {{ item.Key }}
                  </td>

                  <td align="left" style="padding:5px">
                      {{ item.Value }}
                  </td>
              </tr>
            </table>

            <el-form-item label="加速域名:" prop="Domains" label-width="100px">
              <el-form-item v-for="(list,index) in form.Domains" :key="index" style="margin-top: 5px">
                <!-- <el-button @click.stop="get_itemDomain()" size="mini">
                  <i class="fa fa-refresh"></i>
                </el-button> -->
                <el-select v-model="form.Domains[index]" placeholder="请指定加速域名" 
                            value-key="Id"
                            filterable
                            clearable
                            style="width: 400px;"
                            @change="change_domain(index)">
                    <!-- v-for="item in form.Platforms[0].Domains" -->
                  <el-option   
                    value-key
                    v-for="item in itemDomain"
                    :key="item.Name"
                    :label="`${item.Name} / ${item.Domain} / ${(item.Quicken == 0) ? '' : '加速'}`"
                    :value="item">
                  </el-option>
                </el-select>
                <!-- <el-button type="warning" icon="delete" size="mini"   @click="del_addDomain(index)">删除</el-button> -->
                <font v-show="index == 0" style="color: red" size="1">此域名Mqtt加速！</font>
              </el-form-item>
            </el-form-item>

            <!-- <el-button type="primary" icon="plus" size="mini"   @click="add_addDomain()">添加</el-button> -->
            </el-form-item>

            <el-form-item label="非加速域名:" prop="Domains" label-width="100px">
              <div v-for="item in noQuickenDomain" :key="item">{{item.Domain}}</div>
            </el-form-item>
          </div>
        </el-col>

      </el-row>
    </div>

    <div class="panel" style="margin-top:5px;margin-bottom:15px;padding:15px">
    <el-form-item>
      <el-button type="primary" @click="on_submit_form" :loading="on_submit_loading">立即提交
      </el-button>
      <el-button @click="$router.back()">取消</el-button>
    </el-form-item>
    </div>
  </el-form> 
</template>
<script type="text/javascript">
  import {panelTitle} from 'components'
  import {port_platform, port_domain, port_conf, port_host, port_service} from 'common/port_uri'
  import store from 'store'

  export default {
    data() {
      return {
        ProjectId: store.state.user_info.user.ProjectId, 
        form: {
          Id: 0,
          // UserId: store.state.user_info.user.Id,
          Name: "",
          Class: "java",
          ImagePath: null,
          Port: "",
          LogKeyword: "Started .* seconds",
          // IsRegister: 0,
          Health: "",
          DockerName: "",
          // Upload: "",
          Value: "",
          Hosts: [],
          Confs: [],
          Project: {
            Id: store.state.user_info.user.ProjectId,
          },
          Platforms: [],
          Domains: [],
          // DomainId: "",
          // PlatformId: ""
        },
        itemConf: [],
        itemClass: [
          "java", 
          "h5", 
          "merchant",
          "download", 
          "download-share", 
          "agent", 
          "share-agent", 
          "customer", 
          "chat-backend", 
          "other"
        ],

        // valueJson: [{}],
        valueJson:[{
          Value: null,
          HostId: null
        }],
        
        itemHost: [],

        // addPlatform:[null],
        itemPlatform: [],
        // addDomain:[null],
        itemDomain: [],

        noQuickenDomain: [],

        route_id: this.$route.params.id,
        load_data: false,
        on_submit_loading: false,
        rules: {
          Name: [{required: true, message: '项目名称不能为空', trigger: 'blur'}],
          ImagePath: [{required: true, message: '项目地址不能为空', trigger: 'blur'}],
          Port: [{required: true, message: '项目端口不能为空', trigger: 'blur'}],
        }
      }
    },
    created() {
      if (this.route_id) {
        this.get_form_data()
      }
        this.get_itemPlatform()
        this.get_itemConf()
        this.get_itemHost()
    },
    methods: {
      //下拉框获取已有域名
      // get_itemDomain(){
      //   this.load_data = true
      //   this.$http.get(port_domain.platform, {
      //                       params: {
      //                           // project_id: this.ProjectId,
      //                           platform_id: this.form.Platforms[0].Id,
      //                       }
      //                   })
      //     .then(({data: {data}}) => {
      //       this.load_data = true
      //       this.itemDomain = data
      //       this.load_data = false
      //     })
      //     .
      //     catch(() => {
      //       this.load_data = false
      //     })
      // },

      get_itemDomain() {
        this.load_data = true
        if (this.form.Domains.length == 0) {
          this.form.Domains.push(null)
        }
        this.itemDomain = []
        for(let i in this.form.Platforms[0].Domains) {
          if (this.form.Platforms[0].Domains[i].Class == this.form.Class) {
            this.itemDomain.push(this.form.Platforms[0].Domains[i])
          }
        }
        
        this.load_data = false
      },

      get_noQuickenDomain() {
        this.load_data = true
        this.$http.get(port_domain.noquickenlist, {
                            params: {
                                platform_id: this.form.Platforms[0].Id,
                                class: this.form.Class
                            }
                        })
          .then(({data: {data}}) => {
            this.load_data = true
            this.noQuickenDomain = data
            this.load_data = false
          })
          .
          catch(() => {
            this.load_data = false
          })

      },

      add_addDomain(){
        this.load_data = true
        this.form.Domains.push(null);
        this.load_data = false
      },
      
      del_addDomain(index){
        this.load_data = true
        this.form.Domains.splice(index,1)
        this.load_data = false
      },

      //下拉框获取已有类型
      get_itemPlatform(){
        this.load_data = true
        this.$http.get(port_platform.list, {
                            params: {
                                project_id: this.ProjectId,
                            }
                        })
          .then(({data: {data}}) => {
            this.load_data = true
            this.itemPlatform = data
            this.load_data = false
          })
          .
          catch(() => {
            this.load_data = false
          })
      },

      add_addplatform(){
        this.load_data = true
        this.form.Platforms.push(null);
        this.load_data = false
      },
      
      del_addplatform(index){
        this.load_data = true
        this.form.Platforms.splice(index,1)
        this.load_data = false
      },

      change_class() {
        this.load_data = true
        if (this.form.Class != "java") {
          if (this.form.Platforms.length == 0 || this.form.Platforms[0] == "") {
            this.form.Platforms[0] = this.itemPlatform[0]
          }
          this.get_itemDomain()
          this.get_noQuickenDomain()
        } else {
            this.form.Platforms = []
            this.form.Domains = []
        }
        this.load_data = false
      },

      change_platform() {
        this.load_data = true
        this.form.Domains = []
        this.get_itemDomain()
        this.get_noQuickenDomain()
        this.load_data = false
      },

      change_domain(index) {
        this.load_data = true
        let obj = {}
        if(this.form.Domains.length == 1 && this.form.Domains[1] == null) {
          this.load_data = false
          return
        }
        for ( let i in this.form.Domains) {
          obj[this.form.Domains[i].Id] = i
        }

        if ( this.form.Domains.length != Object.keys(obj).length) {
          // this.form.Domains.splice(index, 1)
          this.form.Domains[index] = {}
          this.$message({
            message: "不能选择重复域名",
            type: 'warning'
          })
          this.load_data = false
        }

        this.load_data = false
      },

      //下拉框获取已有公有类型
      get_itemHost(){
        this.load_data = true
        this.$http.get(port_host.project, {
                            params: {
                                project_id: this.form.Project.Id
                            }
                        })
          .then(({data: {data}}) => {
            this.load_data = true
            this.itemHost = data
            this.load_data = false
          })
          .
          catch(() => {
            this.load_data = false
          })
      },

      add_Host(){
        this.load_data = true
        this.form.Hosts.push({});
        this.load_data = false
      },
      
      del_Host(index){
        this.load_data = true
        this.form.Hosts.splice(index,1)
        this.load_data = false
      },

      add_valueJson(){
        this.load_data = true
        this.valueJson.push({
          Value: null,
          HostId: null
        });
        this.load_data = false
      },
      
      del_valueJson(index){
        this.load_data = true
        this.valueJson.splice(index, 1)
        this.load_data = false
      },
            
      //下拉框获取已有公有类型
      get_itemConf(){
        this.load_data = true
        this.$http.get(port_conf.list, {
          params: {
            project_id: this.form.Project.Id
          }
        })
          .then(({data: {data}}) => {
            this.itemConf = data
            // for(var i in this.itemConf) {
            //     // this.item[i].json = JSON.parse(this.itemConf[i].Value);
            //     this.$set(this.itemConf[i], "json", JSON.parse(this.itemConf[i].Value))
            // }
            this.load_data = false
          })
          .
          catch(() => {
            this.load_data = false
          })
      },

      add_Conf(){
        this.load_data = true
        this.form.Confs.push({});
        this.load_data = false
      },
      
      del_Conf(index){
        this.load_data = true
        this.form.Confs.splice(index,1);
        this.load_data = false
      },

      //获取数据
      get_form_data() {
        this.load_data = true
        this.$http.get(port_service.id, {
          params: {
            id: this.route_id
          }
        })
          .then(({data: {data}}) => {
            this.form = data
            if(this.form.Platforms[0]) {
              this.get_itemDomain()
            }
            // for(var i in this.form.Confs) {
            //     // this.item[i].json = JSON.parse(this.itemConf[i].Value);
            //     this.$set(this.form.Confs[i], "json", JSON.parse(this.form.Confs[i].Value))
            // }
            this.valueJson = JSON.parse(this.form.Value)
          })
          .catch(() => {
            this.load_data = false
          })
      },

      //时间选择改变时
      on_change_birthday(val) {
        this.$set(this.form, 'birthday', val)
      },

      //提交
      on_submit_form() {
        this.$refs.form.validate((valid) => {
          if (
            !valid
          )
            return false

          this.load_data = true
          this.on_submit_loading = true

          this.form.Value = JSON.stringify(this.valueJson)
          for (let i in this.form.Domains) {
            if(this.form.Domains[i] === "" || this.form.Domains[i] === null) {
              this.form.Domains.splice(i, 1)
            }
          }
          this.$http.post(port_service.saverelated, this.form)
            .then(({data: {msg}}) => {
              this.$message({
                message: msg,
                type: 'success'
              })
              setTimeout(() => {           
                  this.load_data = true
                  this.$router.back()
                },
                500
              )
            })
            .catch(() => {
              this.load_data = false
              this.on_submit_loading = false
            })

            this.load_data = false
            this.on_submit_loading = false
          
        })
      }
    },
    components: {
      panelTitle
    }
  }
</script>
