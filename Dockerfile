# Build/test environment for gospn (Go + ANTLR4)
FROM golang:1.22-bookworm

ARG ANTLR_VERSION=4.7.2

RUN apt-get update \
 && apt-get install -y --no-install-recommends default-jre-headless curl make zip \
 && rm -rf /var/lib/apt/lists/* \
 && curl -fsSL -o /usr/local/lib/antlr.jar \
      https://www.antlr.org/download/antlr-${ANTLR_VERSION}-complete.jar \
 && printf '#!/bin/sh\nexec java -jar /usr/local/lib/antlr.jar "$@"\n' > /usr/local/bin/antlr4 \
 && chmod +x /usr/local/bin/antlr4

WORKDIR /src
CMD ["make", "test"]
