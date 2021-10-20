
<template>

    <div class="panel">
        <panel-title :title="$route.meta.title">
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
                    width="90">
                </el-table-column>

                <el-table-column
                    label="计划任务ID"
                    prop="Crontab.Id"
                    width="100">
                </el-table-column>

                <el-table-column
                    label="计划任务名"
                    prop="CrontabName"
                    width="140">
                </el-table-column>

                <el-table-column
                    label="是否自动"
                    prop="Auto"
                    width="100">
                    <template scope="props">
                        <el-tag v-if="props.row.Auto == 1" type="primary">自动</el-tag>
                        <el-tag v-if="props.row.Auto == 0" type="gray">手动</el-tag>
                    </template>
                </el-table-column>

                <el-table-column
                    label="执行时间"
                    prop="CreatedAt"
                    width="220">
                </el-table-column>

                <el-table-column
                    label="执行时间"
                    prop="Result">
                </el-table-column>

                <!-- <el-table-column
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
                        <el-button v-if="props.row.Status == '0'" type="success" size="mini" v-popover:result>正常</el-button>
                        <el-button v-else-if="props.row.Status == '1'" type="danger" size="mini" v-popover:result>失败</el-button>
                        <el-button v-else type="warning" size="mini" v-popover:result>异常</el-button>
                    </template>
                </el-table-column> -->


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

            //获取数据
            get_table_data(){
                this.load_data = true
                this.$http.get(port_crontab.listlog, {
                            params: {
                                page: this.currentPage,
                                length: this.length
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

            //复制项目
            execute(id){
                this.$confirm('确定要执行此任务?', '提示', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning'
                }).then(() => {
                    this.load_data = true
                    this.$http.get(port_crontab.execute, {
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


            //复制项目
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
                    this.$http.get(port_crontab.del, {
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
