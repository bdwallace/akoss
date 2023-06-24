<template>

    <div class="panel">
        <panel-title :title="$route.meta.title">
            <div style="float: left;margin-right: 10px">
                <search @search="submit_search"></search>
            </div>
            <el-button @click.stop="on_refresh" size="small">
                <i class="fa fa-refresh"></i>
            </el-button>
        </panel-title>

        <div class="panel-body"  style="clear: both;">
            <el-table
                    :data="table_data"
                    v-loading="load_data"
                    element-loading-text="拼命加载中"
                    border
                    style="width: 100%;">
                <el-table-column
                        prop="Id"
                        label="id"
                        width="80">
                </el-table-column>

                <el-table-column
                        prop="User.Username"
                        label="创建人"
                        width="90">
                </el-table-column>

                <!-- <el-table-column
                        prop="Title"
                        label="上线单标题">
                </el-table-column> -->

                <el-table-column
                        prop="CreatedAt"
                        label="上线时间"
                        width="170"
                        :formatter="dateFormat">
                </el-table-column>

                <el-table-column>
                    <el-table-column 
                        width="220"
                        label="服务名">
                        <template scope="props">
                            <!-- <el-table :data="props.row.Tasks" :show-header=false style="border: none;border-right: none">
                                <el-table-column :border="false" style="padding: 0px;margin: 0px;border: none;"
                                    prop="Service.Name"
                                    label="服务名">
                                </el-table-column> 
                            </el-table> -->
                                <table style="width: 100%" rules="rows">
                                <tr v-for="item in props.row.Tasks" :key="item">
                                    <td>
                                        {{ item.Service.Name }}
                                    </td>
                                </tr>
                                </table>
                        </template>
                    </el-table-column>

                    <el-table-column
                        label="版本"
                        width="260">
                        <template scope="props">
                            <!-- <el-table class="panel-body-none" :data="props.row.Tasks" style="padding: 0px;clear: both;border: none;" :show-header=false>
                                <el-table-column
                                    prop="Tag">
                                </el-table-column> 
                            </el-table> -->
                                <table style="width: 100%" rules="rows">
                                <tr v-for="item in props.row.Tasks" :key="item">
                                    <td>
                                        {{ item.Tag }}
                                    </td>
                                </tr>
                                </table>
                        </template>
                    </el-table-column> 

                    <el-table-column
                        label="主机">
                        <template scope="props">
                            <!-- <span v-for="item in props.row.Tasks" :key="item"> {{ item.Service.Hosts.PrivateIp }}, </span> -->
                            <!-- <el-table :data="props.row.Tasks" border class="inline" style="padding: 0px;margin: 0px;clear: both;border: none;" :show-header=false>
                                <el-table-column
                                    prop="Hosts"
                                    :show-overflow-tooltip="false">
                                    <template scope="propshost">
                                        <span v-for="host in propshost.row.Service.Hosts" :key="host"> {{ host.PrivateIp }}, </span>
                                    </template>
                                </el-table-column> 
                            </el-table> -->
                                <table style="width: 100%" rules="rows">
                                <tr v-for="item in props.row.Tasks" :key="item">
                                    <td width=100%>
                                        <span v-for="host in item.Service.Hosts" :key="host"> {{ host.PrivateIp}}, </span>
                                    </td>
                                </tr>
                                </table>
                        </template>
                    </el-table-column> 
                </el-table-column>

                <el-table-column
                        label="操作"
                        width="160">
                    <template scope="props">
                        <router-link :to="{name: 'deployRelease', params: {id: props.row.Id}}">
                            <el-button type="info" size="small" icon="edit">上线</el-button>
                        </router-link>


                        <el-button type="danger" size="small" icon="delete" v-if="props.row.Status==0" @click="delete_data(props.row.Id)">删除
                        </el-button>
                    </template>
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

<style>
    .el-table {
        border: 0px solid #ffffff; 
    }
    .el-table__body-wrapper {
        overflow: hidden;
    }
    .el-table--border, .el-table--group{
        border: none;
    }
    .el-table__header-wrapper th:nth-last-of-type(2){
        border-right: none;
    }
    .el-table--border::after, .el-table--group::after{
        width: 0;
    }
</style>

<script type="text/javascript">
    import {panelTitle, bottomToolBar, search} from 'components'
    import {port_deploy} from 'common/port_uri'
    import store from "store"

    export default{
        data(){
            return {
                ProjectId: store.state.user_info.user.ProjectId, 
                route_id: this.$route.params.id,
                table_data: [],
                //当前页码
                currentPage: 1,
                //数据总条目
                total: 0,
                //每页显示多少条数据
                length: 15,
                //请求时的loading效果
                load_data: true,
                //批量选择数组
                batch_select: [],
                //批量选择数组
                select_info: ""
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
            submit_search(value) {
                this.select_info = value
                this.$message({
                    message: value,
                    type: 'success'
                })
                this.get_table_data()
            },
            //刷新
            on_refresh(){
                this.select_info = ""
                this.get_table_data()
            },
            //获取数据
            get_table_data(){
                this.load_data = true
                this.$http.get(port_deploy.listuser, {
                            params: {
                                project_id: this.ProjectId,
                                user_id: this.route_id,
                                page: this.currentPage,
                                length: this.length,
                                // select_info: this.select_info
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

            //根据id删除数据
            delete_data(id){
                this.$confirm('此操作将删除该数据, 是否继续?', '提示', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning'
                }).then(() => {
                    this.load_data = true
                this.$http.delete(port_deploy.id, {
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
                this.load_data = false
            }).
                catch(() => {
                    this.load_data = false
            })
            }).
                catch(() => {
                    this.load_data = false
            })
            },
            
            //页码选择
            handleCurrentChange(val) {
                this.currentPage = val
                this.get_table_data()
            },
 
        }
    }
</script>
