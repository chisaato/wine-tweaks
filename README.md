# wine-tweaks

一些个人向的 Wine 辅助工具  
运行环境主要针 Bottles (Flatpak) 且遵循以下偏好

- 避免使用 Deepin 组件,如 deepinwine
- 尽可能使用原生 Wine 或 Bottles 的 Caffe Runtime

目前解决了以下问题

- 移除钉钉带来的透明边框
- 移除微信带来的半透明边框
- 移除网易云的透明边框

接下来?:

- 运行 NtrQQ

## xcb-hide-shadow-go

这是一个通过 XCB 实现的简单的隐藏阴影的工具，可以用于隐藏阴影的窗口，比如钉钉、微信等。  
原理基本上是侦听 X 的 map 窗口事件,检测窗口是否是无标题的 (有 WM_CLASS 无 \_NET_WM_NAME), 如果是则隐藏阴影。

虽然这种方式很 Hack,也不优雅,但是作为一个技术实现已经足够了.未来可以考虑加入更多判定.

## xcb-hide-shadow-rs

同上,但是用 Rust 实现.

首要维护应该还是会选择 Go 或者依据各语言的 Binding 完善程度决定
