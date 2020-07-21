<template>
    <div class="panel">
        <panel-title :title="$route.meta.title">
            <!-- <div style="float: left;margin-right: 10px;margin-top: 5px">
                <search @search="submit_search"></search>
            </div>
            <el-button @click.stop="on_refresh" size="small">
                <i class="fa fa-refresh"></i>
            </el-button> -->
            <router-link :to="{name: 'projectAdd'}" tag="span">
                <el-button type="primary" icon="plus" size="small">创建项目</el-button>
            </router-link>

        </panel-title>
        <div class="panel-body-line" style="clear: both;">
            <el-table
                    :data="table_data"
                    v-loading="load_data"
                    element-loading-text="拼命加载中"
                    border
                    :row-class-name="tableRowClassName"
                    style="width: 100%;">
                <el-table-column
                        prop="Id"
                        label="id"
                        width="70">
                </el-table-column>

                <el-table-column
                        prop="Name"
                        label="项目名称"
                        width="190">
                </el-table-column>

                <el-table-column
                        prop="Alias"
                        label="版本">
                </el-table-column>

                <el-table-column
                        prop="CreatedAt"
                        label="创建时间"
                        width="180"
                        :formatter="dateFormat">
                </el-table-column>

                <el-table-column
                        label="操作">
                <template scope="props">
                        <!-- <el-button type="info" v-popover:popover4 size="small" icon="search" @click="open(props.row.Id)">查看
                        </el-button> -->
                        <!-- <router-link :to="{name: 'confDetection', params: {id: props.row.Id}}" tag="span">
                            <el-button type="success" size="small" icon="setting">检测</el-button>
                        </router-link> -->
                        <router-link :to="{name: 'projectUpdate', params: {id: props.row.Id}}" tag="span">
                            <el-button type="info" size="small" icon="edit">修改</el-button>
                        </router-link>
                        <!-- <el-button type="warning" size="small" icon="document" @click="copy_data(props.row.id)">复制
                        </el-button> -->
                        <el-button style="margin-left:0px" type="danger" size="small" icon="delete" @click="delete_data(props.row.Id)">删除
                        </el-button>
                    </template>
                </el-table-column>

            </el-table>
        </div>
    </div>
</template>

<style>
  .el-table .positive-row {
    background: #13CE66;
  }
</style>

<script type="text/javascript">
    import {panelTitle, bottomToolBar, search} from 'components'
    import {port_project} from 'common/port_uri'
    import {tools_date} from 'common/tools'
    import store from "store"

    export default{
        data(){
            return {
                ProjectId: store.state.user_info.user.ProjectId, 
                table_data: [],
                load_data: false
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
            tableRowClassName(row, index) {
                if (row.Id === this.ProjectId) {
                    return 'positive-row';
                }
                return '';
            },

            dateFormat(row, column, cellValue, format='YY-MM-DD hh:mm:ss'){
                var date = new Date(cellValue);
 
                var year = date.getFullYear(),
                    month = date.getMonth()+1,//月份是从0开始的
                    day = date.getDate(),
                    hour = date.getHours(),
                    min = date.getMinutes(),
                    sec = date.getSeconds();
                var preArr = Array.apply(null,Array(10)).map(function(elem, index) {
                    return '0'+index;
                });//开个长度为10的数组 格式为 ["00", "01", "02", "03", "04", "05", "06", "07", "08", "09"]
 
                var newTime = format.replace(/YY/g,year)
                    .replace(/MM/g,preArr[month]||month)
                    .replace(/DD/g,preArr[day]||day)
                    .replace(/hh/g,preArr[hour]||hour)
                    .replace(/mm/g,preArr[min]||min)
                    .replace(/ss/g,preArr[sec]||sec);
 
                return newTime
            },

            get_table_data() {
                this.load_data = true
                this.$http.get(port_project.list)
                        .then(({data: {data}}) => {
                    this.table_data = data
                    this.load_data = false
                })
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
                    this.$http.delete(port_project.id, {
                            params: {
                                id: id,
                            }
                        })
                    .then(({data: {msg}}) => {
                        this.get_table_data()
                        this.$message({
                            message: msg,
                            type: 'success'
                        })
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
<style lang="scss" type="text/css" rel="stylesheet/scss">

</style>
