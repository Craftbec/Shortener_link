# Shortener_link

This project implements a service providing an API for creating unique shortened links.

## Run

Running with in memory storage
```
make in-memory
```
Running with postgres storage
```
make postgres
```

## Usage

GET - returns the original URL
```
curl -X GET http://localhost:8080/{ShortURL}
```
POST - create short URL
```
curl -X POST http://localhost:8080/?url={OriginalURL}
```
