<template>

    <div class="panel-body" style="padding:15px">
        <span> {{$route.meta.title}} </span>
        <!-- <panel-title :title="$route.meta.title"> -->
        <panel-title style="margin-top: -25px">
            <el-radio-group style="float: left;margin-right: 100px;margin-top: 15px" v-model="ProjectId" @change="get_table_data()">
            <el-radio v-for="item in project_list" :key="item.Id" :label="item.Id">{{item.Name}}</el-radio>
            </el-radio-group>

            <el-button @click.stop="on_refresh" size="small">
                <i class="fa fa-refresh"></i>
            </el-button>
            <router-link :to="{name: 'sgroupSave'}" tag="span">
                <el-button type="primary" icon="plus" size="small">新增白名单</el-button>
            </router-link>
        </panel-title>

        <div class="panel-body" style="clear: both;margin-top:5px;margin-bottom:15px;padding:15px;">
            <el-collapse>
                <!-- <el-collapse-item disabled v-for="(list,index) in table_data" :key="list.ToPort" :title="list.title" :name="index"> -->
            <el-collapse-item disabled v-for="(group,index) in table_data" :key="group.GroupName" :title="group.GroupName" :name="index">
                <el-collapse>
                <el-collapse-item disabled v-for="(list,listindex) in group.IpPermissions" :key="list.ToPort" :title="list.title" :name="listindex" style="margin-left:50px">

                    <el-table :data="list.IpRanges">
                        <el-table-column
                            prop="CidrIp"
                            label="IP范围"
                            width="300">
                        </el-table-column>
                        <el-table-column
                            prop="Description"
                            label="说明"
                            width="300">
                        </el-table-column>
                        <el-table-column
                            label="操作"
                            width="230">
                            <template scope="props">
                                <el-button style="margin-left:0px" type="danger" size="small" icon="delete" @click="set_Sgroup(index, listindex, props.row)">删除
                                </el-button>
                            </template>
                        </el-table-column>
                    </el-table>

                </el-collapse-item>
                </el-collapse>
            </el-collapse-item>
            </el-collapse>
        </div>
    </div>
</template>
<script type="text/javascript">
    import {panelTitle, bottomToolBar, search} from 'components'
    import {port_sgroup, port_project} from 'common/port_uri'
    import store from "store"

    export default{
        data(){
            return {
                project_list: [],
                ProjectId: store.state.user_info.user.ProjectId,
                // project:
                on_submit_loading: false,
                table_data: [],
                //请求时的loading效果
                load_data: true,
                //项目详情
            }
        },
        components: {
            panelTitle,
            bottomToolBar,
            search

        },
        created(){
            this.get_project_list()
            this.get_table_data()
        },
        methods: {
            //刷新
            on_refresh(){
                this.get_table_data()
            },

            get_project_list() {
                this.load_data = true
                this.$http.get(port_project.list)
                        .then(({data: {data}}) => {
                    this.project_list = data
                    this.load_data = false
                })
            },

            //获取数据
            get_table_data(){
                this.load_data = true
                this.$http.get(port_sgroup.list, {
                            params: {
                                projectId: this.ProjectId,
                            }
                        })
                .then(({data: {data}}) => {
                    this.table_data = data.SecurityGroups
                    var title = ""
                    for(var i in this.table_data) {
                        for(var j in this.table_data[i].IpPermissions) {
                            title = this.table_data[i].IpPermissions[j].ToPort + " 端口"
                            this.$set(this.table_data[i].IpPermissions[j], 'title', title)
                        }
                    }
                    this.load_data = false
                })
                .catch(() => {
                    this.$forceUpdate()
                    this.load_data = false
                })
            },

            set_Sgroup(index, listindex, id){
                this.$confirm('此操作将删除该数据, 是否继续?', '提示', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning'
                }).then(() => {
                this.load_data = true
                this.on_submit_loading = true
                let revoke_data = {
                    ProjectId: this.ProjectId,
                    Data: {
                        GroupId: this.table_data[index].GroupId,
                        IpPermissions: [{
                        FromPort: this.table_data[index].IpPermissions[listindex].FromPort,
                        ToPort: this.table_data[index].IpPermissions[listindex].ToPort,
                        IpProtocol: this.table_data[index].IpPermissions[listindex].IpProtocol,
                        IpRanges: [ id ]
                        }]
                    }
                }
                this.$http.post(port_sgroup.del, revoke_data)
                        .then(({data: {data}}) => {
                            this.get_table_data()
                    this.$message({
                        message: "删除安全组成功",
                        type: 'success'
                    })
                    this.load_data = false
                    this.on_submit_loading = false
            })
            .
                catch(() => {
                    this.load_data = false
                    this.on_submit_loading = false
                })
            })
            },

        },

    }
</script>
