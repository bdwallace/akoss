
<template>
    <div class="panel" style="padding:15px">
        <span style="float: left;"> {{$route.meta.title}} </span>
        <panel-title>

            <el-button type="success" size="small" icon="setting" @click="get_service_status_all()" :loading="on_submit_loading">更新全部余额</el-button>
            <el-button type="primary" icon="plus" size="small" @click="on_dialog_save(0)">创建云帐号</el-button>
        </panel-title>

        <div class="panel-body-line" style="clear: both;">
        <!-- <div class="panel-body-line" style="clear: both;" v-for="(item, index) in table_data" :key="item">
            <div style="margin: 10px;margin-top: 10px">
                <b style="color: red;">{{index}} </b>
                环境链接
            </div>
            <el-table
                :data="item" -->
            <el-table
                    :data="table_data"
                    v-loading="load_data"
                    element-loading-text="拼命加载中"
                    border
                    :default-sort = "{prop: 'Class', order: 'ascending'}"
                    style="width: 100%">

                <el-table-column
                    label="ID"
                    prop="Id"
                    width="100">
                </el-table-column>

                <el-table-column
                    label="云平台"
                    prop="Class"
                    width="100">
                </el-table-column>

                <el-table-column
                    label="帐号"
                    prop="Account"
                    width="180">
                </el-table-column>

                <el-table-column
                    label="余额"
                    prop="Balance"
                    width="150">
                    <template scope="props">
                        <span>{{ props.row.Balance }}</span>
                        <el-button type="info" size="mini"  style="float: right;" @click="on_update_balance(props.row, props.$index)">更新
                        </el-button>
                    </template>
                </el-table-column>

                <el-table-column
                        prop="UpdatedAt"
                        label="更新时间"
                        width="180"
                        :formatter="dateFormat">
                </el-table-column>
                
                <el-table-column
                    label="登录地址"
                    prop="Link"
                    width="400">
                    <template scope="props">
                        <a target="_blank" :href="props.row.Link">{{props.row.Link}}</a>
                    </template>
                </el-table-column>

                <el-table-column
                    label="说明"
                    prop="Direction"
                    width="200">
                </el-table-column>

                <el-table-column
                        label="操作"
                         width="220">
                <template scope="props">
                        <el-button type="info" size="mini" icon="edit" @click="on_dialog_save(props.row.Id)">修改</el-button>

                        <el-button type="warning" size="mini" icon="document" @click="copy_data(props.row.Id)">复制
                        </el-button>

                        <el-button type="danger" size="mini" icon="delete" @click="delete_data(props.row.Id)">删除
                        </el-button>
                    </template>
                </el-table-column>
            </el-table>
        </div>
        <el-dialog :title="title" 
            :visible.sync="is_saveDialog"
            :modal="true"
            :show-close="false"
            :modal-append-to-body="false">
            <el-form ref="form" :model="form" :rules="rules">

                <el-form-item label="云类型:" prop="Class" label-width="120px">
                    <el-select 
                        v-model="form.Class" 
                        clearable
                        placeholder="请选择">
                        <el-option
                        v-for="item in cloud_class"
                        :key="item"
                        :label="item"
                        :value="item">
                        </el-option>
                    </el-select>
                </el-form-item>

                <el-form-item label="帐号:" prop="Account" label-width="120px">
                    <el-input v-model="form.Account" placeholder="请输帐号" style="width: 400px;">
                    </el-input>
                </el-form-item>

                <el-form-item label="AccessKeyId:" prop="AccessKeyId" label-width="120px">
                    <el-input v-model="form.AccessKeyId" placeholder="请输AccessKeyId" style="width: 400px;">
                    </el-input>
                </el-form-item>

                <el-form-item label="AccessSecret:" prop="AccessSecret" label-width="120px">
                    <el-input v-model="form.AccessSecret" placeholder="请输AccessSecret" style="width: 400px;">
                    </el-input>
                </el-form-item>

                <el-form-item label="登录地址:" prop="Link" label-width="120px">
                    <el-input v-model="form.Link" placeholder="请输入链接地址，如：www.baidu.com"
                            style="width: 400px;">
                    </el-input>
                </el-form-item>

                <el-form-item label="说明:" prop="Direction" label-width="120px">
                    <el-input v-model="form.Direction" placeholder="请输入链接说明，非必填"
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
    // import {savedialog} from "pages/link"
    import {port_cloud} from 'common/port_uri'
    import store from 'store'
    
    export default{
        data(){
            return {
                table_data: [], 
                
                is_saveDialog: false,
                title: "",

                cloud_class: ["阿里云", "腾讯云", "百度云", "天翼云", "七牛云"],
                
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
                    Name: "",
                    Class: "",
                    Account: "", 
                    AccessKeyId: "",
                    AccessSecret: "",
                    Direction: "",
                    Link: "",
                    Balance: 0
                },

                rules: {
                    // Name: [{required: true, message: '链接名称不能为空', trigger: 'blur'}],
                    // Link: [{required: true, message: '链接地址不能为空', trigger: 'blur'}],
                    Class: [{required: true, message: '云平台类型不能为空', trigger: 'blur'}],
                    Account: [{required: true, message: '帐号不能为空', trigger: 'blur'}],
                    // Direction: [{required: true, message: '链接说明不能为空', trigger: 'blur'}],
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
            on_update_balance(row, index) {
                this.$http.get(port_cloud.sync, {
                            params: {
                                // id: row.Id,
                                id: row.Id,
                            }
                        })
                .then(({data: {data}}) => {
                    this.$message({
                        message: "更新成功",
                        type: 'success'
                    })
                    // 使用这个方法排序，会在表格在有default-sort的情况，导致数据错乱。
                    //this.table_data[index] = data

                    this.$set(row, "Balance", data.Balance)
                    this.$set(row, "UpdatedAt", data.UpdatedAt)

                    // 使用此方法不能动态更新
                    // row = data
                    // this.$forceUpdate();
                }).catch(() => {
                    this.load_data = false
                })
            },

            //获取数据
            get_table_data(){
                this.load_data = true
                this.table_data = []
                this.$http.get(port_cloud.list)
                .then(({data: {data}}) => {
                    this.table_data = data 
                    // console.log("----data---", data)
                    // let temp_data = {}
                    // for (var i in data) {
                    //     if (temp_data[data[i].Class] == null) {
                    //         temp_data[data[i].Class] = [data[i]]
                    //     } else {
                    //         temp_data[data[i].Class].push(data[i])
                    //     }
                         
                    // }
                    // this.table_data = temp_data
                    // console.log("----tab_data---", this.table_data)

                    this.load_data = false
                    this.on_submit_loading = false
                })
                .catch(() => {
                    this.load_data = false
                    this.on_submit_loading = false
                })
            },

            //根据id复制数据
            copy_data(id){
                this.$confirm('是否复制此信息?', '提示', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning'
                }).then(() => {
                    this.load_data = true
                    this.$http.get(port_cloud.copy, {
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

            //根据id删除数据
            delete_data(id){
                this.$confirm('此操作将删除该数据, 是否继续?', '提示', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning'
                }).then(() => {
                    this.load_data = true
                    this.$http.delete(port_cloud.id, {
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
                    this.title = "新增链接"
                    // this.init_form()
                } else {
                    this.title = "修改链接"
                    this.get_form_data(id)
                }
                this.is_saveDialog = true
            },

            // init_form() {
            //     this.form = {
            //         Id: 0,
            //         Name: "",
            //         Class: "",
            //         Account: "", 
            //         AccessKeyId: "",
            //         AccessSecret: "",
            //         Direction: "",
            //         Link: "",
            //         Balance: ""
            //     }
            // },

            //获取数据
            get_form_data(id) {
                this.load_data = true
                this.$http.get(port_cloud.id, {
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
                this.$http.post(port_cloud.save, this.form)
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
        }
    }
</script>
