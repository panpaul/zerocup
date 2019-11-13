<template>
    <div class="home">
        <img alt="Vue logo" src="../assets/logo.png">
        <HelloWorld msg="Welcome to Your Vue.js App"/>
		<Comment article_id=1 />
    </div>
</template>

<script>
    // @ is an alias to /src
    import HelloWorld from '@/components/HelloWorld.vue'
    import Comment from '@/components/Comment.vue'

    export default {
        name: 'Home',
        components: {
			HelloWorld,
			Comment
        }
    }
</script>
