<template>
  <div class="app-container">
    <el-container>
      <el-header>
        <el-col :span="4">
          <el-input placeholder="请输入查找的类目名" v-model.lazy="searchNameInput" clearable>
          </el-input>
        </el-col>
        <el-col :span="1" :offset="0" style="padding-left: 10px;">
          <el-button type="primary" icon="el-icon-search" @click="search">搜索</el-button>
        </el-col>
        <el-col :span="1" :offset="2">
          <el-button type="primary" @click="openAddDialog" round>添加类目</el-button>
        </el-col>
      </el-header>
      <el-main>
        <el-table v-loading="listLoading" :data="list" element-loading-text="Loading" border fit highlight-current-row>
          <!--          <el-table-column align="center" label="ID" width="95">
            <template slot-scope="scope">
              {{ scope.row.id }}
            </template>
          </el-table-column> -->
          <el-table-column label="类目名" width="100" align="center">
            <template slot-scope="scope">
              {{ scope.row.name }}
            </template>
          </el-table-column>
          <el-table-column label="类目介绍">
            <template slot-scope="scope">
              <span>{{ scope.row.intro }}</span>
            </template>
          </el-table-column>
          <el-table-column label="操作" width="150" fixed="right">
            <template slot-scope="scope">
              <el-tooltip class="item" effect="dark" content="查询该类目下的图书" placement="top">
                <el-button icon="el-icon-search"  @click="openBookDialog(scope.row.id)" circle></el-button>
              </el-tooltip>
              <el-popconfirm title="删除类目,同时会删除该类目下的图书" @onConfirm="deleteRow(scope.row.id)" icon="el-icon-info"
                icon-color="red">
                <el-button slot="reference" type="danger" icon="el-icon-delete" circle></el-button>
              </el-popconfirm>
              <el-button type="primary" icon="el-icon-edit"
                @click="openEditDialog(scope.row.id,scope.row.name,scope.row.intro)" circle></el-button>
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

    <el-dialog title="添加类目" :visible.sync="addDialogVisible" width="500px">
      <el-form :model="addForm" :rules="FormRules" ref="addForm" label-width="90px" @submit.native.prevent>
        <el-form-item label="类目名" prop="name">
          <el-input placeholder="请输入类目名" v-model="addForm.name" clearable>
          </el-input>
        </el-form-item>
        <el-form-item label="类目介绍">
          <el-input placeholder="请输入类目介绍" v-model="addForm.intro" clearable></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="addDialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="submitAddForm">确 定</el-button>
      </span>
    </el-dialog>

    <el-dialog title="修改类目" :visible.sync="editDialogVisible" width="500px">
      <el-form :model="editForm" :rules="FormRules" ref="editForm" label-width="90px" @submit.native.prevent>
        <el-form-item label="类目名" prop="name">
          <el-input placeholder="请输入类目名" v-model="editForm.name" clearable>
          </el-input>
        </el-form-item>
        <el-form-item label="类目介绍">
          <el-input placeholder="请输入类目介绍" v-model="editForm.intro" clearable></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="editDialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="submitEditForm">确 定</el-button>
      </span>
    </el-dialog>

    <el-dialog title="该类目下的图书" :visible.sync="bookDialogVisible" width="300px">
      <el-table :data="bookList" element-loading-text="Loading" border fit
        highlight-current-row>
        <el-table-column label="图书名" align="center">
          <template slot-scope="scope">
            {{ scope.row.name }}
          </template>
        </el-table-column>
      </el-table>
      <span slot="footer" class="dialog-footer">
        <el-button @click="bookDialogVisible = false">关闭</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
  import {
    mapActions
  } from 'vuex'
  import {
    getTypes,
    editType,
    deleteType,
    addType
  } from '@/api/typetable'
  import {
    getTypeById
  } from '@/api/search.js'
  export default {
    data() {
      return {
        searchNameInput: '',
        list: null,
        listLoading: true,
        page: 1,
        limit: 10,
        total: 0,
        addDialogVisible: false,
        addForm: {
          name: "",
          intro: ""
        },
        editDialogVisible: false,
        editForm: {
          id: 0,
          name: "",
          intro: ""
        },
        FormRules: {
          name: [{
            required: true,
            message: "请输入书名",
            trigger: "blur"
          }]
        },
        bookDialogVisible: false,
        bookList: []
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
      async deleteRow(id) {
        let res = await deleteType(id)
        this.$message({
          message: res.msg,
          type: 'success'
        });
        this.fetchData({
          "page": this.page,
          "limit": this.limit,
          "name": this.searchNameInput,
        })
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
      search() {
        this.fetchData({
          "page": 1,
          "limit": this.limit,
          "name": this.searchNameInput,
        })
        this.page = 1
      },
      openAddDialog() {
        this.clearAddForm()
        this.addDialogVisible = true
      },
      clearAddForm() {
        this.addForm.name = ""
        this.addForm.intro = ""
      },
      submitAddForm() {
        this.$refs["addForm"].validate(async (valid) => {
          if (valid) {
            let res = await addType(this.addForm)
            if (res != undefined)
              this.$message({
                message: res.msg,
                type: 'success'
              });
            this.addDialogVisible = false
            this.fetchData({
              "page": this.page,
              "limit": this.limit,
              "name": this.searchNameInput,
            })
          }
        })
      },
      openEditDialog(id, name, intro) {
        this.clearEditForm()
        this.editForm = {
          id,
          name,
          intro
        }
        this.editDialogVisible = true
      },
      clearEditForm() {
        this.editForm.id = 0
        this.editForm.name = ""
        this.editForm.intro = ""
      },
      submitEditForm() {
        this.$refs["editForm"].validate(async (valid) => {
          if (valid) {
            let res = await editType(this.editForm.id, this.editForm)
            if (res != undefined)
              this.$message({
                message: res.msg,
                type: 'success'
              });
            this.editDialogVisible = false
            this.fetchData({
              "page": this.page,
              "limit": this.limit,
              "name": this.searchNameInput,
            })
          }
        })
      },
      async openBookDialog(id){
        let res = await getTypeById(id)
        this.bookList = res.data.type.booksInfo
        this.bookDialogVisible = true
      },
      ...mapActions({
        refreshToken: 'user/refreshToken'
      })
    }
  }
</script>
