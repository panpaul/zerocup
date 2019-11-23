<template>
    <div class="container" id="story1">
		<p>朗读文章</p>
		<div class="aud"><audio :src="require('@/static/audios/news.mp3')" controls></audio></div>
        <h2>当AI遇上武大，会擦出什么火花？</h2>
		<img src="@/static/news1.webp"
             style="width:60%;margin-left:20%;">
        <p>
            11月22日，武汉大学人工智能研究院成立大会暨人工智能国际学术研讨会举行。海内外高校近30名人工智能专家学者、小米集团、中科寒武纪、华为公司等业界专家齐聚珞珈山，围绕人工智能进行对话与交流。
        </p>
		<img src="@/static/news2.webp"
             style="width:60%;margin-left:20%;">
        <p>
            窦贤康指出，人工智能研究院的首要任务是做好人才队伍建设和人才培养工作。他希望人工智能院在湖北省、武汉市政府的领导以及兄弟院校和企业的支持下，通过自身不断努力，把人工智能学科发展提升到一个很高的水平。
        </p>
		<img src="@/static/news3.webp"
             style="width:60%;margin-left:20%;">
        <p>
            兄弟院校代表、北京航空航天大学自动化科学与电气工程学院院长吕金虎指出，武汉大学在人工智能各个交叉学科方面有着深厚底蕴，人工智能研究院通过学科的交叉融合，集中人才资源优势，进一步推动人工智能与传统学科融合，成为武汉大学不断创新发展动力源泉。
        </p>
		<img src="@/static/news4.webp"
             style="width:60%;margin-left:20%;">
        <p>
            窦贤康、吴建平、刘经南、葛琳、陈建共、李记泽共同为武汉大学人工智能研究院揭牌。
		</p>
		<img src="@/static/news5.webp"
             style="width:60%;margin-left:20%;">
        <p>
            据了解，据悉武汉大学“人工智能研究院”是依托于计算机学院，联合17个校内单位共同建设的大型跨学科研究平台。研究院将引进和培养一批人工智能研究领域的领军人才，开展多层次专业人才培养，促进科研成果转化和技术落地，使之成为人工智能基础理论研究、多学科交叉应用技术创新、相关领域学者学术交流的平台，为武汉大学各学科发展和科技创新服务。</p>
        <Comment article_id="7"/>
    </div>
</template>

<script>
    import Comment from '@/components/Comment.vue'

    export default {
        name: "news",
        components: {
            Comment
        },
    }
</script>

<style scoped>
p {
	font-size:150%;
}
h2 {
	font-size:150%;
	text-align: center;
}
</style>