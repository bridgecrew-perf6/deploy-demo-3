FROM golang:1 AS builder
COPY . /deploy-demo
WORKDIR /deploy-demo
RUN CGO_ENABLED=0 go build && ls -al /deploy-demo

FROM alpine:3.14
COPY --from=builder /deploy-demo/deploy-demo /usr/bin/deploy-demo
RUN chmod +x /usr/bin/deploy-demo && ls -al /usr/bin/deploy-demo

#ENTRYPOINT ["/bin/sh", "-c", "/usr/bin/deploy-demo"]
CMD /usr/bin/deploy-demo
