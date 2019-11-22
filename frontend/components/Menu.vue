<template>
    <div id="menu">
        <el-menu :default-active=this.$nuxt.$route.name
                 @select="handleSelect"
                 background-color="#B3C0D1"
                 class="menu"
                 mode="horizontal"
                 router
        >
            <el-menu-item index="index">首页</el-menu-item>
            <el-menu-item index="description">介绍</el-menu-item>
            <el-submenu index="Story">
                <template slot="title">AI小故事</template>
                <el-menu-item index="story_1">被重新定义的人类健康</el-menu-item>
                <el-menu-item index="story_2">机器指引下的社交</el-menu-item>
                <el-menu-item index="story_3">空无一人的车间</el-menu-item>
                <el-menu-item index="story_4">机器控制的人生</el-menu-item>
            </el-submenu>
            <el-menu-item index="ai">AI实践</el-menu-item>
            <el-menu-item index="goldtime">黄金时代</el-menu-item>
            <el-submenu index="User" v-if="!$store.state.token">
                <template slot="title">用户中心</template>
                <el-menu-item index="login">登录</el-menu-item>
                <el-menu-item index="register">注册</el-menu-item>
            </el-submenu>

        </el-menu>
    </div>
</template>

<script>
    export default {
        name: "menu",
        methods: {
            handleSelect(key, keyPath) {
                window.console.log(key, keyPath);
            }
        }
    }
</script>

<style>
    .el-menu {
        text-align: center;
    }
</style>