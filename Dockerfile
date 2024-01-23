FROM golang:latest

RUN mkdir -p /app

COPY . /app

WORKDIR /app

RUN chmod +x launch.sh

ENTRYPOINT ["./launch.sh"]