<template>
    <div id="mobileLayout">
        <div id="mobileHeader">
            <mt-header fixed title="AI时代">
                <nuxt-link slot="left" to="/mobile/menu" v-if="this.$nuxt.$route.path!=='/mobile/menu'">
                    <mt-button icon="back">目录</mt-button>
                </nuxt-link>
                <!--<mt-button icon="more" slot="right"></mt-button>-->
            </mt-header>
        </div>
        <br><br>
        <div id="mobileBody">
            <nuxt/>
        </div>
    </div>
</template>

<script>
    export default {
        name: 'mobile',
    }
</script>

<style scoped>

</style>
