package main

import (
	"image/color"
	"machine"

	"tinygo.org/x/drivers/hub75"
)

func main() {
	// SPI 通信をする際に必要となる設定
	spi := machine.SPI0
	spi.Configure(machine.SPIConfig{
		SCK: machine.SPI0_SCK_PIN,
		SDO: machine.SPI0_SDO_PIN,
		SDI: machine.SPI0_SDI_PIN,
	})

	// HUB-75通信をする際に必要となる設定
	// oe,lat,a～dピンには、どの GPIOピンを使用しても良いです(↓では D1～D6 を使用)
	h75 := hub75.New(spi, machine.D1, machine.D2, machine.D3,
		machine.D4, machine.D5, machine.D6)
	h75_config := hub75.Config{
		Width:      64,
		Height:     32,
		FastUpdate: true,
	}
	h75.Configure(h75_config)

	// ディスプレイの描画準備
	h75.ClearDisplay()
	h75.FlushDisplay()
	w, h := h75.Size()
	for y := int16(0); y < h; y++ {
		for x := int16(0); x < w; x++ {
			// 20列目 or 20列目だけは黒(点灯しない)設定にしています。
			if x == 20 || y == 20 {
				h75.SetPixel(x, y, color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xFF})
			} else {
				h75.SetPixel(x, y, color.RGBA{R: 0xFF, G: 0x80, B: 0x00, A: 0xFF})
			}
		}
	}

	// ディスプレイの描画処理
	for {
		err := h75.Display()
		if err != nil {
			println(err.Error())
		}
	}
}
