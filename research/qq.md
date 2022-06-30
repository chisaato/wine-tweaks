# QQ

## 字体文件

目前确认登陆界面字体与 Arial 字体有关,必须定义 Arial 字体去哪里寻找.

同时必须不能只标记字体文件名,必须告知字形. 如下,使用 `,` 分隔文件名和字形名

```reg
[HKEY_LOCAL_MACHINE\Software\Microsoft\Windows NT\CurrentVersion\FontLink\SystemLink]
"Arial"="wqy-microhei.ttc,WenQuanYi Micro Hei"
"Arial Black"="wqy-microhei.ttc,WenQuanYi Micro Hei"
```

## 字体缺失

使用 Caffe 7.7 的时候 QQ 会在一次启动字体正确一次缺失之间循环.  
使用 Caffe 7.10 也有这个问题  
但是一旦在字体正确的情况下登陆,则后续字体都没有问题.

使用 wine-GE 时,每次启动登陆界面字体一定是正常的.  
但是进入之后还是会有少量字体缺失

## 滚动卡顿

使用 Caffe 的时候,直接将 `dwrite.dll` 设定为原生优于内建即可

使用 wine-GE 时,上述解决方案无效. 也许需要自行拷贝 Windows 下的 DLL

## Tahoma 字体

根据试验,在 wincfg 中预览的 Tahoma 字体,与 QQ 字体是否正常也有一定关系.

使用 Caffe 时,Tahoma 字体正常则登陆 QQ 后一定正常(在登陆界面正常时)

使用 wine-GE 时,Tahoma 字体正常与 QQ 字体是否正常暂时没发现强关联.