<template>
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
</template>

<script>
    export default {
        name: 'Home',
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

</style>
