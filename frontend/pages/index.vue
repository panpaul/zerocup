<template>
    <div class="container":style="bg">
		<nuxt-link to="/description">
        <img alt="logo" class="imgWidth" src="../static/logo.png"></nuxt-link>
		<br/><br/><br/><br/><br/><br/><br/><br/><br/><br/><br/><br/><br/><br/><br/><br/><br/><br/><br/><br/>
		<br/><br/><br/><br/><br/><br/><br/><br/><br/><br/><br/><br/><br/><br/><br/>
		<p>点击上方logo跳转简介&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
		&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</p>
		<br/><br/><br/><br/><br/>
    </div>
</template>

<script>
    // @ is an alias to /src
    import HelloWorld from '@/components/HelloWorld.vue'

    export default {
		data () {
    	return {
		    bg: {
                backgroundImage: "url(" + require("../static/background.jpg") + ")",
                backgroundRepeat: "no-repeat",
                backgroundSize: "cover",
            },
    	}
		},
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
		font-size:28px;
		color: #FFFFFF;
	}
</style>
