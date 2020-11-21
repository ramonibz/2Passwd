package main

import (
	"2Passwd/cfg"
	"2Passwd/icons"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/getlantern/systray"
	"github.com/martinlindhe/notify"
	"github.com/ramonibz/totp/totp"
	"os"
	"time"
)

func main() {

	systray.Run(onReady, onExit)
}

func onReady() {
	go updateSystrayIcon()
	config := cfg.Configuration()
	buildItems(config)

}

func buildItems(config *cfg.Config) {
	if len(config.Accounts) > 0 {
		//Starred items
		for _, account := range config.Accounts {
			if account.Star {
				accountItem := systray.AddMenuItem(account.Name, "")
				go updateItem(accountItem, account)
			}
		}

		othersItem := systray.AddMenuItem("Others", "")

		//Other Account
		for _, account := range config.Accounts {
			if !account.Star {
				accountItem := othersItem.AddSubMenuItem(account.Name, "")
				go updateItem(accountItem, account)
			}
		}
	} else {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}

		_ = systray.AddMenuItem("No accounts defined. configure them in " + homeDir + "/.2Passwd/config.yaml", "")
	}


	//Quit Item
	mQuitOrig := systray.AddMenuItem("Quit", "Quit 2Passwd")
	go func() {
		<-mQuitOrig.ClickedCh
		fmt.Println("Requesting quit")
		systray.Quit()
		fmt.Println("Finished quitting")
	}()
}

func updateItem(item *systray.MenuItem, account cfg.Account){
	go func() {
		for {
			<-item.ClickedCh
			_ = clipboard.WriteAll(totp.GetTotpCode(account.Seed))
			notify.Notify("2Passwd", account.Name, "Copied to clipboard", "logo.icns")
		}
	}()
}

func updateSystrayIcon() {
	for {
		second := 60 - time.Now().Second()
		if second > 30 {
			second = second - 30
		}
		switch second {
		case 1:
			systray.SetTemplateIcon(icons.Logo1, icons.Logo1)
		case 2:
			systray.SetTemplateIcon(icons.Logo2, icons.Logo2)
		case 3:
			systray.SetTemplateIcon(icons.Logo3, icons.Logo3)
		case 4:
			systray.SetTemplateIcon(icons.Logo4, icons.Logo4)
		case 5:
			systray.SetTemplateIcon(icons.Logo5, icons.Logo5)
		case 6:
			systray.SetTemplateIcon(icons.Logo6, icons.Logo6)
		case 7:
			systray.SetTemplateIcon(icons.Logo7	, icons.Logo7)
		case 8:
			systray.SetTemplateIcon(icons.Logo8, icons.Logo8)
		case 9:
			systray.SetTemplateIcon(icons.Logo9, icons.Logo9)
		case 10:
			systray.SetTemplateIcon(icons.Logo10, icons.Logo10)
		case 11:
			systray.SetTemplateIcon(icons.Logo11, icons.Logo11)
		case 12:
			systray.SetTemplateIcon(icons.Logo12, icons.Logo12)
		case 13:
			systray.SetTemplateIcon(icons.Logo13, icons.Logo13)
		case 14:
			systray.SetTemplateIcon(icons.Logo14, icons.Logo14)
		case 15:
			systray.SetTemplateIcon(icons.Logo15, icons.Logo15)
		case 16:
			systray.SetTemplateIcon(icons.Logo16, icons.Logo16)
		case 17:
			systray.SetTemplateIcon(icons.Logo17, icons.Logo17)
		case 18:
			systray.SetTemplateIcon(icons.Logo18, icons.Logo18)
		case 19:
			systray.SetTemplateIcon(icons.Logo19, icons.Logo19)
		case 20:
			systray.SetTemplateIcon(icons.Logo20, icons.Logo20)
		case 21:
			systray.SetTemplateIcon(icons.Logo21, icons.Logo21)
		case 22:
			systray.SetTemplateIcon(icons.Logo22, icons.Logo22)
		case 23:
			systray.SetTemplateIcon(icons.Logo23, icons.Logo23)
		case 24:
			systray.SetTemplateIcon(icons.Logo24, icons.Logo24)
		case 25:
			systray.SetTemplateIcon(icons.Logo25, icons.Logo25)
		case 26:
			systray.SetTemplateIcon(icons.Logo26, icons.Logo26)
		case 27:
			systray.SetTemplateIcon(icons.Logo27, icons.Logo27)
		case 28:
			systray.SetTemplateIcon(icons.Logo28, icons.Logo28)
		case 29:
			systray.SetTemplateIcon(icons.Logo29, icons.Logo29)
		case 30:
			systray.SetTemplateIcon(icons.Logo30, icons.Logo30)
		default:
			systray.SetTemplateIcon(icons.Logo30, icons.Logo30)
		}
		time.Sleep(1*time.Second)
	}
}

func onExit() {
}
