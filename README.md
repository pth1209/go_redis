# Create go web server with redis database

## Start redis database
```
cd redis
docker run -d --rm --name redis-0 \
    --net redis \
    -v ${PWD}/redis-0:/etc/redis/ \
    redis:6.0-alpine redis-server /etc/redis/redis.conf
docker run -d --rm --name redis-1 \
    --net redis \
    -v ${PWD}/redis-1:/etc/redis/ \
    redis:6.0-alpine redis-server /etc/redis/redis.conf
docker run -d --rm --name redis-2 \
    --net redis \
    -v ${PWD}/redis-2:/etc/redis/ \
    redis:6.0-alpine redis-server /etc/redis/redis.conf
```

## Start sentinel (handles automatic failover in redis database)
```
cd redis (if already not in redis directory)
docker run -d --rm --name sentinel-0 --net redis \
    -v ${PWD}/sentinel-0:/etc/redis/ \
    redis:6.0-alpine \
    redis-sentinel /etc/redis/sentinel.conf
docker run -d --rm --name sentinel-1 --net redis \
    -v ${PWD}/sentinel-1:/etc/redis/ \
    redis:6.0-alpine \
    redis-sentinel /etc/redis/sentinel.conf
docker run -d --rm --name sentinel-2 --net redis \
    -v ${PWD}/sentinel-2:/etc/redis/ \
    redis:6.0-alpine \
    redis-sentinel /etc/redis/sentinel.conf
```
