<template>
  <div class="dashboard-container">
    <div class="dashboard-text">name: {{ name }}</div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import request from '@/utils/request'

export default {
  name: 'Dashboard',
  created() {
    request({
      url: 'http://127.0.0.1:8888/book',
      method: 'GET',
      params: {
        page:1,
        limit:10
      }
    }).then(response => {
      console.log(response)
    })
  },
  computed: {
    ...mapGetters({
      name: "name"
    })
  }
}
</script>

<style lang="scss" scoped>
.dashboard {
  &-container {
    margin: 30px;
  }
  &-text {
    font-size: 30px;
    line-height: 46px;
  }
}
</style>
