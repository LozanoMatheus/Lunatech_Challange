FROM java:openjdk-8-jre-alpine
LABEL Goal="Lunatech Labs - Challange" \
        Maintainer="Matheus Lozano Sanches" \
        Package="airports-assembly-1.1.0.jar" \
        GitHub="https://github.com/LozanoMatheus/Lunatech_Challange" \
        DockerHub="https://hub.docker.com/r/lozanomatheus/airports-assembly/"

ADD airports-assembly-1.1.0.jar /

CMD [ "java", "-jar", "airports-assembly-1.1.0.jar" ]
