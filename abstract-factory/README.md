# Abstract Facotry(抽象ファクトリー)
## 概念
ref: https://refactoring.guru/ja/design-patterns/abstract-factory
## サンプルコード
https://refactoring.guru/ja/design-patterns/abstract-factory/go/example#example-0
サンプルコードを実行する時は、```go run .```としてください

## いつ使うのか
 > 関連する製品の集まりである様々な変種に対応したいが、 製品の具象クラス （それは設計段階では未知かもしれません） に依存させたくない場合に、 Abstract Factory を使用します

## 実装方法の補足
1. 明確に区別できる製品の型と、 製品型の変種の表を作ります。
  サンプルコードでいうと、製品の型はShirt, Showとなり、製品型はaddidas, nikeとなる

2. 全部の製品型について、 抽象製品インターフェースを宣言します。 そして、 全インターフェースを実装した製品の具象クラスを作ります。

3. 全抽象製品型に対する生成メソッドを含んだ抽象ファクトリー・インターフェースを宣言します。
  ```sports_factory_interface.go```のこと

4. 各変種ごとに、 具象ファクトリー・クラスを実装します。

5. アプリのどこかに、 ファクトリー初期化コードを追加します。 そこでアプリケーションの構成か環境設定に従って具象ファクトリー・クラスのインスタンスを作成します。 製品を構築するすべてのクラスに、 このファクトリー・オブジェクトを渡します。

6. コードをスキャンして、 製品クラスのコンストラクターへの直接の呼び出しを探します。 これらの呼び出しを、 ファクトリー・オブジェクトに対する適切な作成メソッドの呼び出しと置き換えます。

## 注意点
ブランド間で必要な製品のリストは共通である時に使えるパターン、と言う認識。
Addidasの製品にのみCapを追加、Nikeには追加しないとなった場合、対象性が損なわれそう...
現実的に考えると、商品を販売するサイト側で、全てセットアップを選択できるのであれば、対称性が損なわれることはなさそう。
外部のブランドと提携して、うちのブランドだけCapをセットで販売してくれませんか..?となったら厳しそう。
明確に区別できる製品の型と、 製品型の変種を実装側でコントロールできる場合に使うべき(https://refactoring.guru/ja/design-patterns/abstract-factory の疑似コードはサンプルにピッタリだ)
