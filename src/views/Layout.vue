<template>
    <div id="desktop">
        <el-container>
            <el-header>
                <Menu/>
            </el-header>
            <el-container>
                <el-container>
                    <el-main>
                        <router-view/>
                    </el-main>
                    <el-footer>
                        <Footer/>
                    </el-footer>
                </el-container>
                <el-aside width="20%">Aside</el-aside>
            </el-container>
        </el-container>

    </div>
</template>

<script>
    import Menu from '@/components/Menu'
    import Footer from '@/components/Footer'

    export default {
        name: 'desktop',
        components: {Menu, Footer}
    }
</script>

<style>
    .el-footer {
        background-color: #B3C0D1;
        color: #333;
        text-align: center;
        line-height: 60px;
    }

    .el-aside {
        background-color: #D3DCE6;
        color: #333;
        text-align: center;
        line-height: 200px;
    }

    .el-main {
        background-color: #E9EEF3;
        color: #333;
        text-align: center;
        line-height: 160px;
    }

    body > .el-container {
        margin-bottom: 40px;
    }
</style>
