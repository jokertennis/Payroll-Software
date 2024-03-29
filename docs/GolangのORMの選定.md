## 概要
GolangのORMで何を利用するか決める。  
前提として、Golangには標準ライブラリであるdatabase/sqlはアプリケーションのコア機能とは関係ない部分で実装が増えてしまう(例えば、Scan関数で構造体に直接マッピングできないため、DBからデータを取得してGolangの構造体にマッピングする際のコードの量が多くなってしまう)ため、選択肢から排除する方針とする。

## GolangのORMの特徴の整理
以下にそれぞれのORMの特徴を整理する。(golangには他のormも幾つか存在するが、gormとsqlboilerのどちらかで絞る方針とする。)

gormの特徴
- code-firstの設計である。gormにはmigrationの機能がある。migrationを行う際は、事前に構造体を定義し、構造体を指定することで、構造体に対応するテーブルが作成される。例えば、Userという構造体を作ったとして、migration時にUserを指定することでusersテーブルが作成される。[参考文献](https://gorm.io/ja_JP/docs/migration.html)

sqlboilerの特徴
- database-firstの設計である。実装の前にデータベースにあらかじめテーブルが存在している必要がある。sqlboilerにはmigrationの機能がないため、他のmigrationを行うためのツールの導入が必要である。
- データベース内のテーブルからコードが自動生成される。

| ORM  | 長所 | 短所
| ------------- | ------------- | ------------- |
| gorm  | ・構造体を定義するだけで ORM としての機能とマイグレーションの両方が使える。 | ・実行時にリフレクションが内部で行われているため、sqlboilerと比較すると処理が遅くなる。<br>・フィールドに渡す型が誤っていてもコンパイルが通ってしまう。<br>・処理がブラックボックスとなる部分が多いため、予期せぬ挙動が起こりうる。(GORMのv1では主キーがゼロ値の構造体を渡すと更新/削除時のWhere句が生成されないので、引数などを適切に検査しないとテーブルの全データが削除される不具合が入り込んだりもする。v2ではこの問題はない。)<br>リフレクションを使ったORMライブラリでは構造体が実行時に生成されるので、フィールドに渡す型に誤りがあった場合、コンパイル時に気づくことができない。 |
| sqlboiler  | ・自動生成されたコードを利用可能。挙動を把握しやすいため、ブラックボックスとなる部分が少ない。<br>・SQLBoilerは静的に型付けされており、実行時にリフレクションを使う必要がないため、[高速](https://github.com/volatiletech/sqlboiler#benchmarks)に動作しする。<br>・フィールドに渡す型に誤りがあった時、コンパイル時にエラーが出るので、すぐに誤りを修正することできる。  | ・マイグレーション機能はない。<br>・自動生成コードのメンテナンスコストが増える。 |

## 選定結果
sqlboilerを採用することにした。理由として大きいのが以下の3点である。
- gormの予期せぬ挙動が怖い
- 自動生成コードを利用可能なため、コードが属人化しにくい。
- リフレクションがされないことでパフォーマンスが良くなりやすい点やコンパイル時のエラー出力の面でsqlboilerの方がより良いと捉えたため

短所としてマイグレーション機能がないという点があるが、その点は新しいマイグレーションツールを導入すれば解決できるため、問題ないと判断した。自動生成コードのメンテナンスコストが増えてしまう点に関しては、受け入れるべきコストと捉えることにした。

## 参考文献
- [GoにおけるORMと、SQLBoiler入門マニュアル](https://zenn.dev/gami/articles/0fb2cf8b36aa09)
- [What do you think is the best ORM or SQL builder for Go?](https://www.reddit.com/r/golang/comments/t5l7uu/what_do_you_think_is_the_best_orm_or_sql_builder/)
- [ORM を GORM から SQLBoiler に変えた理由](https://tech-blog.optim.co.jp/entry/2021/03/22/100000)
