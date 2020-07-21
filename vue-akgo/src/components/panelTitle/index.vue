<template>
    <div class="panel-title">
        <!-- <span v-if="title" v-text="title"></span> -->
        <span v-if="title">
            {{title}} <font style="color: red" size="5">{{ProjectId}}-{{ProjectName}}</font>
        </span>
        <div class="fr">
            <slot></slot>
        </div>
    </div>
</template>
<script type="text/javascript">
    import store from 'store'
    export default{
        props: {
            title: {
                type: String
            }
        },
        data() {
            return {
                ProjectId: store.state.user_info.user.ProjectId,
                ProjectName: store.state.user_info.user.ProjectName,
            }
        }
    }
</script>
