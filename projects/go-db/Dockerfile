FROM golang:1.15-alpine

# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh make

LABEL maintainer="Alex Mogilevskiy <dark.qvark@gmail.com>"

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY Makefile .
RUN make install-deps

COPY . .

RUN make build

EXPOSE 8080

CMD ["make", "runserver"]

# TODO: Docker readme in my repo !!!
# https://www.digitalocean.com/community/tutorials/how-to-remove-docker-images-containers-and-volumes
# INFO: For test: run deamon in container
#RUN touch test.sh && echo "while true; do ping 8.8.8.8; done" >> test.sh
#CMD ["sh", "test.sh"]
