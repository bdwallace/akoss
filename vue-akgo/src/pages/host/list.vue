
<template>

    <div class="panel">
        <panel-title :title="$route.meta.title">
            <!-- <el-radio-group style="float: left;margin-right: 100px;margin-top: 15px" v-model="level" @change="get_table_data()">
            <el-radio v-for="item in level_data" :key="item.Id" :label="item.Id">{{item.Name}}</el-radio>
            </el-radio-group> -->

            <div style="float: left;margin-right: 10px;margin-top: 5px">
                <el-button plain size="small" v-if="table_data">全部&nbsp;{{table_data.length}}&nbsp;台主机</el-button>
            </div>

            <div style="float: left;margin-right: 10px;margin-top: 5px">
                <el-button type="success" size="small" icon="setting" @click="get_all_host_status()">Aws检测</el-button>
            </div>

            <div style="float: left;margin-right: 10px;margin-top: 5px">
                <el-button type="success" size="small" icon="setting" @click="get_all_docker_status()">docker检测</el-button>
            </div>

            <div style="float: left;margin-right: 10px;margin-top: 5px">
                <search @search="submit_search"></search>
            </div>
            <el-button @click.stop="on_refresh" size="small">
                <i class="fa fa-refresh"></i>
            </el-button>
            <router-link :to="{name: 'hostAdd'}" tag="span">
                <el-button type="primary" icon="plus" size="small">新建主机</el-button>
            </router-link>

            <router-link :to="{name: 'hostStopTime'}" tag="span">
                <el-button type="primary" icon="plus" size="small">批量设置关机时间</el-button>
            </router-link>

        </panel-title>

        <div class="panel-body" style="clear: both">
            <el-table
                    :data="table_data"
                    v-loading="load_data"
                    element-loading-text="拼命加载中"
                    @filter-change="filterChange"
                    border
                    style="width: 100%;">
                <el-table-column
                        prop="Id"
                        label="id"
                        width="50">
                </el-table-column>
                <el-table-column
                        prop="Name"
                        label="主机名">
                </el-table-column>
                <el-table-column
                        prop="ip"
                        label="基本信息"
                        width="150">
                    <template scope="props" style="clear:both;">
                        <div style="clear: both;margin-bottom: 10px"></div>

                        <!-- <div>
                        <el-tag type="gray">{{props.row.name}}</el-tag>
                        </div> -->

                        <div>
                        <el-tag type="primary">{{props.row.PrivateIp}}</el-tag>
                        <i size="mini" v-if="props.row.PrivateIp == props.row.UseIp" class="el-icon-check"></i>
                        </div>

                        <div>
                        <el-tag type="warning">{{props.row.PublicIp}}</el-tag>
                        <i size="mini" v-if="props.row.PublicIp == props.row.UseIp" class="el-icon-check"></i>
                        </div>

                        <div style="clear: both;margin-bottom: 10px;"></div>
                    </template>
                </el-table-column>
                <!-- <el-table-column
                        prop="ip_pub"
                        label="公网地址"
                        width="140">
                </el-table-column> -->

                <!-- <el-table-column
                        label="公网"
                        width="70">
                        <template scope="props">
                            <span v-if="props.row.is_pub == 1">是</span>
                        </template>
                </el-table-column> -->

                <el-table-column
                    label="定时关机"
                    width="115">
                    <template scope="props">
                        <!-- <div>
                        <el-tag type="primary">{{props.row.start_time}}</el-tag>
                        <i size="mini" v-if="props.row.is_pub == 0" class="el-icon-check"></i>
                        </div> -->
                        {{props.row.StartTime}}-{{props.row.EndTime}}
                    </template>
                </el-table-column>

                <el-table-column
                    prop="awsRegion"
                    label="AWS"
                    :filters="awsRegionItem"
                    column-key="awsRegion"
                    width="160">
                    <template scope="props">
                        <!-- <el-popover
                            ref="aws"
                            placement="left-start"
                            width="180"
                            trigger="hover">
                            <p style="color:teal;white-space: pre-wrap;">实例Id: {{props.row.InstanceId}}</p>
                            <p style="color:teal;white-space: pre-wrap;">所在地区: {{props.row.region}}</p>
                        </el-popover>
                        <el-button type="info" style="clear: both;margin:-10px" v-popover:aws>查看</el-button> -->

                        <div style="clear: both;margin-bottom: 10px"></div>

                      <div>
                        <el-tag type="primary">{{props.row.InstanceType}}</el-tag>
                      </div>

                      <div>
                        <el-tag type="primary">{{props.row.InstanceId}}</el-tag>
                        </div>

                        <div>
                        <el-tag type="warning">{{props.row.Region}}</el-tag>
                        </div>

                        <div style="clear: both;margin-bottom: 10px"></div>
                    </template>
                </el-table-column>

                <el-table-column
                        label="状态 . 操作"
                        width="152">
                    <template scope="props" style="clear:both;">
                        <div  v-if="props.row.Status == 'running'">
                            <el-tag type="success">{{props.row.Status}}</el-tag>
                            <!-- <el-button @click="stopMessage(props.$index)" style="margin-left:5px" type="info" size="mini" icon="caret-bottom">关闭</el-button> -->
                            <el-button @click="stop_host_status(props.$index)" style="margin-left:5px" type="info" size="mini" icon="caret-bottom">关闭</el-button>
                        </div>

                        <div  v-else-if="props.row.Status == 'stopped'">
                        <el-tag type="danger">{{props.row.Status}}</el-tag>
                        <el-button @click="start_host_status(props.$index)" style="margin-left:5px" type="info" size="mini" icon="caret-top">启动</el-button>
                        </div>

                        <el-tag type="warning" v-else>
                            {{props.row.Status}}</el-tag>
                        <!-- {{props.row.Status}} -->
                    </template>
                </el-table-column>

                <el-table-column
                    label="项目"
                    width="55">
                    <template scope="props">
                        <el-popover
                            ref="project"
                            placement="left-start"
                            width="80"
                            trigger="hover">
                            <span style="color:teal;white-space: pre-wrap;">
                                <font v-for="item in props.row.Services" :key="item">{{item.Name}}, </font>
                            </span>
                        </el-popover>
                        <el-button type="info" size="mini" v-popover:project>查看</el-button>
                    </template>
                </el-table-column>

                <el-table-column
                        label="操作"
                         width="200">
                    <template scope="props">
                        <router-link :to="{name: 'hostUpdate', params: {id: props.row.Id}}" tag="span">
                            <el-button style="margin-left:0px" type="info" size="mini" icon="edit">修改</el-button>
                        </router-link>

                        <el-button style="margin-left:0px" type="warning" size="mini" icon="document" @click="copy_data(props.row.Id)">复制
                        </el-button>

                        <el-button style="margin-left:0px" type="danger" size="mini" icon="delete" @click="delete_data(props.row.Id)">删除
                        </el-button>
                    </template>
                </el-table-column>

                <el-table-column
                        label="dokcer ps"
                        prop="ProjectStatus"
                        width="570">
                        <template scope="props" style="clear:both;">
                        <div style="clear: both;margin-bottom: 10px"> <span></span> </div>
                        <div v-for="(item, index) in props.row.ProjectStatus" :key="index">
                            <el-input v-if="item.ps_status.substr(0, 2) != 'Up'" class="input_red" readonly style="width: 100px" size="mini" v-model="item.ps_status">
                            </el-input>
                            <el-input v-else readonly style="width: 100px" size="mini" v-model="item.ps_status">
                            </el-input>

                            <el-input readonly style="width: 180px" size="mini" v-model="item.ps_name" >
                            </el-input>
                            <!-- <el-input readonly style="width: 150px" size="mini" v-model="item.ps_created_at">
                            </el-input> -->
                            <el-input readonly style="width: 180px" size="mini" v-model="item.ps_image">
                            </el-input>
                        </div>
                        <div style="clear: both;margin-bottom: 10px;"> <span></span> </div>
                    </template>
                </el-table-column>

            </el-table>
            <!-- <bottom-tool-bar>

                <div slot="page">
                    <el-pagination
                            @current-change="handleCurrentChange"
                            :current-page="currentPage"
                            :page-size="10"
                            layout="total, prev, pager, next"
                            :total="total">
                    </el-pagination>
                </div>
            </bottom-tool-bar> -->
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
    import {port_host} from 'common/port_uri'
    import store from 'store'

    export default{
        data(){
            return {
                awsRegion: "NOT_FILTER",
                awsRegionItem: [
                    { value: "ap-southeast-1", text: "亚太地区(新加坡)" },
                    { value: "ap-east-1", text: "亚太地区(香港)" },
                    { value: "ap-northeast-2", text: "亚太地区(首尔)" },
                    { value: "us-east-1", text: "美国东部(弗吉尼亚北部)" },
                    { value: "us-west-1", text: "美国西部(加利福尼亚北部)" },
                    { value: "", text: "" },
                ],
                filter: {
                    awsRegion: [],
                },
                ProjectId: store.state.user_info.user.ProjectId,
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
                batch_select: [],
                //批量选择数组
                search_text: "",
                //项目详情
                project_data:[]
            }
        },
        components: {
            panelTitle,
            bottomToolBar,
            search

        },
        created(){
            // this.get_level_data()
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

            filterChange(filters) {
                Object.assign(this.filter, filters)
                if (this.filter.awsRegion.length == 0) {
                    this.awsRegion = "NOT_FILTER"
                } else {
                    this.awsRegion = this.filter.awsRegion.join()
                }
                this.get_table_data()
            },

            start_host_status(index){
                this.$confirm('此操作将启动主机，并在10秒后删除nacos关联此主机的信息, 是否继续?', '提示', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning'
                }).then(() => {
                    // this.load_data = true
                    this.$http.post(port_host.hoststartaws, this.table_data[index])
                    .then(({data: {data}}) => {
                        this.get_host_status(index)
                        this.load_data = false
                    })
                    .
                    catch(() => {
                        this.get_host_status(index)
                        this.$forceUpdate()
                        this.load_data = false
                    })
                })
            },

            //根据aws的实例ID停止该实例
            stop_host_status(index){
                this.$confirm('此操作将在nacos增加些主机的关联信息，并在10秒后停止此主机, 是否继续?', '提示', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning'
                }).then(() => {
                    // this.load_data = true
                    this.$http.post(port_host.hoststopaws, this.table_data[index])
                    .then(({data: {data}}) => {
                        this.get_host_status(index)
                        this.load_data = false
                    })
                    .
                    catch(() => {
                        this.get_host_status(index)
                        this.$forceUpdate()
                        this.load_data = false
                    })
                })
            },

            //根据aws的实例ID获取实例的状态
            get_host_status(index){
                // this.load_data = true
                if (this.table_data[index].InstanceId == ""|| this.table_data[index].Region == "") {
                    return
                }
                this.$http.post(port_host.hoststatusaws, this.table_data[index])
                .then(({data: {data}}) => {
                    this.$set(this.table_data[index], 'Status', data)
                    this.load_data = false
                })
                .
                catch(() => {
                    this.$forceUpdate()
                    this.load_data = false
                })
            },

            // 显示所有状态
            get_all_host_status(){
                this.load_data = true
                for(var i in this.table_data){
                    this.get_host_status(i)
                }
                this.load_data = false
            },

            //根据项目ID获取项目的状态
            get_docker_status(index){
                this.load_data = true
                this.$http.get(port_host.dockerps, {
                    params: {
                        use_ip: this.table_data[index].UseIp,
                    }
                })
                .then(({data: {data}}) => {
                    this.$set(this.table_data[index], 'ProjectStatus', data)
                    this.load_data = false
                })
                .catch( res => {
                    // if meg != "" {
                        this.$message({
                            message: res.msg,
                            type: 'warning'
                        })
                    // }
                    // this.$forceUpdate()
                    this.$set(this.table_data[index], 'ProjectStatus', res.data)
                    this.load_data = false
                })
            },

            // 显示所有状态
            get_all_docker_status(){
                this.load_data = true
                for(var i in this.table_data){
                    this.get_docker_status(i)
                }
                this.load_data = false
            },

            //获取数据
            get_table_data(){
                this.load_data = true
                this.$http.get(port_host.list, {
                            params: {
                                project_id: this.ProjectId,
                                // page: this.currentPage,
                                // length: this.length,
                                search_text: this.search_text,
                                aws_region: this.awsRegion
                            }
                        })
                .then(({data: {data}}) => {
                    this.table_data = data
                    // console.log("------data----", this.table_data)
                    // this.table_data = data.table_data
                    // this.currentPage = data.currentPage
                    // this.total = data.total
                    // this.get_all_host_status()
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
                    this.load_data = true
                    this.$http.delete(port_host.id, {
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

            //复制项目
            copy_data(id){
                this.$confirm('是否复制项目?', '提示', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning'
                }).then(() => {
                    this.load_data = true
                    this.$http.get(port_host.copy, {
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

            //页码选择
            handleCurrentChange(val) {
                this.currentPage = val
                this.get_table_data()
            }
        }
    }
</script>
