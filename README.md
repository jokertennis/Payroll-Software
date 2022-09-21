## 開発方針
開発方針を以下に記載する。  

### 導入するソフトウェアアーキテクチャ
オニオンアーキテクチャ,DDD

選定理由：オニオンアーキテクチャを導入する理由は、高凝集・低結合なソフトウェアを目指し、ソフトウェアの保守性を高めるためである。オニオンアーキテクチャの他にもクリーンアーキテクチャ、ヘキサゴナルアーキテクチャなどが存在するが、川田がオニオンアーキテクチャに慣れているため、オニオンアーキテクチャを採用することにした。DDDを導入する理由は、ドメインオブジェクトのロジックを集約させることで保守性を高めるためである。

### API開発で利用する言語・webフレームワーク
golang,[iris](https://github.com/kataras/iris)  

golangを利用する理由は、以下の理由である。
* 型宣言が可能なため
* 実行速度が基本的に速い。(並列処理の実装が比較的容易であるため、パフォーマンス改善もしやすい。)
* docker imageが軽くしやすい。
* golangがシンプルな言語であるため、コードが属人化しにくい。  
  
irisを利用する理由は、以下の理由である。
* 依存性の注入が可能
* ドキュメントが豊富
* 他のgolangのwebフレームワークと比較して高速である。

### 利用するデータベース,sqlライブラリ,マイグレーションツール
データベースはRDBのMySQLを利用する。sqlライブラリはgolangの標準ライブラリである[database/sql](https://pkg.go.dev/database/sql)を利用する。開発を進めていくうちに不便さを感じた場合、ormやサードパーティツールの導入を検討する。マイグレーションツールは[golang-migrate](https://github.com/golang-migrate/migrate)を利用する。

### OpenAPI自動生成ツール
OpenAPI自動生成ツールはgolangのライブラリであるgo-swaggerを利用する方針である。

### テストに関して
APIのテストに関してはc1カバレッジ100%を目指す。エンドポイントからの単体テストで正常系・異常系・認証系のテストを行うこととする。また、実装前に行うべきテストをまとめてから、テストを通るような実装を行うべきである。理由はテスト作成後にテスト失敗を確認し、実装を変更するという作業をなるべく無くすためである。シナリオテストに関しては、ユースケースを満たすAPIの実装や単体テストを全て作成後に作成することとする。

### MR・issueに関して
MRの粒度は以下の方針にする。
* 特定のAPIのIFと実装・正常系テスト・異常系テスト・認証系テストの作成
* シナリオテストの作成

追跡性を考慮してissueとMRの粒度は一致させる。また、タスクの管理はgithubのissueを利用する方針とする。

### ブランチ運用に関して
ブランチ運用は以下の通りである。
* developブランチ：開発用ブランチ
* mainブランチ：本番環境の役割を持つブランチ。定期的にdevelopブランチからmergeする。

タスクに取り組む際はdevelopブランチからブランチを切って取り組む。ブランチ名に関しては、追跡性を考慮して「issues/issueの番号」を推奨する。

### その他
APIの設計思想はREST APIを採用する。また、できる限りマイクロサービスアーキテクチャを採用する。また、golangには[google/wire](https://github.com/google/wire)のようなDIツール(injectorの作成を行なってくれる。)がいくつか存在するが、一旦手動で行う方針とし、運用していく上で必要性を感じた場合に導入を検討する。インスタンスの生成方法はDI(Dependency Injection)とFactory Patternの2つに
分けられるが、インスタンスの柔軟性を考慮して極力DIを利用する方針とする。開発・運用していく中でFactory Patternの方が良いと判断した場合は、Factory Patternに一部のインスタンスの生成方法を変更する。DIの方法はConstructor Injectionで統一する。

### 参考文献
* [Dependency Injection vs Factory Pattern](https://stackoverflow.com/questions/557742/dependency-injection-vs-factory-pattern)
* [Go言語とDependency Injection](https://blog.recruit.co.jp/rtc/2017/12/11/go_dependency_injection/)
