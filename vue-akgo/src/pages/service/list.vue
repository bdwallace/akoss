
<template>

    <div class="panel">
        <panel-title :title="$route.meta.title">
            <div style="float: left;margin-right: 10px;margin-top: 5px">
                <el-button plain size="small">全部&nbsp;{{table_data.length}}&nbsp;个服务</el-button>
            </div>

            <div style="float: left;margin-right: 10px;margin-top: 5px">
                <el-button type="success" size="small" icon="setting" @click="get_service_status_all()" :loading="on_submit_loading">全部检测</el-button>
            </div>

            <div style="float: left;margin-right: 10px;margin-top: 5px;">
                <el-button type="danger" size="small" icon="setting" @click="on_submit_form()" :loading="on_submit_loading">发布</el-button>
            </div>

            <div style="float: left;margin-right: 10px;margin-top: 5px">
                <search @search="submit_search"></search>
            </div>

            <el-button @click.stop="on_refresh" size="small">
                <i class="fa fa-refresh"></i>
            </el-button>

            <router-link :to="{name: 'serviceAdd'}" tag="span">
                <el-button type="primary" icon="plus" size="small">创建服务</el-button>
            </router-link>
        </panel-title>

        <div class="panel-body" style="clear: both;">
            <el-table
                    :data="table_data"
                    v-loading="load_data"
                    element-loading-text="拼命加载中"
                    border
                    :default-sort = "{prop: 'Id', order: 'ascending'}"
                    style="width: 100%;"
                    @selection-change="handleSelectionChange"
                    ref="multipleTable">
                <el-table-column
                    type="selection">
                </el-table-column>

                <el-table-column
                        prop="Id"
                        sortable
                        min-width="3%">
                </el-table-column>

                <el-table-column
                        prop="Name"
                        label="名称"
                        sortable
                        min-width="15%">
                </el-table-column>

                <el-table-column
                        prop="Tag"
                        sortable
                        label="最新版本"
                        min-width="17%">
                </el-table-column>

                <el-table-column
                        prop="Port"
                        label="端口"
                        min-width="4%">
                    <template scope="props">
                        <font style="display: flex" v-for="item in props.row.Port.split(',')" :key="item">{{item}}</font>
                    </template>
                </el-table-column>

                <el-table-column
                    prop="service_data"
                    label="主机"
                    fit=true
                    min-width="46%">
                    <template scope="props">
                        <div style="clear: both;margin-bottom: 10px;"> <span></span> </div>
                        <div v-for="(item, index) in props.row.service_data" :key="item"  style="clear:both">
                            <div style="display: inline-block;margin-right: 10px" v-if="item.ip_show == item.ip">
                                <el-input readonly style="width: 120px" size="mini" v-model="item.ip_pub" >
                                </el-input>
                                <el-input readonly class="input_blue" style="width: 120px" size="mini" v-model="item.ip" >
                                </el-input>
                            </div>

                            <div style="display: inline-block;margin-right: 10px" v-else>
                                <el-input readonly class="input_blue" style="width: 120px" size="mini" v-model="item.ip_pub" >
                                </el-input>
                                <el-input readonly style="width: 120px" size="mini" v-model="item.ip" >
                                </el-input>
                            </div>

                            <el-input v-if="item.ps_status.substr(0, 2) != 'Up'" class="input_red" readonly style="width: 100px" size="mini" v-model="item.ps_status">
                            </el-input>
                            <el-input v-else readonly style="width: 100px" size="mini" v-model="item.ps_status">
                            </el-input>

                            <el-input readonly style="width: 220px" size="mini" v-model="item.ps_image">
                            </el-input>

                            <div style="display: inline-block;margin-right:10px">
                                <el-tooltip effect="light" :content="item.url" placement="left">
                                    <el-tag v-if="item.health=='200'" size="mini" type="primary" >健康</el-tag>
                                    <el-tag v-else size="mini" type="danger">异常</el-tag>
                                </el-tooltip>
                            </div>

                            <el-button size="mini" type="primary" style="margin-right:-12px" @click="restart(props.row, props.row.Hosts[index])" :loading="on_submit_loading">重启
                            </el-button>

                            <el-button size="mini" type="primary" style="margin-right:-12px" @click="stop(props.row, props.row.Hosts[index])" :loading="on_submit_loading">关闭
                            </el-button>

                        </div>
                        <div style="clear: both;margin-bottom: 10px;"> <span></span> </div>
                    </template>
                </el-table-column>

                <el-table-column
                    label="上下线"
                    fit=true
                    min-width="5%">
                  <!--                    width="80">-->
                    <template scope="props">
                        <div v-for="item in props.row.service_data" :key="item"  style="clear:both">
                            <el-switch v-model="item.line" v-if="props.row.Class == 'java'" size="mini"
                                on-color="#13ce66"
                                off-color="#ff4949"
                                on-value="on"
                                off-value="off"
                                on-text="上线"
                                off-text="下线"
                                @change="linechange(props.row, item.ip, item.line, item.service_id)">
                            </el-switch>
                        </div>
                    </template>
                </el-table-column>

                <el-table-column
                        label="操作"
                        min-width="12%">
                <template scope="props">
                        <router-link :to="{name: 'serviceUpdate', params: {id: props.row.Id}}" tag="span">
                            <el-button type="info" size="mini" icon="edit">修改</el-button>
                        </router-link>

                        <el-button type="warning" size="mini" icon="document" @click="copy_data(props.row.Id)">复制
                        </el-button>

                        <el-button style="margin-left:0px" type="danger" size="mini" icon="delete" @click="delete_data(props.row.Id)">删除
                        </el-button>
                    </template>
                </el-table-column>
            </el-table>
        </div>
    </div>
</template>

