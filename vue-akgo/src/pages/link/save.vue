<template>
  <div class="panel">
    <panel-title :title="$route.meta.title"></panel-title>
    <div class="panel-body"
         v-loading="load_data"
         element-loading-text="拼命加载中">
      <el-row>
        <el-col :span="24">
          <el-form ref="form" :model="form" :rules="rules">

            <el-form-item label="链接名称:" prop="Name" label-width="120px">
              <el-input v-model="form.Name" placeholder="请输入链接名称，非必填"
                        style="width: 600px;"></el-input>
            </el-form-item>

            <!-- <el-form-item label="说明:" prop="Direction" label-width="120px">
              <el-input v-model="form.Direction" placeholder="请输入链接说明，非必填"
                        style="width: 600px;"></el-input>
            </el-form-item> -->

            <el-form-item label="地址:" prop="Link" label-width="120px">
              <el-input v-model="form.Link" placeholder="请输入链接地址，如：www.baidu.com"
                        style="width: 600px;"></el-input>
            </el-form-item>

            <!-- <el-form-item label="选择环境:" prop="Projects" label-width="120px">
              <el-checkbox-group v-model="form.Projects" :max="1">
              <el-checkbox v-for="item in itemProject" :key="item.Id" :label="item">{{item.Name}}</el-checkbox>
              </el-checkbox-group>
            </el-form-item> -->

            <el-form-item label="选择环境:" prop="Projects" label-width="120px">
              <el-radio-group v-model="form.Projects[0].Id" :max="1" @change="get_itemPlatform()">
              <el-radio v-for="item in itemProject" :key="item.Id" :label="item.Id">{{item.Name}}</el-radio>
              </el-radio-group>
            </el-form-item>

            <el-form-item v-if="form.Projects[0] && form.Projects[0].Id != null" label="选择平台:" prop="Project" label-width="120px">
              <el-select 
                  v-if="form.Platforms[0]"
                  v-model="form.Platforms[0].Id" 
                  clearable
                  placeholder="请选择"
                  @change="change_itemPlatform()">
                <el-option
                  v-for="item in itemPlatform"
                  :key="item.Id"
                  :label="item.Name"
                  :value="item.Id">
                </el-option>
              </el-select>
              <el-select 
                  v-if="form.Platforms[0] && form.Platforms[0].Id != null"
                  v-model="form.Class" 
                  placeholder="请选择">
                <el-option
                  v-for="item in itemClass"
                  :key="item"
                  :label="item"
                  :value="item">
                </el-option>
              </el-select>
            </el-form-item>

            <!-- <el-radio-group v-model="form.Platforms[0].Id" :max="1">
            <el-radio v-for="item in itemPlatform" :key="item.Id" :label="item.Id">{{item.Name}}</el-radio>
            </el-radio-group> -->


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
  import {port_link, port_project, port_platform} from 'common/port_uri'
  import store from 'store'

  export default {
    data() {
      return {
        form: {
          	Id: 0,
	          Name: null,
	          Direction: null,
            Link: null,
            Class: null,
            Projects: [{Id:null}],
            Platforms: [{Id:null}]
            // Projects: []
        },
        itemProject: [],
        itemPlatform: [],
        itemClass: [
          "h5", 
          "merchant",
          "download", 
          "download-share", 
          "agent", 
          "share-agent", 
          "customer", 
          "chat-backend", 
        ],
        route_id: this.$route.params.id,
        load_data: false,
        on_submit_loading: false,
        rules: {
          // Name: [{required: true, message: '链接名称不能为空', trigger: 'blur'}],
          Link: [{required: true, message: '链接地址不能为空', trigger: 'blur'}],
          // Direction: [{required: true, message: '链接说明不能为空', trigger: 'blur'}],
        }
      }
    },
    created() {
      this.get_itemProject()
      if (this.route_id) {
        this.get_form_data()
      }
    },
    methods: {
        // 下拉框获得所有环境
        get_itemProject() {
            this.$http.get(port_project.list)
                    .then(({data: {data}}) => {
                this.itemProject = data
                this.itemProject.unshift({Id: null, Name: "共同环境"})
            })
        },
        
        // 下拉框获得选择的环境所有平台
        get_itemPlatform(){
          this.load_data = true
          if(this.form.Projects[0] && this.form.Projects[0].Id == null) {
              this.form.Platforms[0] = {Id: null}
              this.load_data = false
              return
          }
          this.$http.get(port_platform.list, {
                              params: {
                                  project_id: this.form.Projects[0].Id,
                              }
                          })
            .then(({data: {data}}) => {
              this.itemPlatform = data
              this.form.Platforms[0].Id = this.itemPlatform[0].Id
              // this.$set(this.form.Platforms[0], "Id", this.itemPlatform[0].Id)
              this.load_data = false
            })
            .catch(() => {
              this.load_data = false
            })
        },

        change_itemPlatform() {
          if (this.form.Class == "") {
            this.form.Class = this.itemClass[0]
          }
        },

        //获取数据
        get_form_data() {
            this.load_data = true
            this.$http.get(port_link.id, {
              params: {
                  id: this.route_id
              }
            })
            .then(({data: {data}}) => {
                this.form = data
                if (this.form.Projects.length == 0) {
                  this.form.Projects.push({Id: null})
                }
                if (this.form.Platforms.length == 0) {
                  this.form.Platforms.push({Id: null})
                }
                this.load_data = false
            })
            .catch(() => {
                this.load_data = false
            })
        },

        //提交
        on_submit_form() {
            this.$http.post(port_link.save, this.form)
            .then(({data: {msg}}) => {
              this.on_submit_loading = false
              this.$message({
                  message: msg,
                  type: 'success'
              })
              // setTimeout(() => {
              //   this.$router.back()
              // }, 500)
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
