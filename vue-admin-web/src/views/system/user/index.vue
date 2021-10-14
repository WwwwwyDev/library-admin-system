<template>
  <div class="app-container">
    <el-container>
      <el-header>
        <el-col :span="4">
          <el-input placeholder="请输入查找的用户名" v-model="input" clearable>
          </el-input>
        </el-col>
        <el-col :span="1" :offset="0" style="padding-left: 10px;">
          <el-button icon="el-icon-search" circle></el-button>
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
          <el-table-column label="头像" width="110" align="center">
            <template slot-scope="scope">
              {{ scope.row.avatar }}
            </template>
          </el-table-column>
          <el-table-column label="用户名">
            <template slot-scope="scope">
              {{ scope.row.username }}
            </template>
          </el-table-column>
          <el-table-column label="用户信息" width="110" align="center">
            <template slot-scope="scope">
              <span>{{ scope.row.info }}</span>
            </template>
                      </el-table-column>
        </el-table>
      </el-main>
      <el-footer>
        <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange"
          :current-page="currentPage4" :page-sizes="[100, 200, 300, 400]" :page-size="100"
          layout="total, sizes, prev, pager, next, jumper" :total="400">
        </el-pagination>
      </el-footer>
    </el-container>

  </div>
</template>

<script>
  import {
    getList
  } from '@/api/table'

  export default {
    filters: {
      statusFilter(status) {
        const statusMap = {
          published: 'success',
          draft: 'gray',
          deleted: 'danger'
        }
        return statusMap[status]
      }
    },
    data() {
      return {
        input: '',
        list: [{
            "id": 1,
            "username": "superadmin",
            "password": "禁止访问该数据",
            "salt": "禁止访问该数据",
            "avatar": "https://p.qqan.com/up/2021-10/16337437846131956.png",
            "info": "超级管理员账号"
          },
          {
            "id": 2,
            "username": "admin1",
            "password": "禁止访问该数据",
            "salt": "禁止访问该数据",
            "avatar": "https://p.qqan.com/up/2021-10/16337437846131956.png",
            "info": "管理员账号"
          },
          {
            "id": 3,
            "username": "admin2",
            "password": "禁止访问该数据",
            "salt": "禁止访问该数据",
            "avatar": "https://p.qqan.com/up/2021-10/16337437846131956.png",
            "info": "管理员账号"
          },
          {
            "id": 4,
            "username": "admin3",
            "password": "禁止访问该数据",
            "salt": "禁止访问该数据",
            "avatar": "https://p.qqan.com/up/2021-10/16337437846131956.png",
            "info": "管理员账号"
          },
          {
            "id": 5,
            "username": "admin4",
            "password": "禁止访问该数据",
            "salt": "禁止访问该数据",
            "avatar": "https://p.qqan.com/up/2021-10/16337437846131956.png",
            "info": "管理员账号"
          },
          {
            "id": 6,
            "username": "admin5",
            "password": "禁止访问该数据",
            "salt": "禁止访问该数据",
            "avatar": "https://p.qqan.com/up/2021-10/16337437846131956.png",
            "info": "管理员账号"
          },
          {
            "id": 7,
            "username": "admin6",
            "password": "禁止访问该数据",
            "salt": "禁止访问该数据",
            "avatar": "https://p.qqan.com/up/2021-10/16337437846131956.png",
            "info": "管理员账号"
          },
          {
            "id": 8,
            "username": "admin7",
            "password": "禁止访问该数据",
            "salt": "禁止访问该数据",
            "avatar": "https://p.qqan.com/up/2021-10/16337437846131956.png",
            "info": "管理员账号"
          },
          {
            "id": 9,
            "username": "admin8",
            "password": "禁止访问该数据",
            "salt": "禁止访问该数据",
            "avatar": "https://p.qqan.com/up/2021-10/16337437846131956.png",
            "info": "管理员账号"
          },
          {
            "id": 10,
            "username": "admin9",
            "password": "禁止访问该数据",
            "salt": "禁止访问该数据",
            "avatar": "https://p.qqan.com/up/2021-10/16337437846131956.png",
            "info": "管理员账号"
          }
        ],
        listLoading: false
      }
    },
    created() {
      // this.fetchData()
    },
    methods: {
      fetchData() {
        this.listLoading = true
        getList().then(response => {
          this.list = response.data.items
          this.listLoading = false
        })
      }
    }
  }
</script>