<style>
    .input_red .el-input__inner {
        color: red;
    }
    .input_blue .el-input__inner {
        color: blue;
    }
    .input_green .el-input__inner {
        color: green;
    }
    .input_gray .el-input__inner {
        color: gray;
    }
    .el-table .cell {
        padding-left: 10px;
        padding-right: 10px;
    }
</style>

<script type="text/javascript">
    import {panelTitle, bottomToolBar, search} from 'components'
    import {port_service, port_deploy, port_walle} from 'common/port_uri'
    import {walle_stop, walle_restart, walle_reload} from 'common/walle'
    import store from "store"

    export default{
        data(){
            return {
                ProjectId: store.state.user_info.user.ProjectId,
                table_data: [],
                form: [],
                //当前页码
                currentPage: 1,
                //数据总条目
                total: 0,
                //每页显示多少条数据
                length: 15,
                //请求时的loading效果
                load_data: false,
                on_submit_loading: false,
                //批量选择数组
                search_text: "",

                stop: walle_stop.stop,
                restart: walle_restart.restart,
                reload: walle_reload.reload,
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
                this.search_text = value
                this.$message({
                    message: value,
                    type: 'success'
                })
                this.get_table_data()
            },

            //刷新
            on_refresh(){
                this.search_text = ""
                this.get_table_data()
            },

            handleSelectionChange(val) {
                this.form = val;
            },

            //获取数据
            get_table_data(){
                this.load_data = true
                this.$http.get(port_service.backend, {
                            params: {
                                project_id: this.ProjectId,
                                search_text: this.search_text
                            }
                        })
                .then(({data: {data}}) => {
                    this.table_data = data
                    this.load_data = false
                })
                .catch(() => {
                    this.load_data = false
                })
            },

            get_service_status_all() {
                this.load_data = true
                for(var i in this.table_data) {
                    this.get_service_status(i)
                }
                this.load_data = false
            },

            //根据项目ID获取项目的状态
            get_service_status(index){
                this.load_data = true
                this.$http.get(port_walle.dockerps, {
                            params: {
                                service_id: this.table_data[index].Id
                            }
                        })
                .then(({data: {data}}) => {
                    this.$set(this.table_data[index], "service_data", data)
                    this.load_data = false
                })
                this.load_data = false
            },

            linechange(service, host, line, service_id) {
                this.on_submit_loading = true
                this.load_data = true
                let host_port = host + ":" + service.Port.split(",")[0].split(":")[0]
                this.$http.get(port_walle.linechange, {
                        params: {
                            nacos: service.Project.Nacos,
                            host_port: host_port,
                            line: line,
                            service_id: service_id
                        }
                    })
                .then(({data: {data}}) => {
                    this.$message({
                        message: "发送上下线成功",
                        type: 'success'
                    })
                    this.on_submit_loading = false
                    this.load_data = false
                })
                this.on_submit_loading = false
                this.load_data = false
            },

            //提交
            on_submit_form() {
                this.load_data = true
                this.on_submit_loading = true
                if (this.form.length == 0) {
                    this.$message({
                        message: "请选择发布的项目",
                        type: 'warning'
                    })
                    this.load_data = false
                    this.on_submit_loading = false
                    return
                }
                this.$http.post(port_deploy.save, this.form)
                .then(({data: {data}}) => {

                    setTimeout(() => {
                        let routeData = this.$router.resolve({
                            name: "deployRelease",
                            params: {
                                id: data,
                            }
                        });
                        window.open(routeData.href, '_blank');
                        this.on_submit_loading = false
                    })
                })
                .catch(() => {
                    this.load_data = false
                    this.on_submit_loading = false
                })

                this.load_data = false
                this.on_submit_loading = false
            },

            //重启
            on_restart(service, host){
                this.$confirm('确定需要重启?', '提示', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning'
                }).then(() => {

                    let walle_form = {
                        User: store.state.user_info.user,
                        Service: service,
                        Host: host
                    }
                    this.$http.post(port_walle.restart, walle_form)
                    .then(({data: {data}}) => {
                        this.on_submit_loading = false
                        this.load_data = false
                        this.$message({
                            message: "发送重启成功",
                            type: 'success'
                        })
                    }).catch(() => {
                        this.on_submit_loading = false
                        this.load_data = false
                    })
                }).catch(() => {
                    this.on_submit_loading = false
                    this.load_data = false
                })
            },

            //关闭
            on_stop(service, host){
                this.$confirm('确定需要关闭?', '提示', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning'
                }).then(() => {

                    let walle_form = {
                        User: store.state.user_info.user,
                        Service: service,
                        Host: host
                    }
                    this.$http.post(port_walle.stop, walle_form)
                    .then(({data: {data}}) => {
                        this.on_submit_loading = false
                        this.load_data = false
                        this.$message({
                           message: "发送关闭成功",
                            type: 'success'
                        })
                    }).catch(() => {
                        this.on_submit_loading = false
                        this.load_data = false
                    })
                }).catch(() => {
                    this.on_submit_loading = false
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
                    this.$http.delete(port_service.id, {
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
                    })
                    .catch(() => {
                        this.load_data = false
                    })
                })
                .catch(() => {
                    this.load_data = false
                })
            },

            //复制项目
            copy_data(id){
                this.$confirm('是否复制项目?', '提示', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning'
                })
                .then(() => {
                    this.load_data = true
                    this.$http.get(port_service.copy, {
                            params: {
                                id: id,
                            }
                        })
                    .then(({data: {msg}}) => {
                        this.$message({
                            message: msg,
                            type: 'success'
                        })
                        this.get_table_data()
                        this.load_data = false
                    })
                    .catch(() => {
                    })
                })
                .catch(() => {
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
