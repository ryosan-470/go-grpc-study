# gRPC study

# ビルドの仕方
## 前提として必要なもの

- Go
- Protocol Bufferをインストールする. Macなら `brew install protobuf` で入る。
- 以下のパッケージを `go get` でインストールする
```console
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/protoc-gen-go
```

## 実際にビルドして動かす
- gRPCスタブを次のようにして生成する
```console
protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. ./proto/alive.proto
```

- リバースプロキシを次のようにして生成する
```console
protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. ./proto/alive.proto
```

- gRPCのサーバーを起動する
```console
go run server/main.go
```

- HTTPサーバーを起動する
```console
go run gateway/main.go
```

- 動くかの確認

```console
curl http://localhost:3000/alive
```
