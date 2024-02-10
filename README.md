<img src="https://github.com/traPtitech/game3-front/blob/main/public/favicon.svg" align="right" width="250" alt="game3 logo"/>

# game3-back

[Game³](https://game3.trap.games)のバックエンドです。

フロントエンド: <https://github.com/traPtitech/game3-front>

[Go](https://github.com/golang/go)のフレームワーク[Echo](https://github.com/labstack/echo)を使用して作成しています。

## ドキュメント
- [OpenAPI](https://github.com/traPtitech/game3-back/blob/main/docs/openapi.yaml): APIの定義
- [DB Schema](https://github.com/traPtitech/game3-back/blob/main/docs/db_schema.mmd): データベースの定義

## 開発環境の実行
- [Docker](https://www.docker.com/)と[Docker Compose(ver2.22以上)](https://docs.docker.com/compose/)が必要です。
- `docker compose watch`でアプリが起動し、以下にアクセスできます。
  - <http://localhost:8080/api/ping> (API)
  - <http://localhost:8081/> (Adminer: DBの管理画面)
- 認証部分を動かすためには`.env`ファイルで環境変数を定義する必要があります。
  - <https://ns.trap.jp/apps/0765eb937f27fdc377eec1/settings/envVars> を参考に、
  - ```
    DISCORD_CLIENT_ID=NeoShowcaseと同じ
    DISCORD_CLIENT_SECRET=NeoShowcaseと同じ
    DISCORD_CLIENT_REDIRECT_URI="http://localhost:8080/api/auth/callback"
    DISCORD_SERVER_ID=NeoShowcaseと同じ
    DISCORD_BOT_TOKEN=NeoShowcaseと同じ
    DISCORD_NORMAL_PARTICIPANT_ROLE_ID=NeoShowcaseと同じ
    DISCORD_EXHIBITOR_PARTICIPANT_ROLE_ID=NeoShowcaseと同じ
    ``` 
    を`.env`に記載すると正常に全てが動作します。

## 構成
**TL;DR**: 大体`main.go`→`internal/handler`(ルーティング)→`internal/repository`(DBの操作)
- https://github.com/ras0q/go-backend-template をベースに作成。
- `main.go`: エントリーポイント
  - 依存ライブラリの初期化など最低限の処理のみを書く
  - ルーティングの設定は`./internal/handler/handler.go`に書く
  - 肥大化しそうなら`./internal/infrastructure/{pkgname}`を作って外部ライブラリの初期化処理を書くのもアリ
- `internal/`: アプリ本体の主実装
  - Tips: Goの仕様で`internal`パッケージは他プロジェクトから参照できない (<https://go.dev/doc/go1.4#internalpackages>)
  - `api/`: 外部APIのラッパー
    - 外部サービス（例：Discord API）へのリクエストを抽象化したメソッドを提供します。
  - `domain/``: ドメインモデル
    - アプリケーションのビジネスロジックをカプセル化したドメインオブジェクトを定義します。
    - 今回は、ほとんどoapi-codegenが生成したmodelを使っているので、出番は少しだけ。
  - `handler/`: ルーティング
    - 飛んできたリクエストを裁いてレスポンスを生成する
    - DBアクセスは`repository/`で実装したメソッドを呼び出す
    - Tips: リクエストのバリデーションがしたい場合は↓のどちらかを使うと良い
      - [go-playground/validator](https://github.com/go-playground/validator)でタグベースのバリデーションをする
      - [go-ozzo/ozzo-validation](https://github.com/go-ozzo/ozzo-validation)でコードベースのバリデーションをする
  - `migration/`: DBマイグレーション
    - DBのスキーマを定義する
    - Tips: マイグレーションツールは[pressly/goose](https://github.com/pressly/goose)を使っている
    - 初期化スキーマは`1_schema.sql`に記述し、運用開始後のスキーマ定義変更等は`2_add_user_age.sql`のように連番を振って記述する
      - Tips: Goでは1.16から[embed](https://pkg.go.dev/embed)パッケージを使ってバイナリにファイルを文字列として埋め込むことができる
  - `repository/`: DBアクセス
    - DBへのアクセス処理
      - 引数のバリデーションは`handler/`に任せる
  - `pkg/`: 汎用パッケージ
    - 複数パッケージから使いまわせるようにする
    - 例: `pkg/config/`: アプリ・DBの設定
    - Tips: 外部にパッケージを公開したい場合は`internal/`の外に出しても良い
    - 定数やenumや使いまわす関数をここで定義している。
- `openapi/`: oapi-codegenがopenapi.yamlから生成したモデルとサーバー
  - https://github.com/deepmap/oapi-codegen を用いて、OpenAPIの定義からGoのコードを生成している。以下のコードで生成できる。
  - ```
    cd docs
    oapi-codegen --config=models.cfg.yaml openapi.yaml
    oapi-codegen --config=server.cfg.yaml openapi.yaml
    ```
- `integration/`: 結合テスト
  - 現状テストを書いていないので、使っていない。//TODO
  - `internal/`の実装から実際にデータが取得できるかテストする
  - DBの立ち上げには[ory/dockertest](https://github.com/ory/dockertest)を使っている
  - 短期開発段階では時間があれば書く程度で良い
  - Tips: 外部サービス(traQ, Twitterなど)へのアクセスが発生する場合は[golang/mock](https://github.com/golang/mock)などを使ってモック(テスト用処理)を作ると良い

## 長期開発に向けた改善点
- 単体テスト・結合テストのカバレッジを上げる
  - カバレッジの可視化には[Codecov](https://codecov.io)(traPだと主流)や[Coveralls](https://coveralls.io)が便利
- ログの出力を整備する
  - ロギングライブラリは好みに合ったものを使うと良い

## TODO
- テストを書く
- DBの中に保存している画像を、オブジェクトストレージに移行する。
