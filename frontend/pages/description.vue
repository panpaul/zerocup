<template>
    <div id="desc">
        <div class="img"><img alt="AI WordCloud" src="../static/wordcloud-alpha.png"/></div>
        <p><strong>AI</strong>就是<strong>人工智能</strong></p>
        <p>它是研究、开发用于模拟、延伸和扩展人的智能的理论、方法、技术及应用系统的一门新的<strong>技术科学</strong>。</p>
        <p>人工智能是<strong>计算机科学</strong>的一个分支，它企图了解智能的实质，并生产出一种新的能以人类智能相似的方式做出反应的智能机器，该领域的研究包括机器人、语言识别、图像识别、自然语言处理和专家系统等。
        </p>
        <p>人工智能从诞生以来，理论和技术日益成熟，应用领域也不断扩大，可以设想，未来人工智能带来的科技产品，将会是人类智慧的“容器”。</p>
    </div>
</template>

<script>
    export default {
        name: "description"
    }
</script>

<style scoped>
    p {
        text-align: left;
    }

    .img {
        float: right;
        width: 40%;
    }

    img {
        width: 100%;
    }
</style>