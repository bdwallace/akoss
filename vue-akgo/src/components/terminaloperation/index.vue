<template>
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

</template>
<script type="text/javascript">
    import {port_recordoperation} from 'common/port_uri'
    export default{
        props: ['serviceId', 'clas'],
        data(){
            return {
                showText: [],
                action: 0,
                time: (Date.parse(new Date()) / 1000)-10
            }
        },
        created () {
            this.get_data()
            var _this = this;
            this.intervalid1 = setInterval(function () {
                _this.get_data()
            }, 4000)
        },
        beforeDestroy () {
            this.time=(Date.parse(new Date()) / 1000);
            clearInterval(this.intervalid1)
        },
        methods: {
            get_data() {
                this.$http.get(port_recordoperation.list, {
                    params: {
                        service_id: this.serviceId,
                        class: this.clas
                    }
                }).then(({data: {data}}) => {
                    console.log("------data---", data)
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

                        if(this.taskMultiId !== undefined) {
                            if ((data[i].Action / 10) === 10) {
                                action_nub = action_nub + 1
                            }
                            task_multi_action = data[i].task_multi_action
                        } else {
                            if (action < (data[i].Action / 10)) {
                                action = data[i].Action / 10
                            }
                        }
                    }

                }).catch(() => {
                })
            },
           formatJson(json, options) {
                var reg = null,
                        formatted = '',
                        pad = 0,
                        PADDING = '    '; // one can also use '\t' or a different number of spaces
                // optional settings
                options = options || {};
                // remove newline where '{' or '[' follows ':'
                options.newlineAfterColonIfBeforeBraceOrBracket = (options.newlineAfterColonIfBeforeBraceOrBracket === true) ? true : false;
                // use a space after a colon
                options.spaceAfterColon = (options.spaceAfterColon === false) ? false : true;
                // begin formatting...
                if (typeof json !== 'string') {
                    // make sure we start with the JSON as a string
                    json = JSON.stringify(json);
                } else {
                    // is already a string, so parse and re-stringify in order to remove extra whitespace
                    json = JSON.parse(json);
                    json = JSON.stringify(json);
                }
                // add newline before and after curly braces
                reg = /([\{\}])/g;
                json = json.replace(reg, '\r\n$1\r\n');
                // add newline before and after square brackets
                reg = /([\[\]])/g;
                json = json.replace(reg, '\r\n$1\r\n');
                // add newline after comma
                reg = /(\,)/g;
                json = json.replace(reg, '$1\r\n');
                // remove multiple newlines
                reg = /(\r\n\r\n)/g;
                json = json.replace(reg, '\r\n');
                // remove newlines before commas
                reg = /\r\n\,/g;
                json = json.replace(reg, ',');
                // optional formatting...
                if (!options.newlineAfterColonIfBeforeBraceOrBracket) {
                    reg = /\:\r\n\{/g;
                    json = json.replace(reg, ':{');
                    reg = /\:\r\n\[/g;
                    json = json.replace(reg, ':[');
                }
                if (options.spaceAfterColon) {
                    reg = /\:/g;
                    json = json.replace(reg, ':');
                }
                for(var index in json.split('\r\n')){
                    var node=json.split('\r\n')[index]
                    var i = 0,
                            indent = 0,
                            padding = '';
                    if (node.match(/\{$/) || node.match(/\[$/)) {
                        indent = 1;
                    } else if (node.match(/\}/) || node.match(/\]/)) {
                        if (pad !== 0) {
                            pad -= 1;
                        }
                    } else {
                        indent = 0;
                    }
                    for (i = 0; i < pad; i++) {
                        padding += PADDING;
                    }
                    formatted += padding + node + '\r\n';
                    pad += indent;
                }
                return formatted;
            }
        }
    }
</script>
