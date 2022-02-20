<template>
  <div class="panel">
    <panel-title :title="$route.meta.title">
    </panel-title>


    <div class="panel" style="margin-top:5px;margin-bottom:15px;padding:15px;weight:500px">
      <H2> 域名信息：</H2>
         <el-input type="textarea" v-model.trim="this.DomainInfo" style="width: 599px;"></el-input>
      <H2> 域名解析信息：</H2>
      <pre>
          {{this.DomainRecord}}
        </pre>
    </div>
  </div>


</template>

<script type="text/javascript">
  import {panelTitle} from 'components'
  import store from 'store'

  export default {
    data() {
      return {
        ProjectId: store.state.user_info.user.ProjectId,
        Domain: "",
        DomainInfo: "",
        DomainRecord: "",

      }
    },
    created() {
      this.getParams()
    },
    methods: {
      getParams(){
        this.Domain = this.$route.query
        this.DomainInfo = JSON.parse(this.Domain.DomainInfo)
        this.DomainRecord = JSON.parse(this.Domain.DomainRecord)


      },
    },
    components: {
      panelTitle,
    }
  }
</script>
