<template>
    <div id="mobile-menu">
        <nuxt-link to="/mobile/index">
            <mt-cell is-link label="首页"/>
        </nuxt-link>
        <nuxt-link to="/mobile/description">
            <mt-cell is-link label="AI介绍"/>
        </nuxt-link>
        <nuxt-link to="/mobile/story">
            <mt-cell is-link label="AI小故事"/>
        </nuxt-link>
        <nuxt-link to="/mobile/news">
            <mt-cell is-link label="武大AI"/>
        </nuxt-link>
        <nuxt-link to="/mobile/goldtime">
            <mt-cell is-link label="黄金时代"/>
        </nuxt-link>
    </div>
</template>

<script>
    export default {
        name: 'mobileMenu',
        layout: 'mobile',
    }
</script>

<style scoped>

</style>