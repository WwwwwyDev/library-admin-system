# 后端部署教程

## 部署相关软件

1. 部署etcd注册中心

2. 部署postgresql

3. 部署redis

## 启动

### 不使用构建脚本

4. 修改service目录下的每个服务的配置文件，并启动每个服务

5. 修改api目录下的配置文件，并启动网关

### 使用构建脚本

4. 修改service目录下的每个服务的配置文件和api目录下的配置文件

5. 执行`sh build.sh`启动服务(保证端口不被占用)

### 数据初始化

启动后，数据库会自动建表，在user服务对应数据库执行sql
```
INSERT INTO "public"."users" VALUES (1, '0001-01-01 08:05:43+08:05:43', '2021-10-18 13:04:39.236154+08', NULL, 'superadmin', '64c5fcb8dffb83eac249afd711921100', 'd986b5b5-0999-4a02-b1d9-bdf2382ddb15', 'http://file.wwywwy.top/FrKwEmE-0z5Ise5BjiINVdAocn6N', '超级管理员，拥有至高无上的权限，密码123456');
INSERT INTO "public"."roles" VALUES (2, NULL, NULL, NULL, 'admin', '管理员');
INSERT INTO "public"."roles" VALUES (3, NULL, NULL, NULL, 'bookadmin', '图书管理员');
INSERT INTO "public"."roles" VALUES (4, NULL, NULL, NULL, 'lendadmin', '借阅管理员');
INSERT INTO "public"."roles" VALUES (5, NULL, NULL, NULL, 'vipadmin', '会员管理员');
INSERT INTO "public"."roles" VALUES (1, NULL, NULL, NULL, 'superadmin', '超级管理员');
INSERT INTO "public"."user_role" VALUES (1, 1);
```

