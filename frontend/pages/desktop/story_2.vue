<template>
    <div id="story2">
        <p>朗读文章</p>
        <div class="aud">
            <audio :src="require('@/static/audios/story_2.mp3')" controls></audio>
        </div>
        <h2>机器指引下的社交</h2>
        <p>
            你是一个不善交际的AI程序员， 早上写了两小时程序， 你的大脑监控仪早已发现了你有效率减低趋势，于是， 它建议你出去走走。
            一辆自动驾驶的清洁能源车早已停在了门口，它决定带你到颐和园里走一趟，当然如果你不满意只要在地点选项上稍微调整一下就好了，于是车辆启动，前往目的地。
        </p>
        <p>
            在公园里， 你的个人助理系统知道你平时很宅但其实喜欢邂逅， 于是它迅速与目前处在公园里的异性的个人AI助理进行沟通， 匹配双方的数据，然后计算得到了一个最佳的邂逅对象， 双方电脑达成共识给主人发送一条信息。
            于是一场欢乐的邂逅就开始了。
        </p>
        <img src="@/static/2.png">
        <p>
            当然不善言辞的你见到她感到有点紧张， 没关系， 个人助理跟据对方的性格， 已经给你做好了聊天提示， 你啥事变得十分健谈。由于这样的浪漫邂逅太多太容易， 我们已经不愿意去和一人结婚终老， 而是享受这种U盘式的邂逅关系。
        </p>
        <p>
            甚至由于助理机器人的功能过于强大， 很多宅男和她聊天就已经不愿意接触人类了。
        </p>
        <p>
            分析： 这个故事是展示AI对人类社交可能产生的影响， 人与人的最佳匹配变得太过容易， 会不会最后动摇了婚姻家庭等各类最基础的人类关系纽带。 倘若一段关系的开始，仅仅是在你的个人助理app里发出一条请求，
            还会有人需要长久的关系吗？
        </p>
        <br/><br/>
        <el-button-group>
            <nuxt-link to="/desktop/story_1">
                <el-button icon="el-icon-arrow-left" type="primary">上一篇</el-button>
            </nuxt-link>
            <nuxt-link to="/desktop/story_3">
                <el-button type="primary">下一篇<i class="el-icon-arrow-right el-icon--right"></i></el-button>
            </nuxt-link>
        </el-button-group>
        <Comment article_id="4"/>
    </div>
</template>

<script>
    import Comment from '../../components/desktop/Comment.vue'

    export default {
        name: "story2",
        components: {
            Comment
        },
    }
</script>

<style scoped>
    img {
        width: 40%;
        margin-left: calc((100% - 300px) / 2);
        max-width: 300px;
    }

    .el-button-group {
        margin-left: calc((68% - 500px) / 2);
    }
</style>
