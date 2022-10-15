# Image in docker hub
docker pull bryan335/http_server:v1.0

docker run -p30080:80 bryan335/http_server:v1.0


root        7513    7492  0 13:17 ?        00:00:00 /app/http_server
cnadmin     8858    1695  0 13:19 pts/0    00:00:00 grep --color=auto http
root@cloudnative01:# nsenter -t 7513 -n ip addr
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
12: eth0@if13: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default
    link/ether 02:42:ac:11:00:02 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 172.17.0.2/16 brd 172.17.255.255 scope global eth0
       valid_lft forever preferred_lft forever
