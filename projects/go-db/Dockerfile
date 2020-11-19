FROM golang:1.15-alpine

# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

# Maintainer Info
LABEL maintainer="Alex Mogilevskiy <dark.qvark@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependancies
# Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 1235

# Run the executable
CMD ["./main"]


# TODO: Docker readme !!!

# https://www.digitalocean.com/community/tutorials/how-to-remove-docker-images-containers-and-volumes

# INFO: For test: run deamon in container
#RUN touch test.sh && echo "while true; do ping 8.8.8.8; done" >> test.sh
#CMD ["sh", "test.sh"]
