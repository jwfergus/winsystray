package main

import (
	"time"

	"github.com/jwfergus/winsystray"
)

func main() {

	ti, err := winsystray.NewTrayIcon()
	if err != nil {
		panic(err)
	}
	defer ti.Dispose()

	/*
		These can be called as frequently as necessary. Changing the
			icon quickly gives the illusion of animation.
	*/
	ti.SetIconFromFile("icon.ico")
	ti.SetTooltip("おはよう世界！")
	time.Sleep(2 * time.Second)
}
