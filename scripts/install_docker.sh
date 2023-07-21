#!/bin/bash

# Update container Linux image APT repository data
yum update

amazon-linux-extras install docker
service docker start
usermod -a -G docker ec2-user
chmod 777 /var/run/docker.sock
chkconfig docker on

curl -L https://github.com/docker/compose/releases/download/1.22.0/docker-compose-$(uname -s)-$(uname -m) -o /usr/local/bin/docker-compose

curl -L https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m) -o /usr/local/bin/docker-compose

chmod +x /usr/local/bin/docker-compose

