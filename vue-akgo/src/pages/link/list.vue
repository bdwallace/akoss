
<template>

    <div class="panel">
        <panel-title :title="$route.meta.title">
            <!-- <router-link :to="{name: 'linkAdd'}" tag="span">
                <el-button type="primary" icon="plus" size="small">创建链接</el-button>
            </router-link> -->
            <el-button type="primary" icon="plus" size="small" @click="on_dialog_save(0)">创建链接</el-button>
        </panel-title>

        <!-- <div class="panel-body-line" style="clear: both;" v-for="(item, index) in table_data" :key="item"> -->
        <div class="panel-body-line" style="clear: both;">
            <div style="margin: 10px"><b style="color: red;">{{table_public.name}} </b>环境链接</div>


            <el-table
                :data="table_public.data"
                    v-loading="load_data"
                    element-loading-text="拼命加载中"
                    border
                    :show-header="false"
                    style="width: 100%;margin-left:100px;">

                <el-table-column
                    label="链接名"
                    prop="Name"
                    width="140">
                </el-table-column>

                <el-table-column
                    label="链接地址"
                    prop="Link"
                    width="600">
                    <template scope="props">
                        <a target="_blank" :href="props.row.Link">{{props.row.Link}}</a>

                        <!-- <router-link :to="{name: 'linkUpdate', params: {id: props.row.Id}}" load_data=false tag="span">
                            <el-button size="mini" type="text" style="margin-right: 1px">修改</el-button>
                        </router-link> -->
                        <el-button type="text" size="mini" @click="on_dialog_save(props.row.Id)">修改</el-button>

                        <el-button size="mini" type="text" @click="delete_data(props.row.Id)">删除</el-button>

                    </template>
                </el-table-column>

            </el-table>

        </div>
        <div class="panel-body-line" style="clear: both;" v-for="(item, index) in table_project_index" :key="item">
            <div style="margin: 10px;margin-top: 50px"><b style="color: red;">{{table_project_index[index]}} </b>环境链接</div>
            <el-table
                :data="table_project_data[index]"
                    v-loading="load_data"
                    element-loading-text="拼命加载中"
                    border
                    :show-header="false"
                    style="width: 100%;margin-left: 100px">
                <el-table-column
                    label="链接名"
                    prop="Name"
                    width="140">
                </el-table-column>

                <el-table-column
                    label="链接地址"
                    prop="Link"
                    width="600">
                    <template scope="props">
                        <a target="_blank" :href="props.row.Link">{{props.row.Link}}</a>

                        <el-button type="text" size="mini" @click="on_dialog_save(props.row.Id)">修改
                        </el-button>

                        <el-button size="mini" type="text" @click="delete_data(props.row.Id)">删除</el-button>

                    </template>
                </el-table-column>

            </el-table>

            <!-- <el-table :data="table_platform_data[index]"
                    border
                    v-loading="load_data"
                    style="width: 100%;margin-top:5px">
                <el-table-column
                    key="name"
                    label="平台"
                    prop="name"
                    width="120px">
                </el-table-column>
                <el-table-column
                        :key="col"
                        :label="col"
                        :prop="col"
                        v-for="col in itemClass">
                    <template scope="props">
                        <div  v-if="props.row[props.column.label]" >
                        <p v-for="item in props.row[props.column.label]" :key="item">
                            <a target="_blank" :href="item.Link">{{item.Link}}</a>

                            <el-button type="text" size="mini" @click="on_dialog_save(item.Id)">修改
                            </el-button>

                            <el-button size="mini" type="text" @click="delete_data(item.Id)">删除</el-button>
                        </p>
                        </div>
                    </template>
                </el-table-column>
            </el-table>  -->
        </div>
        <el-dialog :title="title"
            :visible.sync="is_saveDialog"
            :modal="true"
            :show-close="false"
            :modal-append-to-body="false">
            <!-- <save_link></save_link> -->
            <el-form ref="form" :model="form" :rules="rules">

                <el-form-item label="链接名称:" prop="Name" label-width="120px">
                    <el-input v-model="form.Name" placeholder="请输入链接名称，非必填"
                            style="width: 400px;">
                    </el-input>
                </el-form-item>

                <el-form-item label="说明:" prop="Direction" label-width="120px">
                    <el-input v-model="form.Direction" placeholder="请输入链接说明，非必填"
                            style="width: 400px;">
                    </el-input>
                </el-form-item>

                <el-form-item label="地址:" prop="Link" label-width="120px">
                    <el-input v-model="form.Link" placeholder="请输入链接地址，如：www.baidu.com"
                            style="width: 400px;">
                    </el-input>
                </el-form-item>

                <el-form-item label="选择环境:" prop="Projects" label-width="120px">
                    <el-radio-group v-model="form.Projects[0].Id" :max="1">
                        <!-- @change="get_itemPlatform()"> -->
                        <el-radio v-for="item in itemProject" :key="item.Id" :label="item.Id">{{item.Name}}</el-radio>
                    </el-radio-group>
                </el-form-item>

                <!-- <el-form-item v-if="form.Projects[0].Id" label="选择平台:" prop="Project" label-width="120px">
                    <el-select
                        v-model="form.Platforms[0]"
                        clearable
                        value-key="Id"
                        placeholder="请选择"
                        @change="change_itemPlatform()">
                        <el-option
                        v-for="item in itemPlatform"
                        :key="item"
                        :label="item.Name"
                        :value="item">
                        </el-option>
                    </el-select>
                    <el-select
                        v-if="form.Platforms[0].Id"
                        v-model="form.Class"
                        clearable
                        placeholder="请选择">
                        <el-option
                        v-for="item in itemClass"
                        :key="item"
                        :label="item"
                        :value="item">
                        </el-option>
                    </el-select>
                </el-form-item> -->

                <el-form-item>
                <el-button type="primary" @click="on_submit_form" :loading="on_submit_loading">立即提交</el-button>
                <el-button @click="is_saveDialog=false">取消</el-button>
                </el-form-item>

            </el-form>

            <!-- <div slot="footer" class="dialog-footer">
                <el-button @click="close_dialog(is_saveDialog=false)">取 消</el-button>
                <el-button type="primary" @click="on_submit_form" :loading="on_submit_loading">立即提交
                </el-button>
            </div> -->
        </el-dialog>
    </div>
