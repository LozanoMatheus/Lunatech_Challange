FROM java:openjdk-8-jre-alpine
LABEL Goal="Lunatech Labs - Challange" \
        Maintainer="Matheus Lozano Sanches" \
        Package="countries-assembly-1.0.1.jar" \
        GitHub="https://github.com/LozanoMatheus/Lunatech_Challange" \
        DockerHub="https://hub.docker.com/r/lozanomatheus/countries-assembly/"

ADD countries-assembly-1.0.1.jar /

CMD [ "java", "-jar", "countries-assembly-1.0.1.jar" ]