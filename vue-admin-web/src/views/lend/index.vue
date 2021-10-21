<template>
  <div class="app-container">
    <el-container>
      <el-header>
        <el-col :span="4">
          <el-input placeholder="请输入查找的会员卡号" v-model.lazy="searchVipCardNumberInput" clearable>
          </el-input>
        </el-col>
        <el-col :span="4" :offset="1">
          <el-input placeholder="请输入查找的书名" v-model.lazy="searchBookNameInput" clearable>
          </el-input>
        </el-col>
        <el-col :span="1" :offset="0" style="padding-left: 10px;">
          <el-button icon="el-icon-search" circle @click="search"></el-button>
        </el-col>
        <el-col :span="1" :offset="1">
          <el-button type="primary" @click="openAddDialog" round>添加图书</el-button>
        </el-col>
      </el-header>
      <el-main>
        <el-table v-loading="listLoading" :data="list" element-loading-text="Loading" border fit highlight-current-row>
<!--          <el-table-column align="center" label="ID" width="95">
            <template slot-scope="scope">
              {{ scope.row.id }}
            </template>
          </el-table-column> -->
          <el-table-column label="借阅会员卡号" width="300" align="center">
            <template slot-scope="scope">
              <el-tag>{{ scope.row.vipCardNumber }}</el-tag>
            </template>
          </el-table-column>
          <el-table-column label="借阅书名" width="300" align="center">
            <template slot-scope="scope">
              {{ scope.row.bookName }}
            </template>
          </el-table-column>
          <el-table-column label="借阅时间">
            <template slot-scope="scope">
              <span>{{ new Date(scope.row.lendTime.seconds * 1000).toLocaleString() }}</span>
            </template>
          </el-table-column>
          <el-table-column label="操作" width="150" fixed="right">
            <template slot-scope="scope">
              <el-popconfirm title="确定删除吗?" @onConfirm="deleteRow(scope.row.id)" icon="el-icon-info" icon-color="red">
                <el-button slot="reference" type="danger" icon="el-icon-delete" circle></el-button>
              </el-popconfirm>
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
    <el-dialog title="图书借阅" :visible.sync="addDialogVisible" width="400px">
      <el-form :model="addForm" ref="addForm" :rules="FormRules" label-width="90px" @submit.native.prevent>
        <el-form-item label="借阅会员" prop="vipChecked">
          <el-select v-model="addForm.vipChecked" filterable remote reserve-keyword placeholder="请输入搜索关键词"
            :remote-method="remoteVipMethod" :loading="addForm.vipLoading">
            <el-option v-for="item in addForm.vipOptions" :key="item.id" :label="item.cardNumber" :value="item.id">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="借阅图书" prop="bookChecked">
          <el-select v-model="addForm.bookChecked" filterable remote reserve-keyword placeholder="请输入搜索关键词"
            :remote-method="remoteBookMethod" :loading="addForm.bookLoading">
            <el-option v-for="item in addForm.bookOptions" :key="item.id" :label="item.name" :value="item.id">
            </el-option>
          </el-select>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="addDialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="submitAddForm">确 定</el-button>
      </span>
    </el-dialog>


  </div>
</template>

