<template>
    <div id="story3">
        <p>朗读文章</p>
        <div class="aud">
            <audio :src="require('../../static/audios/story_3.mp3')" controls/>
        </div>
        <h2>空无一人的车间</h2>
        <p>
            我们再从人类社会尺度来看这个美丽新世界。 机器， 已经普遍的存在在了所有以前认为只有个体才完成的职位上。我们看到空无一人的汽车生产线，空无一人的医院， 即使是最精妙的技艺， 配备了最好的视觉和听觉识别系统，
            经过了多轮强化学习的机器也可以娴熟的完成。 那你要问劳动者在什么地方？都在玩啊，除了几个极为聪明的AI工程师。</p>
        <p>
        </p>
        <img alt="worker" src="../../static/story3.jpg">
        <p>
            分析： 这可能是你听到的公号上最多的一个故事， 关于自动化和失业。 我个人认为，首先这样一个故事真正实现的历程会很漫长，再有， 就算有那么多AI占领 了工厂的流水线，
            就有更多的人进入到整个AI的生产线里，这并不仅仅包含程序员， 还包括需要大量检测AI在各种情况下行为合理性的人， 就拿无人驾驶情况看，一辆新的无人驾驶车出来，那么它可能和之前车不一样的就不仅仅是发动机，
            还有行为规则， 比如在各种特殊情况下是选择保守还是效率 ， 那么单纯对AI行为的管理就要有多少设计师呢。还有有些人可能会变成AI分析家， 比如柯洁已经开始想从阿法狗那里探讨围棋的真谛了。</p>
        <p>
            退一步讲， 即使AI真的取代了大量取代了人的工作， 也不是人类末日，人类并非一定为工作而生， 有很多可以给人生活以意义的事情， 那么如何制造一个低工作但是高幸福度的社会或许就是那时候新的课题吧 。</p>

        <br/><br/>
        <el-button-group>
            <nuxt-link to="/desktop/story_2">
                <el-button icon="el-icon-arrow-left" type="primary">上一篇</el-button>
            </nuxt-link>
            <nuxt-link to="/desktop/story_4">
                <el-button type="primary">下一篇<i class="el-icon-arrow-right el-icon--right"/></el-button>
            </nuxt-link>
        </el-button-group>
        <Comment article_id="5"/>
    </div>
</template>

<script>
    import Comment from '../../components/desktop/Comment.vue'

    export default {
        name: "desktopStory3",
        components: {
            Comment
        },
    }
</script>

<style scoped>
    img {
        width: 40%;
        margin-left: calc((100% - 600px) / 2);
        max-width: 600px;
    }

    .el-button-group {
        margin-left: calc((68% - 500px) / 2);
    }
</style>
