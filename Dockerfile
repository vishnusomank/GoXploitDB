FROM golang
ARG DEBIAN_FRONTEND=noninteractive
WORKDIR /GoXploitDB
RUN apt update && \
    apt install -y git sqlite3
COPY . .
RUN go mod tidy
RUN go build -o GoXploitDB main.go
EXPOSE 8080
CMD ["./GoXploitDB"]
