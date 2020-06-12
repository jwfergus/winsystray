# winsystray
 Extremely simple Go windows system tray. 
 
 
 All code taken from hallazzang (below) and moved out into a simpler interface.
 
 Thanks to: https://github.com/hallazzang/go-tray-icons-tutorial

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
  ti.SetIcon("icon.ico")
  ti.SetTooltip("おはよう世界！")
  
  // Do other stuff
}
```
