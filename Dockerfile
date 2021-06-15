FROM alpine:edge as builder
COPY . /app
WORKDIR /app
RUN apk add nodejs npm go
RUN cd www && npm install && npm run build
RUN go build

FROM alpine:edge as app
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/rainy /app/rainy

EXPOSE 4000

CMD [ "./rainy" ]