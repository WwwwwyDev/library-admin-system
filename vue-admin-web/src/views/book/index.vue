<template>
  <div class="app-container">
    <el-container>
      <el-header>
        <el-row>
          <el-col :span="4">
            <el-input placeholder="请输入查找的书名" v-model.lazy="searchNameInput" clearable>
            </el-input>
          </el-col>
          <el-col :span="4" style="padding-left: 10px;">
            <el-input placeholder="请输入查找的作者名" v-model.lazy="searchAuthorInput" clearable>
            </el-input>
          </el-col>
          <el-col :span="1" style="padding-left: 10px;">
            <el-button type="primary" icon="el-icon-search" @click="search">搜索</el-button>
          </el-col>
          <el-col :span="1" :offset="3">
            <el-button type="primary" @click="openAddDialog" round>添加用户</el-button>
          </el-col>
        </el-row>
      </el-header>
      <el-main>
        <el-table v-loading="listLoading" :data="list" element-loading-text="Loading" border fit highlight-current-row>
          <el-table-column align="center" label="ID" width="95">
            <template slot-scope="scope">
              {{ scope.row.id }}
            </template>
          </el-table-column>
          <el-table-column label="书名" width="100" align="center">
            <template slot-scope="scope">
              {{ scope.row.name }}
            </template>
          </el-table-column>
          <el-table-column label="图片" width="150">
            <template slot-scope="scope">
              <el-image :src="scope.row.image">
                <div slot="error" class="image-slot">
                  <i class="el-icon-picture-outline"></i>
                </div>
              </el-image>
            </template>
          </el-table-column>
          <el-table-column label="作者" width="150">
            <template slot-scope="scope">
              <span>{{ scope.row.author }}</span>
            </template>
          </el-table-column>
          <el-table-column label="类目" width="150">
            <template slot-scope="scope">
              <span>{{ scope.row.type }}</span>
            </template>
          </el-table-column>
          <el-table-column label="信息">
            <template slot-scope="scope">
              <span>{{ scope.row.info }}</span>
            </template>
          </el-table-column>
          <el-table-column label="操作" width="150" fixed="right">
            <template slot-scope="scope">
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
          :page-sizes="[5, 10, 20, 100]" :page-size="limit" layout="total, sizes, prev, pager, next, jumper"
          :total="total">
        </el-pagination>
      </el-footer>
    </el-container>

    <el-dialog title="添加用户" :visible.sync="addDialogVisible" width="500px">
      <el-form :model="addForm" :rules="FormRules" ref="addForm" label-width="90px" @submit.native.prevent>
        <el-form-item label="图书名" prop="name">
          <el-input placeholder="请输入书名" v-model="addForm.name" clearable>
          </el-input>
        </el-form-item>
        <el-form-item label="作者">
          <el-input placeholder="请输入作者" v-model="addForm.author" clearable>
          </el-input>
        </el-form-item>
        <el-form-item label="信息">
          <el-input placeholder="请输入信息" v-model="addForm.info" clearable>
          </el-input>
        </el-form-item>
        <el-form-item label="类目" prop="type">
          <el-select v-model="addForm.type" filterable placeholder="请选择">
            <el-option v-for="item in options" :key="item.id" :label="item.name" :value="item.id">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="图片">
          <el-upload class="image-uploader" action="" :show-file-list="false" :auto-upload="false"
            :on-change="changePhotoFile">
            <img v-if="addForm.image" :src="addForm.image" class="image">
            <i v-else class="el-icon-plus image-uploader-icon"></i>
          </el-upload>
          <imageCropper ref="imageCropper" @getFile="getAddFile"></imageCropper>
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
    getBooks
  } from '@/api/booktable'
  import {
    getAllTypes
  } from '@/api/search'
  import imageCropper from '@/components/ImageCropper'
  import {
    getQiniuToken
  } from '@/api/qiniu'
  import {
    domain,
    uploadAsync
  } from '@/utils/qiniu'
  export default {
    components: {
      imageCropper
    },
    data() {
      return {
        imageFile: null,
        searchAuthorInput: '',
        searchNameInput: '',
        list: null,
        listLoading: true,
        page: 1,
        limit: 10,
        total: 0,
        options: [],
        addDialogVisible: false,
        addForm: {
          name: "",
          image: "",
          author: "",
          info: "",
          type: null,
          typeId: 0,
        },
        editDialogVisible: false,
        editForm: {
          id: 0,
          name: "",
          image: "",
          author: "",
          info: "",
          typeId: 0,
        },
        FormRules: {
          name: [{
            required: true,
            message: "请输入书名",
            trigger: "blur"
          }],
          type: [{
            required: true,
            message: "请选择图书类目",
            trigger: 'change'
          }]
        },
      }
    },
    created() {
      this.listLoading = true
      this.fetchOption()
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
        let res = await getBooks(params)
        let {
          books,
          total
        } = res.data
        this.list = books
        this.total = total
        this.listLoading = false
      },
      async fetchOption() {
        let res = await getAllTypes()
        this.options = res.data.types
      },
      handleSizeChange(limit) {
        this.fetchData({
          "page": this.page,
          "limit": limit,
          "name": this.searchNameInput,
          "author": this.searchAuthorInput
        })
        this.limit = limit
      },
      handleCurrentChange(page) {
        this.fetchData({
          "page": page,
          "limit": this.limit,
          "name": this.searchNameInput,
          "author": this.searchAuthorInput
        })
        this.page = page
      },
      search() {
        this.fetchData({
          "page": 1,
          "limit": this.limit,
          "name": this.searchNameInput,
          "author": this.searchAuthorInput
        })
        this.page = 1
      },
      openAddDialog() {
        this.clearAddForm()
        this.addDialogVisible = true
      },
      clearAddForm() {
        this.imageFile = null
        this.addForm.author = ""
        this.addForm.image = ""
        this.addForm.info = ""
        this.addForm.name = ""
        this.addForm.typeId = 0
        this.addForm.type = null
      },
      submitAddForm() {

      },
      openEditDialog() {

      },
      clearEditForm() {

      },
      submitEditForm() {

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
        this.imageFile = file
        this.addForm.image = window.URL.createObjectURL(file)
        this.$refs.imageCropper.close();
      },
      ...mapActions({
        refreshToken: 'user/refreshToken'
      })
    }
  }
</script>
<style>
  .image-uploader .el-upload {
    border: 1px dashed #d9d9d9;
    border-radius: 6px;
    cursor: pointer;
    position: relative;
    overflow: hidden;
  }

  .image-uploader .el-upload:hover {
    border-color: #409EFF;
  }

  .image-uploader-icon {
    font-size: 28px;
    color: #8c939d;
    width: 178px;
    height: 178px;
    line-height: 178px;
    text-align: center;
  }

  .image {
    width: 178px;
    height: 178px;
    display: block;
  }
</style>
