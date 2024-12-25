
docker pull mysql:8.4
mkdir -p $HOME/mount/mysql8.4/data

docker run --name mysql8.4 \
	-v $HOME/mount/mysql8.4/data:/var/lib/mysql \
	-p 3308:3306 \
	-e MYSQL_ROOT_PASSWORD=123456 \
	-e MYSQL_DATABASE=goexample \
	-e MYSQL_ALLOW_PUBLIC_KEY_RETRIEVAL=true \
	-d mysql:8.4