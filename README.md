# userSrv
Storing information of the users

### How to run this app
1. Migrate to userSrv folder
2. Run ---> ./buildapp.zsh to build the app
2. Run ---> ./runapp.zsh "<port_number>" for eg  ./runapp.zsh "9000"
3. Alternatively you can simply run --->  go run main.go <port_number> for eg  go run main.go 9000
4. The app starts to listen for requests on port 9000 by default if no port number is provided
5. Download Postman desktop agent 
6. From Postman from Web Browser run the following apis

### userSrv APIs

1.
| POST   | localhost:9000/users/signup           | Set the user object

Input :

{
    "id": "1",
    "name": "gaurav",
    "signupTime": 1234566
}

Expected Output :
Response HTTP Status Code : 200

2.
| GET    | localhost:9000/users/:user_id         | Get the user object from user_id

Input : mention whatever user_id you want in url for eg "localhost:9000/users/1"

Expected Output :

{
    "id": "1",
    "name": "gaurav",
    "signupTime": 1234566
}

Response HTTP Status Code : 200

3.
| GET    | localhost:9000/users                  | Lists all the users

Input : Nothing

Expected Output : array of users present for eg 

[
    {
        "id": "1",
        "name": "gaurav",
        "signup_time": 1234566
    },
    {
        "id": "2",
        "name": "saurav",
        "signupTime": 1234566
    }
]

Response HTTP Status Code : 200

### how to run the unitests
1. Migrate to project folder
2. Run -->   ./runtests.zsh
3. Alternatively you can run the test with command : go test ./... -v


### General Module Info #WHY
1. Used "gin" framework for writing Apis : Chose this since it provides a intuitive api and built in functionality
2. Used "testing" framework for testing
3. Used "github.com/go-playground/validator/v10" for validation

### Improvements to be made in future #WHY
1. Add basic Authentication
3. Currently we are storing the user objects in memory we can use NoSql database for storing this in a persistent manner as there is no relational data and queries are simple 
4. Add pagination in the GetUsers Api 

