# receipt-processor

An implementation of the Fetch Receipt Processor Challenge

_Based on awesome-compose's `nginx-golang`_

## Instructions

> Build a webservice that fulfils the documented API. The API is described below. A formal definition is provided in the
> api.yml file, but the information in this README is sufficient for completion of this challenge. We will use the
> described API to test your solution.

### Rules

These rules collectively define how many points should be awarded to a receipt.

1. One point for every alphanumeric character in the retailer name.
2. 50 points if the total is a round dollar amount with no cents.
3. 25 points if the total is a multiple of 0.25.
4. 5 points for every two items on the receipt.
5. If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the
   nearest integer. The result is the number of points earned.
6. 6 points if the day in the purchase date is odd.
7. 10 points if the time of purchase is after 2:00pm and before 4:00pm.

## Build and Deploy

The compose file defines an application with two services `proxy` and `backend`.
When deploying the application, docker compose maps port 80 of the frontend service container to the same port of the
host as specified in the file.
Make sure port 80 on the host is not already in use.

## Deploy with docker compose

```
$ docker compose up -d
Creating network "nginx-golang_default" with the default driver
Building backend
Step 1/7 : FROM golang:1.13 AS build
1.13: Pulling from library/golang
...
Successfully built 4b24f27138cc
Successfully tagged nginx-golang_proxy:latest
Creating nginx-golang_backend_1 ... done
Creating nginx-golang_proxy_1 ... done
```

## Expected result

Listing containers must show two containers running and the port mapping as below:

```
$ docker compose ps
NAME                     COMMAND                  SERVICE             STATUS              PORTS
nginx-golang-backend-1   "/code/bin/backend"      backend             running
nginx-golang-proxy-1     "/docker-entrypoint.…"   proxy               running             0.0.0.0:80->80/tcp
```

After the application starts, navigate to `http://localhost:80` in your web browser or run:

```
$ curl localhost:80

          ##         .
    ## ## ##        ==
 ## ## ## ## ##    ===
/"""""""""""""""""\___/ ===
{                       /  ===-
\______ O           __/
 \    \         __/
  \____\_______/


Hello from Docker!
```

Stop and remove the containers

```
$ docker compose down
```
