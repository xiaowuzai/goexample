
# 进入 mysql 容器
docker exec -it mysql8.4 /bin/bash

# 进入 mysql
mysql -p123456

# 创建数据库
create database goexample;

# 授权
GRANT ALL PRIVILEGES ON goexample.* TO 'root'@'%';

# 退出
\q;
exit;