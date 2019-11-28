<template>
    <div id="menu">
        <el-menu :default-active=this.$nuxt.$route.fullPath
                 @select="handleSelect"
                 background-color="#B3C0D1"
                 class="menu"
                 mode="horizontal"
                 router
        >
            <el-menu-item index="/desktop/index">首页</el-menu-item>
            <el-menu-item index="/desktop/description">介绍</el-menu-item>
            <el-submenu index="Story">
                <template slot="title">AI小故事</template>
                <el-menu-item index="/desktop/story_1">被重新定义的人类健康</el-menu-item>
                <el-menu-item index="/desktop/story_2">机器指引下的社交</el-menu-item>
                <el-menu-item index="/desktop/story_3">空无一人的车间</el-menu-item>
                <el-menu-item index="/desktop/story_4">机器控制的人生</el-menu-item>
            </el-submenu>
            <el-menu-item index="/desktop/ai">AI实践</el-menu-item>
            <el-menu-item index="/desktop/news">武大AI</el-menu-item>
            <el-menu-item index="/desktop/goldtime">黄金时代</el-menu-item>
            <el-submenu index="User" v-if="!$store.state.token">
                <template slot="title">用户中心</template>
                <el-menu-item index="/desktop/login">登录</el-menu-item>
                <el-menu-item index="/desktop/register">注册</el-menu-item>
            </el-submenu>
            <el-submenu index="User" v-if="$store.state.token">
                <template slot="title">用户中心</template>
                <el-menu-item @click="clearCookie()" index="/desktop/login">注销</el-menu-item>
            </el-submenu>
        </el-menu>
    </div>
</template>

<script>
    import Cookies from 'js-cookie'

    export default {
        name: "myMenu",
        methods: {
            handleSelect(key, keyPath) {
                window.console.log(key, keyPath);
            },
            clearCookie: function () {
                Cookies.remove('token');
                Cookies.remove('vuex');
//location.reload();
                window.location.href = "/desktop/login";
            },
        }
    }
</script>


<style>
    .el-menu {
        text-align: center;
    }
</style>