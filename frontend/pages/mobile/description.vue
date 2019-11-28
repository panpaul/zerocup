<template>
    <div id="description">
        <h2>AI介绍</h2>

        <div id="container">
            <img alt="AI WordCloud" class="img1" src="../../static/ai_wordcloud.jpg"/>
            <p><strong>AI</strong>就是<strong>人工智能</strong></p>
            <p>它是研究、开发用于模拟、延伸和扩展人的智能的理论、方法、技术及应用系统的一门新的<strong>技术科学</strong>。</p>
            <p>人工智能是<strong>计算机科学</strong>的一个分支，它企图了解智能的实质，并生产出一种新的能以人类智能相似的方式做出反应的智能机器，该领域的研究包括机器人、语言识别、图像识别、自然语言处理和专家系统等。
            </p>
            <p>人工智能从诞生以来，理论和技术日益成熟，应用领域也不断扩大，可以设想，未来人工智能带来的科技产品，将会是人类智慧的“容器”。</p>
            <p>
                人工智能科学的主旨是研究和开发出智能实体，在这一点上它属于工程学。工程的一些基础学科自不用说，数学、逻辑学、归纳学、统计学、系统学、控制学、工程学、计算机科学，还包括对哲学、心理学、生物学、神经科学、认知科学、仿生学、经济学、语言学等其他学科的研究，可以说这是一个集数门学科精华的尖端学科中的尖端学科。所以说人工智能是一门综合学科。</p>
            <p>
                <strong>人工智能的分类</strong><br/>
                强人工智能：指能制造出真正地推理和解决问题的智能机器。比如日本动漫里的“阿童木”、钢铁侠的AI管家“贾维斯”就可以说是一个强人工智能。
                <br/>弱人工智能：只用于解决某些特定领域的问题，比如 AlphaGo，只会下围棋。
            </p>
            <img alt="img" src="../../static/desc.png">
        </div>
        <read :sound="require('../../static/audios/description.mp3')"/>
    </div>
</template>

<script>
    import read from "../../components/mobile/read";

    export default {
        name: "mobileDescription",
        components: {read},
        layout: "mobile",
    }
</script>

<style scoped>
    h2 {
        color: #6d727c;
        text-align: center;
    }

    img {
        width: 100%;
    }

    #container {
        width: 80vw;
        margin-right: auto;
        margin-left: auto;
    }
</style>