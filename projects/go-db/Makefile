SRC = cmd/main.go
RM = rm -rf
NAME = server

download-dependencies:
	go mod download

# INFO: prod run
runserver: build
	./${NAME}

# TODO: "source .env"
run:
	go run ${SRC}

build: download-dependencies
	go build -o $(NAME) $(SRC)

fclean:
	$(RM) $(NAME)

re: fclean build

## TODO
## https://levelup.gitconnected.com/how-to-live-reload-code-for-golang-and-docker-without-third-parties-ee90721ef641
#
## Variable for filename for store running procees id
#PID_FILE = /tmp/my-app.pid
## We can use such syntax to get main.go and other root Go files.
#GO_FILES = "./src/$(wildcard *.go)"
#
## Start task performs "go run main.go" command and writes it's process id to PID_FILE.
#start:
#	go run $(GO_FILES) & echo $$! > $(PID_FILE)
## You can also use go build command for start task
## start:
##   go build -o /bin/my-app . && \
##   /bin/my-app & echo $$! > $(PID_FILE)
#
## Stop task will kill process by ID stored in PID_FILE (and all child processes by pstree).
#
## -kill `pstree -p \`cat $(PID)\` | tr "\n" " " |sed "s/[^0-9]/ /g" |sed "s/\s\s*/ /g"`
#stop:
#	kill -9 $(lsof -t -i:1234)
#
## Before task will only prints message. Actually, it is not necessary. You can remove it, if you want.
#before:
#	@echo "STOPED my-app" && printf '%*s\n' "40" '' | tr ' ' -
#
## Restart task will execute stop, before and start tasks in strict order and prints message.
#restart: stop before start
#	@echo "STARTED my-app" && printf '%*s\n' "40" '' | tr ' ' -
#
## Serve task will run fswatch monitor and performs restart task if any source file changed. Before serving it will execute start task.
## fswatch -or --event=Updated /home/parla-go/src |
#serve: start
#	fswatch -o src | \
#	xargs -n1 -I {} make restart
#
## .PHONY is used for reserving tasks words
#.PHONY: start before stop restart serve

