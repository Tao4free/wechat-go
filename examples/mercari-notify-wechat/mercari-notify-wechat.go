package main

import (
	"log"
	// "github.com/tao4free/wechat-go/plugins/wxweb/cleaner"
	"github.com/tao4free/wechat-go/plugins/wxweb/config"
	// "github.com/tao4free/wechat-go/plugins/wxweb/faceplusplus"
	// "github.com/tao4free/wechat-go/plugins/wxweb/forwarder"
	// "github.com/tao4free/wechat-go/plugins/wxweb/gifer"
	// "github.com/tao4free/wechat-go/plugins/wxweb/joker"
	// "github.com/tao4free/wechat-go/plugins/wxweb/laosj"
	"github.com/tao4free/wechat-go/plugins/wxweb/replier"
	"github.com/tao4free/wechat-go/plugins/wxweb/repeater"
	// "github.com/tao4free/wechat-go/plugins/wxweb/revoker"
	// "github.com/tao4free/wechat-go/plugins/wxweb/share"
	// "github.com/tao4free/wechat-go/plugins/wxweb/switcher"
	"github.com/tao4free/wechat-go/plugins/wxweb/system"
	// "github.com/tao4free/wechat-go/plugins/wxweb/verify"
	// "github.com/tao4free/wechat-go/plugins/wxweb/youdao"
	"github.com/tao4free/wechat-go/wxweb"
	"time"
)

func main() {
	// create session
	session, err := wxweb.CreateSession(nil, nil, wxweb.TERMINAL_MODE)
	if err != nil {
		log.Fatal(err)
		return
	}
	// load plugins for this session
	// faceplusplus.Register(session)
	replier.Register(session)
	repeater.Register(session)
	// switcher.Register(session)
	// gifer.Register(session)
	// cleaner.Register(session)
	// laosj.Register(session)
	// joker.Register(session)
	// revoker.Register(session)
	// forwarder.Register(session)
	system.Register(session)
	// youdao.Register(session)
	// verify.Register(session)
	// share.Register(session)
	config.Register(session)

	// disable by type example
	if err := session.HandlerRegister.DisableByType(wxweb.MSG_SYS); err != nil {
		log.Fatal(err)
		return
	}
	if err := session.HandlerRegister.DisableByType(wxweb.MSG_TEXT); err != nil {
		log.Fatal(err)
		return
	}
	if err := session.HandlerRegister.DisableByType(wxweb.MSG_IMG); err != nil {
		log.Fatal(err)
		return
	}
	session.HandlerRegister.EnableByName("replier")

	for {
		if err := session.LoginAndServe(false); err != nil {
			log.Fatal("session exit, %s", err)
			for i := 0; i < 3; i++ {
				log.Print("trying re-login with cache")
				if err := session.LoginAndServe(true); err != nil {
					log.Fatal("re-login error, %s", err)
				}
				time.Sleep(3 * time.Second)
			}
			if session, err = wxweb.CreateSession(nil, session.HandlerRegister, wxweb.TERMINAL_MODE); err != nil {
				log.Fatal("create new sesion failed, %s", err)
				break
			}
		} else {
			log.Print("closed by user")
			break
		}
	}
}
