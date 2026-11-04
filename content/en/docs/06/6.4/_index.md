---
title: "6.4 Use Chainguard Libraries"
weight: 64
sectionnumber: 6.4
---

## Check Chainguard usage

Before we switch to Chainguard Libraries, let's check how many Chainguard libraries are used in the app we built in [lab 6.1](/docs/06/6.1/).

Start the container in the background:

```bash
docker run -d --rm --name my-spring-jdk my-spring-jdk:latest
```

Copy the jar file from the container to the host:

```bash
docker cp my-spring-jdk:/spring-boot-0.0.1-SNAPSHOT.jar ./spring-boot-jdk.jar
```

Verify the usage of Chainguard libraries:

```bash
chainctl libraries verify --detailed --parent=chainguard.dev-partners ./spring-boot-jdk.jar
```

As expected, the libraries are not from Chainguard.

Stop and remove the container:

```bash
docker stop my-spring-jdk
```


## Get Libraries

Now we will also use Chainguard Libraries.

The [CG doc](https://edu.chainguard.dev/chainguard/libraries/access/) states the following to get access to the [Java libraries](https://edu.chainguard.dev/chainguard/libraries/java/overview/):

```bash
chainctl auth pull-token --repository=java --parent=chainguard.dev-partners --ttl=8670h
```

We will use the eval command that will set the needed env vars:
```bash
eval $(chainctl auth pull-token --output env --repository=java --parent=chainguard.dev-partners)
```

Check the Gradle build configuration (`initial/build.gradle`) for dependencies.

Chainguard provides the `spring-boot-starter-web` library:

https://libraries.cgr.dev/java/org/springframework/boot/spring-boot-starter-web/


Now we have to configure Gradle to use Chainguard libraries that are available through the Chainguard repository: `https://libraries.cgr.dev/java/`


Extend the file `build.gradle` to get Chainguard libraries:

```shell
...
repositories {
  maven {
    url = uri("https://libraries.cgr.dev/java/")
    credentials {
      username = providers.environmentVariable("CHAINGUARD_JAVA_IDENTITY_ID").orNull
      password = providers.environmentVariable("CHAINGUARD_JAVA_TOKEN").orNull
    }
  }
  mavenCentral()
}
...
```

Official documentation: https://edu.chainguard.dev/chainguard/libraries/java/build-configuration/#gradle

Build the application again with Docker or build it local.

### Application build with Docker

We will build the application inside a container, like in [lab 6.2](/docs/06/6.2/).

Create the `Containerfile.cg-lib` file inside `gs-spring-boot/initial` with following content:

```Dockerfile
FROM cgr.dev/chainguard.dev-partners/jdk:openjdk-21 AS build
COPY . .
RUN --mount=type=secret,id=CHAINGUARD_JAVA_IDENTITY_ID,env=CHAINGUARD_JAVA_IDENTITY_ID \
    --mount=type=secret,id=CHAINGUARD_JAVA_TOKEN,env=CHAINGUARD_JAVA_TOKEN \
    ./gradlew build

FROM cgr.dev/chainguard.dev-partners/jre:openjdk-21
COPY --from=build /home/build/build/libs/spring-boot-0.0.1-SNAPSHOT.jar /home/build/
EXPOSE 8080
ENTRYPOINT ["java", "-jar", "/home/build/spring-boot-0.0.1-SNAPSHOT.jar"]
```

The credentials for the Chainguard libraries repository are passed as build secrets. They are only available during the `./gradlew build` step and are not stored in the image.

Build the image, the secrets are read from the environment variables set with the `eval` command above:

```bash
docker build -t my-spring-cg-lib -f Containerfile.cg-lib \
  --secret id=CHAINGUARD_JAVA_IDENTITY_ID,env=CHAINGUARD_JAVA_IDENTITY_ID \
  --secret id=CHAINGUARD_JAVA_TOKEN,env=CHAINGUARD_JAVA_TOKEN \
  .
```

Start the container in the background:

```bash
docker run -d --rm --name my-spring-cg-lib -p 8080:8080 my-spring-cg-lib
```

Open your browser at http://localhost:8080/

You should see: `Greetings from Chainguard!`


Copy the jar file from the running container to the host:

```bash
docker cp my-spring-cg-lib:/home/build/spring-boot-0.0.1-SNAPSHOT.jar ./spring-boot-cg-lib.jar
```

Stop and remove the container:

```bash
docker stop my-spring-cg-lib
```

### Local application build

{{% alert title="Note" color="primary" %}}
Only needed when you did not build the application inside the container.
{{% /alert %}}

We will build the application locally on your computer.

Java 21 is needed and a clean Gradle cache.

Clean the Gradle cache:

```bash
rm -rf .gradle/
rm -rf ~/.gradle/caches/
```

Build the app:

```bash
./gradlew clean build
```

Run the app:

```bash
java -jar build/libs/spring-boot-0.0.1-SNAPSHOT.jar
```

Open your browser at http://localhost:8080/

You should see: `Greetings from Chainguard!`

{{% alert title="Note" color="primary" %}}
Stop the app with `ctrl+c`
{{% /alert %}}

Copy your application jar for comparison:

```bash
cp build/libs/spring-boot-0.0.1-SNAPSHOT.jar ./spring-boot-cg-lib.jar
```

## Verify Chainguard usage

Now we can verify the usage of Chainguard libraries inside our Java app:

```bash
chainctl libraries verify --detailed --parent=chainguard.dev-partners ./spring-boot-cg-lib.jar
```

Compare the result with the verification at the beginning of this lab.

Official documentation: https://edu.chainguard.dev/chainguard/libraries/verification/
