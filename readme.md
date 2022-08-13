# Resource to learn how to apply gRPC with Golang
[course](https://www.youtube.com/watch?v=x8dybRs5q_g&list=PLC4c48H3oDRzLAn-YsHzY306qhuEvjhmh)
[grpc status code](https://grpc.github.io/grpc/core/md_doc_statuscodes.html)

# Calculator Project
## Requirement
- Operating System: Linux/MacOS

## How to run the code in *calculator* folder's project
### Generate the SSL key and cert
- First, browse to *calculator* folder by using the command: `cd calculator`
- Then, run the **SSLCmd.sh** file by using the command: `ssl/SSLCmd.sh`
- After that, **server.crt** and **server.key** are generated in *ssl* folder 

### Run the server
- Browse to the root folder
- Run this command: `make run-cal-server` to run the code binding in the **Makefile** file
- We would see the notification that the server is running on console

### Run the client app
- Browse to the root folder
- Run this command: `make run-cal-client` to run the code binding in the **Makefile** file
- We would see the notification that the client is communicating with the server on console

## How to run the code in *contact* folder's project
### Run the server
- Browse to the root folder
- Run this command: `make run-contact-server` to run the code binding in the **Makefile** file
- We would see the notification that the server is running on console

### Run the client app
- Browse to the root folder
- Run this command: `make run-contact-client` to run the code binding in the **Makefile** file
- We would see the notification that the client is communicating with the server on console

## How to run the code in *contact* folder's project
### Run the server
- Browse to the root folder
- Run this command: `make run-gateway-server` to run the code binding in the **Makefile** file
- We would see the notification that the server is running on console

### Run the reverse proxy
- Browse to the root folder
- Run this command: `make run-gateway-rproxy` to run the code binding in the **Makefile** file
- We would see the notification that the server is running on console
- At this moment, the backend was already available

### Using Curl/Postman to request to the backend
- The URL of the API is http://localhost:8081/api/echo
- The type of API is POST
- JSON format is {"msg": string}
- Using Curl/Postman or any type of app to request to the backend with the above specs, you would get the response, and the server would be logging the process
