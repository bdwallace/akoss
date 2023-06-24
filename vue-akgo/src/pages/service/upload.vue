<template>
    <div class="panel">
        <panel-title :title="$route.meta.title"></panel-title>
        <div class="panel-body"
             v-loading="load_data"
             element-loading-text="拼命加载中">
            <el-form label-width="100px">
                <el-row>
                    <el-col :span="5">
                        <el-form-item label="任务标题:">
                            {{ Title }}
                        </el-form-item>

                        <el-form-item label="项目名称:">
                            {{service.Name}}
                        </el-form-item>

                        <el-form-item label="页面上传:">
                            <el-upload
                                class="upload-demo"
                                action="http://127.0.0.1:8192"
                                accept=".zip"
                                :http-request = "handleGetFile">
                                <el-button size="small" type="primary">点击上传</el-button>
                            </el-upload>
                        </el-form-item>

                        <!-- <el-form-item>
                            <el-button type="primary" @click="on_submit_build" :loading="on_submit_loading">Build
                            </el-button>
                            <el-button type="primary" @click="on_submit_form" :loading="on_submit_loading">快速部署
                            </el-button>
                            <el-button @click="$router.back()">取消</el-button>
                        </el-form-item> -->
                    </el-col>
                    
                    <el-col :span="19">
                        <el-form-item label="部署路径:">
                            {{service.ReleaseTo}}
                        </el-form-item>

                        <el-form-item label="绑定域名:">
                            <span v-for="domain in noQuickenDomain" :key="domain">
                                {{ domain.Domain }}、
                            </span>
                        </el-form-item>

                        <el-form-item label="Build ip:">
                            <!-- <span v-for="host in service.Hosts" :key="host"> -->
                                <div style="display: inline-block;margin-right: 20px" v-if="service.Hosts[0].UseIp == service.Hosts[0].PrivateIp">
                                    <el-input readonly style="width: 120px" size="mini" v-model="service.Hosts[0].PublicIp" >
                                    </el-input>
                                    <el-input readonly class="input_blue" style="width: 120px" size="mini" v-model="service.Hosts[0].PrivateIp" >
                                    </el-input>
                                </div>

                                <div style="display: inline-block;margin-right: 20px" v-else>
                                    <el-input readonly class="input_blue" style="width: 130px" size="mini" v-model="service.Hosts[0].PublicIp" >
                                    </el-input>
                                    <el-input readonly style="width: 130px" size="mini" v-model="service.Hosts[0].Private" >
                                    </el-input>
                                </div>
                            <!-- </span> -->
                        </el-form-item>

                    </el-col>
                </el-row>
                <el-row>
                    <el-col :span="5">
                        <el-form-item>
                            <el-button type="primary" @click="on_submit_build" :loading="on_submit_loading">
                                Build
                            </el-button>
                            <el-button type="primary" @click="on_submit_form" :loading="on_submit_loading">
                                快速部署
                            </el-button>
                        <!-- <el-button @click="$router.back()">取消</el-button> -->

                    <!-- </div> -->
                        </el-form-item>
                    </el-col>
                    <el-col :span="19">
                        <el-form-item label="tag标签:">
                            {{service.Tag}}
                        </el-form-item>
                    </el-col>
                </el-row>
            </el-form>

            <el-tabs type="border-card" v-model="count" @tab-click="check_tag">
                <el-tab-pane :label="`新${Class}`" :name="`0`">
                </el-tab-pane>


                <el-tab-pane v-for="item in itemCount" :key="item" :name="`${parseInt(item)}`" :label="`第${item}次${Class}`">
                </el-tab-pane>
            </el-tabs>
            <!-- <terminaloperation :serviceId="route_id" :clas='Class' v-if="!on_submit_loading"></terminaloperation> -->
                <div   style="margin: 5px 5px 0px;
                            padding: 3px;
                            border: 1px dashed rgb(0, 160, 198);
                            background-color: rgb(0,0,0);">
                    <code style="background-color: rgb(0, 0, 0);color:#00ff00">
                        <br>
                        <span v-for="n in showText" :style="{'color': n.color}" :key="n"> <pre style=" white-space: pre-wrap;" v-html="n.text"></pre> <br></span>
                        <br>
                    </code>
                </div>
        </div>
    </div>
