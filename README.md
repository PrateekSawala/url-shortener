## Overview

The project aims to provide a simple URL shortener service.

## Description

- The API processes the input URL as an argument over a REST API, write it to a JSON file, and returns a shortened URL as a result.
- The case where the API request gets made with a long URL that has already got shortened before, then the output response will be the same short URL.
- The API request made using the shortened URL will redirect the request to the original URL.
- The project is created using Golang, Docker, HTML, and Makefile.

## Installations

The following software needs to be installed in the operating system to operate on the application.
```
1. Golang
2. Make
3. Docker
```

## API Endpoints

Below is the list of API endpoints with their respective input and output. Consider Base_URL to be the exposed application port.
Example: Base_URL=http://localhost:3002 

```
1. To Create a ShortURL -

#### Input Request
GET {Base_URL}/createShortUrl/?url={Long_URL}

Example: 
http://localhost:3002/createShortUrl/?url=https://google.com

#### Output Response
{
    "Url": "{Base_URL}/shortUrl/{shortURL_Path}"
}

Example:
{
    "Url": "http://localhost:3002/shortUrl/2aa789c823074703b7baa8a524eb4aad"
}
```
```
2. To Fetch the originalURL -

#### Input Request
GET "{Base_URL}/shortUrl/{shortURL_Path}"

Example: 
http://localhost:3002/shortUrl/2aa789c823074703b7baa8a524eb4aad

#### Output Response
Redirection to the original URL

Example: 
https://google.com
```

## Useful make commands

###### Compile the application
- make build

###### Run the application
- make run

###### Execute all operations
- make

###### Unit test the application
- make test

###### Create a docker image to package the application
- make image

## Unit test

The unit tests are present in the application. The **make test** command can run unit tests on the application after the application starts serving API requests over a path [Url]({Base_URL}).

## Containerization

A docker file is present in the application with a docker-compose file. The **make image** command can build the docker image. Any of the following two ways can get used for serving the application over a path [Url]({Base_URL}) for the API requests after the image is created.
```
1. Run using docker.
docker run -p 3002:80 --name urlshortener urlshortener:test

2. Run using the docker-compose file.
docker-compose -f docker-compose.yml up
```