## Generate Server
```bash
swagger generate server -A permissions -f swagger.yml
```

## Run Server
```bash
docker-compose up
```

## Example Requests
Get All Permissions
```bash
curl "http://127.0.0.1:8080/permission"
```
Get Specific User's Permissions
```bash
  curl "http://127.0.0.1:8080/permission?user=user_1_id"
```

### Please check the swagger.yml for more details on the API
Go to [Swagger](http://127.0.0.1:8080/docs)
