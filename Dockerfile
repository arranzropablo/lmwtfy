FROM golang
ADD . app/
WORKDIR app/
EXPOSE 80
CMD go run cmd/main.go