# wine-tweaks

一些个人向的 Wine 辅助工具
运行环境主要针对

- Bottles (Flatpak)

目前解决了以下问题

- 移除钉钉带来的透明边框
- 移除微信带来的半透明边框

接下来?:

- 隐藏网易云的半透明边框
- 运行 NtrQQ

## xcb-hide-shadow-go

这是一个通过 XCB 实现的简单的隐藏阴影的工具，可以用于隐藏阴影的窗口，比如钉钉、微信等。  
原理基本上是侦听 X 的 map 窗口事件,检测窗口是否是无标题的 (有 WM_CLASS 无 \_NET_WM_NAME), 如果是则隐藏阴影。

虽然这种方式很 Hack,也不优雅,但是作为一个技术实现已经足够了.未来可以考虑加入更多判定.

## xcb-hide-shadow-rs

同上,但是用 Rust 实现.

首要维护应该还是会选择 Go 或者依据各语言的 Binding 完善程度决定

## 网易云边框问题

可以使用 `research` 目录中的 `hide-shadow-frame.py` 来隐藏,这个脚本也是最开始研究隐藏窗口所用的.

网易云音乐的透明边框与其他应用不同,它是由四个独立的窗口组成的,这可以在窗口管理器中看到. 但是透明边框的属性不太一样,所以可以很明显的分辨出来然后干掉.  
在侦听的窗口 map 事件中,网易云产生的四个独立边框中具有空的 WM_CLASS 或 \_NET_WM_NAME. 但是空的 WM_CLASS 在在事件中很常见,也有许多后台程序会创建这类窗口.  
所以不能直接对所有的空 WM_CLASS 进行 unmap 正在考虑使用其他检测手段

通过 xprop 去检测窗口属性时,是可以检测到边框的 WM_CLASS 的.但是在事件发生的时候,这个属性是空的.  
根据 Reddit 上的一些讨论,以 Sportily 为例子,它就存在先 map 窗口后设定 WM_CLASS 的问题. 所以此处假设网易云也是这种情况.  
但是通过在代码中添加延迟来获取 WM_CLASS 依然没有解决这个问题.

不过在 xprop 检测中,有一些更有意思的情况.  
这四个边框的 ICON 是 Wine 的 Logo,并且具有属性 \_NET_WM_VISIBLE_NAME,这个属性分别为  
**但是同样无法在事件发生时捕获!**

- 空
- <2>
- <3>
- <4>

猜测可能是因为四个无名窗口导致 WM 对窗口自动命名

下面放上正常窗口与边框的 xprop 检测

边框

