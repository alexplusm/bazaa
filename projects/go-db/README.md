### Run

- go run cmd/main.go
- go build cmd/main.go && ./main

app must run from project root for correct creation media_* directories

----------------------

// todo: decompose tasks into little tasks

1) (zip-parsing) Parse ZIP and create tasks in DB
   * task.expertAnswer take from folder structure
   * DONE media_root has plane structure | media_root - folder with images
   * resolve TODO with "io/ioutil.ReadFile"

2) (game-creation) Create API endpoint for game creation
   * uploading zip | choice schedules ids | ... and other data from model

3) Upload images for tasks using SOAP service
   * TODO

...

n) Use codegen for interface generation
   * write openAPI spec for each endpoint

// undefined - правильный ответ (for statisctics) !!! answer

TODO: 
-- Database dumps?

------------

## info
* pgx
https://eax.me/golang-pgx/

* testing
https://eax.me/golang-dockertest/

## todos

* if I change package code and dont run "go build" what happens? "go run main.go"

* авто импорты
go get golang.org/x/tools/cmd/goimports

* work with file 
io/ioutil.ReadFile

How work with them?
* go.sum & go.mod

* Create Makefile with reload

Tasks:
* (Go Kernighan p.29, ex 1.3)
* (Go Kernighan p.39, ex 1.5 1.6) + rainbow color palette
* (Go Kernighan p.43, ex 1.10 1.11)
* (Go Kernighan p.71, ex 2.3 2.4 2.5) // byte operations
* (Go Kernighan p.87, ex 3.1 - 3.4) // svg + math
* (Go Kernighan ch 3.4) // complex numbers
* (Go Kernighan p.103, ex 3.10 - 3.12)

## business todos
* method for adding new external system
    * with exist uuid
    * with generating new uuid
