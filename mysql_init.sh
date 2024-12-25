
docker pull mysql:8.4
mkdir -p $HOME/mount/mysql8.4/data

docker run --name mysql8.4 \
	-v $HOME/mount/mysql8.4/data:/var/lib/mysql \
	-p 3308:3306 \
	-e MYSQL_ROOT_PASSWORD=123456 \
	-e MYSQL_ALLOW_PUBLIC_KEY_RETRIEVAL=true \
	-d mysql:8.4

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