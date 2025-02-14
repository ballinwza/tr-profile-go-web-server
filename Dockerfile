FROM --platform=amd64 golang:1.24.0-alpine3.21 AS build
WORKDIR /app

COPY . ./
RUN go mod download
RUN chmod +x /app/docker-load-dotenv.sh
RUN sh /app/docker-load-dotenv.sh


RUN go build

FROM --platform=amd64 node:20.14-alpine3.20 AS runner
RUN apk add --no-cache libc6-compat
WORKDIR /app
COPY --from=build /app ./

EXPOSE 8080

ENTRYPOINT [ "/bin/sh", "docker-entrypoint.sh" ]