go env -w CGO_ENABLED=0
go build -o GXL.exe
go build -o chain33-cli.exe github.com/33cn/plugin/cli
