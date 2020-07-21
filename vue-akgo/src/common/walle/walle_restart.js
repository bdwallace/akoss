/**
 * @file: tools_date.
 * @intro: uri编码工具类.
 * @author: zzmhot.
 * @email: zzmhot@163.com.
 * @Date: 2017/5/9 14:03.
 * @Copyright(©) 2017 by zzmhot.
 *
 */


 import store from "store"
 import {port_walle} from 'common/port_uri'
 
// var dateFormat = function(cellValue, format='YY-MM-DD hh:mm:ss') {
function restart(service, host) {
    this.$confirm('确定需要重启?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
        let walle_form = {
            User: store.state.user_info.user,
            Service: service,
            Host: host
        }
        this.$http.post(port_walle.restart, walle_form)
        .then(({data: {data}}) => {
            this.on_submit_loading = false
            this.load_data = false                                
            this.$message({
                message: "发送重启成功",
                type: 'success'
            })
        })
    })
}

export default {
    restart
}