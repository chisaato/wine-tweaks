package main

import (
	"fmt"
	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"
	"strings"
)

func registerAtom(X *xgb.Conn, name string) xproto.Atom {
	atomReply, err := xproto.InternAtom(X, true, uint16(len(name)), name).Reply()
	if err != nil {
		fmt.Println(err)
	}
	return atomReply.Atom
}
func main() {
	X, err := xgb.NewConn()
	if err != nil {
		fmt.Println(err)
		return
	}
	setup := xproto.Setup(X)
	root := setup.DefaultScreen(X).Root
	xproto.ChangeWindowAttributes(X, root, xproto.CwEventMask, []uint32{xproto.EventMaskSubstructureNotify | xproto.EventMaskStructureNotify})
	// 注册自己的 Atom
	atomNetWmName := registerAtom(X, "_NET_WM_NAME")
	atomUtf8String := registerAtom(X, "UTF8_STRING")

	//atomNetActiveWindow := registerAtom(X, "_NET_ACTIVE_WINDOW")
	for {
		//X.Sync()
		ev, xerr := X.PollForEvent()
		//if ev == nil && xerr == nil {
		//	fmt.Println("Both event and error are nil. Exiting...")
		//	return
		//}
		if ev != nil {
			//fmt.Println("Event:", ev)
			switch ev.(type) {
			case xproto.ConfigureNotifyEvent:
				//fmt.Println("ConfigureNotifyEvent", ev)
				//fmt.Println("发生窗口焦点切换")
				event := ev.(xproto.ConfigureNotifyEvent)
				wmClassCookie := xproto.GetProperty(X, false, event.Window, xproto.AtomWmClass, xproto.AtomString, 0, (1<<32)-1)
				wmClassReply, err := wmClassCookie.Reply()
				if err != nil {
					// BadWindow 就别看了
					//fmt.Println(err)
					return
				}
				var wmClassName string
				if wmClassReply.Value != nil {
					// 只打印一半字符串
					wmClassName = string(wmClassReply.Value[:len(wmClassReply.Value)/2])
					wmClassName = strings.TrimSuffix(wmClassName, "\x00")
					fmt.Println("WM_CLASS: " + wmClassName)
				}
				atomWindowType := registerAtom(X, "WM_ICON_NAME")
				windowTypeCookie := xproto.GetProperty(X, false, event.Window, atomWindowType, xproto.GetPropertyTypeAny, 0, (1<<32)-1)
				windowTypeReply, err := windowTypeCookie.Reply()
				if err != nil {
					fmt.Println(err)
					continue
				}
				if windowTypeReply.Value != nil {
					//netWmName = strings.TrimSuffix(netWmName, "\x00")
					fmt.Println(windowTypeReply.Value)
					fmt.Println(string(windowTypeReply.Value))
					if len(windowTypeReply.Value) > 0 {
						// 说明是那些透明边框,关闭
						fmt.Println("是透明边框")
						//xproto.UnmapWindow(X, event.Window)
						xproto.DestroyWindow(X, event.Window)
					}
				}
			case xproto.DestroyNotifyEvent:
				//fmt.Println("DestroyNotifyEvent")
			case xproto.MapRequestEvent:
				//fmt.Println("MapRequestEvent")
			case xproto.CreateNotifyEvent:
				//fmt.Println("CreateNotifyEvent")
			case xproto.MapNotifyEvent:
				fmt.Println("有窗口显示")
				// 据说延迟一下,可以让避免窗口已经 Map 但是拿不到属性的情况
				// 不一定有效
				//time.Sleep(1 * time.Second)
				event := ev.(xproto.MapNotifyEvent)
				wmClassCookie := xproto.GetProperty(X, false, event.Window, xproto.AtomWmClass, xproto.AtomString, 0, (1<<32)-1)
				wmClassReply, err := wmClassCookie.Reply()
				if err != nil {
					// BadWindow 就别看了
					//fmt.Println(err)
					return
				}
				var wmClassName string
				if wmClassReply.Value != nil {
					// 只打印一半字符串
					wmClassName = string(wmClassReply.Value[:len(wmClassReply.Value)/2])
					wmClassName = strings.TrimSuffix(wmClassName, "\x00")
					//fmt.Println("WM_CLASS: " + wmClassName)
				}
				netWmNameCookie := xproto.GetProperty(X, false, event.Window, atomNetWmName, atomUtf8String, 0, (1<<32)-1)
				netWmNameReply, err := netWmNameCookie.Reply()
				if err != nil {
					fmt.Println(err)
					return
				}
				var netWmName string
				if netWmNameReply.Value != nil {
					netWmName = string(netWmNameReply.Value)
					netWmName = strings.TrimSuffix(netWmName, "\x00")
					//fmt.Println("_NET_WM_NAME: " + netWmName + "长度：" + strconv.Itoa(len(netWmName)))
				}
				// 这里开始处理这些古怪的窗口,未来考虑拆分代码
				if wmClassName == "wechat.exe" {
					fmt.Println("检测到微信")
					if len(netWmName) == 0 {
						fmt.Println("没有名字, Unmap!")
						xproto.UnmapWindow(X, event.Window)
					}
				}
				if wmClassName == "dingtalk.exe" {
					fmt.Println("检测到钉钉")
					if len(netWmName) == 0 {
						fmt.Println("没有名字, Unmap!")
						xproto.UnmapWindow(X, event.Window)
					}
				}
				if len(wmClassName) == 0 {
					//fmt.Println("检测到可能是网易云音乐")
					// 2s
					atomWindowType := registerAtom(X, "WM_CLIENT_MACHINE")
					windowTypeCookie := xproto.GetProperty(X, false, event.Window, atomWindowType, xproto.GetPropertyTypeAny, 0, (1<<32)-1)
					windowTypeReply, err := windowTypeCookie.Reply()
					if err != nil {
						//fmt.Println(err)
						continue
					}
					if windowTypeReply.Value != nil {
						//netWmName = strings.TrimSuffix(netWmName, "\x00")
						//fmt.Println(windowTypeReply.Value)
					}
				}

			}
		}
		if xerr != nil {
			fmt.Printf("Error: %s\n", xerr)
		}
	}
}

// hideShadow 隐藏窗口的阴影
func hideShadow(X *xgb.Conn, win xproto.Window) {

}
