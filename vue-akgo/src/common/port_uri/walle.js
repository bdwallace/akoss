//获取service的服务列表
exports.taglist = "/api/walle/taglist"
//获致service的host主机状态
exports.dockerps = "/api/walle/dockerps"
//获致下线的主机列表，用于批量比对显示上下线主机
exports.lineget = "/api/walle/lineget"
//改变上下线主机状态
exports.linechange = "/api/walle/linechange"
//上传服务包
exports.upload = "/api/walle/upload"
//上传服务包
exports.download = "/api/walle/download"
//编译上传服务包
exports.build = "/api/walle/build"
//重载docker
exports.reload = "/api/walle/reload"
//重启docker
exports.restart = "/api/walle/restart"
//停止docker
exports.stop = "/api/walle/stop"