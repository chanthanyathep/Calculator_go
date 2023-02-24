FROM golang:1.20

WORKDIR /app

COPY . /app

#RUN go mod tidy
RUN go build -o test main.go  

EXPOSE 8080

RUN ls
CMD ["./test"]