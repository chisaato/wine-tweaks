// we import the necessary modules (only the core X module in this application).
use xcb::{x};
// we need to import the `Xid` trait for the `resource_id` call down there.
use xcb::{Xid};

// Many xcb functions return a `xcb::Result` or compatible result.
fn main() -> xcb::Result<()> {
    // Connect to the X server.
    let (conn, screen_num) = xcb::Connection::connect(None)?;

    // Fetch the `x::Setup` and get the main `x::Screen` object.
    let setup = conn.get_setup();
    let screen = setup.roots().nth(screen_num as usize).unwrap();

    // Generate an `Xid` for the client window.
    // The type inference is needed here.
    let window: x::Window = conn.get_setup().roots().nth(screen_num as usize).unwrap().root();
    // Let's change the window title
    let cookie = conn.send_request_checked(&x::ChangeWindowAttributes {
        window,
        value_list: &[
            x::Cw::EventMask(x::EventMask::STRUCTURE_NOTIFY)
        ],
    });
    let (net_wm_name) = {
        let cookies = (
            conn.send_request(&x::InternAtom {
                only_if_exists: true,
                name: b"_NET_WM_NAME",
            }),
        );
        (
            conn.wait_for_reply(cookies.0)?.atom(),
        )
    };
    // And check for success again
    // conn.check_request(cookie);
    conn.flush();
    // We enter the main event loop
    loop {
        match conn.wait_for_event()? {
            xcb::Event::X(x::Event::MapNotify(ev)) => {
                println!("MapNotify:");
            }
            xcb::Event::X(x::Event::ClientMessage(ev)) => {}
            _ => {}
        }
    }
}