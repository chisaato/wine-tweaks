#!/usr/bin/env python3

import time
import gi

gi.require_version('Wnck', '3.0')
gi.require_version('Gtk', '3.0')

from gi.repository import Wnck, Gtk


class WineAppsFrameHide:
    def __init__(self):
        self.activeWindow = None
        self.screen = Wnck.Screen.get_default()
        self.screen.force_update()
        # self.hide_apps = [
        #     "cloudmusic.exe"
        # ]
        self.screen.connect("active_window_changed",
                            self.onActiveWindowChanged)
        # 当前激活的窗口名字
        self.activeWindowInstanceName = ""

    def onActiveWindowChanged(self, screen, window):
        # 先拿到当前激活的窗口
        self.activeWindow = self.screen.get_active_window()
        try:
            self.activeWindowInstanceName = self.activeWindow.get_class_instance_name()
        except AttributeError:
            # 跳过获取 instance name 失败的情况
            # print("无 instance name 跳过")
            return

        # print(f"window changed - class group {active_window.get_class_group_name()}")
        # print(f"window changed - class instance {self.activeWindowInstanceName}")
        # 当激活的窗口不在隐藏 APP 中的时候，就隐藏
        # if activeWindowInstanceName in self.hide_apps:
        if self.activeWindowInstanceName == "cloudmusic.exe":
            self.hideByTypeDialog()
        if self.activeWindowInstanceName == "dingtalk.exe":
            self.hideDingTalk()
        if self.activeWindowInstanceName == "wechat.exe":
            self.hideDingTalk()

    def hideByTypeDialog(self):
        """
        使用隐藏 dialog 的方式
        适用于网易云音乐
        :return: 
        """
        for win in self.screen.get_windows():
            # print(f"窗口名字 {win.get_class_instance_name()} 标题 {win.get_name()}")
            # 选取相同 CLASS 且没有标题的
            if win.get_class_instance_name() == self.activeWindowInstanceName:
                # print(f"状态 {win.get_state()}")
                # 网易云可以用这个办法筛选出背景窗口
                print(f"窗口名字 {win.get_class_instance_name()} 标题 {win.get_name()} 窗口状态 {win.get_window_type()}")
                # 如果类型是 WNCK_WINDOW_DIALOG 就关掉
                if win.get_window_type() == Wnck.WindowType.DIALOG:
                    print("发现边框,关闭")
                    win.close(time.time())

    def hideDingTalk(self):
        """
        针对钉钉进行隐藏
        :return: 
        """
        for win in self.screen.get_windows():
            # print(f"窗口名字 {win.get_class_instance_name()} 标题 {win.get_name()}")
            # 选取相同 CLASS 且没有标题的
            if win.get_class_instance_name() == self.activeWindowInstanceName:
                print(f"窗口名字 {win.get_class_instance_name()} 标题 {win.get_name()}")
                print(f"是否有标题 {win.has_name()}")
                print(f"PID {win.get_pid()}")
                # 如果类型是 WNCK_WINDOW_DIALOG 就关掉
                # if win.get_window_type() == Wnck.WindowType.DIALOG:
                #     win.close(time.time())

    def hideByEmptyNetName(self):
        """
        通过隐藏空白的窗口来实现
        :return: 
        """
        for win in self.screen.get_windows():
            if win and win.get_class_instance_name() == self.activeWindowInstanceName:
                win.minimize(time.time())

    def run(self):
        Gtk.main()


WineAppsFrameHide().run()
