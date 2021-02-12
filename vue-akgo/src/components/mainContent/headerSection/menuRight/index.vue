<template>
    <div class="menu-right" v-if="get_user_info.login">
        <div class="notification-menu">
            <el-dropdown trigger="click" class="notification-list">
                <div class="notification-btn">
                    <img :src="get_user_info.user.Avatar" :alt="get_user_info.user.Realname"/>
                    <span v-text="get_user_info.user.Realname"></span>
                    <span class="icon"></span>
                </div>
                <el-dropdown-menu slot="dropdown" class="dropdown-menu">
                    <el-dropdown-item v-for="item in project_list" :key="item" class="dropdown-list">
                        <a href="javascript:" class="dropdown-btn" @click="project_change(item)">
                            <i class="icon fa fa-user"></i>
                            <span>{{ item.Name }}</span>
                        </a>
                    </el-dropdown-item>
                    <el-dropdown-item class="dropdown-list">
                        <a href="javascript:" class="dropdown-btn" @click="user_click(1)">
                            <i class="icon fa fa-user"></i>
                            <span>密码修改</span>
                        </a>
                    </el-dropdown-item>
<!--                    <el-dropdown-item class="dropdown-list">-->
<!--                        <a href="javascript:" class="dropdown-btn" @click="user_click(2)">-->
<!--                            <i class="icon fa fa-cog"></i>-->
<!--                            <span>注册用户</span>-->
<!--                        </a>-->
<!--                    </el-dropdown-item>-->
                    <el-dropdown-item class="dropdown-list">
                        <a href="javascript:" class="dropdown-btn" @click="user_click(0)">
                            <i class="icon fa fa-sign-out"></i>
                            <span>安全退出</span>
                        </a>
                    </el-dropdown-item>
                </el-dropdown-menu>
            </el-dropdown>
        </div>
    </div>
</template>
<script type="text/javascript">
    import {port_user, port_project} from 'common/port_uri'
    import {mapGetters, mapActions} from 'vuex'
    import {GET_USER_INFO} from 'store/getters/type'
    import {SET_USER_INFO} from 'store/actions/type'
    import store from 'store'
    const USER_OUT = 0
    const USER_INFO = 1
    const ADMIN_SETTING = 2

    export default{
        data() {
            return {
                project_list: []
            }
        },

        computed: {
            ...
                mapGetters({
                    get_user_info: GET_USER_INFO
                })
        },

        created() {
            this.get_project_list()
        },

        methods: {
        ...
            mapActions({
                set_user_info: SET_USER_INFO
            }),
            //退出
            user_out() {
                this.$confirm('此操作将退出登录, 是否继续?', '提示', {
                        confirmButtonText: '确定',
                        cancelButtonText: '取消',
                        type: 'warning'
                    })
                .then(() => {
                    this.$http.post(port_user.logout)
                        .then(({data: {msg}}) => {
                            this.$message({
                                message: msg,
                                type: 'success'
                            })
                        this.set_user_info(null)
                        setTimeout(() => {
                            this.$router.replace({name: "login"})
                        }, 500)
                    })
                })
                .catch(() => {})
            },

            user_info()
            {
                this.$router.push({path: '/user/changepasswd'})
                //个人信息
            },

            admin_setting()
            {
                // if(store.state.user_info.user.Role==1){
                //     this.$router.push({path: '/user/register'})
                // }
            },

            user_click(type)
            {
                switch (type) {
                    case USER_OUT :
                        //退出
                        this.user_out()
                        break
                    case USER_INFO:
                        //个人信息
                        this.user_info()
                        break
                    case ADMIN_SETTING:
                    //密码修改
                    this.admin_setting()
                        break
                }
            },

            get_project_list() {
                this.$http.get(port_project.list)
                        .then(({data: {data}}) => {
                    this.project_list = data
                })
            },

            project_change(project) {
                this.$http.get(port_user.projectchange, {
                        params: {
                            id: project.Id,
                            name: project.Name
                        }
                    })
                .then(({data: {data}}) => {
                    this.set_user_info({
                        user: data,
                        login: true
                    })
                    // let user_info = this.get_user_info
                    // console.log("-----", user_info)
                    this.$message({
                        message: "切换至： " + data.ProjectId + " " + data.ProjectName,
                        type: 'success'
                    })
                    setTimeout(() => {
                        // 切换环境后，回到主页面。避免出现停留在上一套环境的单页面。
                        this.$router.push({name: 'home'})
                        // 切换后将页面重新刷新一下，在leftslide的title的数据重新更新
                        location.reload()
                    }, 500)
                })
            },

        }
    }
</script>
