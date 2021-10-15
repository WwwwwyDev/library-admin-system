<template>
  <div class="app-container">
    <el-container>
      <el-container>
        <el-aside width="500px">
          <el-card style="background-color: #F2F6FC;">
            <el-row :gutter="20">
              <el-col :span="6" :offset="8">
                <el-avatar :size="100" :src="avatar"></el-avatar>
              </el-col>
            </el-row>
            <el-row :gutter="20">
              <el-divider></el-divider>
            </el-row>
            <el-row :gutter="20">
              <el-form label-width="80px">
                <el-form-item label="用户名">
                  {{name}}
                </el-form-item>
                <el-form-item label="个人信息">
                  <el-input type="textarea" v-model="info" :disabled="true"></el-input>
                </el-form-item>
                <el-form-item label="拥有权限">
                  <el-tag v-for="item in rolesList" :key='item.index'>
                    {{ item }}
                  </el-tag>
                </el-form-item>
              </el-form>
            </el-row>
            <el-row :gutter="20">
              <el-col :span="6" :offset="4">
                <el-button type="primary">编辑信息</el-button>
              </el-col>
              <el-col :span="6" :offset="4">
                <el-button type="primary">修改密码</el-button>
              </el-col>
            </el-row>
          </el-card>
        </el-aside>
      </el-container>
    </el-container>
  </div>
</template>

<script>
  import {
    mapGetters
  } from 'vuex'

  export default {
    data() {
      return {
        rolesList:[]
      }
    },
    created() {
      this.rolesList = this.roles == ""?["无权限"]:this.roles.split(",")
    },
    computed: {
      ...mapGetters({
        id: "user/id",
        name: "user/name",
        avatar: "user/avatar",
        info: "user/info",
        roles: "user/roles"
      })
    }
  }
</script>

<style>

</style>
