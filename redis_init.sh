docker run --name redis5.0 -d -p 6379:6379 \
-v $HOME/mount/redis/redis.conf:/usr/local/etc/redis/redis.conf \
redis:5.0 \
redis-server /usr/local/etc/redis/redis.conf

