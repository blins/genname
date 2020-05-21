FROM golang:1-alpine AS build
COPY . .
RUN go build -ldflags '-s -w' -o /genname

FROM alpine
COPY --from=build /genname /usr/bin/genname
ENTRYPOINT ["/usr/bin/genname"]
EXPOSE 8080/tcp
CMD [""]
