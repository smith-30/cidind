FROM docker:dind

RUN apk add --no-cache py-pip openssl bash
RUN pip install docker-compose

WORKDIR /go/src/github.com/smith-30/cidind


ENTRYPOINT ["/bin/bash", "-c"]