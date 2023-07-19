FROM       golang:1.21-rc-alpine3.18 as builder

ENV        DIR=/scaffold
ADD .      ${DIR}
WORKDIR    ${DIR}

RUN        apk add --update make gcc g++ git sqlite
RUN        make build



FROM       golang:1.21-rc-alpine3.18

RUN        apk add --update make gcc g++ git sqlite
COPY       --from=builder /scaffold/scaffold /usr/local/bin/scaffold

ENTRYPOINT ["scaffold"]
CMD        ["--help"]

