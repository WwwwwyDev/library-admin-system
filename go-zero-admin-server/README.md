# build step

- 部署etcd集群

- 部署postgresql

- 修改每个服务的配置文件

- 启动每个服务

注：启动后，数据库会自动建表，在user服务对应数据库执行这句插入sql
```
INSERT INTO "public"."users" VALUES (1, '0001-01-01 08:05:43+08:05:43', '2021-10-18 13:04:39.236154+08', NULL, 'superadmin', '64c5fcb8dffb83eac249afd711921100', 'd986b5b5-0999-4a02-b1d9-bdf2382ddb15', 'http://file.wwywwy.top/FrKwEmE-0z5Ise5BjiINVdAocn6N', '超级管理员，拥有至高无上的权限，密码123456');
```

- 修改api配置文件

- 启动api网关

