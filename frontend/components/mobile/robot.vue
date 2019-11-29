<template>
    <div id="mobileRobotContainer">
        <ul class="robotContent" style="padding-left: 0;">
            <li :key="index" v-for="(item, index) in messageList">
                <span :class="'span'+(item.myself?'right':'left')">{{item.message}}</span>
            </li>
        </ul>
        <br>
        <div class="robotFooter">
            <label for="text"/>
            <input @keyup.enter='sendBtn' id="text" placeholder="说点什么吧..." type="text" v-model.trim="inputValue"/>
            <span @click="sendBtn()" id="btn">发送</span>
        </div>
    </div>
</template>

<script>
    import $ from 'jquery'

    export default {
        name: "mobileRobot",
        data() {
            return {
                inputValue: "",
                messageList: [],
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
        }
    }
</script>

<style scoped>
    #mobileRobotContainer {
        box-shadow: 10px 10px 25px #777;
        background-color: rgba(179, 192, 209, 0.5);
        position: absolute;
        margin-top: 10px;
        width: 90vw;
        min-height: 200px;
        left: 5vw;
    }

    .robotFooter {
        margin-top: 5px;
        width: 100%;
        height: 35px;
        background: #666;
        position: absolute;
        bottom: 0;
        overflow: hidden;
    }

    .robotFooter input {
        margin-top: 4px;
        width: 85%;
        height: 20px;
        outline: none;

        position: absolute;
        border-radius: 3%;
        margin-left: 5px;
    }

    .robotFooter span {
        margin-top: 5px;
        display: inline-block;
        width: 10%;
        height: 25px;
        background: rgb(200, 221, 8);
        cursor: pointer;
        text-align: center;
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
        margin-top: 0;
        display: block;
        clear: both;
        overflow: hidden;
    }

    .robotContent li span {
        padding: 2%;
        border-radius: 10px;
        max-width: 85vw;
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