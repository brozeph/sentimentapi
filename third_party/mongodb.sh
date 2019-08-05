export container_name=mongodb
export container_host=127.0.0.1

if [ "docker inspect -f '{{.State.Running}}' $container_name" ]; then
    docker stop $container_name && docker rm $container_name;
fi

docker run \
    --name $container_name \
    --restart unless-stopped \
    -v /data/db:/data/db \
    -v /data/backups:/backups \
    -p $container_host:27017:27017/tcp \
    -d mongo \
    --smallfiles