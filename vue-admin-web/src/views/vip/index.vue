<template>
  <div class="app-container">
    <el-container>
      <el-header>
        <el-col :span="4">
          <el-input placeholder="请输入查找的会员卡号" v-model.lazy="searchCardNumberInput" clearable>
          </el-input>
        </el-col>
        <el-col :span="4" :offset="1">
          <el-input placeholder="请输入查找的会员名" v-model.lazy="searchNameInput" clearable>
          </el-input>
        </el-col>
        <el-col :span="1" :offset="0" style="padding-left: 10px;">
          <el-button icon="el-icon-search" circle @click="search"></el-button>
        </el-col>
        <el-col :span="1" :offset="1">
          <el-button type="primary" @click="openAddDialog" round>添加会员</el-button>
        </el-col>
      </el-header>
      <el-main>
        <el-table v-loading="listLoading" :data="list" element-loading-text="Loading" border fit highlight-current-row>
<!--          <el-table-column align="center" label="ID" width="95">
            <template slot-scope="scope">
              {{ scope.row.id }}
            </template>
          </el-table-column> -->
          <el-table-column label="会员卡号" width="300" align="center">
            <template slot-scope="scope">
              <el-tag>{{ scope.row.cardNumber }}</el-tag>
            </template>
          </el-table-column>
          <el-table-column label="会员名" width="100" align="center">
            <template slot-scope="scope">
              {{ scope.row.name }}
            </template>
          </el-table-column>
          <el-table-column label="信息">
            <template slot-scope="scope">
              <span>{{ scope.row.info }}</span>
            </template>
          </el-table-column>
          <el-table-column label="操作" width="150" fixed="right">
            <template slot-scope="scope">
              <el-popconfirm title="确定删除吗?" @onConfirm="deleteRow(scope.row.id)" icon="el-icon-info" icon-color="red">
                <el-button slot="reference" type="danger" icon="el-icon-delete" circle></el-button>
              </el-popconfirm>
              <el-button type="primary" icon="el-icon-edit"
                @click="openEditDialog(scope.row.id,scope.row.name,scope.row.info)" circle></el-button>
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
    <el-dialog title="添加会员" :visible.sync="addDialogVisible" width="500px">
      <el-form :model="addForm" ref="addForm" label-width="90px" @submit.native.prevent>
        <el-form-item label="会员名">
          <el-input placeholder="请输入会员名" v-model="addForm.name" clearable>
          </el-input>
        </el-form-item>
        <el-form-item label="会员信息">
          <el-input placeholder="请输入会员信息" v-model="addForm.info" clearable></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="addDialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="submitAddForm">确 定</el-button>
      </span>
    </el-dialog>

    <el-dialog title="编辑会员" :visible.sync="editDialogVisible" width="500px">
      <el-form :model="editForm" ref="editForm" label-width="90px" @submit.native.prevent>
        <el-form-item label="会员名">
          <el-input placeholder="请输入会员名" v-model="editForm.name" clearable>
          </el-input>
        </el-form-item>
        <el-form-item label="会员信息">
        <el-input placeholder="请输入会员信息" v-model="editForm.info" clearable></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="editDialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="submitEditForm">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
  import {
    mapActions
  } from 'vuex'
  import {
    getVips,
    deleteVip,
    addVip,
    editVip
  } from '@/api/viptable'

  export default {
    data() {
      return {
        searchCardNumberInput: '',
        searchNameInput: '',
        list: null,
        listLoading: true,
        page: 1,
        limit: 10,
        total: 0,
        addDialogVisible: false,
        addForm: {
          name: "",
          info: ""
        },
        editDialogVisible: false,
        editForm: {
          id: 0,
          name: "",
          info: ""
        }
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
        let res = await getVips(params)
        let {
          vips,
          total
        } = res.data
        this.list = vips
        this.total = total
        this.listLoading = false
      },
      handleSizeChange(limit) {
        this.fetchData({
          "page": this.page,
          "limit": limit,
          "cardNumber": this.searchCardNumberInput,
          "name": this.searchNameInput
        })
        this.limit = limit
      },
      handleCurrentChange(page) {
        this.fetchData({
          "page": page,
          "limit": this.limit,
          "cardNumber": this.searchCardNumberInput,
          "name": this.searchNameInput
        })
        this.page = page
      },
      search() {
        this.fetchData({
          "page": 1,
          "limit": this.limit,
          "cardNumber": this.searchCardNumberInput,
          "name": this.searchNameInput
        })
        this.page = 1
      },
      async deleteRow(id) {
        this.$message({
          message: "演示版本，拒绝对数据进行操作",
          type: 'error'
        });
        return
        let res = await deleteVip(id)
        this.$message({
          message: res.msg,
          type: 'success'
        });
        this.fetchData({
          "page": this.page,
          "limit": this.limit,
          "cardNumber": this.searchCardNumberInput,
          "name": this.searchNameInput
        })
      },
      clearAddForm() {
        this.addForm.name = ""
        this.addForm.info = ""
      },
      clearEditForm() {
        this.editForm.id = 0
        this.editForm.name = ""
        this.editForm.info = ""
      },
      openAddDialog() {
        this.clearAddForm()
        this.addDialogVisible = true
      },
      openEditDialog(id, name, info) {
        this.clearEditForm()
        this.editForm.name = name
        this.editForm.id = id
        this.editForm.info = info
        this.editDialogVisible = true
      },
      submitAddForm() {
        this.$message({
          message: "演示版本，拒绝对数据进行操作",
          type: 'error'
        });
        return
        this.$refs["addForm"].validate(async (valid) => {
          if (valid) {
            let res = await addVip(this.addForm)
            if (res != undefined)
              this.$message({
                message: res.msg,
                type: 'success'
              });
            this.addDialogVisible = false
            this.fetchData({
              "page": this.page,
              "limit": this.limit,
              "cardNumber": this.searchCardNumberInput,
              "name": this.searchNameInput
            })
          }
        })
      },
      submitEditForm() {
        this.$message({
          message: "演示版本，拒绝对数据进行操作",
          type: 'error'
        });
        return
        this.$refs["editForm"].validate(async (valid) => {
          if (valid) {
            let res = await editVip(this.editForm.id, this.editForm)
            if (res != undefined)
              this.$message({
                message: res.msg,
                type: 'success'
              });
            this.editDialogVisible = false
            this.fetchData({
              "page": this.page,
              "limit": this.limit,
              "cardNumber": this.searchCardNumberInput,
              "name": this.searchNameInput
            })
          }
        })
      },
      ...mapActions({
        refreshToken: 'user/refreshToken'
      })
    }
  }
</script>
