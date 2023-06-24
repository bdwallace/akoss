/**
 * Created by zzmhot on 2017/3/23.
 *
 * 状态管理器
 *
 * @author: zzmhot
 * @github: https://github.com/zzmhot
 * @email: zzmhot@163.com
 * @Date: 2017/3/23 18:24
 * @Copyright(©) 2017 by zzmhot.
 *
 */

import Vue from 'vue'
import Vuex from 'vuex'

//引入模块
import actions from './actions'
import getters from './getters'
import mutations from './mutations'
import state from './states'
import createPersistedState from "vuex-persistedstate"

Vue.use(Vuex)
import JsonViewer from "vue-json-viewer"
Vue.use(JsonViewer)


export default new Vuex.Store({
    state,
    getters,
    actions,
    mutations,
    plugins: [createPersistedState]
})