/**
 * Created by zzmhot on 2017/3/23.
 *
 * 路由Map
 *
 * @author: zzmhot
 * @github: https://github.com/zzmhot
 * @email: zzmhot@163.com
 * @Date: 2017/3/23 18:30
 * @Copyright(©) 2017 by zzmhot.
 *
 */

import Vue from 'vue'
import VueRouter from 'vue-router'
import store from 'store'
import NProgress from 'nprogress'
import 'nprogress/nprogress.css'

Vue.use(VueRouter)

//使用AMD方式加载
// component: resolve => require(['pages/home'], resolve),
const routes = [{
        path: '/home',
        name: 'home',
        components: {
            // default: require('pages/home'),
            default: require('pages/project/list'),
            menuView: require('components/leftSlide')
        },
        meta: {
            title: "主页",
            auth: true
        }
    }, {
        path: '/project/list',
        name: 'projectList',
        components: {
            default: require('pages/project/list'),
            menuView: require('components/leftSlide')
        },
        meta: {
            title: "项目配置",
            auth: true
        }
    }, {
        path: '/project/add',
        name: 'projectAdd',
        components: {
            default: require('pages/project/save'),
            menuView: require('components/leftSlide')
        },
        meta: {
            title: "项目新增",
            auth: true
        }
    }, {
        path: '/project/update/:id',
        name: 'projectUpdate',
        components: {
            default: require('pages/project/save'),
            menuView: require('components/leftSlide')
        },
        meta: {
            title: "项目修改",
            auth: true
        }
    }, {
        path: '/conf/list',
        name: 'confList',
        components: {
            default: require('pages/conf/list'),
            menuView: require('components/leftSlide')
        },
        meta: {
            title: "参数",
            auth: true
        }
    }, {
        path: '/conf/add',
        name: 'confAdd',
        components: {
            default: require('pages/conf/save'),
            menuView: require('components/leftSlide')
        },
        meta: {
            title: "参数新增",
            auth: true
        }
    }, {
        path: '/conf/update/:id',
        name: 'confUpdate',
        components: {
            default: require('pages/conf/save'),
            menuView: require('components/leftSlide')
        },
        meta: {
            title: "参数修改",
            auth: true
        }
    }, {
        path: '/service/list',
        name: 'serviceList',
        components: {
            default: require('pages/service/list'),
            menuView: require('components/leftSlide')
        },
        meta: {
            title: "服务",
            auth: true
        }
    }, {
        path: '/service/add',
        name: 'serviceAdd',
        components: {
            default: require('pages/service/save'),
            menuView: require('components/leftSlide')
        },
        meta: {
            title: "服务新增",
            auth: true
        }
    }, {
        path: '/service/update/:id',
        name: 'serviceUpdate',
        components: {
            default: require('pages/service/save'),
            menuView: require('components/leftSlide')
        },
        meta: {
            title: "服务修改",
            auth: true
        }
    }, {
        path: '/service/upload/:id',
        name: 'serviceUpload',
        components: {
            default: require('pages/service/upload'),
            menuView: require('components/leftSlide')
        },
        meta: {
            title: "服务上传、编译",
            auth: true
        }
    }, {
        path: '/platform/list',
        name: 'platformList',
        components: {
            default: require('pages/platform/list'),
            menuView: require('components/leftSlide')
        },
        meta: {
            title: "平台",
            auth: true
        }
    }, {
        path: '/platform/add',
        name: 'platformAdd',
        components: {
            default: require('pages/platform/save'),
            menuView: require('components/leftSlide')
        },
        meta: {
            title: "平台新增",
            auth: true
        }
    }, {
        path: '/platform/update/:id',
        name: 'platformUpdate',
        components: {
            default: require('pages/platform/save'),
            menuView: require('components/leftSlide')
        },
        meta: {
            title: "平台修改",
            auth: true
        }
    }, {
        path: '/host/list',
        name: 'hostList',
        components: {
            default: require('pages/host/list'),
            menuView: require('components/leftSlide')
        },
        meta: {
            title: "主机",
            auth: true
        }
    }, {
        path: '/host/add',
        name: 'hostAdd',
        components: {
            default: require('pages/host/save'),
            menuView: require('components/leftSlide')
        },
        meta: {
            title: "新增主机",
            auth: true
        }
    }, {
        path: '/host/update/:id',
        name: 'hostUpdate',
        components: {
            default: require('pages/host/save'),
            menuView: require('components/leftSlide')
        },
        meta: {
            title: "修改主机",
            auth: true
        }
    }, {
        path: '/host/stoptime',
        name: 'hostStopTime',
        components: {
            default: require('pages/host/stoptime'),
            menuView: require('components/leftSlide')
        },
        meta: {
            title: "批量修改主机关机时间",
            auth: true
        }
    }, {
        path: '/domain/list',
        name: 'domainList',
        components: {
            default: require('pages/domain/list'),
            menuView: require('components/leftSlide')
        },
        meta: {
            title: "域名",
            auth: true
        }
    }, {
        path: '/domain/add',
        name: 'domainAdd',
        components: {
            default: require('pages/domain/save'),
            menuView: require('components/leftSlide')
        },
        meta: {
            title: "新增域名",
            auth: true
        }
    }, {
        path: '/domain/update/:id',
        name: 'domainUpdate',
        components: {
            default: require('pages/domain/save'),
            menuView: require('components/leftSlide')
        },
        meta: {
            title: "修改域名",
            auth: true
        }
    }, {
        path: '/deploy/list/:id',
        name: 'deployList',
        components: {
            default: require('pages/deploy/list'),
            menuView: require('components/leftSlide')
        },
        meta: {
            title: "发布列表",
            auth: true
        }
    }, {
        path: '/deploy/release/:id',
        name: 'deployRelease',
        components: {
            default: require('pages/deploy/release'),
            menuView: require('components/leftSlide')
        },
        meta: {
            title: "发布",
            auth: true
        }
    }, {
        path: '/resource/aws',
        name: 'resourceAws',
        components: {
            default: require('pages/resource/aws'),
            menuView: require('components/leftSlide')
        },
        meta: {
            title: "亚马逊云",
            auth: true
        }
    }, {
        path: '/resource/tag',
        name: 'resourceTag',
        components: {
            default: require('pages/resource/tag'),
            menuView: require('components/leftSlide')
        },
        meta: {
            title: "tag列表",
            auth: true
        }
    }, {
        path: '/crontab/list',
        name: 'crontabList',
        components: {
            default: require('pages/crontab/list'),
            menuView: require('components/leftSlide')
        },
        meta: {
            title: "计划任务列表",
            auth: true
        }
    }, {
        path: '/crontab/updata/:id',
        name: 'crontabUpdate',
        components: {
            default: require('pages/crontab/save'),
            menuView: require('components/leftSlide')
        },
        meta: {
            title: "更新计划任务",
            auth: true
        }
    }, {
        path: '/crontab/add',
        name: 'crontabAdd',
        components: {
            default: require('pages/crontab/save'),
            menuView: require('components/leftSlide')
        },
        meta: {
            title: "增加计划任务",
            auth: true
        }
    }, {
        path: '/crontab/log',
        name: 'crontabLog',
        components: {
            default: require('pages/crontab/log'),
            menuView: require('components/leftSlide')
        },
        meta: {
            title: "计划任务日志",
            auth: true
        }
    }, {
        path: '/crontab/grafana',
        name: 'crontabGrafana',
        components: {
            default: require('pages/crontab/grafana'),
            menuView: require('components/leftSlide')
        },
        meta: {
            title: "grafana巡检任务",
            auth: true
        }
    }, {
        path: '/link/list',
        name: 'linkList',
        components: {
            default: require('pages/link/list'),
            menuView: require('components/leftSlide')
        },
        meta: {
            title: "链接列表",
            auth: true
        }
    }, {
        path: '/link/updata/:id',
        name: 'linkUpdate',
        components: {
            default: require('pages/link/save'),
            menuView: require('components/leftSlide')
        },
        meta: {
            title: "更新链接",
            auth: true
        }
    }, {
        path: '/link/add',
        name: 'linkAdd',
        components: {
            default: require('pages/link/save'),
            menuView: require('components/leftSlide')
        },
        meta: {
            title: "增加链接",
            auth: true
        }
    }, {
        path: '/sgroup/list',
        name: 'sgroupList',
        components: {
            default: require('pages/sgroup/list'),
            menuView: require('components/leftSlide')
        },
        meta: {
            title: "安全组白名单",
            auth: true
        }
    }, {
        path: '/sgroup/sava',
        name: 'sgroupSave',
        components: {
            default: require('pages/sgroup/save'),
            menuView: require('components/leftSlide')
        },
        meta: {
            title: "增加安全组白名单",
            auth: true
        }
    }, {
        path: '/user/login',
        name: 'login',
        components: {
            fullView: require('pages/user/login')
        }
    }, {
        path: '/user/register',
        name: 'register',
        components: {
            menuView: require('components/leftSlide'),
            default: require('pages/user/register')

        }
    }, {
        path: '/user/changepasswd',
        name: 'changepasswd',
        components: {
            menuView: require('components/leftSlide'),
            default: require('pages/user/changepasswd')

        },
        meta: {
            title: "修改密码",
            auth: true
        }
    }, {
    path: '/user/list',
    name: 'userList',
    components: {
      menuView: require('components/leftSlide'),
      default: require('pages/user/list')

    },
    meta: {
      title: "用户列表",
      auth: true
    }
  }, {
        path: '',
        redirect: '/home'
    }, {
        path: '*',
        name: 'notPage',
        components: {
            fullView: require('pages/error/404')
        }
    }
]

const router = new VueRouter({
    routes,
    mode: 'hash', //default: hash ,history
    scrollBehavior(to, from, savedPosition) {
        if (savedPosition) {
            return savedPosition
        } else {
            return { x: 0, y: 0 }
        }
    }
})

//全局路由配置
//路由开始之前的操作
router.beforeEach((to, from, next) => {
    NProgress.start()
    let toName = to.name
        // let fromName = from.name
    let is_login = store.state.user_info.login


    if (!is_login && toName === 'searchtaskList') {
        next();
    } else if (!is_login && toName === 'searchtaskRelease') {
        next({});
    } else if (!is_login && toName !== 'login') {
        next({
            name: 'login'
        });
    } else {
        if (is_login && toName === 'login') {
            next({
                path: '/task/list'
            });
        } else {
            next();
        }
    }
})

//路由完成之后的操作
router.afterEach(route => {
    NProgress.done()
})

export default router
