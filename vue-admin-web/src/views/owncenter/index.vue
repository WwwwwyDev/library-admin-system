<template>
  <div class="app-container">
    <el-container>
      <el-aside width="500px">
        <el-card style="background-color: #fcfcfc;">
          <el-row :gutter="20">
            <el-col :span="6" :offset="8">
              <el-avatar :size="100" :src="avatar"></el-avatar>
            </el-col>
          </el-row>
          <el-row :gutter="20">
            <el-divider></el-divider>
          </el-row>
          <el-row :gutter="20">
            <el-form label-width="80px">
              <el-form-item label="用户名">
                {{name}}
              </el-form-item>
              <el-form-item label="个人信息">
                <el-input type="textarea" v-model="info" :disabled="true"></el-input>
              </el-form-item>
              <el-form-item label="拥有权限">
                <el-tag v-for="item in rolesList" :key='item.index'>
                  {{ item }}
                </el-tag>
              </el-form-item>
            </el-form>
          </el-row>
          <el-row :gutter="20">
            <el-col :span="6" :offset="4">
              <el-button type="primary" @click="openEditDialog">编辑信息</el-button>
            </el-col>
            <el-col :span="6" :offset="4">
              <el-button type="primary" @click="openChangePwdDialog">修改密码</el-button>
            </el-col>
          </el-row>
        </el-card>
      </el-aside>
    </el-container>
    <el-dialog title="编辑信息" :visible.sync="editDialogVisible" width="500px">
      <el-form :model="editForm" ref="editForm" label-width="90px" @submit.native.prevent>
        <el-form-item label="用户账号">
          <el-input placeholder="" v-model="editForm.username" :disabled="true">
          </el-input>
        </el-form-item>
        <el-form-item label="用户头像">
          <el-upload class="avatar-uploader" action="" :show-file-list="false" :auto-upload="false"
            :on-change="changePhotoFile">
            <img v-if="editForm.avatar" :src="editForm.avatar" class="avatar">
            <i v-else class="el-icon-plus avatar-uploader-icon"></i>
          </el-upload>
          <imageCropper ref="imageCropper" @getFile="getEditFile"></imageCropper>
        </el-form-item>
        <el-form-item label="用户信息">
          <el-input type="textarea" :rows="2" placeholder="请输入用户信息" v-model="editForm.info">
          </el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="editDialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="submitEditForm">确 定</el-button>
      </span>
    </el-dialog>

    <el-dialog title="修改密码" :visible.sync="changePwdDialogVisible" width="400px">
      <el-form :model="changePwdForm" :rules="FormRules" ref="changePwdForm" label-width="90px" @submit.native.prevent>
        <el-form-item label="原始密码" prop="oldPassword">
          <el-input placeholder="请输入密码" v-model="changePwdForm.oldPassword" show-password>
          </el-input>
        </el-form-item>
        <el-form-item label="新的密码" prop="newPassword">
          <el-input placeholder="请输入密码" v-model="changePwdForm.newPassword" show-password></el-input>
        </el-form-item>
        <el-form-item label="再次输入" prop="newPasswordAgain">
          <el-input placeholder="请输入密码" v-model="changePwdForm.newPasswordAgain" show-password></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="changePwdDialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="submitChangePwdForm">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
  import imageCropper from '@/components/ImageCropper'
  import {
    mapGetters,
    mapActions
  } from 'vuex'
  import {
    getQiniuToken
  } from '@/api/qiniu'
  import {
    uploadAsync,
    domain
  } from '@/utils/qiniu'
  import {
    editUser,
    changePassword
  } from '@/api/user'
  export default {
    components: {
      imageCropper
    },
    data() {
      return {
        avatarFile: null,
        editDialogVisible: false,
        changePwdDialogVisible: false,
        rolesList: [],
        editForm: {
          username: "",
          avatar: "",
          info: ""
        },
        changePwdForm: {
          oldPassword: "",
          newPassword: "",
          newPasswordAgain: ""
        },
        FormRules: {
          oldPassword: [{
              required: true,
              message: "请输入密码",
              trigger: "blur"
            },
            {
              min: 6,
              message: "密码长度不能小于6个字符",
              trigger: "blur"
            }
          ],
          newPassword: [{
              required: true,
              message: "请输入密码",
              trigger: "blur"
            },
            {
              min: 6,
              message: "密码长度不能小于6个字符",
              trigger: "blur"
            }
          ],
          newPasswordAgain: [{
              required: true,
              message: "请输入密码",
              trigger: "blur"
            },
            {
              min: 6,
              message: "密码长度不能小于6个字符",
              trigger: "blur"
            }
          ],
        }
      }
    },
    created() {
      this.rolesList = this.roles == "" ? ["无权限"] : this.roles.split(",")
    },
    methods: {
      openEditDialog() {
        this.clearEditForm()
        this.editForm.username = this.name
        this.editForm.avatar = this.avatar
        this.editForm.info = this.info
        this.editDialogVisible = true
      },
      clearEditForm() {
        this.avatarFile = null
        this.editForm.username = ""
        this.editForm.avatar = ""
        this.editForm.info = ""
      },
      async submitEditForm() {
        if (this.avatarFile != null) {
          let res = await getQiniuToken()
          let {
            token
          } = res.data
          const loading = this.$loading({
            lock: true,
            text: '上传头像中',
            spinner: 'el-icon-loading',
            background: 'rgba(0, 0, 0, 0.7)'
          })
          try {
            this.editForm.avatar = await uploadAsync(this.avatarFile, token)
          } catch (err) {
            this.$message({
              message: "头像上传失败",
              type: 'error'
            });
            loading.close()
            return
          }
          loading.close()
        }
        let res = await editUser(this.editForm)
        this.editDialogVisible = false
        try {
          await this.refreshToken()
          await this.refreshUserInfo()
        } catch (error) {
          location.reload()
        }
      },
      openChangePwdDialog() {
        this.clearChangePwdForm()
        this.changePwdDialogVisible = true
      },
      clearChangePwdForm() {
        this.changePwdForm.oldPassword = ""
        this.changePwdForm.newPassword = "",
          this.changePwdForm.newPasswordAgain = ""
      },
      submitChangePwdForm() {
        this.$refs["changePwdForm"].validate(async (valid) => {
          if (valid) {
            if (this.changePwdForm.newPassword != this.changePwdForm.newPasswordAgain) {
              this.$message({
                message: "两次输入密码不一致",
                type: 'error'
              });
              return
            }
            let res = await changePassword(this.changePwdForm)
            this.$message({
              message: res.msg,
              type: 'success'
            });
            this.changePwdDialogVisible = false
          }
        })
      },
      changePhotoFile(file, fileList) {
        if (fileList.length > 0) {
          this.photoList = [fileList[fileList.length - 1]];
        }
        var size = file.raw.size;
        var type = file.raw.type;
        if (type != "image/jpeg" && type != "image/jpg" && type != "image/png") {
          this.$message({
            message: "仅支持上传jpg、jpeg、png图片",
            type: "error",
          });
          return;
        }
        let sizeKB = 200
        if (size > 1024 * sizeKB) {
          this.$message({
            message: `图片不能超过${sizeKB}k`,
            type: "error",
          });
          return;
        }
        this.handleCrop(file);
      },
      //file:原图
      handleCrop(file) {
        this.$nextTick(() => {
          this.$refs.imageCropper.open(file.raw || file)
        })
      },
      //file:裁剪后的图
      async getEditFile(file) {
        this.avatarFile = file
        this.editForm.avatar = window.URL.createObjectURL(file)
        this.$refs.imageCropper.close();
      },
      ...mapActions({
        refreshUserInfo: 'user/getInfo',
        refreshToken: 'user/refreshTokenNow'
      })
    },
    computed: {
      ...mapGetters({
        id: "user/id",
        name: "user/name",
        avatar: "user/avatar",
        info: "user/info",
        roles: "user/roles"
      })
    }
  }
</script>

<style>
  .avatar-uploader .el-upload {
    border: 1px dashed #d9d9d9;
    border-radius: 6px;
    cursor: pointer;
    position: relative;
    overflow: hidden;
    border-radius: 50%;
  }

  .avatar-uploader .el-upload:hover {
    border-color: #409EFF;
  }

  .avatar-uploader-icon {
    font-size: 28px;
    color: #8c939d;
    width: 178px;
    height: 178px;
    line-height: 178px;
    text-align: center;
  }

  .avatar {
    width: 178px;
    height: 178px;
    display: block;
    border-radius: 50%;
  }
</style>
