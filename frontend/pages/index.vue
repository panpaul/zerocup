<template>
    <div class="bigger">
        <div class="container">
            <nuxt-link to="/description">
                <img alt="logo" class="imgWidth" src="../static/logo.png">
            </nuxt-link>

            <br/><br/><br/><br/><br/><br/><br/><br/>
            <br/><br/><br/><br/><br/><br/><br/><br/>
            <br/><br/><br/><br/><br/><br/><br/><br/>
            <br/><br/><br/><br/><br/><br/><br/><br/>
            <br/><br/><br/><br/><br/><br/>

        </div>
        <div class="container1">
            <el-carousel :interval="4000" height="200px" type="card">
                <el-carousel-item :key="item" v-for="item in items">
                    <nuxt-link :to=item.link>
                        <h2>{{item.name}}</h2>
                        <img :alt=item.name :src=item.img class="carouselImg">
                    </nuxt-link>
                </el-carousel-item>
            </el-carousel>
        </div>
    </div>

</template>

<script>
    export default {
        name: 'Home',
        data() {
            return {
                items: [
                    {
                        img: "http://www.elecfans.com/uploads/allimg/170713/01001263T_0.jpg",
                        name: "介绍",
                        link: "/description"
                    },
                    {
                        img: "http://www.mhealthchina.org/wp-content/uploads/2018/05/2018051608530124.jpg",
                        name: "小故事",
                        link: "/story_1"
                    },
                    {
                        img: require("../static/ai.PNG"),
                        name: "AI实践",
                        link: "/ai"
                    },
                    {
                        img: "http://tc.sinaimg.cn/maxwidth.800/tc.service.weibo.com/mmbiz_qpic_cn/595192735740d68ab174822c2ce3a566.jpg",
                        name: "黄金时代",
                        link: "/goldtime"
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
        background-image: url("../static/background0.5.png");
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
