import axios from "~/plugins/axios";

<template>
	<div class="cc">
    <div class="comment" v-for="item in comments" :key="item.id">
      <hr>
<!--     展示当前评论用户   -->
      <img :src="item.user.head_img">
      <div>{{item.user.username}}</div>
      <span class="date">{{item.created_at}}</span>
      <p>{{item.content}}</p>
<!--    展示回复input框    -->
      <el-tag @click="showCommentInput(item)">回复</el-tag>

    <!--      展示评论下方的回复  -->
        <div v-for="reply in item.replys":key="item.id">
          <img :src="reply.user.head_img">
          <span>{{reply.user.username}}</span><span>: </span>
          <span>@{{reply.reply_to.username}}</span>
          <div class="date">{{reply.created_at}}</div>
          <p>{{reply.content}}</p>
          <el-tag @click="showCommentInput(item, reply)">回复</el-tag>
        </div>

  <!--    展示评论框    -->
        <div v-if="current_root_id===item.id">
          <el-form :model="ruleForm2" :rules="rules2" ref="ruleForm2">
            <el-form-item prop="content">
              <el-input v-model="ruleForm2.content"
                        type="textarea"
                        placeholder="写下你的评论">
              </el-input>
            </el-form-item>
            <el-form-item>
              <el-button style="float: right;" 
                @click="submitReply">确定</el-button>
            </el-form-item>
          </el-form>
        </div>
    </div>
</div>
</template>
<script>
export default {
  data() {
    return {
    // 当点击展示回复框的时候记录下回复了谁、父亲节点、根节点, 方便提交
      current_root_id: 0, 
      current_parent_id: 0,
      current_reply_to_user_id: 0,
      ruleForm2: {
        content: ''
      },
      rules2: {
        content: [
           { required: true, message: '请输入评论内容', trigger: 'submit' },
           { min: 3, message: '至少输入3个字符', trigger: 'submit' },
        ]
      },
    }
  },
  methods: {
      // 展示回复框
      showCommentInput(item, reply) {
        // 如果回复了回复
        if (reply) {
          this.ruleForm2.content = "@" + reply.user.username + " ";
          this.current_reply_to_user_id = reply.user.id;
          this.current_parent_id = reply.id;
        } else {
          // 如果回复了评论
          this.ruleForm2.content = '';
          this.current_parent_id = item.id;
          this.current_reply_to_user_id = item.user.id;
        }
        this.current_root_id = item.id;
      },
    // 提交回复
      submitReply() {
        this.$refs.ruleForm2[0].validate((valid) => {
          if (valid) {
            axios.post("/comment", {
              reply_to_id: this.current_reply_to_user_id,
              video_id: this.video_id,
              parent_id: this.current_parent_id,
              root_id: this.current_root_id,
              content: this.ruleForm2.content,
            }, {
              headers: {
                token: this.$store.state.token
              }
            }).then(rep => {
              // 获取根节点
              const comment = this.comments.find(item => item.id===this.current_root_id);
              comment.replys = comment.replys||[];
              // 插入到根节点的下方
              comment.replys.push(rep.data.data);
              this.ruleForm2.content = ''
            }).catch(err => {
              console.log(err);
            })
          } else {
            console.log('error submit!!');
            return false;
          }
        });
      },
  }
}
</script>