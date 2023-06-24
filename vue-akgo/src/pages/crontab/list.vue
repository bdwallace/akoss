
<template>

    <div class="panel">
        <panel-title :title="$route.meta.title">
            <b style="color: red">新增、修改计划任务需要重新加载计划任务</b>

            <el-button type="success" size="small" icon="setting" @click="reload_crontab()">重新加载</el-button>

            <router-link :to="{name: 'crontabAdd'}" tag="span">
                <el-button type="primary" icon="plus" size="small">创建计划任务</el-button>
            </router-link>
        </panel-title>

        <div class="panel-body-line" style="clear: both;">
            <el-table
                :data="table_data"
                    v-loading="load_data"
                    element-loading-text="拼命加载中"
                    border
                    style="width: 100%;">
                <el-table-column
                    label="Id"
                    prop="Id"
                    width="60">
                </el-table-column>

                <el-table-column
                    label="计划任务名"
                    prop="Name"
                    width="180">
                </el-table-column>

                <el-table-column
                    label="说明"
                    prop="Direction"
                    width="160">
                </el-table-column>

                <el-table-column
                    label="计划周期"
                    prop="Crontab"
                    width="140">
                </el-table-column>

                <el-table-column
                    label="自动"
                    prop="Auto"
                    width="50">
                    <template scope="props">
                        <el-tag v-if="props.row.Auto == 1" type="primary">是</el-tag>
                        <el-tag v-if="props.row.Auto == 0" type="gray">否</el-tag>
                    </template>
                </el-table-column>

                <el-table-column
                    label="脚本"
                    prop="Script">
                </el-table-column>
                
                <el-table-column
                    label="最后执行时间"
                    prop="UpdatedAt"
                    width="220"
                    :formatter="dateFormat">
                </el-table-column>
                
                <el-table-column
                    label="状态"
                    prop="Status"
                    width="80">
                    <template scope="props">
                        <el-popover
                            ref="result"
                            placement="left-start"
                            width="600"
                            trigger="click">
                            <span style="color:teal;white-space: pre-wrap;">{{props.row.Result}}</span>
                        </el-popover>
                        <!-- <el-button v-show="props.row.crt != ''" type="info" style="clear: both;margin:-10px" size="mini" v-popover:crt>查看</el-button> -->
                        <el-button v-if="props.row.Status == '0'" type="success" size="mini" v-popover:result>正常</el-button>
                        <el-button v-else-if="props.row.Status == '1'" type="danger" size="mini" v-popover:result>失败</el-button>
                        <el-button v-else type="warning" size="mini" v-popover:result>异常</el-button>
                    </template>
                </el-table-column>

                <el-table-column
                    label="操作"
                    width="120">
                    <template scope="props">
                        <el-button type="success" size="small" @click="execute(props.row.Id)">立即执行</el-button>

                    </template>
                </el-table-column>

                <el-table-column
                    label="操作"
                    width="240">
                    <template scope="props">
                        <router-link :to="{name: 'crontabUpdate', params: {id: props.row.Id}}" load_data=false tag="span">
                            <el-button type="info" size="small" icon="edit">修改</el-button>
                        </router-link>

                        <el-button type="warning" size="small" icon="document" @click="copy_data(props.row.Id)">复制
                        </el-button>

                        <el-button style="margin-left:0px" type="danger" size="small" icon="delete" @click="delete_data(props.row.Id)">删除
                        </el-button>
                    </template>
                </el-table-column>
            </el-table>


            <!-- <bottom-tool-bar>
                <div slot="page">
                    <el-pagination
                            @current-change="handleCurrentChange"
                            :current-page="currentPage"
                            :page-size="15"
                            layout="total, prev, pager, next"
                            :total="total">
                    </el-pagination>
                </div>
            </bottom-tool-bar> -->
        </div>
    </div>
</template>
<script type="text/javascript">
    import {panelTitle, bottomToolBar, search} from 'components'
    import {port_crontab} from 'common/port_uri'
    import store from 'store'
    
    export default{
        data(){
            return {
                level: store.state.user_info.user.Level,
                // level_data: [],
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
                select_info: "",
                //项目详情
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
            // submit_search(value) {
            //     this.select_info = value
            //     this.$message({
            //         message: value,
            //         type: 'success'
            //     })
            //     this.get_table_data()
            // },

            // on_refresh(){
            //     this.select_info = ""
            //     this.get_table_data()
            // },

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

            //获取数据
            get_table_data(){
                this.load_data = true
                this.$http.get(port_crontab.list)
                            .then(({data: {data}}) => {
                    this.table_data = data 
                    this.load_data = false
                })
                .catch(() => {
                    this.load_data = false
                })
            },

            //复制项目
            execute(id){
                this.$confirm('确定要执行此任务?', '提示', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning'
                }).then(() => {
                    this.load_data = true
                    this.$http.get(port_crontab.exec, {
                                params: {
                                    crontabId: id,
                                }
                            })
                    .then(({data: {msg}}) => {
                        this.$message({
                            message: msg,
                            type: 'success'
                        })
                        this.get_table_data()
                    })
                    .catch(() => {
                        this.load_data = false
                    })
                }).catch(() => {
                    this.load_data = false
                })
            },

            //重载计划任务
            reload_crontab(){
                this.$confirm('确定要重新加载计划任务吗?', '提示', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning'
                }).then(() => {
                    this.load_data = true
                    this.$http.get(port_crontab.reload)
                    .then(({data: {msg}}) => {
                        this.$message({
                            message: msg,
                            type: 'success'
                        })
                        this.load_data = false
                    })
                    .catch(() => {
                        this.load_data = false
                    })
                }).catch(() => {
                    this.load_data = false
                })
            },
 
            //复制项目
            copy_data(id){
                this.$confirm('是否复制项目?', '提示', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning'
                }).then(() => {
                    this.load_data = true
                    this.$http.get(port_crontab.copy, {
                                params: {
                                    crontabId: id,
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
                }).catch(() => {
                    this.load_data = false
                })
            },

            //根据id删除数据
            delete_data(id){
                this.$confirm('此操作将删除该数据, 是否继续?', '提示', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning'
                }).then(() => {
                    this.load_data = true
                    this.$http.delete(port_crontab.id, {
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
                    }).catch(() => {
                        this.load_data = false
                    })
                }).catch(() => {
                    this.load_data = false
                })
            },

            //页码选择
            handleCurrentChange(val) {
                this.currentPage = val
                this.get_table_data()
            }
        }
    }
</script>
