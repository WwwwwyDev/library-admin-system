<template>
  <div class="app-container">
    <el-container>
      <el-header>
        <el-col :span="4">
          <el-input placeholder="请输入查找的用户名" v-model.lazy="searchUsernameInput" clearable>
          </el-input>
        </el-col>
        <el-col :span="1" :offset="0" style="padding-left: 10px;">
          <el-button icon="el-icon-search" circle @click="search"></el-button>
        </el-col>
        <el-col :span="1" :offset="1">
          <el-button type="primary" @click="clear" round>清空日志</el-button>
        </el-col>
      </el-header>
      <el-main>
        <el-table v-loading="listLoading" :data="list" element-loading-text="Loading" border fit highlight-current-row>
<!--          <el-table-column align="center" label="ID" width="95">
            <template slot-scope="scope">
              {{ scope.row.id }}
            </template>
          </el-table-column> -->
          <el-table-column label="用户名" width="300" align="center">
            <template slot-scope="scope">
              {{ scope.row.username }}
            </template>
          </el-table-column>
          <el-table-column label="日志信息">
            <template slot-scope="scope">
              {{ scope.row.info }}
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
    mapActions
  } from 'vuex'
  import {
    getLoginLogs,
    deleteLoginLogs,
  } from '@/api/loginlog'

  export default {
    data() {
      return {
        searchUsernameInput: '',
        list: null,
        listLoading: true,
        page: 1,
        limit: 10,
        total: 0,
      }
    },
    created() {
      this.fetchData({
        "page": this.page,
        "limit": this.limit
      })
    },
    methods: {
      async fetchData(params) {
        this.listLoading = true
        try {
          await this.refreshToken()
        } catch (error) {
          location.reload()
        }
        let res = await getLoginLogs(params)
        let {
          logs,
          total
        } = res.data
        this.list = logs
        this.total = total
        this.listLoading = false
      },
      handleSizeChange(limit) {
        this.fetchData({
          "page": this.page,
          "limit": limit,
          "username": this.searchUsernameInput,
        })
        this.limit = limit
      },
      handleCurrentChange(page) {
        this.fetchData({
          "page": page,
          "limit": this.limit,
          "username": this.searchUsernameInput,
        })
        this.page = page
      },
      search() {
        this.fetchData({
          "page": 1,
          "limit": this.limit,
          "username": this.searchUsernameInput,
        })
        this.page = 1
      },
      async clear(){
        let res = await deleteLoginLogs()
        this.$message({
          message: res.msg,
          type: 'success'
        });
        this.fetchData({
          "page": 1,
          "limit": this.limit,
          "username": this.searchUsernameInput,
        })
      },
      ...mapActions({
        refreshToken: 'user/refreshToken'
      })
    }
  }
</script>
