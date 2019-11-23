<template>
    <div class="container">
        <el-form :model="ruleForm" :rules="rules" ref="ruleForm" style="margin-top: 30px" v-if="$store.state.token">
            <el-form-item prop="content">
                <el-input :rows="3"
                          autofocus="true"
                          maxlength="250"
                          placeholder="写下你的评论"
                          resize="none"
                          show-word-limit
                          type="textarea"
                          v-model="ruleForm.content"></el-input>
            </el-form-item>
            <el-form-item>


                <el-button @click="submitComment" style="float: right" type="primary">提交评论</el-button>
            </el-form-item>
        </el-form>
        <div class="comment" v-for="item in comments" v-if="$store.state.token">

            <hr>
            <img class="comment_user_head" src="../static/user.png">
            <div>{{item.user.username}}</div>
            <span class="date">{{item.created_at}}</span>

            <div class="content" style="margin-left: 70px;margin-bottom: 20px">
                <p>{{item.content}}</p>
                <el-tag @click="showCommentInput(item)">回复</el-tag>
            </div>

            <div style="margin-left: 70px;">


                <div style="padding-bottom: 15px" v-for="reply in item.replys">
                    <img class="comment_user_head" src="../static/user.png">
                    <span>{{reply.user.username}}</span><span>: </span>
                    <span>@{{reply.reply_to.username}}</span>
                    <div class="date">{{reply.created_at}}</div>
                    <div style="margin-left: 70px">
                        <p>{{reply.content}}</p>
                        <el-tag @click="showCommentInput(item, reply)">回复</el-tag>
                    </div>
                </div>

                <div v-if="current_root_id === item.id">
                    <el-form :model="ruleForm2" :rules="rules2" ref="ruleForm2">
                        <el-form-item prop="content">
                            <el-input :placeholder="placeholder"
                                      :rows="3"
                                      autofocus="true"
                                      maxlength="250"
                                      resize="none"
                                      show-word-limit
                                      type="textarea"
                                      v-model="ruleForm2.content">
                            </el-input>
                        </el-form-item>
                        <el-form-item>
                            <el-button @click="submitReply" style="float: right;">确定</el-button>
                        </el-form-item>
                    </el-form>
                </div>
            </div>

        </div>
        <p v-else>要查看评论请先<a href="/login">登录</a></p>
    </div>
</template>

<script>
    import axios from 'axios'

    export default {
        name: "comment",
        props: {
            article_id: Number
        },
        data() {
            return {
                article: 1,  // 对某一资源进行评论
                current_root_id: '',
                current_parent_id: 0,
                current_reply_to_user_id: 0,
                placeholder: '写下你的评论',
                ruleForm: {
                    content: ''
                },
                rules: {
                    content: [
                        {required: true, message: '请输入评论内容', trigger: 'submit'},
                        {min: 3, message: '至少输入3个字符', trigger: 'submit'},
                    ]
                },
                ruleForm2: {
                    content: ''
                },
                rules2: {
                    content: [
                        {required: true, message: '请输入评论内容', trigger: 'submit'},
                        {min: 3, message: '至少输入3个字符', trigger: 'submit'},
                    ]
                },
                count: 0,
                comments: [
                    {
                        id: 0,
                        user: {
                            id: 0,
                            username: '',
                            head_img: ''
                        },
                        article_id: 0,
                        content: '',
                        created_at: '',
                        like_count: 0,
                        is_like: false,
                        replys: [  //回复，或子评论
                            {
                                id: 0,
                                user: {
                                    id: 0,
                                    username: '',
                                    head_img: ''
                                },
                                reply_to: {
                                    id: 0,
                                    username: '',
                                    head_img: ''
                                },
                                parent_id: 0,
                                root_id: 0,
                                content: '',
                                created_at: ''
                            }
                        ]
                    }
                ]
            }
        },
        created() {
		var aid = parseInt(this.article_id);
            window.console.log("/comment/"+aid);
            // 获取所有评论
            axios.get("/comment/"+aid, {
            //axios.get("/comment/" + 1, {
                headers: {token: this.$store.state.token}
            }).then(rep => {
                console.log(rep.data.status);
                if (rep.data.status === 200) {
                    this.comments = rep.data.data.items;
                    this.count = rep.data.data.total;
                }
            })
        },
        methods: {
            submitComment() {
				var aid = parseInt(this.article_id);
                this.$refs.ruleForm.validate((valid) => {
                    if (valid) {
                        axios.post("/comment", {
                            article_id: aid,
                            //article_id: 1,
                            content: this.ruleForm.content,
                        }, {
                            headers: {
                                token: this.$store.state.token
                            }
                        }).then(rep => {
                            console.log(rep);
                            this.comments = this.comments || [];
                            this.comments.unshift(rep.data.data);
                            this.ruleForm.content = ''
                        }).catch(err => {
                            console.log(err.response);
                        })
                    } else {
                        console.log('error submit!!');
                        return false;
                    }
                });
            },
            submitReply() {
				var aid = parseInt(this.article_id);
                this.$refs.ruleForm2[0].validate((valid) => {
                    if (valid) {
                        axios.post("/comment", {
                            reply_to_id: this.current_reply_to_user_id,
                            article_id: aid,
                            //article_id: 1,
                            parent_id: this.current_parent_id,
                            root_id: this.current_root_id,
                            content: this.ruleForm2.content,
                        }, {
                            headers: {
                                token: this.$store.state.token
                            }
                        }).then(rep => {
                            console.log(rep.data.data);
                            const comment = this.comments.find(item => item.id === this.current_root_id);
                            comment.replys = comment.replys || [];
                            comment.replys.push(rep.data.data);
                            this.ruleForm2.content = ''
                        }).catch(err => {
                            console.log(err);
                            console.log(err.response);
                        })
                    } else {
                        console.log('error submit!!');
                        return false;
                    }
                });
            },
            /**
             * 点击评论按钮显示输入框
             * item: 当前大评论
             * reply: 当前回复的评论
             */
            showCommentInput(item, reply) {
                // 如果回复了回复
                if (reply) {
                    this.placeholder = "@" + reply.user.username + " ";
                    this.current_reply_to_user_id = reply.user.id;
                    this.current_parent_id = reply.id;
                } else {
                    // 如果回复了评论
                    this.placeholder = '';
                    this.current_parent_id = item.id;
                    this.current_reply_to_user_id = item.user.id;
                }
                this.current_root_id = item.id;
            },
        }
    }
</script>

<style scoped>
    .container {
        max-width: 1000px;
        margin-left: auto;
        margin-right: auto;
    }

    .content {
        /*min-height: 30px;*/
        word-wrap: break-word;
        word-break: break-all;
        overflow: hidden;
    }

    .comment_user_head {
        width: 50px;
        height: 50px;
        border-radius: 50%;
        float: left;
        margin-right: 20px;
    }

    .el-tag {
        cursor: pointer;
    }

    .date {
        font-size: 13px;
    }
</style>