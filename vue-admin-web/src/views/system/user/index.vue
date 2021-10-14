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
          <el-button type="primary" round>添加用户</el-button>
        </el-col>

      </el-header>
      <el-main>
        <el-table v-loading="listLoading" :data="list" element-loading-text="加载中" border fit highlight-current-row>
          <el-table-column align="center" label="ID" width="95"  sortable>
            <template slot-scope="scope">
              {{ scope.row.id }}
            </template>
          </el-table-column>
          <el-table-column label="头像" width="100" align="center">
            <template slot-scope="scope">
              <el-avatar :size="50" :src="scope.row.avatar"></el-avatar>
            </template>
          </el-table-column>
          <el-table-column label="用户名" width="150"  sortable>
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

  </div>
</template>

<script>
  import {
    getLoginStatus,
    getUsers,
    deleteUser
  } from '@/api/usertable'

  export default {
    data() {
      return {
        usernameSearchInput: '',
        page: 1,
        limit: 10,
        total: 0,
        list: [],
        listLoading: true,
        loginStatus: null
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
      fetchLoginStatus() {
        this.listLoading = true,
          getLoginStatus().then(response => {
            let loginStatusHandle = response.data.loginStatus.map((item) => parseInt(item))
            this.loginStatus = new Set(loginStatusHandle)
            this.listLoading = false
          })
      },
      fetchData(params) {
        this.listLoading = true
        getUsers(params).then(response => {
          let {
            users,
            total
          } = response.data
          for(let i = 0; i < users.length; i++){
            if(this.loginStatus.has(users[i].id)){
              users[i]["isLogin"] = true
            }else{
              users[i]["isLogin"] = false
            }
          }
          this.list = users
          this.total = total
          this.listLoading = false
        })
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
        if(this.loginStatus.has(id)){
          return  this.$message({
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
      }
    }
  }
</script>
