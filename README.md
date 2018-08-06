# Delivering a local-env with a simple command

This project was created with purpose to delivery a local environment running the country- and airport-service.

---

## Index

* [Dependencies](https://github.com/LozanoMatheus/Lunatech_Challange#dependencies) - Dependencies to use/execute this project.
* [Overview](https://github.com/LozanoMatheus/Lunatech_Challange#overview) - Overview and explaining the why's.
  * [What is Two Times](https://github.com/LozanoMatheus/Lunatech_Challange#what-is-two-times) - Explaining why I use two script, one to deploy and other to redeploy (Upgrade/Downgrade).
  * [Why Docker](https://github.com/LozanoMatheus/Lunatech_Challange#why-docker) - Explaining why I use Docker.
  * [Why Golang](https://github.com/LozanoMatheus/Lunatech_Challange#why-golang) - Explaining why I use Golang as the main script.
  * [Why Caddy](https://github.com/LozanoMatheus/Lunatech_Challange#why-caddy) - Explaining why I use Caddy as a Web Server/*.
  * [Why Shell and not Go](https://github.com/LozanoMatheus/Lunatech_Challange#why-shell-and-not-go) - Explaining why I use Shell script and not just the Golang.
* [How to use](https://github.com/LozanoMatheus/Lunatech_Challange#how-to-use) - This section will show how to use this project.
  * [Deploy - The first time](https://github.com/LozanoMatheus/Lunatech_Challange#deploy---the-first-time) - Showing how to execute when is the first time.
  * [Upgrade/Downgrade](https://github.com/LozanoMatheus/Lunatech_Challange#upgradedowngrade) - Executing the upgrade or downgrade an app in running state with zero-downtime.
  * [An example](https://github.com/LozanoMatheus/Lunatech_Challange#an-example) - An pratical exemple for deploying and redeploying process.
* [Technical Assignment](https://github.com/LozanoMatheus/Lunatech_Challange#technical-assignment) - The technical assignment for Lunatech Labs

---

## Dependencies

* __[Docker](https://docs.docker.com/install/)__ - Using 17.02 or above
* __Linux/Bash__ - Tested in Centos 7.x and Debian Stretch/Jessie
  _Can be implemented in other Operational Systems, like Windows._

---

## Overview

This project was developed to be as easy as possible. So to start the first environment you can just clone this repo and execute the delivery_lunatech. It will create a Docker network, pull the imagens, start two container for each app (country- and airport-service) and one for Caddy (It's a web server with a lot of features) and will make a request to each app containers.

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

---

## How to use

### Deploy - The first time

Just execute the delivery_lunatech Go script. It will create the entire stack (Docker network and docker container) and will start the apps.

```bash
./delivery_lunatech
```

### Upgrade/Downgrade

Simple as the deployment, you can execute the delivery_lunatech Go script and pass the app name and version/tag do you wish.

```bash
./delivery_lunatech airports-assembly 1.1.0
```

### An example

This video will show to us how to use it.

[![asciicast](https://asciinema.org/a/WJe9Hcpvm4jtQrIUJiw5bHtiO.png)](https://asciinema.org/a/WJe9Hcpvm4jtQrIUJiw5bHtiO)

---

- [x] Entire stack able to run in a local machine;
- [x] Apps isolated from each other;
  - [x] Deny inter-communication between the applications; and
  - [x] Deny external requests;
- [x] An Proxy/Loadbalancer expose on port 8000;
- [x] The airports-service in version 1.0.1 must be the first deployment; and
- [x] Upgrade/Downgrade cannot be causing a service interruption.
- Bonus points:
  - [ ] Use a CI/CD tool; and
  - [x] Single click/command to deploy the entire stack.

_Note: I do not use the CI / CD because it will create one more dependency/complex and I tried to focus on simplicity and velocity to deploy. But I used the Blue/Green Deployment process, it's a DevOps culture._