<template>
    <div id="robotContainer">
        <ul class="robotContent" ref="charView">
            <li :key="index" v-for="(item, index) in messageList">
                <span :class="'span'+(item.myself?'right':'left')">{{item.message}}</span>
            </li>
        </ul>
        <div class="robotFooter">
            <input @keyup.enter='chat' id="text" placeholder="说点什么吧..." type="text" v-model.trim="inputValue"/>
            <span @click="sendBtn()" id="btn">发送</span>
        </div>
    </div>
</template>

<script>
    import $ from 'jquery'

    var content = $(".robotContent");

    export default {

        name: 'robot',
        data() {
            return {
                //输入的内容,事先约定好的
                inputValue: '',
                //聊天的数组内容
                messageList: []
            }
        },
        methods: {
            sendBtn() {
                this.messageList.push({
                    message: this.inputValue,
                    myself: true
                });
                $.ajax({
                    url: 'http://www.tuling123.com/openapi/api',
                    type: 'post',
                    data: {
                        key: '38b9f17cf42084cafd7b88da1fae3c60',//apikey
                        info: this.inputValue
                    },
                    success: (data) => {
                        this.messageList.push({
                            message: data.text,
                            myself: false
                        });

                    }
                });
                this.inputValue = ""
            }
        },
        mounted() {
            this.box = this.$refs.charView;
        },
    }
</script>
<style scoped="">
    * {
        margin: 0;
        padding: 0;
        list-style: none;
    }

    #robotContainer {
        box-shadow: 10px 10px 25px #777;
        position: fixed;
        bottom: 30px;
        right: 30px;
        opacity: 1;
        width: 400px;
        height: 200px;
    }

    .robotFooter {
        width: 100%;
        height: 45px;
        background: #666;
        position: absolute;
        bottom: 0;
        overflow: hidden;
    }

    .robotFooter input {
        margin-top: 10px;
        width: 85%;
        height: 45%;
        outline: none;
        font-size: 16px;
        text-indent: 10px;
        position: absolute;
        border-radius: 6px;
        margin-left: 5px;
    }

    .robotFooter span {
        margin-top: 12px;
        display: inline-block;
        width: 10%;
        height: 50%;
        background: rgb(200, 221, 8);
        font-weight: 900;
        cursor: pointer;
        text-align: center;
        position: absolute;
        right: 6px;
        border-radius: 6px;
    }

    .robotContent {
        font-size: 20px;
        width: 95%;
        height: 100%;
        overflow: auto;
        margin-left: auto;
        margin-right: auto;
    }

    .robotContent li {
        margin-top: 10px;
        display: block;
        clear: both;
        overflow: hidden;
    }

    .robotContent li span {
        padding: 10px;
        border-radius: 10px;
        max-width: 310px;
        border: 1px solid #ccc;
        box-shadow: 0 0 3px #ccc;
        word-break: break-all
    }

    .robotContent li span.spanleft {
        float: left;
        background: #fff;
    }

    .robotContent li span.spanright {
        float: right;
        background: #7cfc00;
    }


</style>