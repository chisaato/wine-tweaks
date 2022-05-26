#include "windows.h"
LPCTSTR windowClassName = L"popupshadow";

int WINAPI wWinMain(HINSTANCE hInstance, HINSTANCE hPrevInstance, PWSTR pCmdLine, int nCmdShow) {
  for (;;) {
    for (HWND h = FindWindowEx(NULL, NULL, windowClassName, NULL);
        h = FindWindowEx(NULL, h, windowClassName, NULL);
        h != NULL) {
      ShowWindow(h, SW_HIDE);
    }
    Sleep(3000);
  }
  return 0;
}