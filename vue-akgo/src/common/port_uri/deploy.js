//数据列表
exports.list = "/api/deploy/projectId"
//用户数据列表
exports.listuser = "/api/deploy/projectId/userId"
//根据id,get查询,put修改,delete删除数据
exports.id = "/api/deploy/id"
//添加数据
exports.save = "/api/deploy"
//的上传和build快速发布
exports.savebuild = "/api/deploy/build"
//复制数据
exports.copy = "/api/deploy/copy"
//发布
exports.deploy = "/api/deploy/service"
//提交tag
exports.tagcmd = "/api/deploy/tagcmd"