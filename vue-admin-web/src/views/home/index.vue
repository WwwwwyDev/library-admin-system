<template>
  <div class="dashboard-container">
    <el-row type="flex"  justify="center":gutter="20">
      <el-col :span="10">
        <el-card style="background-color: #fcfcfc;">
          <div ref="bookChart"></div>
        </el-card>
      </el-col>
      <el-col :span="10">
        <el-card style="background-color: #fcfcfc;">
          <div ref="lendChart"></div>
        </el-card>
      </el-col>
      <el-col :span="10">
        <el-card style="background-color: #fcfcfc;">
          <div ref="systemChart"></div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
  import {
    mapGetters
  } from 'vuex'
  import {
    getBooks
  } from '@/api/booktable'
  import {
    getLends
  } from '@/api/lendtable'
  import {
    getAllTypes,
    getTypeById
  } from '@/api/search'
  import {
    getLoginStatus,
    getUsers
  } from '@/api/usertable'
  // import * as echarts from 'echarts';
  export default {
    name: 'home',
    data() {
      return {
        bookOptions: {
          title: {
            text: '图书概括',
            subtext: 'Book Summary',
            left: 'center'
          },
          tooltip: {
            trigger: 'item'
          },
          legend: {
            orient: 'vertical',
            left: 'left'
          },
          series: [{
            name: '图书概括',
            type: 'pie',
            radius: '50%',
            data: [
            ],
            emphasis: {
              itemStyle: {
                shadowBlur: 10,
                shadowOffsetX: 0,
                shadowColor: 'rgba(0, 0, 0, 0.5)'
              }
            }
          }]
        },
        lendOptions: {
          title: {
            text: '借阅概括',
            subtext: 'Lend Summary',
            left: 'center'
          },
          tooltip: {
            trigger: 'item'
          },
          legend: {
            orient: 'vertical',
            left: 'left'
          },
          series: [{
            name: '借阅概括',
            type: 'pie',
            radius: '50%',
            data: [{
                value: 0,
                name: '未借阅'
              },
              {
                value: 0,
                name: '已借阅'
              }
            ],
            emphasis: {
              itemStyle: {
                shadowBlur: 10,
                shadowOffsetX: 0,
                shadowColor: 'rgba(0, 0, 0, 0.5)'
              }
            }
          }]
        },
        systemOptions:{
          title: {
            text: '系统概括',
            subtext: 'Lend Summary',
            left: 'center'
          },
          tooltip: {
            trigger: 'item'
          },
          legend: {
            orient: 'vertical',
            left: 'left'
          },
          series: [{
            name: '系统概括',
            type: 'pie',
            radius: '50%',
            data: [{
                value: 0,
                name: '在线人数'
              },
              {
                value: 0,
                name: '离线人数'
              }
            ],
            emphasis: {
              itemStyle: {
                shadowBlur: 10,
                shadowOffsetX: 0,
                shadowColor: 'rgba(0, 0, 0, 0.5)'
              }
            }
          }]
        }
      }
    },
    mounted() {
      this.drawBookChart()
      this.drawLendChart()
      this.drawSystemChart()
    },
    methods: {
      async drawBookChart(){
        let bookChart = echarts.init(this.$refs.bookChart);
        bookChart.resize({
          width: 500,
          height: 400
        });
        bookChart.showLoading()
        let res = await getAllTypes()
        let bookTypes = res.data.types
        for(let i = 0; i < bookTypes.length; i++){
          let temp = await getTypeById(bookTypes[i].id)
          if (temp.data.type.booksInfo !== undefined)
            bookTypes[i]["value"] = temp.data.type.booksInfo.length
          else bookTypes[i]["value"] = 0
        }
        this.bookOptions.series[0].data = bookTypes
        bookChart.setOption(this.bookOptions)
        bookChart.hideLoading();
      },
      async drawLendChart() {
        let lendChart = echarts.init(this.$refs.lendChart);
        lendChart.resize({
          width: 500,
          height: 400
        });
        lendChart.showLoading()
        let res = await getBooks()
        let booktotal = res.data.total
        res = await getLends()
        this.lendOptions.series[0].data[1].value = res.data.total
        this.lendOptions.series[0].data[0].value = booktotal - res.data.total
        lendChart.setOption(this.lendOptions)
        lendChart.hideLoading();
      },
      async drawSystemChart() {
        let systemChart = echarts.init(this.$refs.systemChart);
        systemChart.resize({
          width: 500,
          height: 400
        });
        systemChart.showLoading()
        let res = await getUsers()
        let usertotal = res.data.total
        res = await getLoginStatus()
        this.systemOptions.series[0].data[0].value = res.data.loginStatus.length
        this.systemOptions.series[0].data[1].value = usertotal - res.data.loginStatus.length
        systemChart.setOption(this.systemOptions)
        systemChart.hideLoading();
      }
    },
    computed: {
      ...mapGetters({
        name: "user/name"
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
