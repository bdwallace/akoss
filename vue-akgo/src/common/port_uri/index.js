/**
 * Created by zzmhot on 2017/3/24.
 *
 * @author: zzmhot
 * @github: https://github.com/zzmhot
 * @email: zzmhot@163.com
 * @Date: 2017/3/24 14:56
 * @Copyright(©) 2017 by zzmhot.
 *
 */

const port_code = require("./code")
const port_user = require("./user")
const port_file = require("./file")
const port_conf = require("./conf")
const port_task = require("./task")
const port_record = require("./record")
const port_recordoperation = require("./recordoperation")
const port_git = require("./git")
const port_p2p = require("./p2p")
const port_other = require("./other")
const port_project = require("./project")
const port_service = require("./service")
const port_platform = require("./platform")
const port_host = require("./host")
const port_domain = require("./domain")
const port_deploy = require("./deploy")
const port_walle = require("./walle")
const port_resource = require("./resource")
const port_crontab = require("./crontab")
const port_link = require("./link")
const port_sgroup = require("./sgroup")
const port_grafana = require("./grafana")
const port_cloud = require("./cloud")
module.exports = {
    port_code,
    port_user,
    port_file,
    port_conf,
    port_task,
    port_record,
    port_recordoperation,
    port_git,
    port_p2p,
    port_other,
    port_project,
    port_platform,
    port_service,
    port_host,
    port_domain,
    port_deploy,
    port_walle,
    port_resource,
    port_crontab,
    port_link,
    port_sgroup,
    port_grafana,
    port_cloud
}