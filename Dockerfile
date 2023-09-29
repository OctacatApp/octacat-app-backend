FROM golang:alpine AS igo
WORKDIR /app
COPY . .
RUN dir="src/cmd/" \
	&& go build -o $dir/main $dir/main.go
CMD [ "./src/cmd/main", "-env", "prod" ]