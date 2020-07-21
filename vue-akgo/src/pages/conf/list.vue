
<template>

    <div class="panel">
        <panel-title :title="$route.meta.title">
            <div style="float: left;margin-right: 10px;margin-top: 5px">
                <el-button plain size="small">全部&nbsp;{{table_data.length}}&nbsp;个参数</el-button>
            </div>

            <div style="float: left;margin-right: 10px;margin-top: 5px">
                <search @search="submit_search"></search>
            </div>
            <el-button @click.stop="on_refresh" size="small">
                <i class="fa fa-refresh"></i>
            </el-button>
            <router-link :to="{name: 'confAdd'}" tag="span">
                <el-button type="primary" icon="plus" size="small">创建域名</el-button>
            </router-link>
        </panel-title>

        <div class="panel-body" style="clear: both;">
            <el-table
                    :data="table_data"
                    v-loading="load_data"
                    element-loading-text="拼命加载中"
                    border
                    style="width: 100%;">

                <el-table-column type="expand">
                    <template scope="props">
                        <!-- <el-row>
                        <el-col :span="12"> -->
                        <table align="left">
                            <tr v-for="(item, index) in props.row.json" :key="item" >
                                <td align="right" style="width:300px;border-style:none;padding-right:10px">
                                    <el-button v-show="props.row.json[index].key_show" size="mini"
                                        v-clipboard:copy="item.Key" 
                                        v-clipboard:success="copy" 
                                        v-clipboard:error="onError">
                                        复制
                                    </el-button>

                                    <el-popover
                                            :disabled="true"
                                            @hide="popover_hide(props.$index, index, 'key_show')"
                                            @show="popover_show(props.$index, index, 'key_show')"
                                            trigger="hover"
                                            content="">
                                        <span slot="reference">{{ item.Key }}</span>
                                    </el-popover>
                                </td>

                                <td align="left" style="width:500px;border-style:none;padding:10px">
                                    <!-- <span>{{item.Value}}</span>
                                    <el-button size="mini"
                                        v-clipboard:copy="item.Value" 
                                        v-clipboard:success="copy" 
                                        v-clipboard:error="onError">
                                        复制
                                    </el-button> -->
                                    <el-popover
                                            :disabled="true"
                                            @hide="popover_hide(props.$index, index, 'value_show')"
                                            @show="popover_show(props.$index, index, 'value_show')"
                                            trigger="hover">
                                        <span slot="reference">{{ item.Value }}</span>
                                    </el-popover>

                                    <el-button v-show="props.row.json[index].value_show" size="mini"
                                        v-clipboard:copy="item.Value" 
                                        v-clipboard:success="copy" 
                                        v-clipboard:error="onError">
                                        复制
                                    </el-button>

                                </td>
                            </tr>
                        </table>
                        <!-- </el-col>

                        <el-col :span="12">
                            <el-tag size="mini" type="gray" style="margin:5px" v-for="item in props.row.Services" :key="item">{{item.Name}}</el-tag>
                        </el-col>
                        </el-row> -->
                    </template>
                </el-table-column>

                <el-table-column
                        prop="Id"
                        label="id"
                        width="60">
                </el-table-column>

                <el-table-column
                        prop="Name"
                        label="参数名称"
                        width="200">
                </el-table-column>

                <el-table-column
                        prop="Services"
                        label="所关联服务">
                    <template scope="props">
                        <!-- <el-input size="mini" type="gray" style="margin:2px" v-for="item in props.row.Services" :key="item" v-model="item.Name"></el-input> -->
                        <el-tag size="mini" type="primary" style="margin:2px" v-for="item in props.row.Services" :key="item">{{item.Name}}</el-tag>
                        <!-- <el-tag size="mini" type="gray" style="margin:2px" v-for="item in props.row.Services" :key="item">{{item.Name}}</el-tag> -->
                        <!-- <font v-for="item in props.row.Services" :key="item">{{item.Name}}, </font> -->
                    </template>
                </el-table-column>

                <!-- <el-table-column
                    label="修改时间"
                    prop="UpdatedAt"
                    width="150"
                    :formatter="dateFormat">
                </el-table-column> -->
                
                <!-- <el-table-column
                    label="创建时间"
                    prop="CreatedAt"
                    width="150"
                    :formatter="dateFormat">
                </el-table-column> -->

                <el-table-column
                        label="操作"
                         width="220">
                <template scope="props">
                        <router-link :to="{name: 'confUpdate', params: {id: props.row.Id}}" tag="span">
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
                            :page-size="10"
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
    import {port_conf} from 'common/port_uri'
    import store from "store"

    export default{
        data(){
            return {
                ProjectId: store.state.user_info.user.ProjectId, 
                table_data: [],
                conf_key: null, 
                conf_value: null, 
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
            this.get_table_data()
            
            if(this.conf_key === ''){
                this.$refs.test.focus();
            }
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

            popover_show(table_index, json_index, type) {
                this.$set(this.table_data[table_index].json[json_index], type, true)
            },

            popover_hide(table_index, json_index, type) {
                var _this = this;
                setTimeout(function() {
                    _this.$set(_this.table_data[table_index].json[json_index], type, false)
                }, 2000);
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

            copy(e) {
                this.$message({
                    message: "复制成功!",
                    type: 'success'
                })
            },

            onError(e) {
                this.$message({
                    message: "复制失败!",
                    type: 'warning'
                })
            },

            //获取数据
            get_table_data(){
                this.load_data = true
                this.$http.get(port_conf.list, {
                            params: {
                                project_id: this.ProjectId,
                                search_text: this.search_text,
                            }
                        })
                .then(({data: {data}}) => {
                    this.table_data = data
                    for(var i in this.table_data) {
                        // this.table_data[i].json = JSON.parse(this.table_data[i].ConfValue);
                        this.$set(this.table_data[i], "json", JSON.parse(this.table_data[i].Value))
                    }
                    // console.log("------", this.table_data)
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
                })
                .then(() => {
                    this.load_data = true
                    this.$http.delete(port_conf.id, {
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
                })
                .then(() => {
                    this.load_data = true
                    this.$http.get(port_conf.copy, {
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
