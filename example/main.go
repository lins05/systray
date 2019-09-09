package main

import (
	"fmt"
	"net/http"
	// "io/ioutil"
	// "time"

	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
	// "github.com/skratchdot/open-golang/open"
)

func main() {
	onExit := func() {
		// fmt.Println("Starting onExit")
		// now := time.Now()
		// ioutil.WriteFile(fmt.Sprintf(`on_exit_%d.txt`, now.UnixNano()), []byte(now.String()), 0644)
		fmt.Println("Finished onExit")
	}
	// Should be called at the very beginning of main().
	systray.Run(onReady, onExit)
}

var syncing bool

func toggleIcon(normal bool) {
	if normal {
		systray.SetIcon(icon.IconDataForNormal)
	} else {
		systray.SetIcon(icon.IconDataForSyncing)
	}
}


func onReady() {
	systray.SetIcon(icon.IconDataForNormal)
	// systray.SetTitle("Awesome App")
	systray.SetTooltip("Lantern")
	mQuitOrig := systray.AddMenuItem("Quit", "Quit the whole app")
	go func() {
		<-mQuitOrig.ClickedCh
		fmt.Println("Requesting quit")
		systray.Quit()
		fmt.Println("Finished quitting")
	}()

	// We can manipulate the systray in other goroutines
	go func() {
		// systray.SetTitle("RSYNC")
		mChange := systray.AddMenuItem("Change Me", "Change Me")
		// mChecked := systray.AddMenuItem("Unchecked", "Check Me")
		// mEnabled := systray.AddMenuItem("Enabled", "Enabled")
		// systray.AddMenuItem("Ignored", "Ignored")
		// mUrl := systray.AddMenuItem("Open Lantern.org", "my home")
		// mQuit := systray.AddMenuItem("退出", "Quit the whole app")

		// // Sets the icon of a menu item. Only available on Mac.
		// mQuit.SetIcon(icon.Data)

		systray.AddSeparator()
		// mToggle := systray.AddMenuItem("Toggle", "Toggle the Quit button")
		// shown := true
		for {
			select {
			case <-mChange.ClickedCh:
				mChange.SetTitle("I've Changed")
				toggleIcon(!syncing)
				syncing = !syncing
			}
			// case <-mChecked.ClickedCh:
			// 	if mChecked.Checked() {
			// 		mChecked.Uncheck()
			// 		mChecked.SetTitle("Unchecked")
			// 	} else {
			// 		mChecked.Check()
			// 		mChecked.SetTitle("Checked")
			// 	}
			// case <-mEnabled.ClickedCh:
			// 	mEnabled.SetTitle("Disabled")
			// 	mEnabled.Disable()
			// case <-mUrl.ClickedCh:
			// 	open.Run("https://www.getlantern.org")
			// case <-mToggle.ClickedCh:
			// 	if shown {
			// 		mQuitOrig.Hide()
			// 		mEnabled.Hide()
			// 		shown = false
			// 	} else {
			// 		mQuitOrig.Show()
			// 		mEnabled.Show()
			// 		shown = true
			// 	}
			// case <-mQuit.ClickedCh:
			// 	systray.Quit()
			// 	fmt.Println("Quit2 now...")
			// 	return
			// }
		}
	}()

    http.HandleFunc("/normal", func(w http.ResponseWriter, r *http.Request) {
		toggleIcon(true)
	})
    http.HandleFunc("/syncing", func(w http.ResponseWriter, r *http.Request) {
		toggleIcon(false)
	})

    go http.ListenAndServe(":11111", nil)
}
