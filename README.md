# winsystray
 Extremely simple Go windows system tray. 
 
 
 98% of code taken from hallazzang (Super thanks to: https://github.com/hallazzang/go-tray-icons-tutorial) and moved out into a simpler interface. Added an intermediary function taken from https://github.com/getlantern/systray to allow []byte icons in addition to regular icon files.
 
 

 Worth pointing out this library: https://github.com/getlantern/systray which currently has some issue w/ windows systray stuff randomly crashing, otherwise I would have just used that. https://github.com/getlantern/systray/issues/148
 
 ## example usage
 ```
import "github.com/jwfergus/winsystray"
 
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
  
  // Do other stuff
}
```
