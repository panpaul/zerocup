<template>
    <div id="desktop">
        <el-container>
            <el-header>
                <nuxt-link to="/"><span style="float: left;padding: 3px;"><img alt="logo"
                                                                               src="../static/logo1.png"></span>
                </nuxt-link>
                <Menu style="float: right;"/>
            </el-header>
            <el-container>
                <el-container>
                    <el-main>
                        <nuxt/>
                        <!--<router-view/>-->
                    </el-main>
                    <Robot/>
                    <el-footer>
                        <Footer/>
                    </el-footer>

                </el-container>
            </el-container>
        </el-container>

    </div>
</template>

<script>
    import Menu from '../components/desktop/Menu'
    import Footer from '../components/desktop/Footer'
    import Robot from '../components/desktop/Robot'

    export default {
        name: 'desktop',
        components: {Menu, Footer, Robot}
    }
</script>

<style>
    .el-header {
        background-color: #B3C0D1;
        color: #333;
        line-height: 60px;
    }

    .el-footer {
        background-color: #B3C0D1;
        color: #333;
        text-align: center;
        line-height: 60px;
    }

    .el-main {
        background-color: #E9EEF3;
        color: #333;
    }

    body > .el-container {
        margin-bottom: 40px;
    }

    body {
        width: 100%;
    }

    p {
        max-width: 1100px;
        margin-left: auto;
        margin-right: auto;
    }

    h2 {
        max-width: 1100px;
        margin-left: auto;
        margin-right: auto;
    }

    h3 {
        max-width: 1100px;
        margin-left: auto;
        margin-right: auto;
    }
</style>
