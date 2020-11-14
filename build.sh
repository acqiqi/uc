CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
rm -rf uc
mv main uc