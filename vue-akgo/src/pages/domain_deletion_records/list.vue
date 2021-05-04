
<template>

  <div class="panel">
    <panel-title :title="$route.meta.title">

      <div style="float: left;margin-right: 10px;margin-top: 5px">
        <search @search="submit_search"></search>
      </div>

      <el-button @click.stop="on_refresh" size="small">
        <i class="fa fa-refresh"></i>
      </el-button>

    </panel-title>

    <div class="panel-body-line" style="clear: both;">
      <el-table
        ref="filterTable"
        :data="table_data"
        v-loading="load_data"
        element-loading-text="拼命加载中"
        @filter-change="filterChange"
        border
        style="width: 100%;">
        <el-table-column
          prop="Id"
          label="id"
          width="80"
          column-key="Id">
        </el-table-column>

        <el-table-column
          prop="Domain"
          label="域名"
          column-key="Domain">
        </el-table-column>

        <el-table-column
          prop="Class"
          label="类型"
          column-key="Class"
          :filters="itemClass">
        </el-table-column>

        <el-table-column
          prop="Comment"
          label="备注"
          column-key="Comment">
        </el-table-column>

        <el-table-column
          prop="UserName"
          label="用户"
          column-key="UserName">
        </el-table-column>
      </el-table>

      <bottom-tool-bar>
        <div slot="page">
          <el-pagination
            @current-change="handleCurrentChange"
            :current-page="currentPage"
            :page-size="15"
            layout="total, prev, pager, next"
            :total="total">
          </el-pagination>
        </div>
      </bottom-tool-bar>
    </div>
  </div>
</template>
<script type="text/javascript">
  import {panelTitle, bottomToolBar, search} from 'components'
  import {port_domain} from 'common/port_uri'
  import store from 'store'
  export default{
    data(){
      return {
        ProjectId: store.state.user_info.user.ProjectId,
        // level_data: [],
        table_data: [],
        total: 0,
        //当前页码
        currentPage: 1,
        //数据总条目
        //每页显示多少条数据
        length: 15,
        //请求时的loading效果
        load_data: true,
        //批量选择数组
        search_text: "",
        Class: "NOT_FILTER",
        itemClass: [],
        filter: {
          Class: [],
        }
      }
    },
    components: {
      panelTitle,
      bottomToolBar,
      search
    },
    created(){
      this.get_table_data()
      this.get_domain_class()
    },
    methods: {
      submit_search(value) {
        this.search_text = value
        this.$message({
          message: value,
          type: 'success'
        })
        this.get_table_data()
      },

      on_refresh(){
        this.search_text = ""
        this.get_table_data()
      },

      //获取数据
      get_table_data(){
        this.load_data = true
        this.table_data = []
        this.$http.get(port_domain.domainReco, {
          params: {
            project_id: this.ProjectId,
            page: this.currentPage,
            length: this.length,
            search_text: this.search_text,
            class: this.Class,
          }
        })
          .then(({data: {data}}) => {
            this.table_data = data.table_data
            this.currentPage = data.currentPage
            this.total = data.total
            this.load_data = false
          })
          .catch(() => {
            this.load_data = false
          })
      },

      get_domain_class() {
        this.load_data = true
        this.$http.get(port_domain.domainRecoClass, {
          params: {
            project_id: this.ProjectId,
          }
        })
          .then(({data: {data}}) => {
            for(let i in data) {
              this.itemClass.push({text: data[i], value: data[i]})
            }
            this.load_data = false
          })
          .catch(() => {
            this.load_data = false
          })
      },

      filterChange(filters) {
        Object.assign(this.filter, filters)
        if (this.filter.Class.length == 0) {
          this.Class = "NOT_FILTER"
        } else {
          this.Class = this.filter.Class.join()
        }
        this.get_table_data()
      },

      //页码选择
      handleCurrentChange(val) {
        this.currentPage = val
        this.get_table_data()
      }
    }
  }
</script>
