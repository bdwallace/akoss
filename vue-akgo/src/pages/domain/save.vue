<template>
  <div class="panel">
    <panel-title :title="$route.meta.title"></panel-title>
    <div class="panel-body"
         v-loading="load_data"
         element-loading-text="拼命加载中">
      <el-row>
        <el-col :span="24">
          <el-form ref="form" :model="form" :rules="rules">
            <!-- <el-form-item prop="Level" label-width="120px" >
              <el-radio-group v-if="! route_id" style="margin-top: 15px" v-model="form.Level">
                <el-radio v-for="item in level_data" :key="item.Id" :label="item.Id">{{item.Name}}</el-radio>
              </el-radio-group>

              <el-radio-group v-show="route_id" style="margin-top: 15px" v-model="form.Level">
                <el-radio disabled v-for="item in level_data" :key="item.Id" :label="item.Id">{{item.Name}}</el-radio>
              </el-radio-group>
            </el-form-item> -->

            <el-form-item label="域名:" prop="Domain" label-width="120px">
<!--              <el-input v-model="form.Domain" placeholder="请输入域名"-->
<!--                        style="width: 600px;"></el-input>-->

              <el-input type="textarea" :rows="2" placeholder="请输入域名"
                        v-model="form.Domain" style="width: 600px;"></el-input>

              <font style="color: red" size="1">域名不需要添加www前缀！</font>
              <font style="color: red" size="1">添加多个域名需要使用换行符分隔！</font>
            </el-form-item>

            <el-form-item label="名称\备注:" prop="Name" label-width="120px">
              <el-input v-model="form.Name" placeholder="请域名名称,方便识别域名作用"
                        style="width: 600px;"></el-input>
            </el-form-item>

            <el-form-item label="Crt:" label-width="120px">
                <el-popover
                    ref="Crt"
                    placement="left-start"
                    width="600"
                    trigger="click">
                    <!-- <span style="color:teal;white-space: pre-wrap;">{{form.Crt}}</span> -->
                    <el-input type="textarea" :rows="10" readonly v-model="form.Crt"></el-input>
                </el-popover>

                <el-button v-show="form.Crt != ''" type="info" size="mini" v-popover:Crt>
                  查看
                </el-button>

                <el-button v-show="form.Crt == ''" type="info" size="mini">无证书
                </el-button>

                <el-upload
                  style="display: inline-block; margin-left: 100px"
	                ref="upload"
	                action="alert"
                  :show-file-list="false"
	                :auto-upload="false"
	                :file-list="uploadFilesCrt"
	                :on-change="loadFileCrt">
	                <el-button size="mini" type="primary">选取文件</el-button>
                </el-upload>
            </el-form-item>

            <el-form-item label="Key:" label-width="120px">
                <el-popover
                    ref="Key"
                    placement="left-start"
                    width="600"
                    trigger="click">
                    <!-- <span style="color:teal;white-space: pre-wrap;">{{form.Key}}</span> -->
                    <el-input type="textarea" :rows="10" readonly v-model="form.Key"></el-input>
                </el-popover>

                <el-button v-show="form.Key != ''" type="info" size="mini" v-popover:Key>
                  查看
                </el-button>

                <el-button v-show="form.Key == ''" type="info" size="mini">无私钥
                </el-button>

                <el-upload
                  style="display: inline-block; margin-left: 100px"
	                ref="upload"
	                action="alert"
                  :show-file-list="false"
	                :auto-upload="false"
	                :file-list="uploadFilesKey"
	                :on-change="loadFileKey">
	                <el-button size="mini" type="primary">选取文件</el-button>
                </el-upload>
            </el-form-item>

            <el-form-item label="是/否监控:" label-width="120px">
              <el-radio-group v-model="form.Monitor">
                <el-radio :label="0">否</el-radio>
                <el-radio :label="1">是</el-radio>
              </el-radio-group>
            </el-form-item>

            <el-form-item label="是/否CDN加速:" label-width="120px">
              <el-radio-group v-model="form.Quicken">
                <el-radio :label="0">否</el-radio>
                <el-radio :label="1">是</el-radio>
              </el-radio-group>
            </el-form-item>

            <el-form-item label="端口:" prop="Port" label-width="120px">
              <el-input v-model="form.Port" placeholder="需要运维人员填入"
                        style="width: 200px;"></el-input>
            </el-form-item>

            <el-form-item label="域名类型" label-width="120px">
              <el-select v-model="form.Class" placeholder="选择域名类型"
                          filterable
                          style="width: 300px;">
                <el-option
                  v-for="item in itemClass"
                  :key="item"
                  :label="item"
                  :value="item">
                </el-option>
              </el-select>
            </el-form-item>

            <el-form-item label="关联平台:"  label-width="120px">
              <el-button @click.stop="get_itemplatform" size="mini">
                <i class="fa fa-refresh"></i>
              </el-button>

                          <!-- clearable -->
              <el-select v-model="form.Platforms" placeholder="选择已有平台"
                          filterable
                          multiple
                          value-key="Id"
                          style="width: 300px;">
                          <!-- @change="get_itemplatform_check(0)"> -->
                          <!-- @visible-change="get_itemplatform_check(index)"> -->
                <el-option
                  v-for="item in itemPlatform"
                  :key="item.Name"
                  :label="item.Name"
                  :value="item">
                </el-option>
              </el-select>

              <!-- <el-select v-model="form.Services[0]" placeholder="选择已有项目"
                          filterable
                          value-key="Id"
                          style="width: 300px;">
                <el-option
                  v-for="item in itemService[0]"
                  :key="item.Name"
                  :label="item.Name"
                  :value="item"
                  placeholder="请输入项目参数名">
                </el-option>
              </el-select>

              <el-button type="primary" icon="plus" size="mini"   @click="del_service(index)">删除</el-button> -->
            </el-form-item>

            <!-- <el-form-item label="关联项目:"  label-width="120px">
                <el-button type="primary" icon="plus" size="mini"   @click="add_service()">添加</el-button>
            </el-form-item> -->

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
  import {port_domain, port_level, port_platform, port_service} from 'common/port_uri'
  import {tools_verify} from 'common/tools'
  import store from 'store'

  export default {
    data() {
      return {
        // classData: [],
        // level_data: [],

        uploadFilesCrt: [],
        uploadFilesKey: [],
        crtFile: null,
        keyFile: null,
        crtFileName: null,
        keyFileNmae: null,
        addProject: [],
        // addProject: [{
        //   platform_id: null,
        //   project_id: null
        // }],
        itemPlatform: [],
        itemService: [],
        itemDomain: [],
        itemClass: ["download", "h5", "download-share", "gateway", "mqtt", "other"],

        form: {
          Name: null,
          Domain: [],
          Crt: "",
          Key: "",
          Port: null,
          Project: {
            Id: store.state.user_info.user.ProjectId
          },
          Monitor: 1,
          Quicken: 0,
          Class: null,
          Platforms: [],
          Services: []
        },
        route_id: this.$route.params.id,
        load_data: false,
        on_submit_loading: false,
        rules: {
          Name: [{required: true, message: '名称不能为空', trigger: 'blur'}],
          Domain: [{required: true, message: '域名不能为空', trigger: 'blur'}],
        }
      }
    },
    created() {
      this.get_itemplatform()
      if (this.route_id) {
        this.get_form_data()
      }
    },
    methods: {
      	loadFileCrt(file, fileList) {
          this.uploadFilesCrt = fileList
          // 解析上传的文件
          let reader = new FileReader()
          reader.readAsText(this.uploadFilesCrt[0].raw)
          reader.onload = (e) => {
            // 读取文件内容
            this.form.Crt = e.target.result
          }
          this.$message({
              message: "上传成功",
              type: 'success'
          })
        },

      	loadFileKey(file, fileList) {
          this.uploadFilesKey = fileList
          // 解析上传的文件
          let reader = new FileReader()
          reader.readAsText(this.uploadFilesKey[0].raw)
          reader.onload = (e) => {
            // 读取文件内容
            this.form.Key = e.target.result
          }
          this.$message({
              message: "上传成功",
              type: 'success'
          })
        },

        on_submit_form() {
          this.on_submit_loading = true
          this.$http.post(port_domain.save, this.form)
          .then(({data: {msg}}) => {
                this.on_submit_loading = false
                this.$message({
                    message: msg,
                    type: 'success'
                })
                setTimeout(() => {
                    this.$router.back()
                }, 500)
          })
          .catch(() => {
            this.on_submit_loading = false
          })
        },

        add_service(){
          this.load_data = true
          this.form.Platforms.push({});
          this.form.Services.push({});
          this.load_data = false
        },

        del_service(index){
          this.load_data = true
          this.itemService.splice(index, 1)
          this.form.Platforms.splice(index, 1)
          this.form.Services.splice(index, 1)
          this.load_data = false
        },

        //下拉框获取已有类型
        get_itemplatform_check(index){
          this.itemService[index] = []
          this.$set(this.itemService, index, this.form.Platforms[index].Services)
          // this.$set(this.form.Services, index, this.itemService[index][0])
          this.$forceUpdate()
        },

        // get_itemproject(platform_id, index) {
        //   this.load_data = true

        //   this.$http.get(port_conf.listplatform, {
        //                       params: {
        //                           // level: this.form.Level,
        //                           // type: "channel"
        //                           platformId: platform_id
        //                       }
        //                   })
        //     .then(({data: {data}}) => {
        //       this.$set(this.itemService, index, data)
        //       this.load_data = false
        //     })
        //     .
        //     catch(() => {
        //       this.load_data = false
        //     })
        // },

        // get_addProject(domain_id) {
        //   this.load_data = true
        //   this.$http.get(port_conf.listdomain, {
        //               params: {
        //                 domainId: domain_id
        //               }
        //             })
        //     .then(({data: {data}}) => {
        //       this.itemDomain = data
        //       this.addProject = []
        //       for(var i in data) {
        //         let addProject = {}
        //         if (data[i].platform_id == "") {
        //           addProject = {
        //             platform_id: null,
        //             project_id: data[i].id
        //           }
        //         } else {
        //           addProject = {
        //             platform_id: Number(data[i].platform_id),
        //             project_id: data[i].id
        //           }
        //         }
        //         this.get_itemproject(addProject.platform_id, i)
        //         // this.$set(this.addProject, i, addProject)
        //         this.addProject.push(addProject)

        //       }
        //       this.load_data = false

        //     })
        //     .
        //     catch(() => {
        //       this.load_data = false
        //     })

        // },

        //下拉框获取已有类型
        get_itemplatform(){
          this.load_data = true
          this.$http.get(port_platform.list, {
                              params: {
                                project_id: this.form.Project.Id
                              }
                          })
            .then(({data: {data}}) => {
              this.itemPlatform = data
              this.load_data = false
            })
            .
            catch(() => {
              this.load_data = false
            })
        },

        //获取数据
        get_form_data() {
            this.load_data = true
            this.$http.get(port_domain.id, {
              params: {
                  id: this.route_id
              }
            })
            .then(({data: {data}}) => {
                this.form = data
                // this.get_addProject(this.form.Id)
                this.load_data = false
            })
            .catch(() => {
                this.load_data = false
            })
        },

        handleGetFileCrt(data) {
          var _this = this;
          if (file.raw) {
              let reader = new FileReader()
              reader.onload = function (e) {
                  _this.contentHtml = e.target.result;
              };
              reader.readAsText(file.raw,'gb2312');

          }
          console.log(file, fileList);
        },

        handleGetFileKey(data) {
            this.keyFile = data.target.files[0];
        },

        upload(file, name) {
            this.on_submit_loading = true
            var formdata = new FormData()
            formdata.append("file", file)
            formdata.append("name", name)

            this.$http.post(port_domain.upload, formdata)
                .then(({data: {msg}}) => {
                    this.on_submit_loading = false

                    this.$message({
                        message: msg,
                        type: 'success'
                    })

                    setTimeout(() => {
                    },
                    500
                )
                })
        },
      },


      components: {
        panelTitle
      }
    }
</script>
