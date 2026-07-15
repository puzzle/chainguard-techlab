---
title: "5.2 Use Chainguard Image"
weight: 52
sectionnumber: 5.2
---

## App - the Chainguard way

Go to the [Chainguard console](https://console.chainguard.dev/) and search for a Java image.

The `Organization` tab shows all images that we can use at Puzzle. Thanks to our Chainguard partnership.
The `Chainguard catalog` tab shows all Chainguard Images. We can ask Chainguard to provide us with images from their catalog such that they are available in our org.

Test Chainguard base image candidate and get infos:

```bash
docker run -ti cgr.dev/puzzle-partner.com/jdk:openjdk-21 /bin/sh

# execute inside the image:

# user
whoami

# location
pwd

# java version
java -version
```

As we see, this image runs as user `java` inside `/home/build` and has java version 21.
As you know, running as root in production is a no-go. OpenShift will also prevent the container from starting.

Change the greetings inside `src/main/java/com/example/springboot/HelloController.java` to `Greetings from Chainguard!`.


Create the `Containerfile.cg` file inside `gs-spring-boot/initial` with following content:

```Dockerfile
FROM cgr.dev/puzzle-partner.com/jdk:openjdk-21 AS build
COPY . .
RUN ./gradlew build

FROM cgr.dev/puzzle-partner.com/jre:openjdk-21
COPY --from=build /home/build/./build/libs/spring-boot-0.0.1-SNAPSHOT.jar /home/build/
EXPOSE 8080
ENTRYPOINT ["java", "-jar", "/home/build/spring-boot-0.0.1-SNAPSHOT.jar"]
```


This is also a multistage build using a Chainguard JDK image to build the app and a Chainguard JRE container to run it.

Build:
```bash
docker build -t my-spring-cg -f Containerfile.cg .
```

Attention: Clean local build and .gradle files if you encounter file permission problems.


Test app inside container:
```bash
docker run -ti --entrypoint /bin/sh my-spring-cg
java -jar ./build/libs/spring-boot-0.0.1-SNAPSHOT.jar
```

You shold get an error.
The images are minimal and hardened. This to keep the attack surface as small as possible.


Run container:
```bash
docker run -p 8080:8080 my-spring-cg
```

Open your browser at http://localhost:8080/

You should see: `Greetings from Chainguard!`

Stop contaier with ctrl+c
