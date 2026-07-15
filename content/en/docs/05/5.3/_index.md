---
title: "5.3 Use Chainguard Libraries"
weight: 53
sectionnumber: 5.3
---

## Get Libraries

Now we will also use Chainguard Libraries.

The [CG doc](https://edu.chainguard.dev/chainguard/libraries/access/) states following to get access to the [Java libraries](https://edu.chainguard.dev/chainguard/libraries/java/overview/):

```bash
chainctl auth pull-token --repository=java --parent=puzzle-partner.com --ttl=8670h
```

We will use the eval command that will set the needed env vars:
```bash
eval $(chainctl auth pull-token --output env --repository=java --parent=puzzle-partner.com)
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
   username = "CHAINGUARD_JAVA_IDENTITY_ID"
   password = "CHAINGUARD_JAVA_TOKEN"
  }
 }
 mavenCentral()
}
...
```

Official documentation: https://edu.chainguard.dev/chainguard/libraries/java/build-configuration/#gradle


We will build the application outside of the container.
Java 21 is needed and a clean Gradle cache.

Clean the Gradle cache:

```bash
rm -rf .gradle/caches/
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

Stop the app with ctrl+c


Now we can verify the usage of Chainguard libraries inside our Java app:

```bash
chainctl libraries verify --detailed --parent=puzzle-partner.com ./build/libs/spring-boot-0.0.1-SNAPSHOT.jar
```

Official documentation: https://edu.chainguard.dev/chainguard/libraries/verification/
