FROM golang AS build

# STAGE 1
WORKDIR /app
ADD . /app
COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go build -o binary

# STAGE 2

FROM alpine 
WORKDIR /application
COPY --from=build /app/binary /application

ENTRYPOINT ./binary

