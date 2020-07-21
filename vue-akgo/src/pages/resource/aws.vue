
<template>

    <div class="panel">
        <panel-title :title="$route.meta.title">
        </panel-title>

        <div class="panel-body-line" style="clear: both;">
            <div style="margin:　10px">
                <el-select v-model="awsRegion" style="float: left;margin-right: 10px" placeholder="请选择" @change="get_table_data()">
                    <el-option
                        v-for="item in awsRegionItem"
                        :key="item.name"
                        :label="item.label"
                        :value="item.name">
                    </el-option>
                </el-select>

                <el-radio-group v-model="server" style="float: left;margin-right: 10px" @change="get_table_data()">
                    <el-radio-button v-for="serveritem in serverItem" :key="serveritem" :label="serveritem.name">
                        {{serveritem.label}}
                    </el-radio-button>
                </el-radio-group>

                <!-- <el-button @click.stop="on_refresh" style="float: right">
                    <i class="fa fa-refresh"></i>
                </el-button>

                <div style="margin-right: 10px;width: 280px;float: right">
                    <search @search="submit_search"></search>
                </div> -->
            </div>

            </div>
            <el-table v-if="server == 'ec2'"
                key="ec2"
                :data="table_data"
                style="width: 1000px;">
                <el-table-column type="expand">
                <template scope="props">
                            <el-form label-position="left" class="demo-table-expand">
                            <el-form-item label="实例ID:" label-width="100px" style="margin: 0px">
                                <span>{{ props.row.InstanceID }}</span>
                            </el-form-item>
                            <el-form-item label="子网:" label-width="100px" style="margin: 0px">
                                <span>{{ props.row.SubnetName }}</span>
                            </el-form-item>
                            <el-form-item label="安全组" label-width="100px">
                                <span v-for="ec2group in props.row.SecurityGroups" :key="ec2group">{{ ec2group.SecurityGroupName}}，</span>
                            </el-form-item>
                        </el-form>
                    </template>
                </el-table-column>

                <el-table-column
                    type="index"
                    width="80">
                </el-table-column>

                <el-table-column
                    label="实例名"
                    prop="InstanceName">
                </el-table-column>

                <el-table-column
                    label="私网IP"
                    prop="PrivateIP"
                    width="140">
                </el-table-column>

                <el-table-column
                    label="公网IP"
                    prop="PublicIP"
                    width="140">
                </el-table-column>

                <el-table-column
                    label="实例类型"
                    prop="InstanceType"
                    width="120">
                </el-table-column>

                <el-table-column
                    label="状态"
                    width="90">
                    <template scope="props">
                        <el-tag v-if="props.row.Status == 'running'" type="primary">{{props.row.Status}}</el-tag>
                        <el-tag v-else-if="props.row.Status == 'stopped'" type="danger">{{props.row.Status}}</el-tag>
                        <el-tag v-else type="warning">{{props.row.Status}}</el-tag>
                    </template>
                </el-table-column>
            </el-table>

            <el-table v-else-if="server == 'vpc'"
                key="vpc"
                :data="table_data"
                style="width: 850px;">
                <el-table-column
                    type="index">
                </el-table-column>

                <el-table-column
                    label="ID"
                    prop="ID">
                </el-table-column>

                <el-table-column
                    label="子网名"
                    prop="Name">
                </el-table-column>

                <el-table-column
                    label="网段"
                    prop="CidrBlock">
                </el-table-column>

                <el-table-column
                    label="状态"
                    width="100">
                    <template scope="props">
                        <el-tag v-if="props.row.Status == 'available'" type="primary">{{props.row.Status}}</el-tag>
                        <el-tag v-else type="warning">{{props.row.Status}}</el-tag>
                    </template>
                </el-table-column>
            </el-table>

            <el-table v-else-if="server == 'subnet'"
                key="subnet"
                :data="table_data"
                style="width: 850px;">
                <el-table-column
                    type="index"
                    width="80">
                </el-table-column>

                <el-table-column
                    label="ID"
                    prop="ID">
                </el-table-column>

                <el-table-column
                    label="子网名"
                    prop="Name">
                </el-table-column>

                <el-table-column
                    label="网段"
                    prop="CidrBlock">
                </el-table-column>

                <el-table-column
                    label="状态"
                    width="100px">
                    <template scope="props">
                        <el-tag v-if="props.row.Status == 'available'" type="primary">{{props.row.Status}}</el-tag>
                        <el-tag v-else type="warning">{{props.row.Status}}</el-tag>
                    </template>
                </el-table-column>
            </el-table>

            <el-table v-else-if="server == 'rds'"
                key="rds"
                :data="table_data"
                style="width: 850px;">
                <el-table-column type="expand">
                <template scope="props">
                            <el-form label-position="left" class="demo-table-expand">
                            <el-form-item label="连接地址" label-width="100px" style="margin: 0px">
                                <span>{{ props.row.EndpointAddr}}</span>
                            </el-form-item>
                            <el-form-item label="安全组" label-width="100px">
                                <span v-for="rdsgroup in props.row.SecurityGroups" :key="rdsgroup">{{ rdsgroup.SecurityGroupName}}，</span>
                            </el-form-item>
                        </el-form>
                    </template>
                </el-table-column>

                <el-table-column
                    type="index"
                    width="80">
                </el-table-column>

                <el-table-column
                    label="实例名"
                    prop="InstanceIdentifier">
                </el-table-column>

                <el-table-column
                    label="实例类型"
                    prop="InstanceType">
                </el-table-column>

                <el-table-column
                    label="状态"
                    width="100px">
                    <template scope="props">
                        <el-tag v-if="props.row.InstanceStatus == 'available'" type="primary">{{props.row.InstanceStatus}}</el-tag>
                        <el-tag v-else type="warning">{{props.row.InstanceStatus}}</el-tag>
                    </template>
                </el-table-column>
            </el-table>

            <el-table v-else-if="server == 'redis'"
                key="redis"
                :data="table_data"
                style="width: 850px;">
                <el-table-column type="expand">
                <template scope="props">
                            <el-form label-position="left" class="demo-table-expand">
                            <!-- <el-form-item label="连接地址" label-width="100px" style="margin: 0px">
                                <span>{{ props.row.EndpointAddr}}</span>
                            </el-form-item> -->
                            <el-form-item label="安全组" label-width="100px">
                                <span v-for="redisgroup in props.row.CacheSecurityGroups" :key="redisgroup">{{ redisgroup.SecurityGroupName}}，</span>
                            </el-form-item>
                        </el-form>
                    </template>
                </el-table-column>

                <el-table-column
                    type="index"
                    width="80">
                </el-table-column>

                <el-table-column
                    label="群集名"
                    prop="CacheClusterName">
                </el-table-column>

                <el-table-column
                    label="群集类型"
                    prop="CacheClusterType">
                </el-table-column>

                <el-table-column
                    label="状态"
                    width="100px">
                    <template scope="props">
                        <el-tag v-if="props.row.CacheClusterStatus == 'available'" type="primary" key="redis">{{props.row.CacheClusterStatus}}</el-tag>
                        <el-tag v-else type="warning">{{props.row.CacheClusterStatus}}</el-tag>
                    </template>
                </el-table-column>
            </el-table>

            <el-table v-else-if="server == 'es'"
                key="es"
                :data="table_data"
                style="width: 850px;">
                <el-table-column type="expand">
                <template scope="props">
                            <el-form label-position="left" class="demo-table-expand">
                            <el-form-item label="连接地址" label-width="100px" style="margin: 0px">
                                <span>{{ props.row.EndpointsVPC}}</span>
                            </el-form-item>
                            <el-form-item label="安全组" label-width="100px">
                                <span v-for="esgroup in props.row.ClusterSecurityGroups" :key="esgroup">{{ esgroup.SecurityGroupName}}，</span>
                            </el-form-item>
                        </el-form>
                    </template>
                </el-table-column>

                <el-table-column
                    type="index"
                    width="80">
                </el-table-column>

                <el-table-column
                    label="群集名"
                    prop="DomainName">
                </el-table-column>

                <el-table-column
                    label="群集类型"
                    prop="DomainType">
                </el-table-column>

                <!-- <el-table-column
                    label="状态"
                    width="100px">
                    <template scope="props">
                        <el-tag v-if="props.row.CacheClusterStatus == 'available'" type="primary">{{props.row.CacheClusterStatus}}</el-tag>
                        <el-tag v-else type="warning">{{props.row.CacheClusterStatus}}</el-tag>
                    </template>
                </el-table-column> -->
            </el-table>

            <el-table v-else-if="server == 'kafka'"
                :data="table_data"
                style="width: 850px;">
                <el-table-column type="expand">
                <template scope="props">
                            <el-form label-position="left" class="demo-table-expand">
                            <!-- <el-form-item label="连接地址" label-width="100px" style="margin: 0px">
                                <span>{{ props.row.EndpointAddr}}</span>
                            </el-form-item> -->
                            <el-form-item label="安全组" label-width="100px">
                                <span v-for="kafkagroup in props.row.ClusterSecurityGroups" :key="kafkagroup">{{ kafkagroup.SecurityGroupName}}，</span>
                            </el-form-item>
                        </el-form>
                    </template>
                </el-table-column>

                <el-table-column
                    type="index"
                    width="80">
                </el-table-column>

                <el-table-column
                    label="集群名"
                    prop="ClusterName">
                </el-table-column>

                <el-table-column
                    label="集群类型"
                    prop="ClusterType">
                </el-table-column>

                <el-table-column
                    label="状态"
                    width="100px">
                    <template scope="props">
                        <el-tag v-if="props.row.ClusterStatus == 'ACTIVE'" type="primary">{{props.row.ClusterStatus}}</el-tag>
                        <el-tag v-else type="warning">{{props.row.ClusterStatus}}</el-tag>
                    </template>
                </el-table-column>
            </el-table>


            </div>

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
        <!-- </div>
    </div> -->
