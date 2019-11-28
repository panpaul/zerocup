<template>
    <div>
        <h1 style="text-align: center">Home</h1>
        <mt-swipe>
            <mt-swipe-item>1</mt-swipe-item>
            <mt-swipe-item>2</mt-swipe-item>
            <mt-swipe-item>3</mt-swipe-item>
            <mt-swipe-item>4</mt-swipe-item>
            <mt-swipe-item>5</mt-swipe-item>
            <mt-swipe-item>6</mt-swipe-item>
            <mt-swipe-item>7</mt-swipe-item>
            <mt-swipe-item>8</mt-swipe-item>

        </mt-swipe>
    </div>
</template>

<script>
    export default {
        layout: "mobile"
    };
</script>

<style scoped>
    .mint-swipe {
        margin-right: auto;
        margin-left: auto;
        height: 50vh;
        width: 80vw;
    }

    .mint-swipe-item {
        background-color: aqua;
    }
</style>
