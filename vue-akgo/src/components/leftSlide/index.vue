<template>
    <div class="left-side">
        <div class="left-side-inner">
            <router-link to="/" class="logo block">
              <!-- <div style="color: white"> 自动发布系统</div> -->
              <div style="color: white">
                <font style="color: red" size="5">{{title}}</font>
              </div>
            </router-link>
            <el-menu
                    class="menu-box"
                    theme="dark"
                    router
                    :default-active="$route.path">
                <div
                        v-for="(item, index) in nav_menu_data"
                        :key="index">
                    <el-menu-item
                            class="menu-list"
                            v-if="typeof item.child === 'undefined'"
                            :index="item.path">
                        <i class="icon fa" :class="item.icon"></i>
                        <span v-text="item.title" class="text"></span>
                    </el-menu-item>
                    <el-submenu
                            :index="item.path"
                            v-else>
                        <template slot="title">
                            <i class="icon fa" :class="item.icon"></i>
                            <span v-text="item.title" class="text"></span>
                        </template>
                        <el-menu-item
                                class="menu-list"
                                v-for="(sub_item, sub_index) in item.child"
                                :index="sub_item.path"
                                :key="sub_index">
                            <!--<i class="icon fa" :class="sub_item.icon"></i>-->
                            <span v-text="sub_item.title" class="text"></span>
                        </el-menu-item>
                    </el-submenu>
                </div>
            </el-menu>
        </div>
    </div>
</template>
<script type="text/javascript">
  import store from 'store'
    export default{
        data(){
          var Role = store.state.user_info.user.Role
          if(Role==1){
            return {
              title: store.state.user_info.user.ProjectName,
              nav_menu_data: [{
                title: "主页",
                path: "/home",
                icon: "fa-home"
              }, {
                title: "项目管理",
                path: "/project",
                icon: "el-icon-menu",
                child: [
                  { title: "服务",
                    path: "/service/list" },
                  { title: "平台",
                    path: "/platform/list" },
                  { title: "参数",
                    path: "/conf/list" },
                  { title: "主机",
                    path: "/host/list" },
                  { title: "域名",
                    path: "/domain/list" },
                ]
              }, {
                title: "上线单",
                path: "/deploy",
                icon: "fa-table",
                child: [{
                  title: "全部上线单",
                  path: "/deploy/list/0"
                }, {
                  title: "我的上线单",
                  path: "/deploy/list/" + store.state.user_info.user.Id
                }]
              }, {
                title: "资源管理",
                path: "/resource",
                icon: "el-icon-menu",
                child: [ {
                  title: "AWS",
                  path: "/resource/aws",
                }, {
                  title: "Tag列表",
                  path: "/resource/tag",
                }]
              }, {
                title: "计划任务",
                path: "/crontab",
                icon: "el-icon-menu",
                child: [ {
                  title: "计划任务列表",
                  path: "/crontab/list",
                }, {
                  title: "计划任务日志",
                  path: "/crontab/log",
                }]
              }, {
                title: "快捷操作",
                path: "/fast",
                icon: "el-icon-share",
                child: [ {
                  title: "安全组白名单",
                  path: "/sgroup/list",
                }, {
                  title: "外部链接",
                  path: "/link/list",
                }]
              }, {
                title: "用户管理",
                path: "/user",
                icon: "el-icon-setting",
                child: [ {
                  title: "用户列表",
                  path: "/user/list",
                  icon: "ace-icon fa fa-desktop"
                }]
              }]
            }
          }else if (Role==2){
            return {
              title: store.state.user_info.user.ProjectName,
              nav_menu_data: [{
                title: "主页",
                path: "/home",
                icon: "fa-home"
              }, {
                title: "项目管理",
                path: "/project",
                icon: "el-icon-menu",
                child: [
                  { title: "服务",
                    path: "/service/list" },
                  { title: "平台",
                    path: "/platform/list" },
                  { title: "参数",
                    path: "/conf/list" },
                  { title: "主机",
                    path: "/host/list" },
                  { title: "域名",
                    path: "/domain/list" },
                ]
              }, {
                title: "上线单",
                path: "/deploy",
                icon: "fa-table",
                child: [{
                  title: "全部上线单",
                  path: "/deploy/list/0"
                }, {
                  title: "我的上线单",
                  path: "/deploy/list/" + store.state.user_info.user.Id
                }]
              }, {
                title: "资源管理",
                path: "/resource",
                icon: "el-icon-menu",
                child: [ {
                  title: "AWS",
                  path: "/resource/aws",
                }, {
                  title: "Tag列表",
                  path: "/resource/tag",
                }]
              }, {
                title: "计划任务",
                path: "/crontab",
                icon: "el-icon-menu",
                child: [ {
                  title: "计划任务列表",
                  path: "/crontab/list",
                }, {
                  title: "计划任务日志",
                  path: "/crontab/log",
                }]
              }, {
                title: "快捷操作",
                path: "/fast",
                icon: "el-icon-share",
                child: [ {
                  title: "安全组白名单",
                  path: "/sgroup/list",
                }, {
                  title: "外部链接",
                  path: "/link/list",
                }]
              }]
            }
          } else if (Role==3){
            return {
              title: store.state.user_info.user.ProjectName,
              nav_menu_data: [{
                title: "主页",
                path: "/home",
                icon: "fa-home"
              }, {
                title: "项目管理",
                path: "/project",
                icon: "el-icon-menu",
                child: [
                  { title: "平台",
                    path: "/platform/list" },
                  { title: "域名",
                    path: "/domain/list" },
                ]
              }, {
                title: "快捷操作",
                path: "/fast",
                icon: "el-icon-share",
                child: [ {
                  title: "安全组白名单",
                  path: "/sgroup/list",
                }, {
                  title: "外部链接",
                  path: "/link/list",
                }]
              }]
            }
          }

        }
    }
</script>
