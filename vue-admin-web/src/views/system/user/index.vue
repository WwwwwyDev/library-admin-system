<template>
  <div class="app-container">
    <el-container>
      <el-header>
        <el-col :span="4">
          <el-input placeholder="请输入查找的用户名" v-model.lazy="usernameSearchInput" clearable>
          </el-input>
        </el-col>
        <el-col :span="1" :offset="0" style="padding-left: 10px;">
          <el-button icon="el-icon-search" circle @click="searchUsername"></el-button>
        </el-col>
        <el-col :span="1" :offset="1">
          <el-button type="primary" @click="addDialogVisible=true" round>添加用户</el-button>
        </el-col>

      </el-header>
      <el-main>
        <el-table v-loading="listLoading" :data="list" element-loading-text="加载中" border fit highlight-current-row>
          <el-table-column align="center" label="ID" width="95" sortable>
            <template slot-scope="scope">
              {{ scope.row.id }}
            </template>
          </el-table-column>
          <el-table-column label="头像" width="100" align="center">
            <template slot-scope="scope">
              <el-avatar :size="50" :src="scope.row.avatar"></el-avatar>
            </template>
          </el-table-column>
          <el-table-column label="用户名" width="150" sortable>
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
                <el-button icon="el-icon-setting" circle></el-button>
              </el-tooltip>
              <el-popconfirm title="确定删除吗?" @onConfirm="deleteRow(scope.row.id)" icon="el-icon-info" icon-color="red">
                <el-button slot="reference" type="danger" icon="el-icon-delete" circle></el-button>
              </el-popconfirm>
              <el-button type="primary" icon="el-icon-edit" circle></el-button>
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

    <el-dialog title="添加用户" :visible.sync="addDialogVisible">
      <el-form :model="addForm" label-width="120px">
        <el-form-item label="用户账号">
          <el-input placeholder="请输入账号" v-model="addForm.username" clearable>
          </el-input>
        </el-form-item>
        <el-form-item label="用户密码">
          <el-input placeholder="请输入密码" v-model="addForm.password" show-password></el-input>
        </el-form-item>
        <el-form-item label="用户头像">
          <el-upload class="avatar-uploader" action="" :show-file-list="false" :auto-upload="false"
            :on-change="changePhotoFile">
            <img v-if="imageUrl" :src="imageUrl" class="avatar">
            <i v-else class="el-icon-plus avatar-uploader-icon"></i>
          </el-upload>
          <imageCropper ref="imageCropper" @getFile="getFile"></imageCropper>
        </el-form-item>
        <el-form-item label="用户信息">
          <el-input type="textarea" :rows="2" placeholder="请输入用户信息" v-model="addForm.info">
          </el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="addDialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="addDialogVisible = false">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
  import {
    getLoginStatus,
    getUsers,
    deleteUser
  } from '@/api/usertable'
  import {
    getQiniuToken
  } from '@/api/qiniu'
  import {
    upload,
    domain
  } from '@/utils/qiniu'
  import imageCropper from '@/components/ImageCropper'
  export default {
    components: {
      imageCropper
    },
    data() {
      return {
        imageUrl: "",
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
        }
      }
    },
    created() {
      this.fetchLoginStatus()
      this.fetchData({
        "page": this.page,
        "limit": this.limit
      })
    },
    methods: {
      async fetchLoginStatus() {
        this.listLoading = true
        let res = await getLoginStatus()
        let loginStatusHandle = res.data.loginStatus.map((item) => parseInt(item))
        this.loginStatus = new Set(loginStatusHandle)
        this.listLoading = false
      },
      async fetchData(params) {
        this.listLoading = true
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
      deleteRow(id) {
        if (this.loginStatus.has(id)) {
          return this.$message({
            message: "无法删除在线用户",
            type: 'error'
          });
        }
        deleteUser(id).then(response => {
          this.$message({
            message: response.msg,
            type: 'success'
          });
          this.fetchData({
            "page": this.page,
            "limit": this.limit,
            "username": this.usernameSearchInput
          })
        })
      },
      changePhotoFile(file, fileList) {
        if (fileList.length > 0) {
          this.photoList = [fileList[fileList.length - 1]];
        }
        var size = file.raw.size;
        var type = file.raw.type;
        if (size > 512000) {
          this.$message({
            message: "图片不能超过500k",
            type: "error",
          });
          return;
        }
        if (type != "image/jpeg" && type != "image/jpg" && type != "image/png") {
          this.$message({
            message: "仅支持上传jpg、jpeg、png图片",
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
      async getFile(file) {
        let res = await getQiniuToken()
        let {
          token
        } = res.data
        const observable = upload(file, token)
        var _this = this
        const loading = this.$loading({
          lock: true,
          text: '上传中',
          spinner: 'el-icon-loading',
          background: 'rgba(0, 0, 0, 0.7)'
        });
        const subscription = observable.subscribe({
          next(res) {
            //进度
          },
          error(err) {
            return _this.$message({
              message: "图片上传失败",
              type: 'error'
            });
            loading.close();
          },
          complete(res) {
            _this.imageUrl = domain + res.hash
            loading.close();
          }
        })
        //this.imageUrl = window.URL.createObjectURL(file);
        //this.temp.File = file; //将file保存到表单对象，与表单一起提交
        this.$refs.imageCropper.close();
      }
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
  }
</style>
