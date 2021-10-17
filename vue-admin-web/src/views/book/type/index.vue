<template>
  <div class="app-container">
    <el-container>
      <el-header>
        <el-col :span="4">
          <el-input placeholder="请输入查找的种类名" v-model.lazy="searchNameInput"  clearable>
          </el-input>
        </el-col>
        <el-col :span="1" :offset="0" style="padding-left: 10px;">
          <el-button type="primary" icon="el-icon-search"  @click="search">搜索</el-button>
        </el-col>
        <el-col :span="1" :offset="3">
          <el-button type="primary"  round>添加用户</el-button>
        </el-col>
      </el-header>
      <el-main>
        <el-table v-loading="listLoading" :data="list" element-loading-text="Loading" border fit highlight-current-row>
          <el-table-column align="center" label="ID" width="95">
            <template slot-scope="scope">
              {{ scope.row.id }}
            </template>
          </el-table-column>
          <el-table-column label="种类名" width="100" align="center">
            <template slot-scope="scope">
              {{ scope.row.name }}
            </template>
          </el-table-column>
          <el-table-column label="种类介绍" >
              <template slot-scope="scope">
                <span>{{ scope.row.intro }}</span>
              </template>
          </el-table-column>
          <el-table-column label="操作" width="150" fixed="right">
            <template slot-scope="scope">
              <el-tooltip class="item" effect="dark" content="查询该类目下的图书" placement="top">
                <el-button icon="el-icon-search" circle></el-button>
                </el-tooltip>
              <el-popconfirm title="确定删除吗?" icon="el-icon-info" icon-color="red">
                <el-button slot="reference" type="danger" icon="el-icon-delete" circle></el-button>
              </el-popconfirm>
              <el-button type="primary" icon="el-icon-edit" circle></el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-main>
      <el-footer>
        <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange" :current-page="page"
          :page-sizes="[1, 10, 20, 100]" :page-size="limit" layout="total, sizes, prev, pager, next, jumper"
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
    getTypes
  } from '@/api/typetable'

  export default {
    data() {
      return {
        searchNameInput:'',
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
        let res = await getTypes(params)
        let {
          types,
          total
        } = res.data
        this.list = types
        this.total = total
        this.listLoading = false
      },
      handleSizeChange(limit) {
        this.fetchData({
          "page": this.page,
          "limit": limit,
          "name": this.searchNameInput,
        })
        this.limit = limit
      },
      handleCurrentChange(page) {
        this.fetchData({
          "page": page,
          "limit": this.limit,
          "name": this.searchNameInput,
        })
        this.page = page
      },
      search(){
        this.fetchData({
          "page": 1,
          "limit": this.limit,
          "name": this.searchNameInput,
        })
        this.page = 1
      },
      ...mapActions({
        refreshToken: 'user/refreshToken'
      })
    }
  }
</script>