<script>
  import {
    mapActions
  } from 'vuex'
  import {
    getLends,
    deleteLend,
    addLend
  } from '@/api/lendtable'

  import {
    getVipByCardNumberLike,
    getBookByNameLike
  } from '@/api/search'

  export default {
    data() {
      return {
        searchVipCardNumberInput: '',
        searchBookNameInput: '',
        list: null,
        listLoading: true,
        page: 1,
        limit: 10,
        total: 0,
        addDialogVisible: false,
        addForm: {
          vipOptions: [],
          vipChecked: null,
          vipLoading: false,
          vipId: 0,
          bookOptions: [],
          bookChecked: null,
          bookLoading: false,
          bookId: 0,
          vipCardNumber: "",
          bookName: ""
        },
        FormRules: {
          bookChecked: [{
            required: true,
            message: "请选择借阅的图书",
            trigger: 'change',
            type: 'number'
          }],
          vipChecked: [{
            required: true,
            message: "请选择借阅的会员",
            trigger: 'change',
            type: 'number'
          }]
        },
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
        let res = await getLends(params)
        let {
          lends,
          total
        } = res.data
        this.list = lends
        this.total = total
        this.listLoading = false
      },
      handleSizeChange(limit) {
        this.fetchData({
          "page": this.page,
          "limit": limit,
          "vipCardNumber": this.searchVipCardNumberInput,
          "bookName": this.searchBookNameInput
        })
        this.limit = limit
      },
      handleCurrentChange(page) {
        this.fetchData({
          "page": page,
          "limit": this.limit,
          "vipCardNumber": this.searchVipCardNumberInput,
          "bookName": this.searchBookNameInput
        })
        this.page = page
      },
      search() {
        this.fetchData({
          "page": 1,
          "limit": this.limit,
          "vipCardNumber": this.searchVipCardNumberInput,
          "bookName": this.searchBookNameInput
        })
        this.page = 1
      },
      async deleteRow(id) {
        this.$message({
          message: "演示版本，拒绝对数据进行操作",
          type: 'error'
        });
        return
        let res = await deleteLend(id)
        this.$message({
          message: res.msg,
          type: 'success'
        });
        this.fetchData({
          "page": this.page,
          "limit": this.limit,
          "vipCardNumber": this.searchVipCardNumberInput,
          "bookName": this.searchBookNameInput
        })
      },
      clearAddForm() {
        this.addForm.vipOptions = []
        this.addForm.vipChecked = null
        this.addForm.vipLoading = false
        this.addForm.bookOptions = []
        this.addForm.bookChecked = null
        this.addForm.bookLoading = false
        this.addForm.vipId = 0
        this.addForm.bookId = 0
        this.addForm.vipCardNumber = ""
        this.addForm.bookName = ""
      },
      openAddDialog() {
        this.clearAddForm()
        this.addDialogVisible = true
      },
      submitAddForm() {
        this.$message({
          message: "演示版本，拒绝对数据进行操作",
          type: 'error'
        });
        return
        this.$refs["addForm"].validate(async (valid) => {
          if (valid) {
            if (this.addForm.bookChecked !== null)
              this.addForm.bookId = this.addForm.bookChecked
            if (this.addForm.vipChecked !== null)
              this.addForm.vipId = this.addForm.vipChecked
            for (let i = 0; i < this.addForm.vipOptions.length; i++) {
              let temp = this.addForm.vipOptions[i]
              if (temp.id == this.addForm.vipId) {
                this.addForm.vipCardNumber = temp.cardNumber
                break
              }
            }
            for (let i = 0; i < this.addForm.bookOptions.length; i++) {
              let temp = this.addForm.bookOptions[i]
              if (temp.id == this.addForm.bookId) {
                this.addForm.bookName = temp.name
                break
              }
            }
            let res = await addLend(this.addForm)
            if (res != undefined)
              this.$message({
                message: res.msg,
                type: 'success'
              });
            this.addDialogVisible = false
            this.fetchData({
              "page": this.page,
              "limit": this.limit,
              "vipCardNumber": this.searchVipCardNumberInput,
              "bookName": this.searchBookNameInput
            })
          }
        })
      },
      async remoteBookMethod(query) {
        if (query === "") return
        this.addForm.bookLoading = true
        let res = await getBookByNameLike(query)
        this.addForm.bookOptions = res.data.books
        this.addForm.bookLoading = false
      },
      async remoteVipMethod(query) {
        if (query === "") return
        this.addForm.vipLoading = true
        let res = await getVipByCardNumberLike(query)
        this.addForm.vipOptions = res.data.vips
        this.addForm.vipLoading = false
      },
      ...mapActions({
        refreshToken: 'user/refreshToken'
      })
    }
  }
</script>

