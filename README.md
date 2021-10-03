# grpc-sample

gRPC を理解するためのサンプルです。

## 構成

![image](https://user-images.githubusercontent.com/47747828/135745893-8eb009ca-bc1a-4bf0-83f9-457ce6fb8995.png)

## ディレクトリ

```
.
├── README.md
├── go.mod
├── go.sum
├── pinger
│   ├── client
│   │   └── client.go
│   ├── pinger // protocによって生成されるGoコード
│   │   ├── pinger.pb.go
│   │   └── pinger_grpc.pb.go
│   └── server.go // gRPCサーバー
├── proto
│   └── pinger.proto // IDL
└── server // gRPCスタブ
    ├── Gemfile
    ├── Gemfile.lock
    ├── README.md
    ├── Rakefile
    ├── app
    ├── bin
    ├── config
    ├── config.ru
    ├── lib // protocによって生成されるRubyコード
    │   ├── pinger_pb.rb
    │   └── pinger_services_pb.rb
    ├── log
    ├── public
    ├── test
    ├── tmp
    └── vendor
```

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
grpc-sample　$ tree pinger
pinger
├── pinger
│   ├── pinger.pb.go
│   └── pinger_grpc.pb.go
└── server.go

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

## Rails のスタブ作成

### Rails のインストール

Rails 5 の API モードで作成

```
server $ bundle init
server $ echo 'gem "rails", ">=5.2.2.1"' >> Gemfile
server $ bundle exec rails new . --api -O
server $ bin/rails server
```

### Ruby コードを自動生成

```
server $ echo "gem 'grpc'" >> Gemfile
server $ echo "gem 'grpc-tools'" >> Gemfile
server $ bundle install

server $ bundle exec grpc_tools_ruby_protoc -I ../proto --ruby_out=lib --grpc_out=lib ../proto/pinger.proto

server $ tree lib
lib
├── pinger_pb.rb
├── pinger_services_pb.rb
└── tasks
```

### Controller から呼び出し

```ruby
def ping
    pinger_stub = Pinger::Pinger::Stub.new('localhost:5300', :this_channel_is_insecure)

    pong = pinger_stub.ping(Pinger::Empty.new)
    render json: { pong: pong.text }
end
```

### Rails を起動して URL にアクセス

gRPC サーバー起動

```
pinger $ go run server.go
2021/10/03 15:42:53 Pinger listening at [::]:5300
```

Rails スタブ起動

```
server $ bin/rails server
```

ブラウザからピン！

```
http://localhost:3000/ping
```

ポン

```json
{
  "pong": "pong"
}
```
