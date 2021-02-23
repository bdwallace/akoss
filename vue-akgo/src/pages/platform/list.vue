<template>
  <div class="panel">
    <panel-title :title="$route.meta.title">
      <!-- <div style="float: left;margin-right: 10px;margin-top: 5px">
          <search @search="submit_search"></search>
      </div>
      <el-button @click.stop="on_refresh" size="small">
          <i class="fa fa-refresh"></i>
      </el-button> -->
      <div style="float: left;margin-right: 10px;margin-top: 5px">
        <el-button plain size="small">全部&nbsp;{{table_data.length}}&nbsp;个平台</el-button>
      </div>


      <div style="float: left;margin-right: 10px;margin-top: 5px">
        <el-button type="success" size="small" icon="setting" @click="get_service_status_all()" :loading="on_submit_loading">全部检测</el-button>
      </div>

      <div style="float: left;margin-right: 10px;margin-top: 5px;">
        <el-button type="warning" size="small" icon="setting" @click="multi_reload()" :loading="on_submit_loading">批量重载</el-button>
      </div>

      <div style="float: left;margin-right: 10px;margin-top: 5px;">
        <el-select v-model="Class" clearable placeholder="服务类型选择" style="width:120px">
          <el-option
            v-for="item in itemClass"
            :key="item"
            :label="item"
            :value="item">
          </el-option>
        </el-select>
      </div>

      <div style="float: left;margin-right: 10px;margin-top: 5px;">
        <el-button type="danger" size="small" icon="setting" @click="on_dialog()" :loading="on_submit_loading">发布</el-button>
      </div>

      <router-link :to="{name: 'serviceAdd'}" tag="span">
        <el-button type="primary" icon="plus" size="small">创建平台服务</el-button>
      </router-link>

      <router-link :to="{name: 'platformAdd'}" tag="span">
        <el-button type="primary" icon="plus" size="small">创建平台</el-button>
      </router-link>

    </panel-title>
    <div class="panel-body-line" style="clear: both;">
      <el-table
        :data="table_data"
        v-loading="load_data"
        @expand="expand"
        element-loading-text="拼命加载中"
        :default-sort = "{prop: 'ChannelId', order: 'ascending'}"
        border
        :row-key="getRowKeys"
        :expand-row-keys="expands"
        style="width: 100%;">
        <el-table-column type="expand">
          <template scope="props">

            <!-- @select="select"
            @select-all="select_all(props.$index)" -->
            <el-table
              :data="props.row.Services"
              v-loading="load_data"
              element-loading-text="拼命加载中"
              @selection-change="(selection)=>select_change(selection, props.$index)"
              border
              :default-sort = "{prop: 'Id', order: 'ascending'}"
              style="width: 100%;"
              :ref="checkTable[props]">
              <el-table-column
                type="selection">
              </el-table-column>

              <el-table-column
                prop="Id"
                sortable
                width="50">
              </el-table-column>

              <el-table-column
                prop="Name"
                label="名称"
                sortable
                width="150">
              </el-table-column>

              <!-- <el-table-column
                      prop="Tag"
                      label="最后发布版本"
                      width="170">
              </el-table-column> -->

              <el-table-column
                prop="Tag"
                sortable
                label="最新版本">
              </el-table-column>

              <el-table-column
                prop="Port"
                label="端口"
                width="115">
                <template scope="service_props">
                  <font style="display: flex" v-for="item in service_props.row.Port.split(',')" :key="item">{{item}}</font>
                </template>
              </el-table-column>

              <el-table-column
                prop="service_data"
                label="主机"
                fit=true
                width="740">
                <template scope="service_props">
                  <div style="margin-top: -15px"> <span></span> </div>
                  <!-- <div v-for="(item, index) in table_data[props.$index].Services[service_props.$index].service_data" :key="index"> -->
                  <div v-for="(item, index) in service_props.row.service_data" :key="index">
                    <div style="display: inline-block;" v-if="item.ip_show == item.ip">
                      <el-input readonly style="width: 120px" size="mini" v-model="item.ip_pub" >
                      </el-input>
                      <el-input readonly class="input_blue" style="width: 120px" size="mini" v-model="item.ip" >
                      </el-input>
                    </div>
                    <div style="display: inline-block;" v-else>
                      <el-input readonly class="input_blue" style="width: 120px" size="mini" v-model="item.ip_pub" >
                      </el-input>
                      <el-input readonly style="width: 120px" size="mini" v-model="item.ip" >
                      </el-input>
                    </div>

                    <el-input v-if="item.ps_status.substr(0, 2) != 'Up'" class="input_red" readonly style="width: 100px" size="mini" v-model="item.ps_status">
                    </el-input>
                    <el-input v-else readonly style="width: 100px" size="mini" v-model="item.ps_status">
                    </el-input>

                    <!-- <el-input readonly style="width: 150px" size="mini" v-model="item.ps_created_at">
                    </el-input> -->

                    <el-input readonly style="width: 220px" size="mini" v-model="item.ps_image">
                    </el-input>

                    <div style="display: inline-block;margin-right:10px">
                      <el-tooltip effect="light" :content="item.url" placement="left">
                        j<el-tag v-if="item.health=='200'" size="mini" type="primary">健康</el-tag>
                        <el-tag v-else size="mini" type="danger">异常</el-tag>
                      </el-tooltip>
                    </div>

                    <el-button style="margin-right:-12px" size="mini" type="primary" @click="reload(service_props.row, service_props.row.Hosts[index])" :loading="on_submit_loading">重载
                    </el-button>

                    <el-button style="margin-right:-12px" size="mini" type="primary" @click="stop(service_props.row, service_props.row.Hosts[index])" :loading="on_submit_loading">关闭
                    </el-button>

                    <el-button style="margin-right:-12px" v-if="service_props.row.ReleaseTo !== ''" size="mini" type="primary" @click="download(props.row.Project, service_props.row, service_props.row.Hosts[index])">下载</el-button>
                  </div>
                  <div style="margin-top: -15px"> <span></span> </div>
                </template>
              </el-table-column>

              <el-table-column
                label=""
                fit=true
                width="55">
                <template scope="service_props">
                  <router-link v-if="service_props.row.ReleaseTo" :to="{name: 'serviceUpload', params: {id: service_props.row.Id}}" tag="span">
                    <el-button  type="primary" size="mini" >上传</el-button>
                  </router-link>
                </template>
              </el-table-column>

              <el-table-column
                label="操作"
                width="185">
                <template scope="service_props">
                  <router-link :to="{name: 'serviceUpdate', params: {id: service_props.row.Id}}" tag="span">
                    <el-button type="info" size="mini" icon="edit">修改</el-button>
                  </router-link>

                  <el-button type="warning" size="mini" icon="document" @click="copy_data(props.$index, service_props.row.Id)">复制
                  </el-button>

                  <el-button style="margin-left:0px" type="danger" size="mini" icon="delete" @click="service_delete_data(props.$index, service_props.row.Id)">删除
                  </el-button>
                </template>
              </el-table-column>
            </el-table>
          </template>
        </el-table-column>

        <el-table-column
          prop="ChannelId"
          sortable
          label="平台id"
          width="100">
        </el-table-column>

        <el-table-column
          prop="Name"
          label="平台名称"
          width="190">
        </el-table-column>

        <el-table-column label="">
          <template scope="props">
            <el-button plain size="small" v-if="props.row.Services != null">全部&nbsp;{{props.row.Services.length}}&nbsp;个服务</el-button>
          </template>
        </el-table-column>
        <!-- <el-table-column
                prop="Value"
                label="平台参数">
            <template scope="props">
                <el-button readonly size="mini" v-for="(item, index) in JSON.parse(props.row.Value)" :key="index" style="margin-right: 5px">{{item.Key}}={{item.Value}}</el-button>
            </template>
        </el-table-column> -->

        <!-- <el-table-column
                label="服务操作"
                width="250">
            <template scope="props">
                <el-button type="info" size="small" icon="setting" @click="get_platform_status(props.$index)" :loading="on_submit_loading">全部检测</el-button>
            </template>
        </el-table-column> -->

        <el-table-column
          label="平台操作"
          width="185">
          <template scope="props">
            <router-link :to="{name: 'platformUpdate', params: {id: props.row.Id}}" tag="span">
              <el-button type="info" size="small" icon="edit">修改</el-button>
            </router-link>
            <!-- <el-button type="warning" size="small" icon="document" @click="copy_data(props.row.id)">复制
            </el-button> -->
            <el-button style="margin-left:0px" type="danger" size="small" icon="delete" @click="delete_check(props.row.Id)">删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <!-- :close-on-press-escape="true"
    @update:visible="close_dialog()"
    @close="close_dialog()"
    :close-on-click-modal="true" -->
    <el-dialog title="确认发布以下服务"
               :visible.sync="is_classDialog"
               :modal="true"
               :show-close="false"
               :modal-append-to-body="false">
      <!-- <div  v-for="(item, index) in dialogForm" :key="item" style="width: 200px;display: inline-block;">
          <el-tag :closable="true" type="primary"
              @close="dialogForm.splice(index, 1)"
              style="align:online">
              {{item.Name}}
          </el-tag>
      </div> -->
      <el-checkbox-group v-model="dialogForm">
        <el-checkbox v-for="item in dialogFormAll" :label="item" :key="item">{{item.Name}}</el-checkbox>
      </el-checkbox-group>

      <div slot="footer" class="dialog-footer">
        <el-button @click="close_dialog()">取 消</el-button>
        <el-button type="primary" @click="on_submit_form(dialogForm)">提 交</el-button>
      </div>
    </el-dialog>

    <el-dialog title="以下有未解除关联的域名，是否同时删除！"
               :visible.sync="is_delDomainDialog"
               :modal="true"
               :show-close="false"
               :modal-append-to-body="false">
      <p v-for="item in del_Domain" :key="item"> {{item.Domain}} </p>

      <div slot="footer" class="dialog-footer">
        <el-button @click="is_delDomainDialog = false">取 消</el-button>
        <el-button type="primary" @click="delete_data(del_Platform_Id)">提 交</el-button>
      </div>
    </el-dialog>

  </div>
