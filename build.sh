# Oldest
docker build -t takitake/gradle-alpine:openjdk8-gradle2 openjdk8/gradle2

docker run \
  -v $PWD/test:/app \
  -w /app \
  -u $UID:$GID \
  takitake/gradle-alpine:openjdk8-gradle2 \
  gradle clean build

# Newest
docker build -t takitake/gradle-alpine:openjdk10-gradle4 openjdk8/gradle4

docker run \
  -v $PWD/test:/app \
  -w /app \
  -u $UID:$GID \
  takitake/gradle-alpine:openjdk10-gradle4 \
  gradle clean build
