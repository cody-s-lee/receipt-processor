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

### Implementation Notes

1. I used [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen) for code generation of the basic types from the api.yml rather than manual creation. This is built into the build process automatically and works from a fresh clone. In a production environment it would require some additional futzing to configure it to handle updates to api.yml gracefully.
2. Rules #2 and #3 could be simplified because if a total is a round dollar amount it is a multiple of 0.25. Howver, I think that would be a mistake because it would cause other difficulties in verifying non round dollar amount multiples of 0.25. It would also be more difficult if either rule changes because they would need to be untangled from each other.
3. I used the `govalues/decimal` library because it provided the features I needed. I considered the more popular `ericlagergren/decimal` library but was unsatisfied with its lack of a ceiling function. An alternative would be to represent the `total` and `price` fields as a pair of ints (dollars and cents) or a single int (cents). I chose the decimal library to minimize writing my own parsing code.
4. Rules #3 and #5 I used `Decimal` kinda-constant values and multiplication to handle. I felt this was more readable than constructing the magic numbers inline. I didn't deeply handle edge cases of very large `total` or `price` values potentially overloading the scale of the decimals. Had I used a different method of representing currency this would be unnecessary.
5. Rule #4 has an order of operations restriction, which I noted in a comment.
6. Managing `purchaseDate`/`purchaseTime` as separate simples with no timezone isn't my favorite thing to do in go because it means that the receipt parser is fairly permissive. Rather than strictly validating those fields I went with being permissive in what I accept and strict in what I emit. With more time and autonomy I'd run down the OpenAPI spec to try to make the formatting restrictions part of the api.yml spec rather than handling it in code.
7. Rule #7 has some ambiguity, primarily around "after 2:00pm and before 4:00pm". In ordinary parlance this may mean any of:
    1. 2 pm inclusive - 4 pm inclusive, so 14:00:00.000 and 16:00:00.000 are both valid
    2. 2 pm inclusive - 4 pm exclusive, so 14:00:00.000 is valid but 16:00:00.000 is invalid
    3. 2 pm exclusive - 4 pm exclusive, so 14:00:00.000 and 16:00:00.000 are invalid but 14:00:00.001 and 15:59:59.999 are valid

    I went with the third option and read it literally as excluding exactly 2 pm and 4 pm. In the real world I'd follow up on purpose and requirements to nail down the correct semantics. Also, since we only have minute values on our receipts everything is truncated to that value. This meant that implementation was done by calculating constants for how many minutes past start of day 2 pm and 4 pm represent.
8. There are 3 utility functions used internally on `Receipt`
    1. `validate` - Because the api.yml is fairly permissive an additional validation function is useful in order to throw a 400 bad request upon call to receipt processing.
        Validation is not as strict as it could be. I omitted validation that `price` and `total` have 2 decimal places because it doesn't make the code harder given the library I used.  
    3. `datetime` - Merges the `purchaseDate` and `purchaseTime` into a single datetime object. Given more time I may have tried to alter the api.yml to do this automatically if possible.
    4. `score` - This one simply calculates the receipt score based on the rules. I broke it out in order to minimize clutter of the processing endpoint function.
9. Receipt id generation uses UUID v1 because it guarantees no collisions in tradeoff of locking generation. For a real system I would more deeply consider the tradeoff of locking during UUID generation versus ensuring no collisions.
10. Receipt id -> score mapping is stored simply on the server. This made testing easier. In the real world this would likely better fit into a cached database or a queue depending on purpose.
11. I implemented some basic tests. Given more time I'd put more energy in feeding tests from example json files and running parsing edge cases through testing.

## Build and Deploy

_Based on awesome-compose's `nginx-golang`_

The compose file defines an application with two services `proxy` and `backend`.
When deploying the application, docker compose maps port 80 of the frontend service container to the same port of the host as specified in the file.
Make sure port 80 on the host is not already in use.

## Deploy with docker compose

```
$ docker compose up -d
...
[+] Running 3/3
 ✔ Network receipt-processor_default      Created                                                                0.0s
 ✔ Container receipt-processor-backend-1  Started                                                                0.2s
 ✔ Container receipt-processor-proxy-1    Started                                                                0.3s
```

## Expected result

Listing containers must show two containers running and the port mapping as below:

```
$ docker compose ps
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
