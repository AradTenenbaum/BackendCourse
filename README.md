# Simple Bank API
### Api built with Golang that implements bank functionality

The api is built using - Golang, Sqlc, PostgreSql, Redis, Github Actions CI, Docker, Kubernetes, gRPC and HTTP.
You can dynamically switch between gRPC server and HTTP server by the SERVER_TYPE in the environments variables.
The system is following the ACID principles

 - gRPC Functions Implemented:
	 - Create User: Create a user in the db, Insert a process of sending verification email into redis queue. Sending an email with a url to verify the email
	 - Create Account: Create an account in the db for the requested user by the auth token
	 - Login: Login in to the system, receiving a bearer token for the security, implements both jwt and paseto tokens
	 - Get Account: Get a specific account
	 - Update User: Update user's details
 - Directories order:
	 - api: http server built with gin library
	 - db: migrations, mock store, all the queries and sqlc functions
	 - doc: documentation for the db using db docs, and for the api using swagger ui
	 - eks: all the yaml files for deploying to aws kubernetes service
	 - gapi: gRPC server all functions and rpcs
	 - grpc_test_client: example of a simple client build with nodejs
	 - mail: all the email sender functions
	 - pb: functions created by protoc for the gRPC server
	 - proto: the proto files contains the rpcs and the service
	 - token: the functions needed for creating and verifying the auth token both with jwt and paseto
	 - util: helpful functions that are in use all over the system
	 - val: validation for the user's requests input
	 - worker: redis connection, enqueue procceses, email verification task
