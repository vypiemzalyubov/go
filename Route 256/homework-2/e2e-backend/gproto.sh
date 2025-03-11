#!/bin/bash

INPUT=$1

echo "${INPUT}"

LOCAL_PROTO_DIR="./proto"

GITLAB_TOKEN="$GITLAB_TOKEN"

GITLAB_PROJECT="$GITLAB_PROJECT"

PROTO_PATH=$(echo "${INPUT}" | sed 's/[/]/%2F/g')

array=($(echo $INPUT | tr "/" "\n"))
array_size=${#array[@]}
last_index=$(( array_size-1))

PROTO_NAME=${array[$last_index]}

echo $LOCAL_PROTO_DIR
echo $PROTO_PATH
echo $PROTO_NAME


curl --header "PRIVATE-TOKEN: $GITLAB_TOKEN" \
  --url "https://gitlab.ozon.ru/api/v4/projects/${GITLAB_PROJECT}/repository/files/${PROTO_PATH}/raw?ref=master" > ./${LOCAL_PROTO_DIR}/${PROTO_NAME}

