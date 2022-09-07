# Docker+Go+GraphQL+Firestoreで簡単なバックエンド作ってみた

## dockerコマンド
- ビルド
`docker build -t gqlgen-todos .`
- コンテナにアタッチする場合
`docker run -it --rm --name go-test -p 8080:8080 gqlgen-todos`
- バックグラウンド起動(CMDのコメントアウトを復活させる)
`docker run --name go-test -d -p 8080:8080 gqlgen-todos`

## docker compose コマンド
- `docker compose up -d --build`

## firestore setup
1. firebaseのプロジェクトを作成
2. プロジェクト内にCloud Firestore を作成
3. secret key を生成する(プロジェクト設定のサービスアカウント)
4. secret key を`firestore/secret.json`に配置
5. ~~`firestore/connect.go`のProjectIDを変更~~

## 動かし方
1. `docker compose up -d --build`
2. http://localhost:8080 を開く
3. 色々リクエストを発行してみる
- todoの一覧取得
```
query {
  todos {
    id
    text
    done
    userId
  }
}
```
- todoの追加
```
mutation {
  createTodo(input: { text: "todo description", userId: "12345" }) {
    id
    text
    text
    done
  }
}
```
- userの一覧取得
```
query {
  users {
    id
    name
  }
}
```
- userの追加
```
mutation {
  createUser(input: { name: "hello"}) {
    id
    name
  }
}
```

## schema, resolver の生成
1. graph/schema/[name].graphqls にスキーマかく
1. `go run github.com/99designs/gqlgen generate`
1. graph/resolver/[name].resolvers.go の後ろの方のMutaionとかのメソッドを消す
1. graph/model/models_gen.go に関連するコードが生成される
1. 生成されたコードを [name].go にうつす(今の所NewUserみたいな生成時に使うstructだけmodels_gen.goにおきっぱ)
1. gqlgen.yml にモデル自動生成をしないようにするため下のように記述
```
models:
  User:
    model: gqlgen-todos/graph/model.User
```
7. `go run github.com/99designs/gqlgen generate`
1. graph/resolver/[name].resolvers.go の後ろの方のMutaionとかのメソッドを消す
1. resolver/[name].resolvers.go に中身を書く