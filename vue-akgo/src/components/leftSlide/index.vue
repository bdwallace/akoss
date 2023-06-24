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
                        v-for="(item, index) in user_menus"
                        :key="index">
                    <el-menu-item
                            class="menu-list"
                            v-if="typeof item.child === 'null'"
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
  export default {
    data(){
      return{
        title: store.state.user_info.user.ProjectName,
        user_menus: [],
      }
    },
    created() {
      this.get_user_menus()
    },
    methods: {
     get_user_menus() {
       this.user_menus = JSON.parse(store.state.user_info.user.UserMenus)
     }
    }
  }

</script>
