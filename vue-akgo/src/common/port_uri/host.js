//数据列表
// exports.list = "/api/host/projectId"
//含搜索的数据列表
exports.list = "/api/host/list"
//关联主机时的专用列表
exports.project = "/api/host/projectId"
//根据id,get查询,put修改,delete删除数据
exports.id = "/api/host/id"
//添加数据
exports.save = "/api/host"
//复制数据,但IP与公网IP都有唯一属性,并且后期，需要展示主机的docker情况，需要会不太够用
exports.copy = "/api/host/copy"
//获取host表IP对应的aws实例信息
exports.hostrsyncaws = "/api/awsHost/awsrsync"
//获取host表实例对应的aws实例状态
exports.hoststatusaws = "/api/awsHost/awsStatus"
//获取host表实例Id停止aws的实例
exports.hoststopaws = "/api/awsHost/awsStop"
//获取host表实例Id启动aws的实例
exports.hoststartaws = "/api/awsHost/awsStart"
//批量保存主机关闭时间
exports.hostsavestoptime = "/api/awsHost/awsSetStopTime"
//批量保存主机关闭时间的页面主机关联Project
exports.hostlistproject = "/api/awsHost/listproject"
//关闭主机和删除主机前，查询nacos上的主机的下线信息
exports.linestop = "/api/awsHost/linestop"
//docker ps
exports.dockerps = "/api/host/dockerPs"