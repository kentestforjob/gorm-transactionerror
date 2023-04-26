# For Testing #

sql: transaction has already been committed or rolled back error when use transaction in gorm. How to handle it?

### Requirements ###

* installed docker 

### How do start the program? ###

1. Need to create docker network: `docker network create dummy-network`
2. Start the program by execute `docker composer build`
3. Start the program by execute `docker composer up`

### How to simulate the problem ###
1. call api/dummy/update api which will perform transaction 
```
curl -X 'POST' \
  'http://127.0.0.1:8003/api/dummy/update' \
  -H 'accept: */*' \
  -H 'Content-Type: application/json' \
  -d '{
  "user_id": 1,
  "email": "kentestforjob@gmail.com"
}'
```

2. call api/dummy/list api
```
curl -X 'GET' \
  'http://127.0.0.1:8003/api/dummy/list' \
  -H 'accept: */*'
```