</template>
<script type="text/javascript">
    import {panelTitle, bottomToolBar, search} from 'components'
    // import {savedialog} from "pages/link"
    import {port_link, port_project, port_platform} from 'common/port_uri'
    import store from 'store'

    export default{
        data(){
            return {
                table_project_index: [],
                table_project_data: [],
                table_public: {
                    name: "共同",
                    data: []
                },

                // cols: [],
                table_platform_index: [],
                table_platform_data: [],
                itemClass: [
                    "h5",
                    "h5-proxy",
                    "h5-site",
                    "merchant",
                    "download",
                    "download-share",
                    "agent",
                    "share-agent",
                    "customer",
                    "chat-backend",
                ],
                // cols: [
                //     {label:"h5", prop: "h5"},
                //     {label:"merchant", prop: "merchant"},
                //     {label:"download", prop: "download"},
                //     {label:"download-share", prop: "download-share"},
                //     {label:"agent", prop: "agent"},
                //     {label:"share-agent", prop: "share-agent"},
                //     {label:"customer", prop: "customer"},
                //     {label:"chat-backend", prop: "chat-backend"},
                // ],
                is_saveDialog: false,
                title: "",

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
	                Direction: null,
                    Link: null,
                    Class: null,
                    Projects: [{Id:null}],
                    // Platforms: [{Id:null}]
                },
                itemProject: [],
                itemPlatform: [],
                rules: {
                    // Name: [{required: true, message: '链接名称不能为空', trigger: 'blur'}],
                    Link: [{required: true, message: '链接地址不能为空', trigger: 'blur'}],
                    // Direction: [{required: true, message: '链接说明不能为空', trigger: 'blur'}],
                },
                on_submit_loading: false,
            }
        },
        components: {
            panelTitle,
            bottomToolBar,
            search,
            // savedialog
        },
        created(){
            this.get_table_data()
        },
        methods: {
            test(p) {
            },

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
                this.$http.get(port_link.list)
                            .then(({data: {data}}) => {
                    // this.table_data = data
                    this.table_data = [[]]
                    for (var i in data) {
                        // console.log("---i---", data[i])
                        if (data[i].Projects.length == 0) {
                            this.table_public.data.push(data[i])
                        } else {
                            // console.log("---p0---", data[i])
                            let table_project_index_item = data[i].Projects[0].Id + "_" + data[i].Projects[0].Alias
                            let index = this.table_project_index.indexOf(table_project_index_item)

                            if (index < 0) {
                                this.table_project_index.push(table_project_index_item)
                                this.table_project_data.push([data[i]])
                            } else {
                                this.table_project_data[index].push(data[i])
                            }
                            // if (data[i].Platforms.length == 0) {
                            //     if (index < 0) {
                            //         this.table_project_index.push(table_project_index_item)
                            //         this.table_project_data.push([data[i]])
                            //     } else {
                            //         this.table_project_data[index].push(data[i])
                            //     }
                            // } else if(data[i].Class == "") {
                            //     if (index < 0) {
                            //         this.table_project_index.push(table_project_index_item)
                            //         this.table_project_data.push([data[i]])
                            //     } else {
                            //         this.table_project_data[index].push(data[i])
                            //     }

                            // }else {
                            //     let platform_index_item = data[i].Platforms[0].Name
                            //     if (index < 0) {
                            //         this.table_project_index.push(table_project_index_item)
                            //         this.table_project_data.push([])

                            //         this.table_platform_index.push([platform_index_item])
                            //         let platform_data = {name: platform_index_item}
                            //         this.$set(platform_data, data[i].Class, [data[i]])
                            //         this.table_platform_data.push([platform_data])
                            //     } else {
                            //         let platform_index
                            //         if(this.table_platform_index[index] == undefined) {
                            //             platform_index = -1
                            //             this.table_platform_index[index] = []
                            //             this.table_platform_data[index] = []
                            //         } else {
                            //             platform_index = this.table_platform_index[index].indexOf(platform_index_item)

                            //         }

                            //         if (platform_index < 0) {
                            //             this.table_platform_index[index].push(platform_index_item)
                            //             let platform_data = {name: platform_index_item}
                            //             this.$set(platform_data, data[i].Class, [data[i]])
                            //             this.table_platform_data[index].push(platform_data)
                            //         } else {
                            //             if(data[i].Class in this.table_platform_data[index][platform_index]) {
                            //                 this.table_platform_data[index][platform_index][data[i].Class].push(data[i])
                            //             } else {
                            //                 this.$set(this.table_platform_data[index][platform_index], data[i].Class, [data[i]])
                            //             }

                            //         }

                            //     }
                                // if (this.cols.indexOf(data[i].Class) < 0) {
                                //     this.cols.push(data[i].Class)
                                // }
                            // }

                        }
                    }

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
                    this.$http.delete(port_link.id, {
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
                this.get_itemProject()
                if (id == 0) {
                    this.title = "新增链接"
                    this.init_form()
                } else {
                    this.title = "修改链接"
                    this.get_form_data(id)
                }
                this.is_saveDialog = true
            },

            init_form() {
                this.form = {
          	        Id: 0,
	                Name: null,
	                Direction: null,
                    Link: null,
                    Class: null,
                    Projects: [{Id:null}],
                    // Platforms: [{Id:null}]
                }
            },

            // 下拉框获得所有环境
            get_itemProject() {
                this.$http.get(port_project.list)
                        .then(({data: {data}}) => {
                    this.itemProject = data
                    this.itemProject.unshift({Id: null, Name: "共同环境"})
                    this.$forceUpdate();
                })
            },

            // 下拉框获得选择的环境所有平台
            get_itemPlatform(){
                // this.load_data = true
                if(this.form.Projects[0] && this.form.Projects[0].Id == null) {
                    this.form.Platforms[0] = {Id: null}
                    this.load_data = false
                    return
                }
                this.$http.get(port_platform.list, {
                                    params: {
                                        project_id: this.form.Projects[0].Id,
                                    }
                                })
                .then(({data: {data}}) => {
                    this.itemPlatform = data
                    if (this.itemPlatform.length == 0) {
                        this.form.Platforms[0].Id = null
                    }
                    if (this.form.Platforms[0].Id != null) {
                        this.form.Platforms[0].Id = this.itemPlatform[0].Id
                    }
                    // this.$set(this.form.Platforms[0], "Id", this.itemPlatform[0].Id)
                    this.load_data = false
                    })
                .catch(() => {
                    this.load_data = false
                })
            },

            change_itemPlatform() {
                if (this.form.Class == "") {
                    this.form.Class = this.itemClass[0]
                }
            },

            //获取数据
            get_form_data(id) {
                // this.load_data = true
                this.$http.get(port_link.id, {
                params: {
                    id: id
                }
                })
                .then(({data: {data}}) => {
                    this.form = data
                    if (this.form.Projects.length == 0) {
                        this.form.Projects.push({Id: null})
                    }
                    // if (this.form.Platforms.length == 0) {
                    //     this.form.Platforms.push({Id: null})
                    // }
                    this.load_data = false
                })
                .catch(() => {
                    this.load_data = false
                })
            },

            //提交
            on_submit_form() {
                if (this.form.Projects[0].Id == null) {
                    this.form.Projects = []
                }
                this.$http.post(port_link.save, this.form)
                .then(({data: {msg}}) => {
                    this.form.Projects.push({Id: null})
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
            }
        }
    }
</script>
