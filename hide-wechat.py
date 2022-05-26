#!/usr/bin/env python3

import time
import gi

gi.require_version('Wnck', '3.0')
gi.require_version('Gtk', '3.0')

from gi.repository import Wnck, Gtk


class WineAppsFrameHide():
    def __init__(self):
        self.screen = Wnck.Screen.get_default()
        self.screen.force_update()
        # self.hide_apps = [
        #     "cloudmusic.exe"
        # ]
        self.screen.connect("active_window_changed", self.onActiveWindowChanged)
        # 当前激活的窗口名字
        self.activeWindowInstanceName = ""

    def onActiveWindowChanged(self, screen, window):
        # 先拿到当前激活的窗口
        active_window = self.screen.get_active_window()
        try:
            self.activeWindowInstanceName = active_window.get_class_instance_name()
        except AttributeError:
            # 跳过获取 instance name 失败的情况
            # print("无 instance name 跳过")
            return

        # print(f"window changed - class group {active_window.get_class_group_name()}")
        # print(f"window changed - class instance {self.activeWindowInstanceName}")
        # 当激活的窗口不在隐藏 APP 中的时候，就隐藏
        # if activeWindowInstanceName in self.hide_apps:
        if self.activeWindowInstanceName == "cloudmusic.exe":
            self.hideCloudMusic()

    def hideCloudMusic(self):
        # 针对网易云音乐进行隐藏
        for win in self.screen.get_windows():
            # print(f"窗口名字 {win.get_class_instance_name()} 标题 {win.get_name()}")
            # 选取相同 CLASS 且没有标题的
            if win.get_class_instance_name() == self.activeWindowInstanceName:
                print(f"窗口名字 {win.get_class_instance_name()} 标题 {win.get_name()} 窗口状态 {win.get_window_type()}")
                # 如果类型是 WNCK_WINDOW_DIALOG 就关掉
                if win.get_window_type() == Wnck.WindowType.DIALOG:
                    win.close(time.time())
                # if win.is_above():
                # print(f"hide window {win.get_class_instance_name()}")
                # win.minimize()
                # win.close(time.time())

    def run(self):
        Gtk.main()


WineAppsFrameHide().run()
