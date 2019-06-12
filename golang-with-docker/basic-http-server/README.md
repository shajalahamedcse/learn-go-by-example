## A Simple Golang Server Deployed Using Docker 

There are many guides out theere on how to deploy a golang server using `Docker`, but the aim of this post is to provide a basic starting point and use a simple `Dockerfile` . It’s also aimed at getting it running locally for development purposes quickly. But this not production ready.


Example Go HTTP Server: main.go



    package main

    import (
        "io"
        "log"
        "net/http"
        "os"
    )

    func main() {

        http.HandleFunc("/", healthCheck)

        port := os.Getenv("PORT")
        if port == "" {
            port = "8080"
        }

        log.Println(" Service started on port: " + port)
        if err := http.ListenAndServe(":"+port, nil); err != nil {
            log.Fatal(err)
        }

    }

    func healthCheck(w http.ResponseWriter, r *http.Request) {
        w.Header().Add("Content-Type", "application/json")
        io.WriteString(w, `{"status":"ok"}`)
    }


#### Dockerfile

    FROM golang:1.11.5

    ENV APP_NAME server
    ENV PORT 8080

    COPY . /go/src/${APP_NAME}
    WORKDIR /go/src/${APP_NAME}

    RUN go get ./
    RUN go build -o ${APP_NAME}

    CMD ./${APP_NAME}

    EXPOSE ${PORT}

#### docker-compose.yml

    version: '3'
    services:
    web:
        build: .
        ports:
        - "8080:8080"


Once we’re all setup is done, we can use `docker-compose` to get this server up and running. Note that this implementation doesn’t have any form of live reloading system, so any changes made on the codebase we'll need to stop the current container and rebuild on up with the command below.

    docker-compose up --build
