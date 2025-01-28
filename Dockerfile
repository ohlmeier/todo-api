# Build stage
FROM golang:1.23.5-alpine AS build
WORKDIR /app
COPY . .
RUN go build -o main .

# Run stage  
FROM scratch
WORKDIR /app
COPY --from=build /app/main .
CMD ["/app/main"]