</template>
<script type="text/javascript">
    import {panelTitle, terminaloperation} from 'components'
    import {port_walle, port_service, port_deploy, port_recordoperation, port_domain} from 'common/port_uri'
    import {tools_verify} from 'common/tools'
    import store from 'store'
    export default{
        data(){
            return {
                // host_in: this.$route.params.host_id,
                service: {},
                Title: "",
                Class: "build",
                check: 1,
                route_id: this.$route.params.id,
                load_data: false,
                on_submit_loading: false,

                noQuickenDomain: [],

                intervalid1: null,
                itemCount: [],
                count: 0,
                showText: [],

                rules: {
                    Branch: [{required: true, message: '分支不能为空', trigger: 'blur'}],
                    CommitId: [{required: true, message: 'Commit不能为空', trigger: 'blur'}],
                    Title: [{required: true, message: '标题不能为空', trigger: 'blur'}]
                }
            }
        },
        created(){
            if (this.route_id) {
                this.get_service()
                this.get_itemCount()
            } else {
                this.$message({
                    message: "任务id不存在",
                    type: 'warning'
                })
                setTimeout(() => {
                    this.$router.back()
            },
                500
            )
            }
        },
        methods: {
            get_data() {
                this.$http.get(port_recordoperation.list, {
                    params: {
                        serviceId: this.route_id,
                        class: this.Class,
                        count: this.count
                    }
                }).then(({data: {data}}) => {
                    // console.log("------data---", data)
                    this.showText = [];
                    var action = 0
                    var action_nub = 0
                    var task_multi_action = 0
                    for (var i = 0; i < data.length; i++) {
                        //var text=data[i].command+"<br>"+data[i].memo
                        // var color = "#00ff00"
                        var color = "#00ff00"
                        if (data[i].Status == "0") {
                            color = "red"
                        }
                        this.showText.unshift({text: data[i].Command, "color": "#FDFEFE"})
                        var text=data[i].Memo;
                        try{
                            var text= JSON.parse(data[i].Memo);
                        }catch (e){
                        }
                        if(typeof text == "string"){
                            this.showText.unshift({text: data[i].Memo, "color": color})
                            this.showText.unshift({text: "=============", "color": "#FDFEFE"})
                        }else if(Object.prototype.toString.call(text)=='[object Array]') {
                            for(var j=0;j<text.length;j++){
                                try{
                                    this.showText.unshift({text: "IP:"+text[j].Host, "color": color})
                                    if(text[j].ErrorInfo) {
                                        this.showText.unshift({text: "错误结果:\n" + text[j].Result, "color": color})
                                    }else{
                                        this.showText.unshift({text: "执行结果:\n" + text[j].Result, "color": color})
                                    }
                                    if(text[j].ErrorInfo){
                                        this.showText.unshift({text: "错误:"+text[j].ErrorInfo, "color": color})
                                    }
                                    this.showText.unshift({text: "=============", "color": "#FDFEFE"})
                                }catch (e){
                                }
                            }
                            // 打所有信息
                            //this.showText.unshift({text: this.formatJson(text), "color": color})
                        }else{
                            this.showText.unshift({text: "\t执行结果:\n"+text.Result, "color": color})
                            if(text.ErrorInfo){
                                this.showText.unshift({text: "\t错误:"+text.ErrorInfo, "color": color})
                            }
                            this.showText.unshift({text: "=============", "color": "#FDFEFE"})
                            // 打所有信息
                            // this.showText.unshift({text: this.formatJson(text), "color": color})
                        }

                        if (data[i].Action >= 100) {
                            this.clear_interval()
                            this.showText.unshift({text: "发布完成", "color": color})
                        }
                        
                    }

                }).catch(() => {
                })
            },
            
            check_tag(tag, event) {
                this.clear_interval()
                if (this.count == 0) {
                    // 准备时，清屏
                    this.showText = []
                } else {
                    // 查看历史记录时，开启间隔刷新日志
                    this.set_interval()
                    // 并马上刷新日志
                    this.get_data()
                }

            },

            set_interval() {
                var _this = this;
                this.intervalid1 = setInterval(function () {
                    _this.get_data()
                }, 3000)
            },

            clear_interval() {
                // setTimeout( () => clearInterval(this.intervalid1), 5000)
                clearInterval(this.intervalid1)
            },

            get_itemCount() {
                this.load_data = true

                this.$http.get(port_recordoperation.count, {
                    params: {
                        serviceId: this.route_id,
                        class: this.Class,
                    }
                })
                .then(({data: {data}}) => {
                    this.itemCount = data

                    this.load_data = false
                }).catch(() => {
                        this.load_data = false
                })

            },

            handleGetFile(data) {
                this.on_submit_loading = true
                var name = this.service.Name + ".zip"
                var formdata = new FormData()
                formdata.append("file", data.file)
                formdata.append("name", name)

                this.$http.post(port_walle.upload, formdata)
                    .then(({data: {msg}}) => {
                        this.check = 0
                        this.on_submit_loading = false

                        this.$message({
                            message: msg,
                            type: 'success'
                        })

                        setTimeout(() => {
                        },
                        500
                    )}
                    )
            },

            get_service(){
                this.load_data = true
                this.$http.get(port_service.id, {
                    params: {
                        id: this.route_id,
                    }
                })
                .then(({data: {data}}) => {
                    this.service = data

                    if( this.service.Hosts.length == 0){
                        this.$message.error('项目主机为空！请联系管理员为项目添加主机！')
                        this.$router.back()
                    }

                    this.Title = this.service.Name + "_页面上传" 
                    this.get_noQuickenDomain()

                    this.load_data = false
                })
                .catch(() => {
                    this.load_data = false
                })
            },
            
            get_noQuickenDomain() {
                this.load_data = true
                this.$http.get(port_domain.noquickenlist, {
                                    params: {
                                        platform_id: this.service.Platforms[0].Id,
                                        class: this.service.Class
                                    }
                                })
                .then(({data: {data}}) => {
                    this.load_data = true
                    this.noQuickenDomain = data
                    this.load_data = false
                })
                .catch(() => {
                    this.load_data = false
                })

            },

            //build
            on_submit_build(){
                if( this.check == 1) {
                    this.$message({message: '请先上传！', type: 'warning'});
                    return
                }
                // this.on_submit_loading = true
                // this.load_data = true
                let newCount = 1
                if(this.itemCount == null ) {
                    this.itemCount = [newCount]
                } else {
                    newCount = this.itemCount[0] + 1
                    this.itemCount.splice(0, 0, newCount)
                }

                // 开启下一次发布，即开始刷新日志
                this.count = newCount + ""
                this.set_interval()

                let walle = {
                    Class: "build",
                    Count: newCount,
                    User: store.state.user_info.user,
                    Service: this.service,
                    Host: this.service.Hosts[0]}
                this.$http.post(port_walle.build, walle)
                .then(({data: {data}}) => {
                    this.service.Tag = data
                    this.on_submit_loading = false
                    this.load_data = false                                
                    this.$message({
                        message: "发送build成功",
                        type: 'success'
                    })
                })
                .catch(() => {
                    this.on_submit_loading = false
                    this.load_data = false
                })
            },

            //提交
            on_submit_form(){
                this.load_data = true
                this.on_submit_loading = true
                this.$http.post(port_deploy.savebuild, this.service)
                .then(({data: {data}}) => {
                    // this.$message({
                    //     message: msg,
                    //     type: 'success'
                    // })
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

        },

        components: {
            panelTitle,
            terminaloperation
        }
    }
</script>
