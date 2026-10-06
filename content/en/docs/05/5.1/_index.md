---
title: "5.1 Example Application"
weight: 51
sectionnumber: 5.1
---

## App - the standard way

Using example [Spring Boot app](https://github.com/spring-guides/gs-spring-boot/tree/main/initial)

```bash
git clone https://github.com/spring-guides/gs-spring-boot.git
cd gs-spring-boot/initial
```

Create the `.dockerignore` file inside `gs-spring-boot/initial` with following content:

```
Containerfile*
```

This should make builds faster. Changes in Containerfiles have no impact on build context and caching.


Test base image candidate and get infos:
```bash
docker run -ti docker.io/eclipse-temurin:21-jdk-ubi9-minimal /bin/bash

# execute inside the image:

# user
whoami

# location
pwd

# java version
java -version
```

As we see, this image runs as user `root` inside root `/` and has java version 21.


Create the `Containerfile.jdk` file inside `gs-spring-boot/initial` with following content:

```Dockerfile
FROM docker.io/eclipse-temurin:21-jdk-ubi9-minimal AS build
# user: root!
COPY . .
RUN ./gradlew build

FROM docker.io/eclipse-temurin:21-jre-ubi9-minimal
COPY --from=build /build/libs/spring-boot-0.0.1-SNAPSHOT.jar /
EXPOSE 8080
ENTRYPOINT ["java", "-jar", "/spring-boot-0.0.1-SNAPSHOT.jar"]
```

This is a multistage build using a JDK to build the app and a JRE container to run it.


Build image (includes app build):
```bash
docker build -t my-spring-jdk -f Containerfile.jdk .
```

Attention: Clean local build and .gradle files if you encounter file permission problems.


Test app inside container:
```bash
docker run -ti --entrypoint /bin/bash my-spring-jdk
java -jar /spring-boot-0.0.1-SNAPSHOT.jar
```

Exit with ctrl+d


Run container:
```bash
docker run -p 8080:8080 my-spring-jdk
```

Open your browser at http://localhost:8080/

You should see: `Greetings from Spring Boot!`

Stop container with ctrl+c
