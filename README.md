# receipt-processor

An implementation of the Fetch Receipt Processor Challenge

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

_Based on awesome-compose's `nginx-golang`_

The compose file defines an application with two services `proxy` and `backend`.
When deploying the application, docker compose maps port 80 of the frontend service container to the same port of the host as specified in the file.
Make sure port 80 on the host is not already in use.

## Deploy with docker compose

```
λ docker compose up -d
...
[+] Running 3/3
 ✔ Network receipt-processor_default      Created                                                                0.0s
 ✔ Container receipt-processor-backend-1  Started                                                                0.2s
 ✔ Container receipt-processor-proxy-1    Started                                                                0.3s
```

## Expected result

Listing containers must show two containers running and the port mapping as below:

```
λ docker compose ps
NAME                          IMAGE                       COMMAND                  SERVICE   CREATED          STATUS          PORTS
receipt-processor-backend-1   receipt-processor-backend   "/code/bin/backend"      backend   47 seconds ago   Up 46 seconds
receipt-processor-proxy-1     nginx                       "/docker-entrypoint.…"   proxy     47 seconds ago   Up 46 seconds   0.0.0.0:80->80/tcp
```

After the application starts, test via curl:

```
$ curl -XPOST localhost:80/receipts/process -d '{
  "retailer": "M&M Corner Market",
  "purchaseDate": "2022-03-20",
  "purchaseTime": "14:33",
  "items": [
    {
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    }
  ],
  "total": "9.00"
}'
{"id":"80ca0306-642b-11ef-939f-0242ac120002"}
$ curl localhost:80/receipts/80ca0306-642b-11ef-939f-0242ac120002/points
{"points":109}
```

Stop and remove the containers

```
$ docker compose down
```
