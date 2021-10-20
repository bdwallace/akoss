<template>
  <el-form ref="form" :model="form" >
    <div class="panel">
      <panel-title :title="$route.meta.title">
      </panel-title>
    </div>


    <div class="panel" style="margin-top:5px;margin-bottom:15px;padding:15px;weight:500px">
<!--      <el-form-item label=$itemConf.conf_full_name prop= label-width="100px">-->
<!--      </el-form-item>-->
      <h3>{{itemConf.conf_file_name}}</h3>
      <el-input type="textarea" autosize v-model="itemConf.conf_info" style="width: 1000px">
      </el-input>

    </div>

    <div class="panel" style="margin-top:5px;margin-bottom:15px;padding:15px">


      <el-form-item>
        <el-button type="primary" @click="submitConf" :loading="on_submit_loading">立即提交
        </el-button>
        <el-button @click="hiddenShow">取消</el-button>
      </el-form-item>
    </div>
  </el-form>
</template>
<script type="text/javascript">
  import {panelTitle} from 'components'
  import {port_platform, port_domain, port_conf, port_host, port_service,port_project} from 'common/port_uri'
  import store from 'store'

  export default {
    data() {

      return {
        ProjectId: store.state.user_info.user.ProjectId,
        form: {
          name: "",
          path: "",
          info: "",
          service_id: 0,
          host_ip: "",
        },

        itemConf: [],

        load_data: false,
        on_submit_loading: false,

      }
    },
    created() {
      this.getParams()

    },
    methods: {
      getParams(){
        this.itemConf = this.$route.query
      },

      submitConf() {
        // this.on_submit_loading = true

        this.form.name = this.itemConf.conf_file_name,
          this.form.path = this.itemConf.conf_full_path,
          this.form.info = this.itemConf.conf_info,
          this.form.host_ip = this.itemConf.host_use_ip,
          this.form.service_id = this.itemConf.conf_service_id,

          this.$confirm('此操作将修改配置文件并reload nginx,是否继续？', '提示', {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning'
          }).then(() => {
            this.load_data = true
            this.$http.post("/api/confManage/nginx", this.form)
              .then(({data: {msg}}) => {
                this.on_submit_loading = false
                this.$message({
                  message: msg,
                  type: 'success'
                })
/*                setTimeout(() => {
                    this.load_data = true
                    this.$router.back()
                  },
                  500
                )*/
              }).catch(() => {
              this.load_data = false
              this.on_submit_loading = false
            })
            this.load_data = false
            this.on_submit_loading = false
            window.close()
          })
      },


      hiddenShow: function () {
        window.close()
      },

    },

  components: {
      panelTitle
    }
  }

function sleep(d) {
  for(var t = Date.now(); Date.now() - t <= d;);
}

</script>

