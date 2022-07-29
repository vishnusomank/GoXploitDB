FROM golang

LABEL version="1.0"
LABEL maintainer="Vishnu Soman <vishnu@accuknox.com>"

ARG DEBIAN_FRONTEND=noninteractive
WORKDIR /GoXploitDB

RUN apt update && \
    apt install -y git sqlite3

COPY . .

RUN go mod tidy
RUN CGO_CFLAGS="-g -O2 -Wno-return-local-addr" go build -o GoXploitDB main.go

EXPOSE 8080
CMD ["./GoXploitDB"]
