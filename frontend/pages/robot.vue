<template>
    <div id="container">
        <ul class="content" ref="charView">
            <li v-for="(item, index) in messageList" :key="index">
                <span :class="'span'+(item.myself?'right':'left')">{{item.message}}</span>
            </li>
        </ul>
        <div class="footer">
            <input id="text" type="text" v-model.trim="inputValue" @keyup.enter='chat' placeholder="说点什么吧...">
            <span id="btn" @click="btn2()">发送</span>
        </div>
    </div>
</template>

<script>
import $ from 'jquery'
var content  = $(".content")

export default{
	
    name:'robot',
    data() {
        return {
             //输入的内容,事先约定好的
            inputValue: '',
            //聊天的数组内容
            messageList: []
        }
    },
    methods: {
            btn2(){
                this.messageList.push({
                    message: this.inputValue,
                    myself: true
                })
                $.ajax({
                    url: 'http://www.tuling123.com/openapi/api',
                    type: 'post',
                    data: {
                        key: '',//apikey
                        info: this.inputValue
                    },
                    success: (data) => {
                        this.messageList.push({
                            message: data.text,
                            myself: false
                        });
                        
                    }
                })                             
                this.inputValue=""
            }
         },
         mounted() {
             this.box=this.$refs.charView;
         },
}
</script>
<style scoped="">
     * {
            margin: 0;
            padding: 0;
            list-style: none;
            font-family: '微软雅黑'
        }
        #container {
            width: 40vw;
            height: 70vh;
            background: #eee;
            position: relative;
            box-shadow: 10px 10px 25px #777;
        }
        .footer {
            width: 40vw;
            height: 70px;
            background: #666;
            position: absolute;
            bottom: 0;
            overflow: hidden;
            
        }
        .footer input {
            margin-top: 10px;
            width: 33vw;
            height: 5vh;
            outline: none;
            font-size: 16px;
            text-indent: 10px;
            position: absolute;
            border-radius: 6px;
           margin-left: 5px;
        }        
        .footer span {
            margin-top: 10px;
            display: inline-block;
            width: 5vw;
            height: 5vh;
            background: rgb(200, 221, 8);
            font-weight: 900;
            line-height: 45px;
            cursor: pointer;
            text-align: center;
            position: absolute;
            right: 10px;
            border-radius: 6px;
        }
        .content {
            font-size: 20px;
            width: 40vw;
            height: 61vh;
            overflow: auto;
        }
        .content li {
            margin-top: 10px;
            display: block;
            clear: both;
            overflow: hidden;
        }
        .content li span{
            padding: 10px;
            border-radius: 10px;
            max-width: 310px;
            border: 1px solid #ccc;
            box-shadow: 0 0 3px #ccc;
            word-break: break-all 
        }
        .content li span.spanleft { 
            float: left;
            background: #fff;
        }
        .content li span.spanright {
            float: right;
            background: #7cfc00;
        } 
</style>