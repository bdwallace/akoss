<template>
  <el-form ref="form" :model="form" >
    <div class="panel">
      <panel-title :title="$route.meta.title">
      </panel-title>
    </div>
      <div class="panel" style="margin-top:5px;margin-bottom:15px;padding:15px;weight:500px">
        <el-form-item label="选择主机/服务类型:" label-width="150px">
            <el-button @click.stop="get_itemHost()" size="mini">
              <i class="fa fa-refresh"></i>
            </el-button>
            <el-select v-model="form.Host" placeholder="选择主机"
                       filterable
                       clearable
                       value-key="Id"
                       style="width: 400px;">
              <el-option
                v-for="item in itemHost"
                :key="item"
                :label="`${item.Name} / ${item.UseIp}`"
                :value="item">
              </el-option>
            </el-select>

          <el-select v-model="form.ServiceClass" placeholder="选择服务类型"
                     filterable
                     clearable
                     value-key="Id"
                     style="width: 300px;">
            <el-option
              v-for="item in itemClass"
              :key="item"
              :label="item"
              :value="item">
            </el-option>
          </el-select>
          <el-button type="info" size="medium" icon="search" @click="get_servcie_and_host()">查询</el-button>
        </el-form-item>
      </div>


<!--    <div class="panel-body">-->

      <div class="panel-body-line" style="clear: both;">
      <el-table
        :data="this.ResConf"
        v-loading="load_data"
        @expand="expand"
        element-loading-text="拼命加载中"
        :default-sort = "{prop: 'ServiceId', order: 'ascending'}"
        border
        :row-key="getRowKeys"
        :expand-row-keys="expands"
        style="width: 100%;">
        <el-table-column type="expand">
          <template scope="props">

            <!-- @select="select"
            @select-all="select_all(props.$index)" -->
            <el-table
              :data="props.row.Confs"
              v-loading="load_data"
              element-loading-text="拼命加载中"
              @selection-change="(selection)=>select_change(selection, props.$index)"
              border
              style="width: 100%;"
              :ref="checkTable[props]">

              <el-table-column
                width="120">

              </el-table-column>
              <el-table-column
                prop="conf_file_name"
                label="文件">
              </el-table-column>

              <el-table-column
                label="操作"
                align="center"
                width="185">
                <template scope="conf_item_props">
                  <el-button type="warning" size="small" icon="edit" @click="on_submit_form(conf_item_props.row)" >修改</el-button>
                </template>
              </el-table-column>

            </el-table>
          </template>
        </el-table-column>

        <el-table-column
          prop="ServiceId"
          label="服务ID"
          align="center"
          width="80">
        </el-table-column>

        <el-table-column
          prop="ServiceName"
          label="服务名称"
          align="center">
        </el-table-column>

        <el-table-column
          prop="HostUseIp"
          label="主机IP"
          align="center">
        </el-table-column>


        <el-table-column
          label="操作"
          align="center"
          width="200">
<!--          <template scope="props">
            <router-link :to="{name: 'register', query:  {id: props.row.id}}"
                         tag="span">
              <el-button type="info" size="medium" >查看</el-button>
            </router-link>-->
<!--          </template>-->
        </el-table-column>
      </el-table>
    </div>

  </el-form>
</template>


<script type="text/javascript">
  import {panelTitle} from 'components'
  import {port_platform, port_domain, port_conf, port_host, port_service,port_project,port_confmanage} from 'common/port_uri'
  import store from 'store'
  import log from "../crontab/log";

  export default {
    data() {
      return {
        ProjectId: store.state.user_info.user.ProjectId,
        form: {
          Project: {
            Id: store.state.user_info.user.ProjectId,
          },
          Host: "",
          ServiceClass: "",

        },
        edit: [],
        ResConf:[],
        on_submit_loading: false,
        itemClass: [
          "h5",
          "h5-proxy",
          "h5-site",
          "merchant",
          "download",
          "download-share",
          "agent",
          "share-agent",
          "customer",
          "chat-backend",
          "other"
        ],

        itemHost: [],
        load_data: true,
        expands:[],
        checkTable: [],

      }
    },
    created() {
      this.get_itemHost()
    },
    methods: {

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
          .catch(() => {
            this.load_data = false
          })
      },

      get_servcie_and_host(){
        this.load_data = true
        // this.$http.get(port_confmanage.search, {
        this.$http.get("/api/confManage/searchServiceHost", {
          params: {
            project_id: this.form.Project.Id,
            host: this.form.Host.UseIp,
            service_class: this.form.ServiceClass,
          }
        })
          .then(({data: {data}}) => {
            this.load_data = true
            this.ResConf = data
            this.load_data = false

          })
          .catch(() => {
            this.load_data = false
          })
      },

      expand(row,expanded) {
        // let index = this.ResConf.indexOf(row)
        if (expanded){
        } else {
          let expands_index = this.expands.indexOf(row)
          this.expands.splice(expands_index, 1)
        }
      },

      getRowKeys(row) {
        return row.Id;
      },
      select_change(selection, index) {
        this.checkTable[index] = selection
      },

      on_submit_form(row) {
        let router_url = this.$router.resolve({
          path: "/confManage/nginx",
          query: row,
        })
        window.open(router_url.href, '_blank');
        }
    },
    components: {
      panelTitle,
    }
  }
</script>
