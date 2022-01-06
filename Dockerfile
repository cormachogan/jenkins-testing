#
# The Dockerfile is a multistage one to keep the image size as small as possible. 
# It starts with a build image based on golang:alpine. 
# The resulting binary is used in the second image, which is just a scratch one. 
# A scratch image contains no dependencies or libraries, 
# just the binary file that starts the application.
#

FROM golang:alpine AS build-env
RUN mkdir /go/src/app && apk update && apk add git
ADD main.go /go/src/app/
WORKDIR /go/src/app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o app .

FROM scratch
WORKDIR /app
COPY --from=build-env /go/src/app/app .
ENTRYPOINT [ "./app" ]
