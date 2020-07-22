# Simple Todo Frontend

## Configuring the frontend

Modify `setup.go` to specify the client ID and the RestEndpoint.

## Get the vugu tools

```
go get -u github.com/vugu/vgrun
```

## Run the application using the vugu tools

```
vgrun devserver.go
```

Then browse to the running server: http://localhost:8844/. There, you will be
presented with a basic login screen which will redirect you to the Cognito
service.

To login, user the credentials

```
Username: testuser
Password: Passw0rd!
```

This will log in and redirect you back to localhost; it should now have a token
which can be used to communicate with the REST API.


