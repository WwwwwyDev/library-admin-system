<template>
  <div class="app-container">
    <el-backtop :bottom="60"></el-backtop>
    <el-container>
      <el-header>
        <el-row>
          <el-col :span="4">
            <el-input placeholder="请输入查找的用户名" v-model.lazy="usernameSearchInput" clearable>
            </el-input>
          </el-col>
          <el-col :span="1" :offset="0" style="padding-left: 10px;">
            <el-button icon="el-icon-search" circle @click="searchUsername"></el-button>
          </el-col>
          <el-col :span="1" :offset="1">
            <el-button type="primary" @click="openAddDialog" round>添加用户</el-button>
          </el-col>
        </el-row>
      </el-header>
      <el-main>
        <el-table v-loading="listLoading" :data="list" element-loading-text="加载中" border fit highlight-current-row>
          <!--          <el-table-column align="center" label="ID" width="95">
            <template slot-scope="scope">
              {{ scope.row.id }}
            </template>
          </el-table-column> -->
          <el-table-column label="头像" width="100" align="center">
            <template slot-scope="scope">
              <el-avatar :size="50" :src="scope.row.avatar"></el-avatar>
            </template>
          </el-table-column>
          <el-table-column label="用户名" width="150">
            <template slot-scope="scope">
              {{ scope.row.username }}
            </template>
          </el-table-column>
          <el-table-column label="用户信息">
            <template slot-scope="scope">
              <span>{{ scope.row.info }}</span>
            </template>
          </el-table-column>
          <el-table-column label="在线状态" width="70">
            <template slot-scope="scope">
              <template v-if="scope.row.isLogin">
                <el-tag type="success">在线</el-tag>
              </template>
              <template v-else>
                <el-tag type="danger">离线</el-tag>
              </template>
            </template>
          </el-table-column>
          <el-table-column label="操作" width="150" fixed="right">
            <template slot-scope="scope">
              <el-tooltip class="item" effect="dark" content="设置用户权限" placement="top">
                <el-button icon="el-icon-setting" circle @click="openRolesDialog(scope.row.id)"></el-button>
              </el-tooltip>
              <el-popconfirm title="确定删除吗?" @onConfirm="deleteRow(scope.row.id)" icon="el-icon-info" icon-color="red">
                <el-button slot="reference" type="danger" icon="el-icon-delete" circle></el-button>
              </el-popconfirm>
              <el-button type="primary" icon="el-icon-edit" circle
                @click="openEditDialog(scope.row.id,scope.row.username,scope.row.avatar,scope.row.info)"></el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-main>
      <el-footer>
        <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange" :current-page="page"
          :page-sizes="[5, 10, 20, 100]" :page-size="limit" layout="total, sizes, prev, pager, next, jumper"
          :total="total">
        </el-pagination>
      </el-footer>
    </el-container>

    <el-dialog title="添加用户" :visible.sync="addDialogVisible" width="500px">
      <el-form :model="addForm" :rules="FormRules" ref="addForm" label-width="90px" @submit.native.prevent>
        <el-form-item label="用户账号" prop="username">
          <el-input placeholder="请输入账号" v-model="addForm.username" clearable>
          </el-input>
        </el-form-item>
        <el-form-item label="用户密码" prop="password">
          <el-input placeholder="请输入密码" v-model="addForm.password" show-password></el-input>
        </el-form-item>
        <el-form-item label="用户头像">
          <el-upload class="avatar-uploader" action="" :show-file-list="false" :auto-upload="false"
            :on-change="changePhotoFile">
            <img v-if="addForm.avatar" :src="addForm.avatar" class="avatar">
            <i v-else class="el-icon-plus avatar-uploader-icon"></i>
          </el-upload>
          <imageCropper ref="imageCropper" @getFile="getAddFile"></imageCropper>
        </el-form-item>
        <el-form-item label="用户信息">
          <el-input type="textarea" :rows="2" placeholder="请输入用户信息" v-model="addForm.info">
          </el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="addDialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="submitAddForm">确 定</el-button>
      </span>
    </el-dialog>

    <el-dialog title="编辑用户" :visible.sync="editDialogVisible" width="500px">
      <el-form :model="editForm" :rules="FormRules" ref="editForm" label-width="90px" @submit.native.prevent>
        <el-form-item label="用户账号">
          <el-input placeholder="" v-model="editForm.username" :disabled="true">
          </el-input>
        </el-form-item>
        <el-form-item label="修改密码" prop="password">
          <el-input placeholder="请输入密码" v-model="editForm.password" show-password></el-input>
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

    <el-dialog title="权限设置" :visible.sync="rolesDialogVisible" width="450px">
      <el-form ref="rolesForm" label-width="90px" @submit.native.prevent>
        <el-form-item label="用户账号">
          <el-input placeholder="" v-model="rolesForm.username" :disabled="true">
          </el-input>
        </el-form-item>
        <el-form-item label="用户权限">
          <el-select v-model="rolesForm.checkedRolesValues" multiple placeholder="请选择">
            <el-option v-for="item in rolesOption" :key="item.value" :label="item.label" :value="item.value">
            </el-option>
          </el-select>
          <el-popover placement="right" width="300" trigger="click">
            <el-table :data="roles">
              <el-table-column width="100" property="name" label="权限标签"></el-table-column>
              <el-table-column width="150" property="info" label="介绍"></el-table-column>
            </el-table>
            <i slot="reference" class="el-icon-info"></i>
          </el-popover>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="rolesDialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="submitRolesForm">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
  import {
    mapActions
  } from 'vuex'
  import {
    getLoginStatus,
    getUsers,
    deleteUser,
    addUser,
    editUser,
    getRoles,
    editUserRoles
  } from '@/api/usertable'
  import {
    getUserById
  } from '@/api/user'
  import {
    getQiniuToken
  } from '@/api/qiniu'
  import {
    domain,
    uploadAsync
  } from '@/utils/qiniu'
  import {
    userAvatar
  } from '@/settings'
  import imageCropper from '@/components/ImageCropper'
  export default {
    components: {
      imageCropper
    },
    data() {
      return {
        avatarFile: null,
        usernameSearchInput: '',
        page: 1,
        limit: 10,
        total: 0,
        list: [],
        listLoading: true,
        loginStatus: new Set(),
        addDialogVisible: false,
        addForm: {
          username: "",
          password: "",
          avatar: "",
          info: ""
        },
        editDialogVisible: false,
        editForm: {
          id: 0,
          username: "",
          password: "",
          avatar: "",
          info: ""
        },
        FormRules: {
          password: [{
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
          username: [{
              required: true,
              message: "请输入用户名",
              trigger: "blur"
            },
            {
              min: 6,
              message: "用户名长度不能小于6个字符",
              trigger: "blur"
            }
          ],
        },
        rolesDialogVisible: false,
        roles: [],
        rolesOption: [],
        rolesForm: {
          id: 0,
          username: "",
          checkedRolesValues: [],
        }
      }
    },
    async created() {
      this.listLoading = true
      await this.fetchRoles()
      await this.fetchLoginStatus()
      await this.fetchData({
        "page": this.page,
        "limit": this.limit
      })
    },
    methods: {
      async fetchRoles() {
        let res = await getRoles()
        this.roles = res.data.roles
        for (let i = 0; i < this.roles.length; i++) {
          let temp = {
            "value": this.roles[i].id,
            "label": this.roles[i].name
          }
          this.rolesOption.push(temp)
        }
      },
      async fetchLoginStatus() {
        let res = await getLoginStatus()
        let loginStatusHandle = res.data.loginStatus.map((item) => parseInt(item))
        this.loginStatus = new Set(loginStatusHandle)
      },
      async fetchData(params) {
        this.listLoading = true
        try {
          await this.refreshToken()
        } catch (error) {
          location.reload()
        }
        let res = await getUsers(params)
        let {
          users,
          total
        } = res.data
        for (let i = 0; i < users.length; i++) {
          if (this.loginStatus.has(users[i].id)) {
            users[i]["isLogin"] = true
          } else {
            users[i]["isLogin"] = false
          }
          if (users[i].avatar === undefined) {
            users[i].avatar = userAvatar
          }
        }
        this.list = users
        this.total = total
        this.listLoading = false
      },
      handleSizeChange(limit) {
        this.fetchData({
          "page": this.page,
          "limit": limit,
          "username": this.usernameSearchInput
        })
        this.limit = limit
      },
      handleCurrentChange(page) {
        this.fetchData({
          "page": page,
          "limit": this.limit,
          "username": this.usernameSearchInput
        })
        this.page = page
      },
      searchUsername() {
        this.fetchData({
          "page": 1,
          "limit": this.limit,
          "username": this.usernameSearchInput
        })
        this.page = 1
      },
      async deleteRow(id) {
        this.$message({
          message: "演示版本，拒绝对数据进行操作",
          type: 'error'
        });
        return
        if (this.loginStatus.has(id)) {
          return this.$message({
            message: "无法删除在线用户",
            type: 'error'
          });
        }
        let res = await deleteUser(id)
        this.$message({
          message: res.msg,
          type: 'success'
        });
        this.fetchData({
          "page": this.page,
          "limit": this.limit,
          "username": this.usernameSearchInput
        })
      },
      openAddDialog() {
        this.clearAddForm()
        this.addDialogVisible = true
      },
      openEditDialog(id, username, avatar, info) {
        this.clearEditForm()
        this.editForm.id = id
        this.editForm.avatar = avatar
        this.editForm.username = username
        this.editForm.password = ""
        this.editForm.info = info
        this.editDialogVisible = true
      },
      async openRolesDialog(id) {
        this.clearRolesForm()
        let res = await getUserById(id)
        let {
          user
        } = res.data
        this.rolesForm.id = id
        this.rolesForm.username = user.username
        let userRoles = user.roles === undefined ? [] : user.roles
        for (let i = 0; i < userRoles.length; i++) {
          this.rolesForm.checkedRolesValues.push(userRoles[i].id)
        }
        this.rolesDialogVisible = true
      },
      clearAddForm() {
        this.avatarFile = null
        this.addForm.avatar = ""
        this.addForm.username = ""
        this.addForm.password = ""
        this.addForm.info = ""
      },
      clearEditForm() {
        this.avatarFile = null
        this.editForm.avatar = ""
        this.editForm.username = ""
        this.editForm.password = ""
        this.editForm.info = ""
      },
      clearRolesForm() {
        this.rolesForm.id = 0
        this.rolesForm.username = ""
        this.rolesForm.checkedRolesValues = []
      },
      submitAddForm() {
        this.$message({
          message: "演示版本，拒绝对数据进行操作",
          type: 'error'
        });
        return
        this.$refs["addForm"].validate(async (valid) => {
          if (valid) {
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
                this.addForm.avatar = await uploadAsync(this.avatarFile, token)
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
            let res = await addUser(this.addForm)
            if (res != undefined)
              this.$message({
                message: res.msg,
                type: 'success'
              });
            this.addDialogVisible = false
            this.fetchData({
              "page": this.page,
              "limit": this.limit,
              "username": this.usernameSearchInput
            })
          }
        });
      },
      submitEditForm() {
        this.$message({
          message: "演示版本，拒绝对数据进行操作",
          type: 'error'
        });
        return
        this.$refs["editForm"].validate(async (valid) => {
          if (valid) {
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
            let res = await editUser(this.editForm.id, this.editForm)
            if (res != undefined)
              this.$message({
                message: res.msg,
                type: 'success'
              });
            this.editDialogVisible = false
            this.fetchData({
              "page": this.page,
              "limit": this.limit,
              "username": this.usernameSearchInput
            })
          }
        });
      },
      async submitRolesForm() {
        this.$message({
          message: "演示版本，拒绝对数据进行操作",
          type: 'error'
        });
        return
        let res = await editUserRoles(this.rolesForm.id, {
          "roleIds": this.rolesForm.checkedRolesValues
        })
        if (res != undefined)
          this.$message({
            message: res.msg,
            type: 'success'
          });
        this.rolesDialogVisible = false
        this.fetchData({
          "page": this.page,
          "limit": this.limit,
          "username": this.usernameSearchInput
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
      async getAddFile(file) {
        this.avatarFile = file
        this.addForm.avatar = window.URL.createObjectURL(file)
        this.$refs.imageCropper.close();
      },
      //file:裁剪后的图
      async getEditFile(file) {
        this.avatarFile = file
        this.editForm.avatar = window.URL.createObjectURL(file)
        this.$refs.imageCropper.close();
      },
      ...mapActions({
        refreshToken: 'user/refreshToken'
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
