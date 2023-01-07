# Factory Method
共通のインターフェースを持つ異なる具象型を生成する関数を通して、オブジェクトを作成する

# 実例
- 君はGunマニア向けのサイトの運営者だ
- アプリをリリースする前に事前にユーザーヒアリングを実施した結果、登録したいGunの種類は1つでAK47のみだったとしよう
- 君は以下のような構成で、gunのインスタンス、gunを表示する関数を作成した

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
当初想定していたGunはAK47以外にも登録したいコードが出てきた。
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
「musketのpowerには + をつけて欲しい表示して欲しい！！」

さあ、どう実装する。安易に実装するとしたら、if文をクライアントコードに追加することになるだろう
```golang

func printDetails(g Gun) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
  if g.name == "musket" {
    fmt.Printf("Power: +%d", g.getPower())
  } else {
    fmt.Printf("Power: %d", g.getPower())
  }
	fmt.Println()
}
```
これは良くない実装だ。クライアントのコードはGunについて知りすぎている。
Gunの種類が増えるたびに条件分岐が足されていき、見通しの悪いコードになってしまう。

Tell, Don`t Ask!を思いだそう。(ifでオブジェクトの中身をaskしている、ただ伝えるだけの構成にしよう)


# Abstract Facotryとの使い分け
