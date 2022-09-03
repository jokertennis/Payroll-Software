## サービス名
給与管理ソフト

## サービスの内容
本サービスは企業向けの給料管理サービスである。現状、企業の規模は想定していないが、サービスをより良くしていくためにも検討していく必要がある。

## サービスの目的
本サービスを利用していただく目的を以下に記載する。
* 企業に所属する人事の方々のbackoffice業務の効率化
* 社員の方々の自身の給料管理の支援

## サービスの機能
サービスの機能に関して、共通する機能、管理者向けの機能、一般社員向けの機能の３つにグループ分けして以下にまとめる。

### 共通する機能
* 認証機能(管理者としてのログインと一般社員としてのログインは別で考慮しなければいけない。)

### 管理者向けの機能
* 指定された年と月の全社員の給料一覧
* 特定の社員の今まで支給された給料一覧
* 特定の社員の指定された年と月の給料詳細(給料明細が表示される)
* 特定の社員の指定された年と月の給料明細を作成
* 本サービスの決済機能
* csvを利用して複数社員の給料一括登録

### 一般社員向けの機能
* 今まで支給された給料一覧
* 特定の月の給料詳細

## 開発方針
開発方針を以下に記載する。  

### 導入するソフトウェアアーキテクチャ
クリーンアーキテクチャ,DDD

選定理由：クリーンアーキテクチャを導入する理由は、責務をより細分化させることでソフトウェアの保守性を高めるためである。クリーンアーキテクチャの他にもオニオンアーキテクチャ、ヘキサゴナルアーキテクチャなどが存在するが、現状世の中で多くの開発チームがクリーンアーキテクチャを利用していることから、クリーンアーキテクチャを採用することにした。DDDを導入する理由は、ドメインオブジェクトのロジックを集約させることで保守性を高めるためである。

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

### データベース
データベースはRDBのMySQLを利用する。

### テストに関して
APIのテストに関してはc1カバレッジ100%を目指す。エンドポイントからのミディアムテストで正常系・異常系・認証系のテストを行うこととする。また、実装前に行うべきテストをまとめてから、テストを通るような実装を行うべきである。理由はテスト作成後にテスト失敗を確認し、実装を変更するという作業をなるべく無くすためである。自動e2eテスト、手動e2eテストも行う予定であるが、詳細は未定である。

### その他
APIの設計思想はREST APIを採用する。また、OpenAPI(RESTful APIを記述するためのフォーマット)の作成にはSwaggerを利用する。
