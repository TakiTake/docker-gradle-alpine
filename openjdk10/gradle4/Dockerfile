FROM adoptopenjdk/openjdk10:alpine
LABEL MAINTAINER="Takeshi Takizawa <takitake.create@gmail.com>"

RUN mkdir /usr/lib/gradle /app
ENV GRADLE_VERSION 4.10.2
ENV GRADLE_HOME /usr/lib/gradle/gradle-${GRADLE_VERSION}
ENV PATH ${PATH}:${GRADLE_HOME}/bin

WORKDIR /usr/lib/gradle
RUN set -x \
  && curl -L -O https://services.gradle.org/distributions/gradle-${GRADLE_VERSION}-bin.zip \
  && unzip gradle-${GRADLE_VERSION}-bin.zip \
  && rm gradle-${GRADLE_VERSION}-bin.zip
