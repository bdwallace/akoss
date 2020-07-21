
<template>

    <div class="panel">
        <panel-title :title="$route.meta.title">
            <!-- <el-radio-group style="float: left;margin-right: 100px;margin-top: 15px" v-model="level" @change="get_table_data()">
            <el-radio v-for="item in level_data" :key="item.Id" :label="item.Id">{{item.Name}}</el-radio>
            </el-radio-group> -->

            <!-- <div style="float: left;margin-right: 10px;margin-top: 5px">
                <el-button plain size="small" @click="resetPlatformFilter">所有平台</el-button>
                <el-button plain size="small" @click="resetClassFilter">所有类型</el-button>
            </div> -->

            <div style="float: left;margin-right: 10px;margin-top: 5px">
                <el-button plain size="small">全部&nbsp;{{total}}&nbsp;个域名</el-button>
            </div>

            <div style="float: left;margin-right: 10px;margin-top: 5px">
                <search @search="submit_search"></search>
            </div>

            <el-button @click.stop="on_refresh" size="small">
                <i class="fa fa-refresh"></i>
            </el-button>

            <router-link :to="{name: 'domainAdd'}" tag="span">
                <el-button type="primary" icon="plus" size="small">新建域名</el-button>
            </router-link>
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
                        sortable
                        width="80">
                </el-table-column>

                <el-table-column
                        prop="Name"
                        label="名称"
                        sortable
                        width="160">
                </el-table-column>

                <el-table-column
                        prop="Domain"
                        label="域名"
                        sortable
                        column-key="Domain"
                        width="160">
                </el-table-column>

                        <!-- :filter-method="filterClass" -->
                        <!-- :filtered-value="filterValue" -->
                        <!-- :filter-method="filterMethod" -->
                        <!-- :filter-multiple="false" -->
                <el-table-column
                        prop="Class"
                        label="类型"
                        column-key="Class"
                        :filters="itemClass"
                        width="100">
                </el-table-column>

                        <!-- :filter-method="filterPlatform" -->
                <el-table-column
                        prop="Platforms"
                        label="平台"
                        column-key="Platforms"
                        :filters="itemPlatform"
                        width="200">
                    <template scope="props">
                        <font v-for="item in props.row.Platforms" :key="item">{{item.Name}}, </font>
                    </template>
                </el-table-column>

                        <!-- :filter-method="filterService"> -->
                <el-table-column
                        prop="Services"
                        label="加速服务"
                        column-key="Services"
                        :filters="itemService">
                    <template scope="props">
                        <font v-for="item in props.row.Services" :key="item">{{item.Name}}, </font>
                    </template>
                </el-table-column>

                <el-table-column
                        label="监控"
                        width="70">
                        <template scope="props">
                            <span v-if="props.row.Monitor == 1">是</span>
                        </template>
                </el-table-column>

                <el-table-column
                        prop="Crt"
                        label="crt证书"
                        width="110">
                    <template scope="props">
                        <el-popover
                            ref="crt"
                            placement="left-start"
                            width="600"
                            trigger="click">
                            <!-- <span style="color:teal;white-space: pre-wrap;">{{props.row.Crt}}</span> -->
                            <el-input type="textarea" :rows="20" readonly v-model="props.row.Crt"></el-input>
                        </el-popover>
                        <el-button v-show="props.row.Crt != ''" type="info" size="mini" v-popover:crt>查看</el-button>
                        <el-button v-show="props.row.Crt != ''" type="info" size="mini" @click="download(props.$index, 'crt')">下载</el-button>
                    </template>
                </el-table-column>

                <el-table-column
                        prop="Key"
                        label="key私钥"
                        width="110">
                    <template scope="props">
                        <el-popover
                            ref="key"
                            placement="left-start"
                            width="600"
                            trigger="click">
                            <!-- <span style="color:teal;white-space: pre-wrap;">{{props.row.Key}}</span> -->
                            <el-input type="textarea" :rows="20" readonly v-model="props.row.Key"></el-input>
                        </el-popover>
                        <el-button v-show="props.row.Key != ''" type="info" size="mini" v-popover:key>查看</el-button>
                        <el-button v-show="props.row.Key != ''" type="info" size="mini" @click="download(props.$index, 'key')">下载</el-button>
                    </template>
                </el-table-column>

                <el-table-column
                        label="操作"
                         width="165">
                    <template scope="props">
                        <router-link :to="{name: 'domainUpdate', params: {id: props.row.Id}}" tag="span">
                            <el-button type="info" size="small" icon="edit">修改</el-button>
                        </router-link>
                        
                        <el-button style="margin-left:0px" type="danger" size="small" icon="delete" @click="delete_data(props.row.Id)">删除
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
<script type="text/javascript">
    import {panelTitle, bottomToolBar, search} from 'components'
    import {port_platform, port_domain, port_service} from 'common/port_uri'
    import store from 'store'
    import FileSaver from "file-saver" 
    export default{
        data(){
            return {
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
                // filter: {},
                filter: {
                    Class: [],
                    Platforms: [],
                    Services: []
                },

                itemClass: [],
                itemPlatform: [],
                itemService: [],

                Class: "NOT_FILTER",
                Platforms: "NOT_FILTER", 
                Services: "NOT_FILTER", 
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
            this.get_platform_name()
            this.get_service_name()
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

            filterClass(value, row) {
                return row.Class === value;
            },

            filterPlatform(value, row) {
                for (var i in row.Platforms) {
                    return row.Platforms[i].Name === value;
                }
            },

            filterService(value, row) {
                for (var i in row.Services) {
                    return row.Services[i].Name === value;
                }
            },

            filterChange(filters) {
                Object.assign(this.filter, filters)
                if (this.filter.Class.length == 0) {
                    this.Class = "NOT_FILTER"
                } else {
                    this.Class = this.filter.Class.join()
                }

                if (this.filter.Platforms.length == 0) {
                    this.Platforms = "NOT_FILTER"
                } else {
                    this.Platforms = this.filter.Platforms.join()
                }

                if (this.filter.Services.length == 0) {
                    this.Services = "NOT_FILTER"
                } else {
                    this.Services = this.filter.Services.join()
                }
                this.get_table_data()
            },

            //获取数据
            get_table_data(){
                this.load_data = true
                this.table_data = []
                this.$http.get(port_domain.list, {
                            params: {
                                project_id: this.ProjectId,
                                page: this.currentPage,
                                length: this.length,
                                search_text: this.search_text,
                                class: this.Class,
                                platforms: this.Platforms,
                                services: this.Services
                            }
                        })
                .then(({data: {data}}) => {
                    this.table_data = data.table_data
                    this.currentPage = data.currentPage
                    this.total = data.total
                    // this.filter()
                    this.load_data = false
                })
                .catch(() => {
                    this.load_data = false
                })
            },

            get_service_name() {
                this.load_data = true
                this.$http.get(port_service.name, {
                            params: {
                                project_id: this.ProjectId,
                            }
                        })
                .then(({data: {data}}) => {
                    for(let i in data) {
                        this.itemService[i] = {text: data[i], value: data[i]}
                    }
                    // this.itemService.push({text: "", value: ""})
                    this.load_data = false
                })
                .catch(() => {
                    this.load_data = false
                })
            },
            
            get_platform_name() {
                this.load_data = true
                this.$http.get(port_platform.name, {
                            params: {
                                project_id: this.ProjectId,
                            }
                        })
                .then(({data: {data}}) => {
                    for(let i in data) {
                        this.itemPlatform[i] = {text: data[i], value: data[i]}
                    }
                    // this.itemPlatform.push({text: "", value: ""})
                    this.load_data = false
                })
                .catch(() => {
                    this.load_data = false
                })
            },

            get_domain_class() {
                this.load_data = true
                this.$http.get(port_domain.class, {
                            params: {
                                project_id: this.ProjectId,
                            }
                        })
                .then(({data: {data}}) => {
                    for(let i in data) {
                        // if(data[i] === "") {
                        //     continue
                        // }
                        this.itemClass.push({text: data[i], value: data[i]})
                    }
                    this.load_data = false
                })
                .catch(() => {
                    this.load_data = false
                })
            },


            // filter() {
            //     let class_list = []
            //     let platform_list = []
            //     let service_list = []
            //     for(var i in this.table_data) {
            //         // if (domain_list.indexOf(this.table_data[i].Domain) == -1) {
            //         //     domain_list.push(this.table_data[i].Domain)
            //         // }

            //         if (class_list.indexOf(this.table_data[i].Class) == -1) {
            //             class_list.push(this.table_data[i].Class)
            //         }

            //         for(var j in this.table_data[i].Platforms) {
            //             if (platform_list.indexOf(this.table_data[i].Platforms[j].Name) == -1) {
            //                 platform_list.push(this.table_data[i].Platforms[j].Name)
            //             }
            //         }

            //         for(var j in this.table_data[i].Services) {
            //             if (service_list.indexOf(this.table_data[i].Services[j].Name) == -1) {
            //                 service_list.push(this.table_data[i].Services[j].Name)
            //             }
            //         }
            //     }

            //     // for (var i in domain_list) {
            //     //     this.itemDomain[i] = {text: domain_list[i], value: domain_list[i]}
            //     // }

            //     for (var i in class_list) {
            //         this.itemClass[i] = {text: class_list[i], value: class_list[i]}
            //     }

            //     for (var i in platform_list) {
            //         this.itemPlatform[i] = {text: platform_list[i], value: platform_list[i]}
            //     }

            //     for (var i in service_list) {
            //         this.itemService[i] = {text: service_list[i], value: service_list[i]}
            //     }
            // },



            //根据id删除数据
            delete_data(id){
                this.$confirm('此操作将删除该数据, 是否继续?', '提示', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning'
                }).then(() => {
                    this.load_data = true
                this.$http.delete(port_domain.id, {
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
            }).
                catch(() => {
                    this.load_data = false
            })
            }).
                catch(() => {
                    this.load_data = false
            })
            },

            download(index, etx) {
                let fileName = this.table_data[index].Name + "." + etx
                if(etx == "crt") {
                    var blob = new Blob([this.table_data[index].Crt], {type: "text/plain;charset=utf-8"})
                } 
                if(etx == "key") {
                    var blob = new Blob([this.table_data[index].Key], {type: "text/plain;charset=utf-8"})
                }
                FileSaver.saveAs(blob, fileName)
            },
            
            //页码选择
            handleCurrentChange(val) {
                this.currentPage = val
                this.get_table_data()
            }
        }
    }
</script>
