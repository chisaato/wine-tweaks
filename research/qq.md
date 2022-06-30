# QQ

## 字体文件

目前确认登陆界面字体与 Arial 字体有关,必须定义 Arial 字体去哪里寻找.

同时必须不能只标记字体文件名,必须告知字形. 如下,使用 `,` 分隔文件名和字形名

```reg
[HKEY_LOCAL_MACHINE\Software\Microsoft\Windows NT\CurrentVersion\FontLink\SystemLink]
"Arial"="wqy-microhei.ttc,WenQuanYi Micro Hei"
"Arial Black"="wqy-microhei.ttc,WenQuanYi Micro Hei"
```

目前确认 Arial 字体不可以使用其他替代,暂时只能用 Caffe 默认提供的那个值

## 字体缺失

使用 Caffe 7.7 的时候 QQ 会在一次启动字体正确一次缺失之间循环. 使用 Caffe 7.10 也有这个问题  
但是一旦在字体正确的情况下登陆,则后续字体都没有问题.

在 Caffe 下使用 `QQSclauncher.exe` 而不是 `QQ.exe` 可以确保每次字体都正常.  
~~但是一旦进入 QQ,仍然会出现中文字体缺失~~ 参考下方 Tahoma 字体的解决方案即可正常

使用 wine-GE 时,每次启动登陆界面字体一定是正常的.  
但是进入之后还是会有少量字体缺失

## 滚动卡顿

使用 Caffe 的时候,直接将 `dwrite.dll` 设定为原生优于内建即可

使用 wine-GE 时,上述解决方案无效. 也许需要自行拷贝 Windows 下的 DLL

## Tahoma 字体

根据试验,在 wincfg 中预览的 Tahoma 字体,与 QQ 字体是否正常也有一定关系.

使用 Caffe 时,Tahoma 字体正常则登陆 QQ 后一定正常(在登陆界面正常时)

使用 wine-GE 时,Tahoma 字体正常与 QQ 字体是否正常暂时没发现强关联.

暂时发现,Tahoma 字体配置必须是指定字体文件名,不可以带字形名. 此时才可以正常在 winecfg 显示  
Tahoma Bold 字体则暂时没做研究,使用 Caffe 默认配置
