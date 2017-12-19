# Book-Server by Dep

## Commands to get the dependecies

- `dep init`
- `dep ensure`
  
## Example commands to run

- The following command will run the book server at port 10001 and it requires the authentication from user,

  `go run main.go --port=10001 --logIn=true`

- these don't require authentication

  `go run main.go --port=10001 --logIn=false`

  or
  
  `go run main.go --port=10001`
