<template>
    <div>
        <h1 style="text-align: center">AI时代</h1>
        <mt-swipe>
            <mt-swipe-item>
                <nuxt-link to="/mobile/description">
                    <div class="textBox">AI介绍</div>
                    <img alt="description" src="../../static/desc.png">
                </nuxt-link>
            </mt-swipe-item>

            <mt-swipe-item>
                <nuxt-link to="/mobile/story">
                    <div class="textBox">AI小故事</div>
                    <img alt="description" src="../../static/story1.jpg">
                </nuxt-link>
            </mt-swipe-item>

            <mt-swipe-item>
                <nuxt-link to="/mobile/news">
                    <div class="textBox">武大AI</div>
                    <img alt="description" src="../../static/news1.webp">
                </nuxt-link>
            </mt-swipe-item>

            <mt-swipe-item>
                <nuxt-link to="/mobile/goldtime">
                    <div class="textBox">黄金时代</div>
                    <img alt="description" src="../../static/goldtime1.png">
                </nuxt-link>
            </mt-swipe-item>


        </mt-swipe>
    </div>
</template>

<script>
    export default {
        name: "mobileIndex",
        layout: "mobile"
    };
</script>

<style scoped>
    .mint-swipe {
        margin-right: auto;
        margin-left: auto;
        height: 30vh;
        width: 80vw;
    }

    .mint-swipe-item {
        background-color: aqua;
    }

    img {
        width: 100%;
        height: 100%;
        position: absolute;
    }

    .textBox {
        position: absolute;
        margin-left: auto;
        margin-right: auto;
        left: 0;
        right: 0;
        bottom: 50%;
        z-index: 100;

        text-align: center;
        color: orange;
        background-color: rgba(180, 190, 200, 0.7);

        font-size: 30px;
    }

</style>