```bash
_NET_WM_DESKTOP(CARDINAL) = 0
WM_STATE(WM_STATE):
                window state: Normal
                icon window: 0x0
_NET_WM_ALLOWED_ACTIONS(ATOM) = _NET_WM_ACTION_MOVE, _NET_WM_ACTION_MINIMIZE, _NET_WM_ACTION_CHANGE_DESKTOP, _NET_WM_ACTION_CLOSE
_KDE_NET_WM_ACTIVITIES(STRING) = "5c41c030-8b10-424d-b545-fd7e1d93cd39"
_NET_WM_VISIBLE_NAME(UTF8_STRING) = " <4>"
_NET_WM_ICON(CARDINAL) =        Icon (32 x 32):

                   ░░░░░░░░░░
                  ░░       ░░░
                  ░░░░░  ░░░░░
                  ░          ░
                  ░          ░
                             ░
                 ░         ░ ░░
                 ░      ░░░░░ ░
                 ░    ░▒▒▒░░░ ░
                 ░▒░░░▓▓▓▓▓▓▓░░
                 ▒▓░░░▓▓▓▒▒▒▒▒░
                 ▒▓░░░▒▒░░░▒▒▒░
                 ▒▓░░▒▓▓▓▓▓▓▓▒░
                 ░▓▓░▒▓▓▓▓▓▓▓░░
                 ░▒▓▒▒▓▓▓▓▓▓▒░░
                  ░▒▓▓▓▓▓▓▓▒░░
                   ░▒▒▓▓▓▒░░░
                    ░░░▒░░░░
                     ░░░░░░
                      ░░▒░
                      ░░░░
                       ░░
                       ░░
                       ░░
                       ░░
                     ░░░▒░░░
                   ░   ░░   ░
                  ░   ░░░░░  ░
                  ░░        ░░
                   ░░░░░░░░░░



_NET_WM_STATE(ATOM) = _NET_WM_STATE_ABOVE, _NET_WM_STATE_STAYS_ON_TOP, _NET_WM_STATE_SKIP_TASKBAR, _NET_WM_STATE_SKIP_PAGER
_NET_WM_NAME(UTF8_STRING) =
WM_ICON_NAME(STRING) =
WM_NAME(STRING) =
WM_HINTS(WM_HINTS):
                Client accepts input or input focus: False
                Initial state is Normal State.
                bitmap id # to use for icon: 0x8000128
                bitmap id # of mask for icon: 0x800012a
                window id # of group leader: 0x8400006
_NET_WM_WINDOW_TYPE(ATOM) = _NET_WM_WINDOW_TYPE_DIALOG
WM_TRANSIENT_FOR(WINDOW): window id # 0x8400006
_MOTIF_WM_HINTS(_MOTIF_WM_HINTS) = 0x3, 0x24, 0x0, 0x0, 0x0
WM_NORMAL_HINTS(WM_SIZE_HINTS):
                program specified location: 2258, 313
                program specified minimum size: 29 by 662
                program specified maximum size: 29 by 662
                window gravity: Static
_NET_WM_USER_TIME_WINDOW(WINDOW): window id # 0x800000d
XdndAware(ATOM) = BITMAP
_NET_WM_PID(CARDINAL) = 87
WM_LOCALE_NAME(STRING) = "zh_CN.UTF-8"
WM_CLIENT_MACHINE(STRING) = "hidden"
WM_CLASS(STRING) = "cloudmusic.exe", "cloudmusic.exe"
WM_PROTOCOLS(ATOM): protocols  WM_DELETE_WINDOW, _NET_WM_PING, WM_TAKE_FOCUS
```

主界面

