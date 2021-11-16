FROM golang as build 

WORKDIR /app

ADD . .

RUN go version

RUN CGO_ENABLED=0 GOOS=linux go build -o counterapi

# RUN CGO_ENABLED=1 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o msgraph-resetpassword
# RUN CGO_ENABLED=1 GOOS=linux go build -a -ldflags '-linkmode external -extldflags "-static"' -o msgraph-resetpassword

FROM scratch as production
# FROM alpine as production

COPY --from=build /app/counterapi .

ENV TZ=Asia/Bangkok

CMD ["./counterapi"]