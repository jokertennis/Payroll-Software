## Payroll-Software
Payroll-Softwareは給与管理システムで利用されるAPIを開発するためのリポジトリである。
現状以下の4つの機能が存在する。
- 認証・認可機能(Basic認証で実装)
- 従業員が自身の今まで受け取った給料明細の一覧を取得する
- 従業員が特定の年月の給料明細を取得する
- 管理者が自身の所属する企業の従業員の給料明細を個別データを利用して作成する。

## ソースコードをビルド・実行・テストするために必要な技術
- docker
- dockercompose
- golang : version1.20.4以上
- VSCode

## golangのインストール方法
[golangのinstallドキュメント](https://go.dev/doc/install)からgolangをローカルにインストールする。インストールはmacの場合/usr/local/goにされるため、pathを通す必要がある。ターミナルを開いて以下のコマンドを実行する。
```
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
```
~/.bashrc の代わりに ~/.zshrc や ~/.bash_profile などの設定ファイルを使っている場合は、適切なファイル名に変更する必要がある。windowsの場合はコントロールパネルを開き、システム > システムの詳細設定 > 環境変数でgolangのインストールされたディレクトリのbinフォルダのパスを追加する必要がある。

## マイグレーションの実行・アプリケーションの実行方法
1. リポジトリをcloneする。(clone済みならこのステップは必要なし)
`git clone https://github.com/jokertennis/Payroll-Software.git`
2. Payroll-Softwareディレクトリ直下で以下のコマンドを実行。
`docker compose up db phpmyadmin`  
以下の状態になれば良い。dbコンテナのstatusがhealthyになっていないといけない。(health: starting)だと待つ必要がある。
```
kenjikawata@KenjinoMacBook-Air Payroll-Software % docker ps
CONTAINER ID   IMAGE            COMMAND                  CREATED         STATUS                   PORTS                               NAMES
976ba39ddb8b   phpmyadmin:5.2   "/docker-entrypoint.…"   2 minutes ago   Up 2 minutes             0.0.0.0:7890->80/tcp                phpmyadmin_container
db6024452d73   mysql:8.0        "docker-entrypoint.s…"   2 minutes ago   Up 2 minutes (healthy)   0.0.0.0:3306->3306/tcp, 33060/tcp   db_container
```
3. 続いてPayroll-Softwareディレクトリ直下でgo run main.goを実行する。(golangのライブラリをローカル環境にインストールする必要がある。[golangのインストール方法](#golangのインストール方法))
```
kenjikawata@KenjinoMacBook-Air Payroll-Software % go run main.go
2023/05/07 11:58:22 Serving swagger at http://[::]:8080
```
4. phpmyadminを開く。[http://localhost:7890](http://localhost:7890)からアクセス可能である。問題なく以下のようにdevelop_dbが作成されていて、データが入っていれば問題ない。
<img width="1440" alt="スクリーンショット 2023-05-07 12 04 16" src="https://user-images.githubusercontent.com/48274379/236655478-251c2f7c-48c6-49a6-a47b-56dbf88f97ea.png">
作成したAPIもcurlで実行可能な状態である。employee向けのapi(employeeがpathの中にあるapi)を実行したいときはemployeeテーブルに登録されているデータのメールアドレスとパスワードを認証ヘッダーに渡して実行する必要がある。administrator向けのapi(administratorがpathの中にあるapi)を実行したいときはadministratorテーブルに登録されているデータのメールアドレスとパスワードを認証ヘッダーに渡して実行する必要がある。

## ローカル環境でテストを実行する方法
テストを実行する前に[マイグレーションの実行・アプリケーションの実行方法](#マイグレーションの実行・アプリケーションの実行方法)をする必要がある。VSCodeでプロジェクトを開き、以下のようなメッセージが右下に表示されたらall installを選択する。
<img width="1512" alt="スクリーンショット 2023-05-07 13 15 26" src="https://user-images.githubusercontent.com/48274379/236659202-0ea9d5c1-760c-4fc5-85b0-d0a5102010b7.png">
そして、vscode再起動後に全テストを実行し、以下のような状態になれば良い。
<img width="1512" alt="スクリーンショット 2023-05-07 13 30 10" src="https://user-images.githubusercontent.com/48274379/236659225-bb9bad27-8998-49f3-b745-99427a8cce49.png">

## 資料に関して
資料に関しては./docsに置いている。  
- Payroll-Softwareのテーブル設計 : [テーブル設計.md](./docs/%E3%83%86%E3%83%BC%E3%83%96%E3%83%AB%E8%A8%AD%E8%A8%88.md)
- ディレクトリ構成 ： [ディレクトリ構成.md](./docs/%E3%83%87%E3%82%A3%E3%83%AC%E3%82%AF%E3%83%88%E3%83%AA%E6%A7%8B%E6%88%90.md)
- ドメインモデル図 : [ドメインモデル図](./docs/%E3%83%89%E3%83%A1%E3%82%A4%E3%83%B3%E3%83%A2%E3%83%87%E3%83%AB%E5%9B%B3/)
- 開発方針 : [開発方針に関して.md](./docs/%E9%96%8B%E7%99%BA%E6%96%B9%E9%87%9D%E3%81%AB%E9%96%A2%E3%81%97%E3%81%A6.md)
- サービスに関して : [サービスに関して.md](./docs/%E3%82%B5%E3%83%BC%E3%83%93%E3%82%B9%E3%81%AB%E9%96%A2%E3%81%97%E3%81%A6.md)

また、APIの一覧を確認する資料に関しては./swagger/swagger.ymlファイルを開き、VSCodeのOpenAPI (Swagger) Editorプラグインをインストールしてpreviewを表示させる必要がある。
<img width="1440" alt="スクリーンショット 2023-05-07 15 49 48" src="https://user-images.githubusercontent.com/48274379/236662406-7e4d4421-0c37-4f52-8140-24ce04f0d73e.png">
