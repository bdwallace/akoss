<template>
    <div class="panel">
        <panel-title :title="$route.meta.title"></panel-title>
        <div class="panel-body"
             v-loading="load_data"
             element-loading-text="拼命加载中">
            <div style="clear: both;margin: 10px;">
                <el-tag v-if="form != null"> Id:&nbsp;&nbsp;{{ form.Id }} </el-tag>
                <el-tag v-if="date != ''"> 创建时间:&nbsp;&nbsp;{{ date }} </el-tag>

                <div  v-if="!tag_post" style="display: inline">
                    <el-input v-model="unity" placeholder="统一版本" 
                        style="margin-left: 40px;width: 340px">
                    </el-input>

                    <el-button type="primary" size="mini"
                        @click="unity_tag" :loading="on_submit_loading">统一版本
                    </el-button>

                    <el-button type="success" size="mini"
                        @click="set_tasks_tag_all">全部回退
                    </el-button>
                </div>
            </div>

            <el-table 
                    v-if="form != null"
                    :data="form.Tasks"
                    v-loading="load_data"
                    element-loading-text="拼命加载中"
                    border
                    style="width: 100%;">
                <!-- <el-table-column
                        prop="Id"
                        label="id"
                        width="80">
                </el-table-column> -->

                <el-table-column
                    type="index"
                    width="40">
                </el-table-column>

                <el-table-column
                        prop="Service.Name"
                        label="项目名称"
                        width="170">
                </el-table-column>

                <el-table-column
                        prop="Service.Port"
                        label="端口"
                        width="90">
                    <template scope="props">
                        <font style="display: flex" v-for="item in props.row.Service.Port.split(',')" :key="item">{{item}}</font>
                    </template>
                </el-table-column>

                        <!-- :render-header="renderHeaderTag" -->
                <el-table-column
                        label="发布版本"
                        width="390">
                    <template scope="props">
                        <el-tag v-if="check_data[props.$index]" style="margin-left: -5px" type="success">
                            <i class="el-icon-circle-check"></i>
                        </el-tag>
                        <el-tag v-if="!check_data[props.$index]" style="margin-left: -5px" type="danger">
                            <i class="el-icon-circle-cross"></i>
                        </el-tag>

                        <el-input v-if="tag_post" v-model="props.row.Tag" size="mini" disabled readonly style="width: 280px">
              
                        </el-input>
                        <el-select v-model="props.row.Tag" 
                                v-else
                                filterable
                                size="mini" 
                                style="width: 280px"
                                @change="check(props.$index)">
                            <el-option
                                v-for="item in tag_data[props.$index]"
                                style="font-size: 5px;"
                                :key="item"
                                :label="item"
                                :value="item">
                            </el-option>
                        </el-select>

                        <el-popover
                            ref="tag"
                            placement="bottom"
                            width="80"
                            trigger="hover">
                            <span style="color:teal;white-space: pre-wrap;">
                                <font>{{props.row.Service.LastTag}}</font>
                            </span>
                        </el-popover>

                        <el-button v-if="!tag_post" type="success" size="mini" @click="props.row.Tag = props.row.Service.LastTag" v-popover:tag>回退版本</el-button>
                    </template>
                </el-table-column>

                <el-table-column
                    label="检测"
                    :render-header="renderHeaderStatus"
                    width="55">
                    <template scope="props">
                        <el-button type="success" size="mini" @click="get_service_status(props.$index)">检测</el-button>
                    </template>
                </el-table-column>

                <el-table-column
                    label="主机"
                    prop="ProjectStatus"
                    fit=true>
                    <template scope="props">
                        <div style="clear: both;margin-bottom: 10px;"> <span></span> </div>
                        <div v-for="(item, host_index) in service_data[props.$index]" :key="item"  style="clear:both">
                            <!-- <el-checkbox v-model="item.deploy">发布</el-checkbox> -->

                            <el-switch v-model="item.line" v-if="props.row.Service.Class == 'java'" size="mini"
                                on-color="#13ce66"
                                off-color="#ff4949"
                                on-value="on"
                                off-value="off"
                                on-text="上线"
                                off-text="下线"
                                @change="linechange(props.$index, item.ip, item.line)">
                            </el-switch>

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
                            <!-- <el-input readonly style="width: 140px" size="mini" v-model="item.ip_show" >
                            </el-input> -->
                            <!-- <el-input readonly style="width: 110px" size="mini" v-model="item.ps_status">
                            </el-input> -->

                            <el-input v-if="item.ps_status.substr(0, 2) != 'Up'" class="input_red" readonly style="width: 100px" size="mini" v-model="item.ps_status">
                            </el-input>
                            <el-input v-else readonly style="width: 100px" size="mini" v-model="item.ps_status">
                            </el-input>
                            <!-- <el-input readonly style="width: 150px" size="mini" v-model="item.ps_created_at">
                            </el-input> -->
                            <el-input readonly style="width: 250px" size="mini" v-model="item.ps_image">
                            </el-input>

                            <div style="display: inline-block;margin-right:10px">
                                <el-tooltip effect="light" :content="item.url" placement="left">
                                    <el-tag v-if="item.health=='200'" size="mini" type="primary" >健康</el-tag>
                                    <el-tag v-else size="mini" type="danger">异常</el-tag>
                                </el-tooltip>
                            </div>

                            <el-button v-if="props.row.Service.Class == 'java'" size="mini" type="primary" style="margin-right:-12px" @click="restart(props.row.Service, props.row.Service.Hosts[host_index])" :loading="on_submit_loading">重启
                            </el-button>    

                            <el-button v-else size="mini" type="primary" style="margin-right:-12px" @click="reload(props.row.Service, props.row.Service.Hosts[host_index])" :loading="on_submit_loading">重载
                            </el-button>    

                            <el-button size="mini" type="primary" style="margin-right:-12px" @click="stop(props.row.Service, props.row.Service.Hosts[host_index])" :loading="on_submit_loading">关闭
                            </el-button>  
                    </div>
                        <div style="clear: both;margin-bottom: 10px;"> <span></span> </div>
                    </template>
                </el-table-column>

                <el-table-column
                        prop="Cmd"
                        label="命令"
                        width="50">
                    <template scope="props">
                        <el-popover
                            ref="cmd"
                            placement="left-start"
                            width="800"
                            trigger="click">
                            <!-- <span style="color:teal;white-space: pre-wrap;">{{props.row.Cmd}}</span> -->
                            <el-input type="textarea" :rows="20" readonly v-model="props.row.Cmd"></el-input>
                        </el-popover>
                        <el-button v-show="props.row.Cmd != ''" type="info" size="mini" style="clear: both;margin:-5px" v-popover:cmd>查看</el-button>
                    </template>
                </el-table-column>
            </el-table>

            <el-tabs type="border-card" v-model="count" @tab-click="check_tag">
                <el-tab-pane label="新发布" :name="`0`">
                    <div style="clear: both;margin: 10px">
                        <el-button v-if="!tag_post" type="success" size="small" 
                            @click="post_tag" :loading="on_submit_loading">确定tag版本
                        </el-button>
                        <div v-else style="display:inline">
                            <el-button type="primary" v-if="form != null && form.Class == 'java'" @click="on_submit_form(true)" :loading="on_submit_loading" style="margin-right: 50px">灰度发布
                            </el-button>
                            <el-button @click="$router.back()">取消</el-button>

                            <el-button type="danger" @click="is_backDialog = true" style="float: right">
                                回退发布单
                            </el-button>

                            <el-button type="danger" v-if="form != null" style="float: right;margin-right:30px"
                                @click="on_submit_form(false)" :loading="on_submit_loading">全部发布
                            </el-button>
                        </div>

                    </div>
                </el-tab-pane>


                <el-tab-pane v-for="item in itemCount" :key="item" :name="`${parseInt(item)}`" :label="`第${item}次发布`">
                </el-tab-pane>
            </el-tabs>

            <el-dialog title="确认发布以下服务" 
                :visible.sync="is_backDialog"
                :modal="true"
                :show-close="false"
                :modal-append-to-body="false">
                <!-- <div  v-for="(item, index) in dialogForm" :key="item" style="width: 200px;display: inline-block;">
                    <el-tag :closable="true" type="primary" 
                        @close="dialogForm.splice(index, 1)"
                        style="align:online">
                        {{item.Name}}
                    </el-tag>
                </div> -->
                <el-checkbox-group v-if="form != null" v-model="dialogForm">
                    <el-checkbox v-for="item in form.Tasks" :label="item.Service" :key="item.Service">{{item.Service.Name}}</el-checkbox>
                </el-checkbox-group>

                <div slot="footer" class="dialog-footer">
                    <el-button @click="is_backDialog = false">取 消</el-button>
                    <el-button type="primary" @click="on_submit_form_rollback(dialogForm)">提 交</el-button>
                </div>
            </el-dialog>

            <div>
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
            
            <!-- <terminal v-if="count != 0" :deployId="route_id"></terminal> -->
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
    import {panelTitle, terminal} from 'components'
    import {port_deploy, port_walle, port_record} from 'common/port_uri'
    import {tools_date} from 'common/tools'
    import {walle_stop, walle_restart, walle_reload} from 'common/walle'
    import store from 'store'

    export default{
        data(){
            return {
                // level_data: [],
                unity: "",
                date: "",
                line_data: "",
                table_data: [],
                tag_data: [],
                tag_post: true,
                check_data: [],
                service_data: [],
                form: null,
                itemCount: [],
                count: 0,

                route_id: this.$route.params.id,
                load_data: false,
                on_submit_loading: false,

                intervalid1: null,
                intervalid2: null,

                showText: [],

                is_backDialog: false,
                dialogForm: [],

                health: [],
                health_all: true, 

                stop: walle_stop.stop,
                restart: walle_restart.restart,
                reload: walle_reload.reload,
            }
        },

        beforeDestroy () {
            clearInterval(this.intervalid1)
            clearInterval(this.intervalid2)
        },

        created(){
            if (this.route_id) {
                this.get_form()
                this.get_itemCount()
                // this.get_record()
                // this.set_interval_1()
                // var _this = this;
                // this.intervalid1 = setInterval(function () {
                //     if (_this.count != 0) {
                //         _this.get_record()
                //     }
                // }, 2000)
            } else {
                this.$message({
                    message: "任务id不存在",
                    type: 'warning'
                })
                setTimeout(() => {
                    this.$router.back()
                }, 500)
            }
        },

        methods: {
            check_tag(tag, event) {
                if (this.count == 0) {
                    // 准备时，清屏
                    this.showText = []
                } else {
                    // 查看历史记录时，开启间隔刷新日志
                    this.set_interval_1()
                    // 并马上刷新日志
                    this.get_record()
                }

            },

            set_interval_1() {
                var _this = this;
                this.intervalid1 = setInterval(function () {
                    _this.get_record()
                }, 4000)
            },

            clear_interval_1() {
                // setTimeout( () => clearInterval(this.intervalid1), 5000)
                clearInterval(this.intervalid1)
            },

            set_interval_2() {
                var _this = this;
                this.clear_interval_2()
                this.intervalid2 = setInterval(function () {
                    _this.get_service_status_all()
                }, 4000)
            },

            clear_interval_2() {
                // setTimeout( () => clearInterval(this.intervalid1), 5000)
                clearInterval(this.intervalid2)
            },
            
            get_record() {
                this.$http.get(port_record.listcount, {
                    params: {
                        deployId: this.route_id,
                        count: this.count
                    }
                }).then(({data: {data}}) => {
                    var action_nub = 0
                    var action = 0
                    var deploy = true
                    var failure_service = ""
                    this.showText = [];
                    console.log("--get_record--")
                    for (var i = 0; i < data.length; i++) {
                        // console.log("-----", data[i])
                        //var text=data[i].command+"<br>"+data[i].memo
                        // var color = "#00ff00"
                        var color = "#00ff00"
                        if (data[i].Status == 0) {
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

                        if (data[i].Action == 100) {
                            action_nub = action_nub + 1
                            let servcie_name = this.search_service_name(data[i].Task.Id)
                            this.showText.unshift({text: servcie_name+" 发布完成", "color": color})
                            this.showText.unshift({text: "=============", "color": "#FDFEFE"})
                        }

                        if (data[i].Action > 100) {
                            action_nub = action_nub + 1
                            let servcie_name = this.search_service_name(data[i].Task.Id)
                            this.showText.unshift({text: servcie_name+" 发布失败", "color": "red"})
                            this.showText.unshift({text: "=============", "color": "#FDFEFE"})
                            deploy = false
                            failure_service = failure_service + servcie_name + ", "
                        }
    
                    }
                    if(this.form.Tasks.length == action_nub ){
                        // setTimeout( () => clearInterval(this.intervalid1), 5000)
                        this.clear_interval_1()
                        if (deploy) {
                            this.showText.unshift({text: "发布结束, 状态成功。", "color": color})
                        } else {
                            this.showText.unshift({text: "发布结束, 状态失败！失败服务：" + failure_service, "color": "red"})
                        }
                    } else {
                        //如果判断没有发布完，自动去检察服务的状态
                        this.get_service_status_all()
                    }
                }).catch(() => {
                })
            },

            search_service_name(task_id) {
                for(let i in this.form.Tasks) {
                    if(this.form.Tasks[i].Id == task_id) {
                        return this.form.Tasks[i].Service.Name
                    }
                }
                return "task_id:" + task_id 
            },

            get_itemCount() {
                this.load_data = true

                this.$http.get(port_record.count, {
                    params: {
                        deployId: this.route_id,
                    }
                })
                .then(({data: {data}}) => {
                    
                    this.itemCount = data

                    this.load_data = false
                }).catch(() => {
                        this.load_data = false
                })

            },

            get_form(){
                this.load_data = true

                this.$http.get(port_deploy.id, {
                    params: {
                        id: this.route_id,
                    }
                })
                .then(({data: {data}}) => {
                    this.form = data
                    this.date = tools_date.dateFormat(this.form.CreatedAt)
                    for(var i in this.form.Tasks) {
                        this.get_tag_data(i)
                        this.get_service_status(i)
                        if( this.form.Tasks[i].Cmd == "") {
                            this.tag_post = false
                        }
                        //给dialogForm的参数默认全选
                        this.dialogForm.push(this.form.Tasks[i].Service)
                    }

                    this.load_data = false
                }).catch(() => {
                        this.load_data = false
                })
            },

            get_tag_data(index){
                this.$http.get(port_walle.taglist, {
                    params: {
                        image_path: this.form.Tasks[index].Service.ImagePath
                    }
                })
                .then(({data: {data}}) => {
                    this.$set(this.tag_data, index, data)
                    this.check(index)
                })
            },

            unity_tag() {
                if(this.unity == "") {
                    this.$message({
                        message: "请输入统一版本!",
                        type: "warning"
                    })
                    this.on_submit_loading = false
                    return
                }
                for(var i in this.form.Tasks) {
                    this.form.Tasks[i].Tag = this.unity
                }
            },

            check(index) {
                if (this.tag_data[index].indexOf(this.form.Tasks[index].Tag) == -1 ) {
                    this.$set(this.check_data, index, false)
                } else {
                    this.$set(this.check_data, index, true)
                }
            },

            // renderHeaderTag(h) {
            //     return (
            //         <div>
            //             <el-button type="success" size="small" 
            //             onClick={this.post_tag}>确定版本
            //             </el-button>
            //         </div>
            //     )
            // },

            // v-if并没有起作用
            //  <el-button v-if={ ! this.tag_post } type="success" size="small" style="float: right"
            renderHeaderTag(h) {
                return (
                    <div>
                        <span style="float: left">发布版本</span>
                        <el-button type="success" size="small" style="float: right"
                            onClick={this.set_tasks_tag_all}>全部回退
                        </el-button>
                    </div>
                )
            },

            set_tasks_tag_all() {
                this.load_data = true
                // 因为在renderHeaderTag方法中v-if无效，加入此判断，使在已经确定版本的情况下，重置为回退版本无效
                if (this.tag_post) {
                    this.load_data = false
                    return
                }
                for(var i in this.form.Tasks) {
                    this.form.Tasks[i].Tag = this.form.Tasks[i].Service.LastTag
                    // this.$set(this.form.Tasks[i], "Tag", this.form.Tasks[i].LastTag)
                }
                this.$forceUpdate();
                this.load_data = false
            },


            renderHeaderStatus(h) {
                return (
                    <div>
                        <el-button type="success" size="mini" 
                        onClick={this.get_service_status_all}>检测
                        </el-button>
                    </div>
                )
            },

            get_service_status_all() {
                this.load_data = true
                console.log("--get_service------")
                this.health_all = true
                for(var i in this.form.Tasks) {
                    this.get_service_status(i)
                    if(this.health[i] == false) {
                        this.health_all = false
                    }
                }
                //每次检查完状态，所有状态都健康后，就清理循环任务。
                if(this.health_all) {
                    this.clear_interval_2()
                }
                this.load_data = false
            },

            //根据项目ID获取项目的状态
            get_service_status(index){
                this.load_data = true
                // 如果是java项目，line_data为空就是需要检测上下线
                // 默认line_data "----"，就是前端项目，就不需要上下线。
                let line_data = "----"
                if (this.form.Tasks[index].Service.Class == "java") {
                    line_data = ""
                }
                this.$http.get(port_walle.dockerps, {
                            params: {
                                service_id: this.form.Tasks[index].Service.Id,
                                line_data: line_data
                            }
                        })
                .then(({data: {data}}) => {
                    this.$set(this.service_data, index, data)
                    for (let i in this.service_data[index]) {
                        //如果service有一个主机不健康就,就开始循环刷新整个发布的所以服务的状态
                       if (this.service_data[index][i].health != "200") {
                           this.set_interval_2()
                           this.health[index] = false
                       } else (
                           this.health[index] = true
                       )
                    }
                    this.load_data = false
                })
                this.load_data = false
            },

            linechange(index, host, line) {
                this.on_submit_loading = true
                this.load_data = true
                let host_port = host + ":" + this.form.Tasks[index].Service.Port.split(",")[0].split(":")[0]
                this.$http.get(port_walle.linechange, { 
                        params: {
                            nacos: this.form.Project.Nacos,
                            host_port: host_port,
                            line: line
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

            //提交tag
            post_tag(){
                this.on_submit_loading = true
                for (var i in this.form.Tasks) {
                    if (this.form.Tasks[i].Tag == "") {
                        this.$message({
                            message: "有服务没有选择版本",
                            type: 'warning'
                        })
                        this.on_submit_loading = false
                        return 
                    }
                }
                this.$http.post(port_deploy.tagcmd, this.form)
                .then(({data: {data}}) => {
                    this.$message({
                        message: "已提交版本",
                        type: 'success'
                    })
                    this.tag_post = true
                    this.get_form()
                    this.on_submit_loading = false
                })
                .catch(() => {
                    this.on_submit_loading = false
                })
            },

            //提交
            on_submit_form(gray){
                let confirmStr = ""
                if(gray) {
                    for (var i in this.service_data) {
                        this.form.Tasks[i].Hosts = []
                        for (var j in this.service_data[i]) {
                            if (this.service_data[i][j].line == "off") {
                                this.form.Tasks[i].Hosts.push(this.select_host(i, this.service_data[i][j].host_id))
                            }
                        }
                    }
                    confirmStr = "确定需要 灰度发布 吗？"
                } else {
                    for (var i in this.service_data) {
                        this.form.Tasks[i].Hosts = []
                        this.form.Tasks[i].Hosts = this.form.Tasks[i].Service.Hosts
                    }
                    confirmStr = "确定需要 全部发布 吗？"
                }

                this.$confirm(confirmStr, '提示', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning'
                }).then(() => {
                    this.on_submit_loading = true

                    let newCount = 1
                    if(this.itemCount == null ) {
                        this.itemCount = [newCount]
                    } else {
                        newCount = this.itemCount[0] + 1
                        this.itemCount.splice(0, 0, newCount)
                    }

                    // 开启下一次发布，即开始刷新日志
                    this.count = newCount + ""
                    this.set_interval_1()

                    this.form.Count = newCount
                    this.$http.post(port_deploy.deploy, this.form)
                    .then(({data: {msg}}) => {
                        // this.$message({
                        //     message: "已经开始发布!",
                        //     type: "success"
                        // })
                        this.on_submit_loading = false
                    })
                    .catch(() => {
                        // this.$message({
                        //     message: "修改发单失败!",
                        //     type: "error"
                        // })
                        this.on_submit_loading = false
                    })
                })
            },

            //找出主机下线要发布的主机
            select_host(index, host_id) {
                for (var i in this.form.Tasks[index].Service.Hosts) {
                    if (this.form.Tasks[index].Service.Hosts[i].Id == parseInt(host_id)) {
                        return this.form.Tasks[index].Service.Hosts[i]
                    }
                }
            },

            // set_rollback() {
            //     for

            // },
            
            //提交
            on_submit_form_rollback(deployform) {
                this.load_data = true
                this.on_submit_loading = true
                if (deployform.length == 0) {
                    this.$message({
                        message: "请选择发布的项目",
                        type: 'warning'
                    })
                    this.load_data = false
                    this.on_submit_loading = false
                    return
                }
                this.$http.post(port_deploy.save, deployform)
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
        },
        components: {
            panelTitle,
            terminal
        }
    }
</script>
