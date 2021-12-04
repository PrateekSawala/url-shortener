## Overview

The project aims to provide a simple URL shortener service that will accept a URL as an argument over a REST API and return a shortened URL as a result.

### Description

- The API processes the input URL as an argument over a REST API, write it to a JSON file, and returns a shortened URL as a result.
- The case where the API request gets made with a long URL that has already got shortened before, then the output response will be the same short URL.
- The API request made using the shortened URL will redirect the request to the original URL.
- The project is created using Golang, Docker, HTML, and Makefile.

### Installations

The following software needs to be installed on the system to operate on the application-
```
1. Golang
2. Make
3. Docker
```

## API Endpoints

Below is the list of API endpoints with their respective input and output.
```
1. To Create a ShortURL -

#### Input Request
GET {Base_URL}/createShortUrl/?url={Long_URL}

#### Output Response
{
    "Url": "{Base_URL}/shortUrl/{shortURL_Path}"
}
```
```
2. To Fetch the originalURL -

#### Input Request
GET "{Base_URL}/shortUrl/{shortURL_Path}"

#### Output Response
Redirection to the original URL
```

## Useful make commands

###### Compile the application
- make build

###### Run the application
- make run

###### Unit test the application
- make test

###### Create a docker image to package the application
- make image

###### Execute all operations
- make

## Containerization

A docker file is present in the application with a docker-compose file. The **make image** command can build the docker image, which later can be used by the docker-compose file to serve the application over the path [Url]({Base_URL}) for the API requests.