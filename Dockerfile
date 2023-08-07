FROM golang:latest

ARG BUILD_TYPE

ARG EXECUTABLE

WORKDIR /app

RUN apt update && apt -y upgrade

RUN apt install -y make


COPY ./ /app
RUN go env -w GOPROXY="https://goproxy.cn,direct/"

RUN go mod tidy

RUN make build BUILD_TYPE="${BUILD_TYPE}"

ENTRYPOINT [ "./build.dev/${EXECUTABLE}" ]