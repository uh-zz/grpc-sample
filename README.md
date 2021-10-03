# grpc-sample

this repo is sample to help you understand gRPC.

# System

![image](https://user-images.githubusercontent.com/47747828/135745893-8eb009ca-bc1a-4bf0-83f9-457ce6fb8995.png)

# Directory

```
.
├── README.md
├── go.mod
├── go.sum
├── pinger
│   ├── client
│   │   └── client.go
│   ├── lib // protocによって生成されるGoコード
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

# Install

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

# VSCode Extention

If you want syntax highlighting or code-formatter

- vscode-proto3
- Clang-Format

```
grpc-sample $ brew install clang-format
```

# Definition `.proto`

# Create a gRPC server in Go

FYI: https://developers.google.com/protocol-buffers/docs/reference/go-generated#package

For Go 1.16 or later, you need to install the Go code generation plugin with the following command

```
grpc-sample　$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
grpc-sample　$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
```

Also, add the following options to `pinger.proto`.

```proto
option go_package = "./lib";
```

# Auto Go code generation

```
grpc-sample　$ protoc -I ./proto pinger.proto --go_out=./pinger --go-grpc_out=./pinger
grpc-sample　$ tree pinger
pinger
├── lib
│   ├── pinger.pb.go
│   └── pinger_grpc.pb.go
└── server.go

```

# Start gRPC server in Go

```
pinger $ go run server.go
2021/10/03 15:42:53 Pinger listening at [::]:5300
```

# Create a client to check if the gRPC server is running

Keep it running gRPC server

```
pinger $ go run server.go
2021/10/03 15:42:53 Pinger listening at [::]:5300
```

Access to client

```
client $ go run client.go
Pong: text:"pong"
```

# Create Stub in Rails

# Install Rails

# Create in Rails 5 api mode

```
server $ bundle init
server $ echo 'gem "rails", ">=5.2.2.1"' >> Gemfile
server $ bundle exec rails new . --api -O
server $ bin/rails server
```

# Auto Ruby code generation

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

# Invoke from Application Controller

```ruby
def ping
    pinger_stub = Pinger::Pinger::Stub.new('localhost:5300', :this_channel_is_insecure)

    pong = pinger_stub.ping(Pinger::Empty.new)
    render json: { pong: pong.text }
end
```

# Start Rails server and Access to URL

Start gRPC server in Go

```
pinger $ go run server.go
2021/10/03 15:42:53 Pinger listening at [::]:5300
```

Start Stub in Rails

```
server $ bin/rails server
```

Ping from Browser!

```
http://localhost:3000/ping
```

Pong

```json
{
  "pong": "pong"
}
```
