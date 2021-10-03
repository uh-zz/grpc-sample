# grpc-sample

this repo is sample to help you understand gRPC

## インストール

- Ruby
- Go
- protobuf

```
grpc-sample $ ruby --version
ruby 2.6.6p146 (2020-03-31 revision 67876) [x86_64-darwin19]

grpc-sample $ go version
go version go1.17 darwin/amd64

grpc-sample $ brew install protobuf
grpc-sample $ protoc --version
libprotoc 3.17.3
```

### VSCode 拡張

シンタックスハイライト、フォーマッターを入れる場合

- vscode-proto3
- Clang-Format

```
grpc-sample $ brew install clang-format
```

## proto ファイル定義

## Go の gRPC サーバー作成

FYI: https://developers.google.com/protocol-buffers/docs/reference/go-generated#package

Go 1.16 以降は以下のコマンドにて、Go コード生成プラグインをインストールする必要がある

```
grpc-sample　$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
grpc-sample　$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
```

また、`pinger.proto`に以下のオプションを追加

```proto
option go_package = "./pinger";
```

### Go コードを自動生成

```
grpc-sample　$ protoc -I ./proto pinger.proto --go_out=./pinger --go-grpc_out=./pinger
```

### Go サーバー起動

```
pinger $ go run server.go
2021/10/03 15:42:53 Pinger listening at [::]:5300
```

### Go のテストクライアントを作成してサーバー起動を確認

サーバーを起動させておく

```
pinger $ go run server.go
2021/10/03 15:42:53 Pinger listening at [::]:5300
```

クライアントからアクセス

```
client $ go run client.go
Pong: text:"pong"
```
