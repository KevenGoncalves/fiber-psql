FROM golang:1.20 as builder 

# changing directory
WORKDIR /build

# copying the mod and sum to download it
COPY go.mod go.sum ./
RUN go mod download 

# copying others content
COPY . .

# installing taskfile and goose
RUN go install github.com/go-task/task/v3/cmd/task@latest
RUN go install github.com/pressly/goose/v3/cmd/goose@latest 

# arguments to run the build
ARG DB_URI 
ARG DB_DRIVER  

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOOSE_DRIVER=${DB_DRIVER} GOOSE_DBSTRING=${DB_URI}

# changing directory to the migrations folder and migrate up
WORKDIR /build/internal/databases/postgres/migrations
RUN goose up

# changing directory to the build and build it
WORKDIR /build 
RUN task build

# changing image and copying the build to the new image
FROM scratch

COPY --from=builder ["/build/http-server","/http-server"]

# env to run the prod
ENV GO_ENV=production
ENV SERVER_PORT=:3000

CMD [ "/http-server" ]
