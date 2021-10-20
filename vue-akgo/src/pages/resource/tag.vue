<template>
    <div class="panel">
        <panel-title :title="$route.meta.title">
        </panel-title>
        <div style="margin: 20px">
            <el-radio-group v-model="service_class" style="margin-right: 200px">
                <el-radio key="html" label="html">前端项目</el-radio>
                <el-radio key="java" label="java">后端项目</el-radio>
            </el-radio-group>

            <el-button type="primary" @click="get_table_data()" :loading="on_submit_loading">加载
            </el-button>

            <el-button @click="$router.back()">取消</el-button>

            <el-button v-if="download_path != ''" @click="download">导出数据</el-button>
        </div>
        <div class="panel-body">
            <el-table
                    id="mytable"
                    :data="table_data"
                    v-loading="load_data"
                    element-loading-text="拼命加载中"
                    border
                    style="width: 100%">
                <el-table-column
                    prop="0.ServiceName"
                    label="项目">
                </el-table-column>
                <el-table-column
                    prop="0.ServiceTag"
                    label="第一套Tag">
                </el-table-column>
                <el-table-column
                    prop="1.ServiceTag"
                    label="第二套Tag">
                </el-table-column>
<!--                <el-table-column-->
<!--                    prop="2.ServiceTag"-->
<!--                    label="第三套Tag">-->
<!--                </el-table-column>-->


<!--                <el-table-column-->
<!--                        prop="PsImage3"-->
<!--                        label="第三套Tag">-->
<!--                </el-table-column>-->
<!--                <el-table-column-->
<!--                        prop="ProjectStatus1"-->
<!--                        label="第一套详情"-->
<!--                        width="110">-->
<!--                    <template scope="props">-->
<!--                        <el-popover-->
<!--                            ref="ProjectStatus1"-->
<!--                            placement="left-start"-->
<!--                            width="600"-->
<!--                            trigger="click">-->
<!--                            <span style="color:teal;white-space: pre-wrap;">{{props.row.ProjectStatus1}}</span>-->
<!--                        </el-popover>-->
<!--                        <el-button v-show="props.row.ProjectStatus1 != ''" type="info" v-popover:ProjectStatus1>查看</el-button>-->
<!--                    </template>-->
<!--                </el-table-column>-->
<!--                <el-table-column-->
<!--                        prop="ProjectStatus2"-->
<!--                        label="第二套详情"-->
<!--                        width="110">-->
<!--                    <template scope="props">-->
<!--                        <el-popover-->
<!--                            ref="ProjectStatus2"-->
<!--                            placement="left-start"-->
<!--                            width="600"-->
<!--                            trigger="click">-->
<!--                            <span style="color:teal;white-space: pre-wrap;">{{props.row.ProjectStatus2}}</span>-->
<!--                        </el-popover>-->
<!--                        <el-button v-show="props.row.ProjectStatus2 != ''" type="info" v-popover:ProjectStatus2>查看</el-button>-->
<!--                    </template>-->
<!--                </el-table-column>-->
<!--                <el-table-column-->
<!--                        prop="ProjectStatus3"-->
<!--                        label="第三套详情"-->
<!--                        width="110">-->
<!--                    <template scope="props">-->
<!--                        <el-popover-->
<!--                            ref="ProjectStatus3"-->
<!--                            placement="left-start"-->
<!--                            width="600"-->
<!--                            trigger="click">-->
<!--                            <span style="color:teal;white-space: pre-wrap;">{{props.row.ProjectStatus3}}</span>-->
<!--                        </el-popover>-->
<!--                        <el-button v-show="props.row.ProjectStatus3 != ''" type="info" v-popover:ProjectStatus3>查看</el-button>-->
<!--                    </template>-->
<!--                </el-table-column>-->
            </el-table>
        </div>
    </div>
</template>
<script type="text/javascript">
    import FileSaver from "file-saver";
    import XLSX from "xlsx";
    import {panelTitle, bottomToolBar, search} from 'components'
    import {port_resource} from 'common/port_uri'
    export default{
        data(){
            return {
                on_submit_loading: false,
                service_class: "java",
                level: 1,
                date: [new Date(new Date() - 3600 * 1000 * 24 * 30), new Date()],
                index: 0,
                download_path: "",
                table_data: [],
                //请求时的loading效果
                load_data: false,
                //项目详情
                prometheus: "",
                pList: [],
            }
        },
        // created(){
        // },
        methods: {

            // 加载项目的tag列表
            get_table_data(){
                this.load_data = true
                this.$http.get(port_resource.tag, {
                            params: {
                                class: this.service_class
                            }
                        })
                .then(({data: {data}}) => {
                    this.table_data = data
                    // this.download_path = data.backPath
                    this.load_data = false
                })
                .catch(() => {
                    this.$forceUpdate()
                    this.load_data = false
                })

            },

            output() {
                const defaultCellStyle =  {
                    font: { name: "Verdana", sz: 11, color: "FF00FF88"},
                    // fill: {fgColor: {rgb: "FFFFAA00"}},
                    border: {color: {auto: 1}},
                    alignment: {wrapText: 1, horizontal: "center", vertical: "center", indent: 0}
                };
                /* generate workbook object from table */
                var wb = XLSX.utils.table_to_book(document.querySelector("#mytable"));//mytable为表格的id名
                /* get binary string as output */
                var wbout = XLSX.write(wb, {
                    bookType: "xlsx",
                    bookSST: true,
                    type: "array"
                }, { defaultCellStyle: defaultCellStyle });
                try {
                    FileSaver.saveAs(
                        new Blob([wbout], { type: "application/octet-stream" }),
                        "sheet.xlsx"
                    );
                } catch (e) {
                    if (typeof console !== "undefined") console.log(e, wbout);
                }
                return wbout;
            },

            download() {
                this.on_submit_loading = true
                window.location.href = this.download_path
                this.on_submit_loading = false
            },


        },

        components: {
            panelTitle,
            bottomToolBar,
            search

        },
    }
</script>
