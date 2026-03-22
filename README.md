# TinyGo drivers hub75 sample

## 使用する機材・ソフト

* 機材
  * LED マトリクス
    * https://www.waveshare.com/rgb-matrix-p3-64x32.htm
  * マイコンボード
    * Wio Terminal(※)
      * https://www.switch-science.com/products/6360
* ソフトウェア
  * OS
    * windows11
  * TinyGo
    * tinygo version 0.39.0 windows/amd64 (using go version go1.24.5 and LLVM version 19.1.2)

## 注意点

上記の(※)について、本ソースコードは Wio Terminal 用の物となっています。
他のマイコンボードを使用する場合は、以下の作業が必要になります。

* 必要に応じてソースコードで適切な GPIO, SPI のピンを指定する
* ターゲットとして、自分の使用するマイコンボードを指定して build する

以下、Raspberry Pi Pico を使用する場合の例です。

* GPIO は `machine.DX`ではなく、`machine.GPIOX` という形で指定します(`X` = 任意の数字)
  * https://tinygo.org/docs/reference/microcontrollers/pico/#pins
* build する際は、ターゲットに `pico` を指定します
  * https://tinygo.org/docs/reference/microcontrollers/pico/#cli-flashing

## ビルド方法

Makefile 参照

## 機材接続情報

Data input | 接続先
--- | ---
R1 | マイコンボード:SPI_MOSI
G1 | Data output:R2
B1 | Data output:G2
GND | 接続不要(二つあるため、どちらか or 両方をマイコンボードに繋ぐ)
R2 | Data output:R1
G2 | Data output:G1
B2 | Data output:B1
E | 接続不要
A | マイコンボード:D3
B | マイコンボード:D4
C | マイコンボード:D5
D | マイコンボード:D6
CLK | マイコンボード:SPI_SCLK
LAT/STB | マイコンボード:D1
OE | マイコンボード:D2
GND | マイコンボード:GND

![Pin layout diagram](.\images\pin_layout.png)

## 動作結果

![Result](.\images\result.png)
