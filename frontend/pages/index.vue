<template>
    <div class="container":style="bg">
		<nuxt-link to="/description">
        <img alt="logo" class="imgWidth" src="../static/logo.png"></nuxt-link>
		<br/><br/><br/><br/><br/><br/><br/><br/><br/><br/><br/><br/><br/><br/><br/><br/><br/><br/><br/><br/>
		<br/><br/><br/><br/><br/><br/><br/><br/>
		<br/><br/><br/><br/><br/>
		<p>点击上方logo跳转简介</p>
		<br/><br/><br/><br/><br/>
    </div>
</template>

<script>
	// @ is an alias to /src

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
		margin-right: 10%;
		font-size:28px;
		color: #FFFFFF;
		max-width: 100%;
	}
</style>
