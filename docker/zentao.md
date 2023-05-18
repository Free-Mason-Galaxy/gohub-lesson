
## 禅道安装步骤

[官方安装文档](https://www.zentao.net/book/zentaopmshelp/40.html)

1. [安装 docker](https://docs.docker.com/engine/install/centos/)

2. [安装 docker-compose](https://docs.docker.com/compose/install/linux/)

3. 拉取镜像

   `docker pull easysoft/zentao:latest`

4. 创建网络

    `docker network create --subnet=172.172.172.0/24 zentaonet`

5. 创建容器命令

   `docker run --name=zentao -p 80:80 -p 3306:3306 -p 6379:6379 --network=zentaonet --ip 172.172.172.2 --mac-address 02:42:ac:11:00:01 -v /www/zentao/pms:/www/zentaopms -v /www/zentao/mysqldata:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=123456 -d easysoft/zentao`