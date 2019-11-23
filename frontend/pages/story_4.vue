<template>
    <div id="story4">
				<p>朗读文章</p>
		<div class="aud"><audio :src="require('@/static/audios/story_4.mp3')" controls></audio></div>
        <h2>机器控制的人生</h2>
        <p>
            既然大部分的产业， 甚至政府的主要功能， 都已经被AI去掉，那我们自然的会问， 人类在干什么。事实上， 这个时候最有用的职业可能是哲学家， 因为他们最喜欢的讨论是人生的意义何在， 这可是在这个时代非常有意义的话题，
            因为我们的生活的主要决策既然都可以大量的交给AI算法， 人存在的意义又何在，除了那些极个别的设计算法的天才外， 你， 作为一个普通人， 存在的意义何再？ 在几十年前， 你起码可以说我就是我的自由意志， 我思故我在，
            我动故我在，而现在，你大脑里的每一个脉冲都被机器记录， 然后它比你预先的知道你想要去哪里， 会和谁dating， 如果机器想控制你，那么只要按照一定的方法引导你的多巴胺信号，就会让你不知不觉的成为你不想成为的人。
            这样的世界不是太可怕了吗？ 你存在和意义何在？ 所以在这样的一个世界里， 哲学家或心理学家成为了非常好的职业，他或许会不停的引导你如果你去发挥你在某个方面的才干， 你会创作出一些好的艺术品， 比如你可以去写小说，
            去画画， 你连续的创造和提升， 或许就是你的生命意义。
        </p>
        <p>
            分析：这里的观点认为创作永远属于人的领地， 但从AI算法发展角度看，不一定。当下的gan技术， 已经能够模拟人创作一些具有固定艺术风格的画，如果未来结合了阿法狗里用到的强化学习，那么说不定超越了简单的摹仿大家作品，
            而是在不同的大家作品里寻找出人类创作的进化规律， 去不停的产生新的能够取悦于人的艺术风格。
        </p>
        <p>
            那么科学创造呢？ 是的， 我们说科学创造是人类历史的长河里， 建立在大量经验之上的灵感火花带来的突破性创新，波普尔认为这些科技成果的出现甚至是没有太多的可预测性，由于科技进步的偶然性，带来了人类历史进程的不确定性。
            然而在真正的AI时代， 这也可能发生改变， 除了相对论这一类的非常接近哲学的研究，大量的实验学课，比如药物的设计， 会被机器取代， 因为这一类学科的探索取决于大量试错， 这个试错过程让AI来做远比人高的多，
            所以未来，大量的高效药物， 可能会是AI设计出来的， 而你所做的汽车，飞机从形状曲线， 到材料选择， 也都来自于AI。
        </p>
        <p>
            当然， 那些相对论，进化论级别的发现呢？ 这也许并非AI一时所能级，因为本质上这些东西需要AI建立对世界的整个认知系统， 和最关键的，对自我的认知， 即使是21世纪初期的脑科学，也对这些一无所知，
            而如果这些都有一天被突破， 那么我们所面临的就已经面临的不是人的存在意义与否的问题， 而是人类本身是否还需要存在了。 既然这个基本问题一时是没有答案了，我们也就不需要太过焦虑这个问题了。
        </p>
        <img src="http://images.ofweek.com/Upload/News/2018-08/22/zhouxiyao/1534897287247057782.jpg">
        <br/><br/>
        <el-button-group>
            <nuxt-link to="/story_3">
                <el-button icon="el-icon-arrow-left" type="primary">上一篇</el-button>
            </nuxt-link>
            <nuxt-link to="/story_4">
                <el-button disabled type="primary">下一篇<i class="el-icon-arrow-right el-icon--right"></i></el-button>
            </nuxt-link>
            <nuxt-link to="/ai">
                <el-button round type="primary">看看其它</el-button>
            </nuxt-link>
        </el-button-group>
        <Comment article_id="6"/>
    </div>
</template>

<script>
    import Comment from '@/components/Comment.vue'

    export default {
        name: "story4",
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