</template>
<script type="text/javascript">
    import {panelTitle, bottomToolBar, search} from 'components'
    import {port_resource} from 'common/port_uri'
    import store from 'store'
    
    export default{
        data(){
            return {
                awsRegion: "ap-southeast-1",
                awsRegionItem: [
                    { name: "ap-southeast-1", label: "亚太地区(新加坡)" },
                    { name: "ap-east-1", label: "亚太地区(香港)" },
                    { name: "ap-northeast-2", label: "亚太地区(首尔)" },
                    { name: "us-east-1", label: "美国东部(弗吉尼亚北部)" },
                    { name: "us-west-1", label: "美国西部(加利福尼亚北部)" },
                ],
                server: "ec2",
                serverItem: [
                    { name: "ec2", label: "云主机"},
                    { name: "vpc", label: "虚拟网络"},
                    { name: "subnet", label: "子网"},
                    { name: "rds", label: "RDS"},
                    { name: "redis", label: "REDIS"},
                    { name: "es", label: "ES"},
                    { name: "kafka", label: "KAFKA"},
                    // { name: "elb", label: "负载均衡"},
                ],
                project: store.state.user_info.user.Project,
                // project_data: [],
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
            submit_search(value) {
                this.select_info = value
                this.$message({
                    message: value,
                    type: 'success'
                })
                this.get_table_data()
            },

            on_refresh(){
                this.select_info = ""
                this.get_table_data()
            },

            //获取数据
            get_table_data(){
                this.load_data = true
                this.table_data = []
                this.table_data.splice(0, this.table_data.length)
                this.$http.get(port_resource.aws, {
                            params: {
                                select_info: this.select_info,
                                projectId: this.project,
                                awsRegion: this.awsRegion,
                                server: this.server
                            }
                        })
                        .then(({data: {data}}) => {
                this.table_data = data
                this.load_data = false
            })
            .
                catch(() => {
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
