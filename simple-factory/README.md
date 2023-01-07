# Simple Facotryパターンの目的
- オブジェクトの生成とオブジェクトが持つ機能を分離する
- オブジェクトの生成が容易になる

## 登場人物
# Simple Factory
単純ファクトリー （simple factory） パターン は、 巨大な条件分岐のついた生成メソッドを持つクラスのことです。 メソッドのパラメーターに基づき、 どのプロダクトのクラスをインスタンス化すべきかを決め、 生成し、 返します。
## 登場人物
- Product: 生成するオブジェクトの抽象
- ConcreateProduct: 生成するオブジェクトの実態
- Creater: オブジェクトを生成する機構

## 処理の概要
- Simple Factoryを使うクライアントは、どのプロダクトを生成するか判別するパラーメータを、生成メソッドに渡す
- 生成メソッドはProductの具象、ConcreateProductを返す
- クライアント側は抽象(Product)を起点にロジックを組めるので、具象1つ1つの処理を必要としない。

## Simple Factoryを使って得られるメリット
Simple Factoryを実装すると、データが抽象と具象に分かれる。それによりどのようなメリットがあるか考える。
Simple Facotry化の直接的なメリットについては、Simple Facotryパターンの目的を参照

- 君はGunマニア向けのサイトの運営者だ
- アプリをリリースする前に事前にユーザーヒアリングを実施した結果、登録したいGunの種類は1つでAK47のみだったとしよう
- 君は以下のような構成で、gunのインスタンス、gunを表示する関数を作成した

### simple fatoryを使う前の実装

```golang
// Gun構造体の定義
type Gun struct {
  name string
  power int
}

func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) setName(name string) {
	g.name = name
}

func (g *Gun) getPower() int {
	return g.power
}

func (g *Gun) setPower(power int) {
	g.power = power
}

// クライアントコード
g := Gun{
  name: "AK47",
  power: 123,
}

printDetails(g)

func printDetails(g Gun) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}
```

時は流れ、たくさんのマニアがこのサイトに登録するようになった。
当初想定していたGunはAK47以外にも登録したいGunが出てきた。
じゃあ、以下のようにGun構造体を新たに作成すればよい！

```golang
g := Gun{
  name: "musket",
  power: 123,
}
printDetails(g)
```
これでAK47と同じように表示できるようになった。

ここでユーザーから要望があった。
「ak47のnameには + をつけて欲しい表示して欲しい！！」

さあ、どう実装する。安易に実装するとしたら、if文をクライアントコードに追加することになるだろう
```golang

func printDetails(g Gun) {
  if g.name == "musket" {
    fmt.Printf("Name: +%d", g.getName())
  } else {
    fmt.Printf("Name: %d", g.getName())
  }
}
```
これは良くない実装だ。クライアントのコードはGunについて知りすぎている。
Gunの種類が増えるたびに条件分岐が足されていき、見通しの悪いコードになってしまう。

Tell, Don`t Ask!を思いだそう。
クライアントコードのifでオブジェクトの中身をaskせず、必要な情報をただ伝えるだけの構成に変更しよう

### 変更後
```golang
package main

type AK47 struct {
	Gun
}

type iGun inteface {
  getName() string
  setName(name string)
  getPower() int
  setPower(power int)
}

func newAK47() iGun {
  return &AK47{Gun: Gun{name: "AK47", power: 123}}
}

// メソッドを上書きする
func (g *AK47) getName() string {
	return "+" + g.name
}

func printDetails(g iGun) { // interfaceに変更
  fmt.Printf("Name: +%d", g.getName())
}
```
## Simple Factoryのデメリット
- 具象の種類が増えた時に、生成メソッドに条件を足す必要がある
- 生成メソッドには、既存の具象、新規の具象を生成するコードを書く必要がある。
- 生成メソッドが1つであるため、「開放閉鎖の原則」を違反している、という見方ができる。(※「開放閉鎖の原則」はプロダクトの新しい型を導入しても、 既存のクライアント・コードの機能に影響しないことを表す)
- 具象の数に比例して生成メソッドはfatになっていく。
# Facotry Method
- 具象型ごとにFacotryを作成することで、Simple Factoryのデメリットをカバーする
https://stackoverflow.com/questions/49712313/golang-factory-method の一番上の回答がFactory Methodっぽいことをやっている
