#!/bin/bash

cd /home/ec2-user/app

docker-compose -f docker-compose.yml build
docker-compose -f docker-compose.yml up -d
# docker image rm -f $(docker images -f dangling=true -q)