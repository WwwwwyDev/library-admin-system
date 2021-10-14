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
          <el-table-column align="center" label="ID" width="95">
            <template slot-scope="scope">
              {{ scope.row.id }}
            </template>
          </el-table-column>
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
          <el-table-column label="操作" fixed="right">
            <template slot-scope="scope">
              <el-tooltip class="item" effect="dark" content="编辑用户权限" placement="top">
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
        listLoading: true
      }
    },
    created() {
      this.fetchData({
        "page": this.page,
        "limit": this.limit
      })
    },
    methods: {
      fetchData(params) {
        this.listLoading = true
        getUsers(params).then(response => {
          let {
            users,
            total
          } = response.data
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
