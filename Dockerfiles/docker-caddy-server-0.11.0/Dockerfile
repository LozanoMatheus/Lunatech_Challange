FROM alpine:3.8
LABEL Goal="Lunatech Labs - Challange" \
        Maintainer="Matheus Lozano Sanches" \
        Package="caddy_v0.11.0_linux_amd64.tar.gz" \
        GitHub="https://github.com/LozanoMatheus/Lunatech_Challange" \
        DockerHub="https://hub.docker.com/r/lozanomatheus/caddy-server/"

ADD caddy_v0.11.0_linux_amd64.tar.gz /etc/

ENV CADDYPATH=/etc/caddyserver

CMD [ "/etc/caddyserver/caddy", "-http2", "-agree", "-conf", "/etc/caddyserver/Caddyfile" ]