FROM golang:1.18-ubuntu

WORKDIR /app

COPY go.mod ./
COPY *.go ./
COPY static ./static

RUN go build -o /belajar-cicd-pemula

EXPOSE 3000

CMD ["/belajar-cicd-pemula"]
