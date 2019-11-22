<template>
    <div id="menu">
        <el-menu :default-active=activeIndex
                 @select="handleSelect"
                 background-color="#B3C0D1"
                 class="menu"
                 mode="horizontal"
				 router
        >
            <el-menu-item index="/">首页</el-menu-item>
            <el-menu-item index="description">介绍</el-menu-item>
            <el-menu-item index="ai">AI实践</el-menu-item>
            <el-menu-item index="about">关于</el-menu-item>
            <el-submenu index="User">
                <template slot="title">用户中心</template>
                <el-menu-item index="login">登录</el-menu-item>
                <el-menu-item index="register">注册</el-menu-item>
            </el-submenu>

        </el-menu>
    </div>
</template>

<script>
    export default {
        data() {
            return {
                activeIndex: "/",
            };
        },
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