</template>

<style>
  .input_red .el-input__inner {
    color: red;
  }
  .input_blue .el-input__inner {
    color: blue;
  }
  .input_green .el-input__inner {
    color: green;
  }
  .input_gray .el-input__inner {
    color: gray;
  }
  .el-table__expanded-cell {
    padding-top: 0px;
    padding-right: 10px;
    padding-bottom: 10px;
    padding-left: 10px;
  }
  .el-table .cell {
    padding-top: 0px;
    padding-right: 10px;
    padding-bottom: 0px;
    padding-left: 10px;
  }
</style>

<script type="text/javascript">
  import {panelTitle, bottomToolBar, search} from 'components'
  import {port_service, port_platform, port_deploy, port_walle} from 'common/port_uri'
  import {walle_stop, walle_restart, walle_reload, walle_multi_reload} from 'common/walle'
  import store from "store"
  export default{
    data(){
      return {
        ProjectId: store.state.user_info.user.ProjectId,
        table_data: [],
        checkTable: [],
        form: [],
        dialogForm: [],
        dialogFormAll: [],
        load_data: false,
        on_submit_loading: false,
        // 已经展开的平台
        expands: [],
        itemClass: [],
        Class: "",
        Services: [],

        is_classDialog: false,
        is_delDomainDialog: false,
        del_Domain: [],

        stop: walle_stop.stop,
        restart: walle_restart.restart,
        reload: walle_reload.reload,
      }
    },

    components: {
      panelTitle,
      bottomToolBar,
      search
    },

    created(){
      this.get_table_data()
    },

    methods: {
      get_table_data() {
        this.load_data = true
        // this.$http.get(port_platform.list,  {
        this.$http.get(port_platform.project,  {
          params: {
            project_id: this.ProjectId,
          }
        })
          .then(({data: {data}}) => {
            this.table_data = data
            this.initItemClass()
            this.load_data = false
          })
      },

      // handleSelectionChange(val) {
      //     this.form = val;
      // },

      initItemClass() {
        this.itemClass = []
        this.checkTable = []
        for(let i in this.table_data) {
          this.checkTable.push(null)
          for(let j in this.table_data[i].Services) {
            if (this.itemClass.indexOf(this.table_data[i].Services[j].Class) == -1) {
              this.itemClass.push(this.table_data[i].Services[j].Class)
            }
          }
        }
      },

      getRowKeys(row) {
        return row.Id;
      },

      select(val, row) {
        // this.form = val;
        // console.log("---val---", val)
        // console.log("---row---", row)
        // let v = val.indexOf(row)
        // if (v == -1) {
        //     this.form.splice(v)
        // } else {
        //     this.form.push(row)
        // }
        let v = this.form.indexOf(row)
        if (v == -1) {
          this.form.push(row)
        } else {
          this.form.splice(v)
        }
        // console.log("---form---", this.form)
      },

      select_change(selection, index) {
        this.checkTable[index] = selection
      },

      select_all(index) {
        // console.log("---i---", index)
        if(val.length == 0) {
          this.form = []
        } else {
          for(var i in val) {
            if ( this.form.indexOf(val[i]) == -1 ) {
              this.form.push(val[i])
            }
          }
        }
      },

      expand(row, expanded) {
        let index = this.table_data.indexOf(row)
        if (expanded) {
          // this.get_platform_service(index)
          this.get_platform_status(index)
        } else {
          let expands_index = this.expands.indexOf(row.Id)
          this.expands.splice(expands_index, 1)
        }
      },

      get_service_status_all() {
        this.load_data = true
        for(var i in this.table_data) {
          // this.get_platform_service(i)
          this.get_platform_status(i)
        }
        this.load_data = false
      },

      get_platform_status(index) {
        for(var i in this.table_data[index].Services) {
          this.get_service_status(index, i)
        }

        if (this.expands.indexOf(this.table_data[index].Id) < 0) {
          this.expands.push(this.table_data[index].Id);
          this.expands.sort()
        }
      },

      //根据项目ID获取项目的状态
      get_platform_service(index){
        this.load_data = true
        this.$http.get(port_service.platform, {
          params: {
            platform_id: this.table_data[index].Id
          }
        })
          .then(({data: {data}}) => {
            this.$set(this.table_data[index], "Services", data)
            this.get_platform_status(index)
            this.load_data = false
          })
        this.load_data = false
      },

      //根据项目ID获取项目的状态
      get_service_status(index, serviceindex){
        this.load_data = true
        this.$http.get(port_walle.dockerps, {
          params: {
            service_id: this.table_data[index].Services[serviceindex].Id,
            line_data: "----"
          }
        })
          .then(({data: {data}}) => {
            this.$set(this.table_data[index].Services[serviceindex], "service_data", data)
            this.load_data = false
          })
        this.load_data = false
      },

      close_dialog() {
        this.is_classDialog = false
        this.dialogFormAll = []
        this.dialogForm = []
      },

      multi_reload() {
        if(this.Class != "") {
          this.dialogFormAll = []
          for(let i in this.table_data) {
            for(let j in this.table_data[i].Services) {
              if (this.table_data[i].Services[j].Class == this.Class) {
                this.dialogFormAll.push(this.table_data[i].Services[j])
              }
            }
          }
          //设置初始值为全部
          this.dialogForm = this.dialogFormAll
          this.$confirm('确定需要重载?', '提示', {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning'
          }).then(() => {
            let walle_form = {
              User: store.state.user_info.user,
              Services: this.dialogForm,
            }
            if (walle_form.Services.length == 0) {
                this.on_submit_loading = false
                this.load_data = false
                this.$message({
                  message: "请选择重载的服务或选择类型批量重载",
                  type: 'warning'
                })
            }else {
              this.$http.post(port_walle.multireload, walle_form)
                .then(({data: {data}}) => {
                  this.on_submit_loading = false
                  this.load_data = false
                  this.$message({
                    message: "发送重载成功",
                    type: 'success'
                  })
                })
            }
          })
        } else {
          this.form = []
          for(let i in this.checkTable) {
            if(this.checkTable[i] != null) {
              this.form = this.form.concat(this.checkTable[i])
            }
          }
          this.$confirm('确定需要批量重载?', '提示', {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning'
          }).then(() => {
            let walle_form = {
              User: store.state.user_info.user,
              Services: this.form,
            }
            if (walle_form.Services.length == 0) {
              this.on_submit_loading = false
              this.load_data = false
              this.$message({
                message: "请选择重载的服务或选择类型批量重载",
                type: 'warning'
              })
            }else {
              this.$http.post(port_walle.multireload, walle_form)
                .then(({data: {data}}) => {
                  this.on_submit_loading = false
                  this.load_data = false
                  this.$message({
                    message: "发送重载成功",
                    type: 'success'
                  })
                })
            }
          })
        }
      },

      on_dialog() {
        if(this.Class != "") {
          this.dialogFormAll = []
          for(let i in this.table_data) {
            for(let j in this.table_data[i].Services) {
              if (this.table_data[i].Services[j].Class == this.Class) {
                this.dialogFormAll.push(this.table_data[i].Services[j])
              }
            }
          }
          //设置初始值为全部
          this.dialogForm = this.dialogFormAll
          this.is_classDialog = true
        } else {
          this.form = []
          for(let i in this.checkTable) {
            if(this.checkTable[i] != null) {
              this.form = this.form.concat(this.checkTable[i])
            }
          }
          this.on_submit_form(this.form)
        }

      },

      //提交
      on_submit_form(deployform) {
        this.load_data = true
        this.on_submit_loading = true
        if (deployform.length == 0) {
          this.$message({
            message: "请选择发布的项目",
            type: 'warning'
          })
          this.load_data = false
          this.on_submit_loading = false
          return
        }
        this.$http.post(port_deploy.save, deployform)
          .then(({data: {data}}) => {
            // this.$message({
            //     message: msg,
            //     type: 'success'
            // })
            setTimeout(() => {
              let routeData = this.$router.resolve({
                name: "deployRelease",
                params: {
                  id: data,
                }
              });
              window.open(routeData.href, '_blank');
              this.on_submit_loading = false
            })
          })
          .catch(() => {
            this.load_data = false
            this.on_submit_loading = false
          })

        this.load_data = false
        this.on_submit_loading = false
      },

      download(project, service, host) {
        this.on_submit_loading = true
        this.load_data = true
        let walle_form = {
          User: store.state.user_info.user,
          Service: service,
          Host: host
        }
        walle_form.Service.Project = project

        this.$http.post(port_walle.download, walle_form)
          .then(({data: {data}}) => {
            window.location.href = data
            // this.$message({
            //    message: "发送关闭成功",
            //     type: 'success'
            // })
            this.on_submit_loading = false
            this.load_data = false
          }).catch(() => {
          this.on_submit_loading = false
          this.load_data = false
        })
        this.on_submit_loading = false
        this.load_data = false
      },

      //重载
      on_reload(project, service, host){
        this.$confirm('确定需要重启?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          // this.on_submit_loading = true
          // this.load_data = true
          let walle_form = {
            User: store.state.user_info.user,
            Service: service,
            Host: host
          }
          walle_form.Service.Project = project

          this.$http.post(port_walle.reload, walle_form)
            .then(({data: {data}}) => {
              this.on_submit_loading = false
              this.load_data = false
              this.$message({
                message: "发送重载成功",
                type: 'success'
              })
            }).catch(() => {
            this.on_submit_loading = false
            this.load_data = false
          })
        }).catch(() => {
          this.on_submit_loading = false
          this.load_data = false
        })
      },

      //关闭
      on_stop(project, service, host){
        this.$confirm('确定需要关闭?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          // this.on_submit_loading = true
          // this.load_data = true
          let walle_form = {
            User: store.state.user_info.user,
            Service: service,
            Host: host
          }
          walle_form.Service.Project = project

          this.$http.post(port_walle.stop, walle_form)
            .then(({data: {data}}) => {
              this.on_submit_loading = false
              this.load_data = false
              this.$message({
                message: "发送关闭成功",
                type: 'success'
              })
            }).catch(() => {
            this.on_submit_loading = false
            this.load_data = false
          })
        }).catch(() => {
          this.on_submit_loading = false
          this.load_data = false
        })
      },

      //根据id删除数据
      delete_check(id){
        // this.$confirm('此操作将删除该数据, 是否继续?', '提示', {
        //     confirmButtonText: '确定',
        //     cancelButtonText: '取消',
        //     type: 'warning'
        // })
        // .then(() => {
        // this.load_data = true
        // this.$http.delete(port_platform.id, {
        this.$http.get(port_platform.relatedcheck, {
          params: {
            id: id,
          }
        })
          .then(({data: {data}}) => {
            if (data != null) {
              this.is_delDomainDialog = true
              this.del_Platform_Id = id
              this.del_Domain = data
            } else {
              this.delete_data(id)
            }
          })
          .catch(({msg: {msg}}) => {
            this.load_data = false
          })
        // })
        // .catch(() => {
        //     this.load_data = false
        // })
      },

      //根据id删除数据
      delete_data(id){
        this.$confirm('此操作将删除该数据, 是否继续?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        })
          .then(() => {
            this.load_data = true
            this.$http.delete(port_platform.id, {
              params: {
                id: id,
              }
            })
              .then(({data: {msg}}) => {
                this.$message({
                  message: msg,
                  type: 'success'
                })
                this.is_delDomainDialog = false
                this.get_table_data()
                // this.get_platform_service(index)
                this.load_data = false
              })
              .catch(({msg: {msg}}) => {
                this.load_data = false
              })
          })
          .catch(() => {
            this.load_data = false
          })
      },

      //根据id删除数据
      service_delete_data(index, id){
        this.$confirm('此操作将删除该数据, 是否继续?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        })
          .then(() => {
            // this.load_data = true
            this.$http.delete(port_service.id, {
              params: {
                id: id,
              }
            })
              .then(({data: {msg}}) => {
                this.get_platform_service(index)
                this.$message({
                  message: msg,
                  type: 'success'
                })
                // this.get_table_data()
                this.load_data = false
              })
              .catch(() => {
                this.load_data = false
              })
          })
          .catch(() => {
            this.load_data = false
          })
      },

      //复制项目
      copy_data(index, id){
        this.$confirm('是否复制项目?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        })
          .then(() => {
            this.load_data = true
            this.$http.get(port_service.copy, {
              params: {
                id: id,
              }
            })
              .then(({data: {msg}}) => {
                this.get_platform_service(index)
                this.$message({
                  message: msg,
                  type: 'success'
                })
                // this.get_table_data()
                this.load_data = false
              })
              .catch(() => {
                this.load_data = false
              })
          })
          .catch(() => {
            this.load_data = false
          })
      },

    },

  }
</script>
