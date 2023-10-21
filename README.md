# userSrv
Storing information of the users
TODO: cleanup this file
TODO: write unitests for the validators as well

### userSrv APIs

| POST   | localhost:9000/users/signup           | Create the user entry
| GET    | localhost:9000/users                  | List the users
| GET    | localhost:9000/users/:user_id         | Get the user from user_id

### how to build the app
go to project folder in mac
run the following commands
go build main.go
mv main ./bin/users

### how to run the app
Migrate to project folder
There is already a "users" binary in bin folder
Run with command : ./bin/users

### how to run the unitests
Migrate to project folder
Run the test with command : go test . -v

### Input
{
    "id": "1",
    "name": "gaurav",
    "signup_time": 1234566
}


The requirements are 

Rest Apis :(correctly validated and the response status code must match the error type , validation and error handling , return error as well , errors should have some reference to the function call)
1. POST {{base_url}/users} : body contains a user Object Json
2. GET  {{base_url}/users/:user_id} : returns a user object
3. GET  {{base_url}/users} : lists all the users // this is bonus , we should take care of listing also and pagination also

write proper covering tests , validation and error handling 
error handlers can be put in a separte file

BONUS
requires basic authentication for any operation(the credentials can be hardcoded for sake of testing)
the listening port of the server can be configured with command line argument 

// reference on how to run and test the app


We are going to use 
- NoSql database (mongo db : why ? - easily scalable , we only require basic queries , there are know relations )

// signup is same as setUser