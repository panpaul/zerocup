<template>
    <div id="app">
        <Layout/>
    </div>
</template>

<script>
    import Layout from "@/views/Layout";

    export default {
        name: 'app',
        components: {Layout},
        data() {
            return {
                fullscreenLoading: false
            }
        },
        mounted: function () {
            const loading = this.$loading({
                lock: true,
                text: 'Loading',
                spinner: 'el-icon-loading',
                background: 'rgba(0, 0, 0, 0.7)'
            });
            setTimeout(() => {
                loading.close();
            }, 2000);
        },
    }
</script>