# go-server-practice
goのバックエンドサーバーの練習用

## 起動方法
以下を実行します。
```
go mod tidy
```

##　ディレクトリ構成

クリーンなアーキテクチャを採用しています。ディレクトリ構成は以下の通り

```
.
├── cmd/
│   └── main.go              # DI配線・サーバー起動のみ
├── internal/
│   ├── domain/
│   │   └── todo.go          # Entity構造体だけ(ロジックなし)
│   ├── usecase/
│   │   ├── todo_repository.go # Repositoryインターフェース定義
│   │   └── todo_usecase.go    # 業務ロジック(interfaceにのみ依存)
│   ├── repository/
│   │   └── todo_repository.go # Postgres実装(usecase.Repositoryを満たす)
│   └── controllerß/
│       └── todo_handler.go    # Ginハンドラ+ルーティング
├── docker-compose.yml
├── Dockerfile
└── go.mod
```
ß
## Gin を使う最小構成

```bash
go get github.com/gin-gonic/gin@latest
go run ./cmd
```

起動後に以下を確認できます。

```bash
curl http://localhost:8080/
```

レスポンス例:

```json
{"message":"hello from gin"}
```

`cmd/main.go` に Gin のルーティングを定義しています。


