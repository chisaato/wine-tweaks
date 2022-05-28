// we import the necessary modules (only the core X module in this application).
use xcb::{x};
use std::fmt::Debug;
xcb::atoms_struct! {
    #[derive(Debug)]
    struct Atoms {
        net_wm_name    => b"_NET_WM_NAME",
        at_spi_bus   => b"AT_SPI_BUS",
        wm_name        => b"WM_NAME",
        wm_class        => b"WM_CLASS",
        wm_hints   => b"WM_HINTS",
        wm_state_maxh   => b"_NET_WM_STATE_MAXIMIZED_HORZ",
    }
}
fn main() -> xcb::Result<()> {
    // Connect to the X server.
    let (conn, screen_num) = xcb::Connection::connect(None).unwrap();
    let setup = conn.get_setup();
    let screen = setup.roots().nth(screen_num as usize).unwrap();
    let window: x::Window = screen.root();
    let cookie = conn.send_request_checked(&x::ChangeWindowAttributes {
        window,
        value_list: &[
            x::Cw::EventMask(x::EventMask::SUBSTRUCTURE_NOTIFY | x::EventMask::EXPOSURE)
        ],
    });
    conn.flush()?;
    println!("已切换窗口属性");

    let atoms = Atoms::intern_all(&conn)?;
    // println!("atoms = {:#?}", atoms);
    // let cookie = conn.send_request(&x::GetProperty {
    //     delete: false,
    //     window,
    //     property: atoms.at_spi_bus,
    //     r#type: x::ATOM_STRING,
    //     long_offset: 0,
    //     long_length: 1024,
    // });
    // let reply = conn.wait_for_reply(cookie)?;
    // let title = std::str::from_utf8(reply.value()).expect("The WM_NAME property is not valid UTF-8");
    // println!("{:#?}", reply);
    // println!("{}", &title);
    conn.flush()?;
    println!("已完成注册");
    loop {
        let event = match conn.wait_for_event() {
            Err(xcb::Error::Connection(err)) => {
                panic!("unexpected I/O error: {}", err);
            }
            Err(xcb::Error::Protocol(xcb::ProtocolError::X(x::Error::Font(err), _req_name))) => {
                // may be this particular error is fine?
                continue;
            }
            Err(xcb::Error::Protocol(err)) => {
                panic!("unexpected protocol error: {:#?}", err);
            }
            Ok(event) => event,
        };
        // println!("Received event {:#?}", event);
        match event {
            xcb::Event::X(x::Event::MapNotify(ev)) => {
                println!("有窗口创建!: {:#?}", ev);
                // 先来拿数据
                let cookie = conn.send_request(&x::GetProperty {
                    delete: false,
                    window: ev.window(),
                    property: atoms.wm_class,
                    r#type: x::ATOM_STRING,
                    long_offset: 0,
                    long_length: 32,
                });
                let reply = conn.wait_for_reply(cookie)?;
                println!("Got");
                println!("{:#?}", reply);
                // let title = std::str::from_utf8(reply.unwrap().value()).expect("The WM_NAME property is not valid UTF-8");
                // println!("{}", title);
            }
            // xcb::Event::X(x::Event::MapRequest(ev)) => {
            //     println!("MapRequest:");
            // }

            _ => {
                // println!("223:");
                // break Ok(())
            }
        }
    }
}