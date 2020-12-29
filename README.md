# my-golang-service-template
my template using echo

#### Running Development

step 1: check and cd in your path

```
$ pwd 
$ cd {yourPath}
```

step 2: run docker-compose up for create mysql / redis

```
$ docker-compose up
```

step 3: run my service :)

```
$ go run main.go
```

step 4: test my service use curl or import postman_collection 

```
$ curl --location --request POST 'http://localhost:1323/inquiry-data' \
--header 'Content-Type: application/json' \
--data-raw '{
    "customerID": "10001"
}'
```