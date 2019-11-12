<template>
    <div id="menu">
        <el-menu :default-active="activeIndex"
                 @select="handleSelect"
                 class="menu"
                 mode="horizontal"
                 router
        >
            <el-row :gutter="10">
                <el-col :lg="2" :md="3" :sm="3" :xl="1" :xs="3">
                    <el-menu-item index="Home">首页</el-menu-item>
                </el-col>
                <el-col :lg="2" :md="3" :sm="3" :xl="1" :xs="3">
                    <el-menu-item index="Description">介绍</el-menu-item>
                </el-col>
                <el-col :lg="2" :md="3" :sm="4" :xl="1" :xs="0">
                    <el-submenu index="ShowCase">
                        <template slot="title">AI实践</template>
                        <el-menu-item index="2-1">选项1</el-menu-item>
                        <el-submenu index="2-2">
                            <template slot="title">选项2</template>
                            <el-menu-item index="2-2-1">选项1</el-menu-item>
                        </el-submenu>
                    </el-submenu>
                </el-col>
                <el-col :lg="2" :md="3" :sm="3" :xl="1" :xs="3">
                    <el-menu-item index="About">关于</el-menu-item>
                </el-col>
                <el-col :lg="2" :md="3" :sm="3" :xl="1" :xs="0">
                    <el-menu-item index="Ref"><a href="https://github.com/panpaul/zerocup" target="_blank">Github</a>
                    </el-menu-item>
                </el-col>
            </el-row>
        </el-menu>
    </div>
</template>

<script>
    export default {
        data() {
            return {
                activeIndex: "Home",
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