/**
 * https://www.cnblogs.com/Seiyagoo/p/3496945.html
 * https://blog.csdn.net/ice__snow/article/details/79979298
 * https://blog.csdn.net/lvbian/article/details/23031635
 * 
 * This code is collected and generalized by hytzongxuan (hytzongxuan@gmail.com)
 */

#include <X11/Xlib.h>
#include <X11/Xatom.h>
#include <sys/types.h>
#include <dirent.h>
#include <iostream>
#include <cstring>
#include <cstdio>
#include <list>
#include <map>

#define BUF_SIZE 1024
using namespace std;

string PID[50];
int tot;

class WindowsMatchingPid
{
  public:
    WindowsMatchingPid(Display *display, Window wRoot, unsigned long pid)
        : _display(display), _pid(pid)
    {
        _atomPID = XInternAtom(display, "_NET_WM_PID", True);
        if (_atomPID == None)
        {
            cout << "No such atom" << endl;
            return;
        }

        search(wRoot);
    }

    const list<Window> &result() const { return _result; }

  private:
    unsigned long _pid;
    Atom _atomPID;
    Display *_display;
    list<Window> _result;

    void search(Window w)
    {
        Atom type;
        int format;
        unsigned long nItems;
        unsigned long bytesAfter;
        unsigned char *propPID = 0;
        if (Success == XGetWindowProperty(_display, w, _atomPID, 0, 1, False, XA_CARDINAL,
                                          &type, &format, &nItems, &bytesAfter, &propPID))
        {
            if (propPID != 0)
            {
                if (_pid == *((unsigned long *)propPID))
                    _result.push_back(w);

                XFree(propPID);
            }
        }

        Window wRoot;
        Window wParent;
        Window *wChild;
        unsigned nChildren;
        if (0 != XQueryTree(_display, w, &wRoot, &wParent, &wChild, &nChildren))
        {
            for (unsigned i = 0; i < nChildren; i++)
                search(wChild[i]);
        }
    }
};

void getPidByName(char *task_name)
{
    DIR *dir;
    struct dirent *ptr;
    FILE *fp;
    char filepath[50];
    char cur_task_name[50];
    char buf[BUF_SIZE];

    dir = opendir("/proc");

    if (NULL != dir)
    {
        while ((ptr = readdir(dir)) != NULL)
        {

            if ((strcmp(ptr->d_name, ".") == 0) || (strcmp(ptr->d_name, "..") == 0))
                continue;
            if (DT_DIR != ptr->d_type)
                continue;

            sprintf(filepath, "/proc/%s/status", ptr->d_name);
            fp = fopen(filepath, "r");

            if (NULL != fp)
            {
                if (fgets(buf, BUF_SIZE - 1, fp) == NULL)
                {
                    fclose(fp);
                    continue;
                }
                sscanf(buf, "%*s %s", cur_task_name);

                if (!strcmp(task_name, cur_task_name))
                {
                    PID[++tot] = ptr->d_name;
                }

                fclose(fp);
            }
        }
        closedir(dir);
    }
}

struct CloudMusicWindow
{
    Window window;
    int height, width;

    CloudMusicWindow() {}
    CloudMusicWindow(Window _win, int _h, int _w)
    {
        window = _win;
        height = _h;
        width = _w;
    }
};

CloudMusicWindow cloudmusicwindow[50];
map<pair<int, int>, int> Map;
int window_cnt;

void findWindow(string x)
{
    int pid = atoi(x.c_str());

    Display *display = XOpenDisplay(0);

    WindowsMatchingPid match(display, XDefaultRootWindow(display), pid);

    const list<Window> &result = match.result();
    for (list<Window>::const_iterator it = result.begin(); it != result.end(); it++)
    {
        XWindowAttributes attrs;
        XGetWindowAttributes(display, (*it), &attrs);

        Map[make_pair(attrs.height, attrs.width)]++;
        cloudmusicwindow[++window_cnt] = CloudMusicWindow((*it), attrs.height, attrs.width);
    }
}

int main(int argc, char **argv)
{
    Display *display;
    Window window;
    XEvent event;
    int s;

    display = XOpenDisplay(NULL);
    if (display == NULL)
    {
        fprintf(stderr, "Cannot open display\n");
        exit(1);
    }

    getPidByName("cloudmusic.exe");

    for (register int i = 1; i <= tot; i++)
    {
        findWindow(PID[i]);
    }

    s = DefaultScreen(display);

    window = XCreateSimpleWindow(display, RootWindow(display, s), 10, 10, 1, 1, 1,
                                 BlackPixel(display, s), WhitePixel(display, s));

    XSelectInput(display, window, ExposureMask | KeyPressMask);

    XMapWindow(display, window);

    for (register int i = 1; i <= window_cnt; i++)
    {
        if (Map[make_pair(cloudmusicwindow[i].height, cloudmusicwindow[i].width)] == 2)
        {
            XReparentWindow(display, cloudmusicwindow[i].window, window, 1, 1);
        }
    }

    XCloseDisplay(display);
    return 0;
}