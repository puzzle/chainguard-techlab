---
title: "6.2 Use Chainguard Image"
weight: 62
sectionnumber: 6.2
---

## App - the Chainguard way

Go to the [Chainguard console](https://console.chainguard.dev/) and click on the `Images` menu:

The `Organization` tab shows all images that we can use in the lab organization `chainguard.dev-partners`.

The `Chainguard catalog` tab shows all Chainguard Images.

Search for a JDK (Java) image.


Test the found Chainguard base image candidate (`jdk`) and get infos:

```bash
docker run -ti cgr.dev/chainguard.dev-partners/jdk:openjdk-17 /bin/sh

# execute inside the image:

# user
whoami

# location
pwd

# java version
java -version
```

As we see, this image runs as user `java` inside `/home/build` and has java version 17.
As you know, running as root in production is a no-go. OpenShift will also prevent the container from starting.

{{% alert title="Note" color="primary" %}}
Exit the container with `ctrl+d`
{{% /alert %}}

Change the greetings inside `src/main/java/com/example/springboot/HelloController.java` to `Greetings from Chainguard!`.


Create the `Containerfile.cg` file inside `gs-spring-boot/initial` with following content:

```Dockerfile
FROM cgr.dev/chainguard.dev-partners/jdk:openjdk-21 AS build
COPY . .
RUN ./gradlew build

FROM cgr.dev/chainguard.dev-partners/jre:openjdk-21
COPY --from=build /home/build/./build/libs/spring-boot-0.0.1-SNAPSHOT.jar /home/build/
EXPOSE 8080
ENTRYPOINT ["java", "-jar", "/home/build/spring-boot-0.0.1-SNAPSHOT.jar"]
```


This is also a multistage build using a Chainguard JDK image to build the app and a Chainguard JRE container to run it.

Build:
```bash
docker build -t my-spring-cg -f Containerfile.cg .
```

{{% alert title="Note" color="primary" %}}
Attention: Clean local build and .gradle files if you encounter file permission problems.
{{% /alert %}}


Test app inside container:
```bash
docker run -ti --entrypoint /bin/sh my-spring-cg
java -jar ./build/libs/spring-boot-0.0.1-SNAPSHOT.jar
```

You should get an error.
The images are minimal and hardened, to keep the attack surface as small as possible.


Run container:
```bash
docker run -p 8080:8080 my-spring-cg
```

Open your browser at http://localhost:8080/

You should see: `Greetings from Chainguard!`

{{% alert title="Note" color="primary" %}}
Stop the container with `ctrl+c`
{{% /alert %}}
