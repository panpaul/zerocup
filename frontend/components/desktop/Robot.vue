<template>
    <div class="floatImg">
        <transition name="el-fade-in">
            <div id="robotContainer" v-show="show">
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
        </transition>
        <div class="floatImg"><img @click="show = !show" alt="robot" src="../../static/1.gif"/></div>

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
                messageList: [],
                show: false,
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
                        key: '94f8d12688f7488fb0ebb5ccd89ff3da',//apikey
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
<style scoped>
    * {
        margin: 0;
        padding: 0;
        list-style: none;
    }

    .floatImg {
        bottom: -70px;
        right: -40px;
        position: fixed;
        float: bottom;
        z-index: 999;
    }

    #robotContainer {
        box-shadow: 10px 10px 25px #777;
        position: fixed;
        bottom: 220px;
        right: 30px;
        opacity: 1;
        width: 40%;
        max-width: 700px;
        height: 30%;
    }

    .robotFooter {
        width: 100%;
        height: 14%;
        background: #666;
        position: absolute;
        bottom: 0;
        overflow: hidden;
    }

    .robotFooter input {
        margin-top: 5px;
        width: 85%;
        height: 47%;
        outline: none;

        position: absolute;
        border-radius: 3%;
        margin-left: 5px;
    }

    .robotFooter span {
        margin-top: 6px;
        display: inline-block;
        width: 10%;
        height: 54%;
        background: rgb(200, 221, 8);
        cursor: pointer;
        text-align: center;
        font-size: 12px;
        position: absolute;
        right: 1%;
        border-radius: 6px;
    }

    .robotContent {
        font-size: 20px;
        width: 100%;
        height: 90%;
        overflow: auto;
        margin-left: auto;
        margin-right: auto;
    }

    .robotContent li {
        margin-top: 0%;
        display: block;
        clear: both;
        overflow: hidden;
    }

    .robotContent li span {
        padding: 2%;
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