```bash
_NET_WM_DESKTOP(CARDINAL) = 0
WM_STATE(WM_STATE):
                window state: Normal
                icon window: 0x0
_NET_WM_ICON_GEOMETRY(CARDINAL) = 720, 1396, 52, 44
_NET_WM_ALLOWED_ACTIONS(ATOM) = _NET_WM_ACTION_MOVE, _NET_WM_ACTION_RESIZE, _NET_WM_ACTION_MINIMIZE, _NET_WM_ACTION_MAXIMIZE_VERT, _NET_WM_ACTION_MAXIMIZE_HORZ, _NET_WM_ACTION_FULLSCREEN, _NET_WM_ACTION_CHANGE_DESKTOP, _NET_WM_ACTION_CLOSE
_KDE_NET_WM_ACTIVITIES(STRING) = "5c41c030-8b10-424d-b545-fd7e1d93cd39"
_NET_WM_ICON(CARDINAL) =        Icon (32 x 32):
            ░▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒░
          ░▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒░
         ░▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒░
         ▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒
        ░▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒░░░▒▒▒▒▒▒▒▒▒░
        ▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒░     ▒▒▒▒▒▒▒▒▒
        ▒▒▒▒▒▒▒▒▒▒▒▒░▒▒▒▒      ░▒▒▒▒▒▒▒▒
        ▒▒▒▒▒▒▒▒▒▒░   ▒▒░  ░▒ ░▒▒▒▒▒▒▒▒▒
        ▒▒▒▒▒▒▒▒▒    ░▒▒░  ▒▒▒▒▒▒▒▒▒▒▒▒▒
        ▒▒▒▒▒▒▒▒    ░▒▒▒░  ░▒▒▒▒▒▒▒▒▒▒▒▒
        ▒▒▒▒▒▒▒    ▒▒▒▒       ░▒▒▒▒▒▒▒▒▒
        ▒▒▒▒▒▒░   ▒▒▒░          ▒▒▒▒▒▒▒▒
        ▒▒▒▒▒▒   ▒▒▒░            ▒▒▒▒▒▒▒
        ▒▒▒▒▒▒  ░▒▒▒   ░▒▒  ░░   ░▒▒▒▒▒▒
        ▒▒▒▒▒░  ░▒▒▒  ░▒▒▒   ▒▒   ▒▒▒▒▒▒
        ▒▒▒▒▒░  ▒▒▒░  ▒▒▒▒   ▒▒░  ░▒▒▒▒▒
        ▒▒▒▒▒░  ▒▒▒░  ▒▒▒▒   ▒▒▒  ░▒▒▒▒▒
        ▒▒▒▒▒░  ▒▒▒░   ▒▒▒   ▒▒▒  ░▒▒▒▒▒
        ▒▒▒▒▒░  ▒▒▒▒   ░░    ▒▒▒  ░▒▒▒▒▒
        ▒▒▒▒▒▒  ░▒▒▒░       ▒▒▒▒  ░▒▒▒▒▒
        ▒▒▒▒▒▒   ▒▒▒▒░     ▒▒▒▒░  ░▒▒▒▒▒
        ▒▒▒▒▒▒░   ▒▒▒▒▒░░░▒▒▒▒▒   ▒▒▒▒▒▒
        ▒▒▒▒▒▒▒   ░▒▒▒▒▒▒▒▒▒▒▒   ░▒▒▒▒▒▒
        ▒▒▒▒▒▒▒▒    ▒▒▒▒▒▒▒▒▒    ▒▒▒▒▒▒▒
        ▒▒▒▒▒▒▒▒░    ░░▒▒▒░     ▒▒▒▒▒▒▒▒
        ▒▒▒▒▒▒▒▒▒▒             ▒▒▒▒▒▒▒▒▒
        ▒▒▒▒▒▒▒▒▒▒▒░         ▒▒▒▒▒▒▒▒▒▒▒
        ░▒▒▒▒▒▒▒▒▒▒▒▒▒░░░░░▒▒▒▒▒▒▒▒▒▒▒▒░
         ▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒
         ░▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒░
          ░▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒░
            ░▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒░


_NET_WM_STATE(ATOM) =
_NET_WM_NAME(UTF8_STRING) = "I LOVE U - 阿良良木健/洛天依"
WM_ICON_NAME(COMPOUND_TEXT) = "I LOVE U - 阿良良木健/洛天依"
WM_NAME(COMPOUND_TEXT) = "I LOVE U - 阿良良木健/洛天依"
WM_HINTS(WM_HINTS):
                Client accepts input or input focus: False
                Initial state is Normal State.
                bitmap id # to use for icon: 0x8000164
                bitmap id # of mask for icon: 0x8000166
                window id # of group leader: 0x8400006
_NET_WM_WINDOW_TYPE(ATOM) = _NET_WM_WINDOW_TYPE_NORMAL
_MOTIF_WM_HINTS(_MOTIF_WM_HINTS) = 0x3, 0x3e, 0x0, 0x0, 0x0
WM_NORMAL_HINTS(WM_SIZE_HINTS):
                program specified location: 1240, 309
                window gravity: Static
_NET_WM_USER_TIME_WINDOW(WINDOW): window id # 0x800000d
XdndAware(ATOM) = BITMAP
_NET_WM_PID(CARDINAL) = 87
WM_LOCALE_NAME(STRING) = "zh_CN.UTF-8"
WM_CLIENT_MACHINE(STRING) = "hidden"
WM_CLASS(STRING) = "cloudmusic.exe", "cloudmusic.exe"
WM_PROTOCOLS(ATOM): protocols  WM_DELETE_WINDOW, _NET_WM_PING, WM_TAKE_FOCUS
```

不知道在启动wine的时候强制WM_CLASS是否有效?