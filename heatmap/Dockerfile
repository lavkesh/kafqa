FROM golang:alpine
RUN apk update
RUN apk add git
RUN mkdir /app 
ADD . /app
WORKDIR /app 
ENV GO111MODULE=on
RUN go build -o heatmap .
CMD ["./heatmap"]