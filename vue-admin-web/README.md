# 前端部署教程

## 开发环境

1. 修改vue.config.js下的反向代理配置

2. `yarn` 或者 `npm i`下载对应的包

3. `yarn run dev` 或者 `npm run dev`启动

## 生产环境

1. `yarn` 或者 `npm i`下载对应的包

2. `yarn run build` 或 `npm run build`打包静态资源

3. 使用nginx或者apache服务器映射静态资源，并配置接口'/admin/api'的反向代理到你的后端服务
