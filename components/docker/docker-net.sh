#!/bin/bash

# 定义网络名称和子网（IP范围）
NETWORK_NAME="im"
SUBNET="10.120.0.0/16"
GATEWAY="10.120.0.1"

echo "开始初始化 Docker网络, name: ${NETWORK_NAME} subnet: ${SUBNET} gateway: ${GATEWAY}"

# 检查网络是否已存在
NETWORK_EXISTS=$(docker network ls --filter "name=${NETWORK_NAME}" -q)

if [ -n "$NETWORK_EXISTS" ]; then
  echo "网络 '${NETWORK_NAME}' 已经存在，跳过创建。"
  exit 0
fi

# 创建 Docker 网络
docker network create --driver bridge --subnet=${SUBNET} --gateway=${GATEWAY} ${NETWORK_NAME}

if [ $? -eq 0 ]; then
  echo "Docker 网络创建成功 name: ${NETWORK_NAME} subnet: ${SUBNET} gateway: ${GATEWAY}"
else
  echo "Docker 网络创建失败 name: ${NETWORK_NAME} subnet: ${SUBNET} gateway: ${GATEWAY}"
fi

