FROM alpine:3.8
LABEL Goal="Lunatech Labs - Challange" \
        Maintainer="Matheus Lozano Sanches" \
        Package="airports-assembly-1.0.1.jar"

ADD caddy_v0.11.0_linux_amd64.tar.gz /etc/

ENV CADDYPATH=/etc/caddyserver

CMD [ "/etc/caddyserver/caddy", "-http2", "-agree", "-conf", "/etc/caddyserver/Caddyfile" ]