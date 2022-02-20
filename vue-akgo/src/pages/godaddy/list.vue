<template>
  <el-form ref="form" :model="Form" >
    <div class="panel">
      <panel-title :title="$route.meta.title">
      </panel-title>
    </div>

    <div class="panel" style="margin-top:5px;margin-bottom:15px;padding:15px;weight:500px">
      <el-form-item label="godaddy key / secret:" label-width="200px" >
        <el-input v-model="Form.GodaddyKey" placeholder="godaddy key" style="width: 400px;"></el-input>
        <el-input v-model="Form.GodaddySecret" placeholder="godaddy secret" style="width: 400px;"></el-input>
      </el-form-item>

      <el-form-item label="godaddy 域名搜索:" label-width="200px">
        <el-input v-model="Form.SearchDomain" placeholder="domain name" style="width: 400px;"></el-input>
        <el-button type="info" size="medium" @click="domain_search()"> 查询 </el-button>
      </el-form-item>

      <div class="panel-body-line" style="clear: both;">
        <el-table
          :data="this.Domains"
          @expand="expand"
          element-loading-text="拼命加载中"
          :default-sort = "{prop: 'Id', order: 'ascending'}"
          border
          :row-key="getRowKeys"
          :expand-row-keys="expands"
          @selection-change="handleSelectionChange"
          style="width: 100%;">

          <el-table-column
            type="selection"
            width="40">
          </el-table-column>

          <el-table-column
            prop="Id"
            label="ID"
            align="center"
            width="80">
          </el-table-column>

          <el-table-column
            prop="Domain"
            label="域名"
            align="center">
          </el-table-column>

          <el-table-column
            prop="Expires"
            label="过期时间"
            align="center">
          </el-table-column>

          <el-table-column
            prop="Status"
            label="状态"
            align="center">
          </el-table-column>

          <el-table-column
            label="操作"
            align="center"
            width="185">
            <template scope="item_props">
              <el-button type="info" size="small" @click="get_info()">查看</el-button>
            </template>
          </el-table-column>

        </el-table>
      </div>

      <bottom-tool-bar>
        <div slot="page">
          <el-pagination
            @current-change="handleCurrentChange"
            :current-page="PageNum"
            :page-size="PageSize"
            layout="total, prev, pager, next"
            :total="Total">
          </el-pagination>
        </div>
      </bottom-tool-bar>

    </div>
  </el-form>
</template>


<script type="text/javascript">
  import {panelTitle} from 'components'
  import store from 'store'

  export default {
    data() {
      return {
        ProjectId: store.state.user_info.user.ProjectId,
        Form: {
          Project: {
            Id: store.state.user_info.user.ProjectId,
          },
          GodaddyKey: "",
          GodaddySecret: "",

          SearchDomain: "",
        },
        selection: [],
        load_data: true,
        expands:[],

        PageNum: 1,
        PageSize: 50,
        Total: 0,
        Domains: [],

        BackupRes: "",

      }
    },
    created() {
    },
    methods: {
      domain_search(){
        this.load_data = true
        this.$http.get("/api/godaddy", {
          params: {
            godaddy_key: this.Form.GodaddyKey,
            godaddy_secret: this.Form.GodaddySecret,
            search_domain: this.Form.SearchDomain,
            page_size: this.PageSize,
            page_number: this.PageNum,
          }
        })
          .then(({data: {data}}) => {
            this.load_data = true
            this.Domains = data.domains
            this.Total = data.item_total
            this.PageNum = data.page_num
            this.load_data = false
          })
          .catch(() => {
            this.load_data = false
          })
      },

      get_info(row){
        let router_url = this.$router.resolve({
          path: "/godaddy/save",
          query: row,
        })
        window.open(router_url.href, '_blank');
      },

      expand(row,expanded) {
        if (expanded){
        } else {
          let expands_index = this.expands.indexOf(row)
          this.expands.splice(expands_index, 1)
        }
      },

      getRowKeys(row) {
        return row.Id;
      },

      //页码选择
      handleCurrentChange(val) {
        this.PageNum = val
        this.domain_search()
      },

      handleSelectionChange(val) {
        this.selection = val;
      },

    },
    components: {
      panelTitle,
    }
  }
</script>
