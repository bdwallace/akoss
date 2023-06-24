
<template>

    <div class="panel">
        <panel-title :title="$route.meta.title">

            <el-button type="primary" icon="plus" size="small" @click="on_dialog_save(0)">创建巡检任务</el-button>
        </panel-title>

        <!-- <div class="panel-body-line" style="clear: both;" v-for="(item, index) in table_data" :key="item"> -->
        <div class="panel-body-line" style="clear: both;">
            <el-table
                :data="table_data"
                    v-loading="load_data"
                    element-loading-text="拼命加载中"
                    border
                    style="width: 100%">
                    
                <el-table-column
                    label="巡检名称"
                    prop="Name"
                    width="160">
                </el-table-column>

                <el-table-column
                    label="巡检地址"
                    prop="Url">
                </el-table-column>

                <el-table-column
                    label="巡检会话"
                    prop="Sessions"
                    width="380">
                </el-table-column>

                <el-table-column
                    label="等待标识"
                    prop="WaitVisible"
                    width="160">
                </el-table-column>
                
                <el-table-column
                    label="链接地址"
                    prop="Link"
                    width="100">
                    <template scope="props">
                        <a target="_blank" :href="props.row.Link">{{props.row.Link}}</a>
    
                        <el-button type="text" size="mini" @click="on_dialog_save(props.row.Id)">修改</el-button>

                        <el-button size="mini" type="text" @click="delete_data(props.row.Id)">删除</el-button>
    
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
   
        <el-dialog :title="dialog_title" 
            :visible.sync="is_saveDialog"
            :modal="true"
            :show-close="false"
            :modal-append-to-body="false">
            <!-- <save_link></save_link> -->
            <el-form ref="form" :model="form" :rules="rules">

                <el-form-item label="巡检名称:" prop="Name" label-width="120px">
                    <el-input v-model="form.Name" placeholder="请输入名称，非必填"
                            style="width: 400px;">
                    </el-input>
                </el-form-item>

                <el-form-item label="巡检地址:" prop="Url" label-width="120px">
                    <el-input v-model="form.Url" placeholder="请输入地址，非必填"
                            style="width: 400px;">
                    </el-input>
                </el-form-item>

                <el-form-item label="巡检会话:" prop="Sessions" label-width="120px">
                    <el-input v-model="form.Sessions" 
                            placeholder="请输入会话，如：grafana_session=7f713c439d00a7e94a98de8cf61d41b3"
                            style="width: 400px;">
                    </el-input>
                </el-form-item>

                <el-form-item label="巡检等待标识:" prop="WaitVisible" label-width="120px">
                    <el-input v-model="form.WaitVisible" 
                            placeholder="请输入等待标识，如：.flot-overlay"
                            style="width: 400px;">
                    </el-input>
                </el-form-item>

                <el-form-item>
                <el-button type="primary" @click="on_submit_form" :loading="on_submit_loading">立即提交</el-button>
                <el-button @click="is_saveDialog=false">取消</el-button>
                </el-form-item>

            </el-form>

        </el-dialog>
    </div>
</template>
<script type="text/javascript">
    import {panelTitle, bottomToolBar, search} from 'components'
    import {port_grafana} from 'common/port_uri'
    import store from 'store'
    
    export default{
        data(){
            return {
                table_data: [],
                
                dialog_title: "",
                is_saveDialog: false,
                
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

                form: {
          	        Id: 0,
	                Name: null,
	                Url: null,
                    Sessions: null,
                    WaitVisible: null,
                },

                rules: {
                    Name: [{required: true, message: '链接名称不能为空', trigger: 'blur'}],
                    Url: [{required: true, message: '链接地址不能为空', trigger: 'blur'}],
                },
                on_submit_loading: false,
            }
        },
        components: {
            panelTitle,
            bottomToolBar,
            search,
        },
        created(){
            this.get_table_data()
        },
        methods: {

            init_data() {
                this.table_project_index = []
                this.table_project_data = []
                this.table_public = {
                    name: "共同",
                    data: []
                }
                
                // cols: [],
                this.table_platform_index = []
                this.table_platform_data = []
            },

            //获取数据
            get_table_data(){
                this.load_data = true
                this.init_data()
                this.$http.get(port_grafana.list, {
                        params: {
                            page: this.currentPage,
                            length: this.length,
                        }
                })
                .then(({data: {data}}) => {
                    // this.table_data = data 
                    this.table_data = data.table_data
                    this.currentPage = data.currentPage
                    this.total = data.total

                    this.load_data = false
                })
                .catch(() => {
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
                    // this.load_data = true
                    this.$http.delete(port_grafana.id, {
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

            on_dialog_save(id) {
                if (id == 0) {
                    this.dialog_title = "新增链接"
                    this.init_form()
                } else {
                    this.dialog_title = "修改链接"
                    this.get_form_data(id)
                }
                this.is_saveDialog = true
            },

            init_form() {
                this.form = {
          	        Id: 0,
	                Name: null,
	                Url: null,
                    Sessions: null,
                    WaitVisible: null,
                }
            },

            //获取数据
            get_form_data(id) {
                // this.load_data = true
                this.$http.get(port_grafana.id, {
                        params: {
                            id: id
                        }
                })
                .then(({data: {data}}) => {
                    this.form = data
                    this.load_data = false
                })
                .catch(() => {
                    this.load_data = false
                })
            },

            //提交
            on_submit_form() {
                this.$http.post(port_grafana.save, this.form)
                .then(({data: {msg}}) => {
                    this.$message({
                        message: msg,
                        type: 'success'
                    })
                    this.get_table_data()
                    this.is_saveDialog = false
                })
                .catch(() => {
                    // this.is_saveDialog = false
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
