FROM golang:1.16-alpine

WORKDIR /app

ENV TZ Asia/Seoul
RUN apk add curl
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s \
        && cp ./bin/air /bin/air
RUN curl -LJO https://raw.githubusercontent.com/eficode/wait-for/master/wait-for \
     && chmod +x ./wait-for \
     && mv ./wait-for /usr/local/bin/wait-for
COPY . .
RUN go install

EXPOSE 8000

CMD ["go", "run", "main.go"]