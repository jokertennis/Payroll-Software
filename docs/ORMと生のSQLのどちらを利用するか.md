## 概要
ORM or 生のSQLのどちらを利用するかまとめる。

## 結論
川田としてはormを利用すると良いと考えた。ただし、ormでは実現不可能な処理の場合は、その処理のみsqlを利用することを検討する。基本的にパフォーマンスの悪さを発見した場合、詳細設計やテーブル設計を見直すことで解決を試みる方針とする。

## 長所短所の整理
以下にormとsqlの長所短所を整理する。
| 方法  | 長所 | 短所
| ------------- | ------------- | ------------- |
| sql  | ・sqlに関して深い知見がある場合、パフォーマンスはormより高い。 | ・sqlが複雑になると保守性が下がる。<br>・再利用性が低い。<br>・アプリケーションとデータベースが密結合になる。 |
| orm  | ・一度実装したormを再利用しやすい。<br>・アプリケーションとデータベースが疎結合になる。  | ・ormがsqlよりも抽象レベルが高いため、sqlと比較すると複雑である。<br>・ormの使い方を学ぶコストが発生する。<br>・ormはsqlと比較すると遅くなるため、大規模なデータ量を処理する上で時間がかかる。<br>・ormだとクエリチューニングが難しい。 |

川田としては、ormを利用すると良いと考えた。最も大きな理由は再利用性が高い(保守性が高い)ことである。プロジェクトの内容が給与管理システムであり、複雑なビジネスロジックが想定されるため、保守性の高さはより重要視しなければいけない。ormの短所に関してそれぞれ対処法を以下にまとめる。  
・ormがsqlよりも抽象レベルが高いため、sqlと比較すると複雑である。  
→必要なコストと受け止める。  
・ormの使い方を学ぶコストが発生する。  
→必要なコストと受け止める。  
・ormはsqlと比較すると遅くなるため、大規模なデータ量を処理する上で時間がかかる。  
→sqlでの改善は最終手段として利用し、詳細設計やテーブル設計を見直すことで解決を試みる。  
・ormだとクエリチューニングが難しい。   
→必要なコストと受け止める。もしもormでクエリチューニングが不可能な場合、その処理のみ生のsqlを利用することを検討する。

## 参考文献
・[【プログラミング学習】 ORMの真のメリット](https://note.com/kazunori_kuroco/n/na2ce4e208137)  
・[Comparing ORM vs SQL: What To Know](https://axiomq.com/blog/comparing-orm-vs-sql-what-to-know/)  
・[Using an ORM or plain SQL? [closed]](https://stackoverflow.com/questions/494816/using-an-orm-or-plain-sql)