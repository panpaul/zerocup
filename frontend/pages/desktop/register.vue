<!--
 * @Date: 2020-10-17 20:05:10
 * @LastEditors: QiuJhao
 * @LastEditTime: 2020-10-22 09:46:27
-->
<template>
    <div class="register">
        <el-form :model="ruleForm" :rules="rules" label-width="100px" ref="ruleForm">
            <el-form-item label="用户名" prop="username">
                <el-input v-model="ruleForm.username"></el-input>
            </el-form-item>
            <el-form-item label="密码" prop="password">
                <el-input autocomplete="off" type="password" v-model="ruleForm.password"></el-input>
            </el-form-item>
            <el-form-item label="确认密码" prop="passwordConfirm">
                <el-input autocomplete="off" type="password" v-model="ruleForm.passwordConfirm"></el-input>
            </el-form-item>
            <el-form-item>
                <el-button @click="submitForm('ruleForm')" type="primary">立即注册</el-button>
                <el-button @click="resetForm('ruleForm')">重置</el-button>
            </el-form-item>
        </el-form>
    </div>

</template>
<script>
    import axios from 'axios'

    export default {
        name: 'register',
        data() {
            const validatePass = (rule, value, callback) => {
                if (value !== this.ruleForm.password) {
                    callback(new Error('两次输入密码不一致!'));
                } else {
                    callback();
                }
            };
            return {
                ruleForm: {
                    username: '',
                    password: '',
                    passwordConfirm: ''
                },
                rules: {
                    username: [
                        {required: true, message: '请输入用户名', trigger: 'blur'},
                        {min: 3, max: 20, message: '长度在 3 到 20 个字符', trigger: 'blur'}
                    ],
                    password: [
                        {required: true, message: '请输入密码', trigger: 'blur'},
                        {min: 6, max: 20, message: '长度在 6 到 20 个字符', trigger: 'blur'},
                    ],
                    passwordConfirm: [
                        {required: true, message: '请再次输入密码', trigger: 'blur'},
                        {min: 6, max: 20, message: '长度在 6 到 20 个字符', trigger: 'blur'},
                        {validator: validatePass, trigger: 'blur'},
                    ]
                }
            };
        },
        methods: {
            submitForm(formName) {
                this.$refs[formName].validate((valid) => {
                    if (valid) {
                        axios.post("/register", {
                            username: this.ruleForm.username,
                            password: this.ruleForm.password,
                            passwordConfirm: this.ruleForm.passwordConfirm
                        }).then(rep => {
                            console.log(rep);
                            if (rep.data.status === 200) {
                                this.$message.success("注册成功");
                                setTimeout(() => {
                                    location.href = "/desktop/login"
                                }, 1000)
                            } else if (rep.data.status === 40002) {
                                this.$message.error("用户名已被注册")
                            }
                        })
                    } else {
                        console.log('error submit!!');
                        return false;
                    }
                });
            },
            resetForm(formName) {
                this.$refs[formName].resetFields();
            }
        }
    }
</script>

<style scoped>
    .register {
        margin: 100px auto;
        width: 600px;
    }
</style>
