#!/bin/bash

. linux/includes/deploy_functions.shinc
. linux/includes/common_variables.shinc
. linux/includes/common_functions.shinc

while  read -r service image network subnet iprange ipaddr; do
  ## Creating docker container for Caddy Proxy/Loadbalancer Server
  if [[ $service == *"caddy-server"* ]]; then
    startContainer $service $image $network "0.11.0"
    continue
  fi

  ## Checking/creating docker network
  checkNetwork $network
  if [[ $? -ne 0 ]]; then
    createNetwork $network $subnet $iprange
  else
    dateTime "Network $network already exists"
  fi

  ## Creating docker container for countries-and airports-service
  checkTag $service $image "1.0.1"
  run 2 startContainer $service $image $network "1.0.1"
done <<< $(echo -e "${COUNTRIES_ENV[@]:0:6}\n${AIRPORTS_ENV[@]:0:6}\n${CADDY_ENV[@]:0:6}")
