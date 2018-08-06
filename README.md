# Delivering a local-env with a simple command

This project was created with purpose to delivery a local environment running the country- and airport-service.

## Index

> TODO

## Dependencies

* __[Docker](https://docs.docker.com/install/)__ - Using 17.02 or above
* __Linux/Bash__ - Tested in Centos 7.x and Debian Stretch/Jessie
  _Can be implemented in other Operational Systems, like Windows._

## Overview

This project was developed to be as easy as possible. So to start the first environment you can just clone this repo and execute the delilunatech. It will create a Docker network, pull the imagens, start two container for each app (country- and airport-service) and one for Caddy (It's a web server with a lot of features) and will make a request to each app containers.

### What is Two Times

Deploy and Redeploy. It's a simple way to distinguish the first delivery and the upgrade/downgrade of the environment.

### Why Docker

Docker it's a very simple to use/make something, integrated with a lot things and can be useful many environments (e.g.: Delivery an app with zero-downtime)

### Why Golang

This language can be beyond the walls between the dependencies and performance. It's simple to use/learn and it's a very flexible. You can run the same code in anywhere (Or something like that :D) just compiling the code.

For this project I chose the Go to allow a delivery in multi-OS.

### Why Caddy

Caddy is a lightweight web server (Not just it) with a lot of features and very simple to configure/understand.

For this project I chose the Caddy because it's very light, can use a proxy/load balance and cache. Beyond that, we can delivery an app without downtime and we can be a few minutes few no apps running.

### Why Shell and not Go

Just because I have more experience in Bash/Shell scripts than with Go and in a short time would not possible to delivery this project with a good quality. But in the future my intentions it's to improve somethings and rewrite the codes.

## How to use

### Deploy - The first time

```