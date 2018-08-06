#!/bin/bash +x

APP=$1
VERSION=$2

. linux/includes/redeploy_functions.shinc
. linux/includes/common_variables.shinc
. linux/includes/common_functions.shinc

while  read -r service image network subnet iprange ipaddr; do

  if [[ $APP == *"$service"* ]] && [[ "$service" != *"caddy-service"* ]] ; then
    checkTag $service $image $VERSION
    run 2 blueGreen $service $image $network $VERSION
  fi

done <<< $(echo -e "${COUNTRIES_ENV[@]:0:6}\n${AIRPORTS_ENV[@]:0:6}\n${CADDY_ENV[@]:0:6}")
