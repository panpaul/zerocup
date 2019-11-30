<template>
    <div class="bigger">
        <div class="container">
            <nuxt-link to="/desktop/description">
                <img alt="logo" class="imgWidth" src="../../static/logo.png">
            </nuxt-link>

            <br/><br/><br/><br/><br/><br/><br/><br/>
            <br/><br/><br/><br/><br/><br/><br/><br/>
            <br/><br/><br/><br/><br/><br/><br/><br/>
            <br/><br/><br/><br/><br/><br/><br/><br/>
            <br/><br/><br/><br/><br/><br/>

        </div>
        <div class="container1">
            <br><br>
            <el-carousel :interval="4000" height="300px" type="card">
                <el-carousel-item :key="i.name" v-for="i in myItems">
                    <nuxt-link :to=i.link>
                        <h2>{{i.name}}</h2>
                        <img :alt=i.name :src=i.img class="carouselImg">
                    </nuxt-link>
                </el-carousel-item>
            </el-carousel>
        </div>
    </div>

</template>

<script>
    export default {
        name: 'desktopHome',
        data() {
            return {
                myItems: [
                    {
                        img: require("../../static/desc.png"),
                        name: "介绍",
                        link: "/desktop/description"
                    },
                    {
                        img: require("../../static/story1.jpg"),
                        name: "小故事",
                        link: "/desktop/story_1"
                    },
                    {
                        img: require("../../static/ai.PNG"),
                        name: "AI实践",
                        link: "/desktop/ai"
                    },
                    {
                        img: require("../../static/goldtime1.png"),
                        name: "黄金时代",
                        link: "/desktop/goldtime"
                    },
                ]
            }
        }
    }
</script>

<style scoped>
    .imgWidth {
        width: 40%;
        float: right;
    }

    p {
        text-align: right;
        margin-right: 10%;
        font-size: 28px;
        color: #FFFFFF;
        max-width: 100%;
    }


    .container {
        background-image: url("../../static/background0.5.png");
        background-repeat: no-repeat;
        background-size: cover;
    }

    .container1 {
        background-color: rgba(4, 77, 112, 0.5);
    }

    .el-carousel__item h2 {
        color: #000000;
        font-size: 30px;
        line-height: 200px;
        margin: 0;
        position: absolute;
    }

    .el-carousel__item:nth-child(2n) {
        background-color: #99a9bf;
    }

    .el-carousel__item:nth-child(2n+1) {
        background-color: #d3dce6;
    }

    .carouselImg {
        opacity: 0.3;
        width: 100%;
        position: relative;
    }
</style>
