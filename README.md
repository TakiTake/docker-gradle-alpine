# docker-gradle-alpine
Dockerfile for Gradle

DIRECTORY STRUCTURE WAS CHANGED.

If you want to use same version as previous use `openjdk8-gradle2` tag please.

## Build

```sh
$ cd /path/to/gradle/project

$ docker run \
  -v $PWD:/app \
  -w /app \
  -u $UID:$GID \
  takitake/gradle-alpine \
  gradle wrapper && ./gradlew clean build
```
