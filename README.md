# ipmsgo

A Go-based IP Messenger Client.

This project is based on a new ipmsg library that I wrote, since mattn's ipmsg library was outdated and incompatible with go modules. The library is located in the `/pkg` folder, and the command line client is in the `/cmd` folder. I also plan to add a Flask-based web GUI in the `/gui` folder.

I have not yet decided whether to create a WebAssembly version of the web GUI or not. [go-app](<https://go-app.dev/>) might be an option.

**The library is still under development and not ready for use.**

## Credits

This project is inspired by [mattn/go-ipmsg](<https://github.com/mattn/go-ipmsg>) and [zeromake/ipmsg-protocol(Chinese Translation)](<https://blog.zeromake.com/pages/ipmsg-protocol/>)

## License

This project is licensed under the Mizumoto General Public License, Version 1.4. You can find the full license text in [LICENSE](./LICENSE/Mizumoto.General.Public.License.v1.4.md).

The _entities_\* on the [Disqualified Entities List](./LICENSE/Disqualified.Entities.List.md) are prohibited from using _files_\*\* from this project in any way.

---
> \*/\*\*: See chapter [#Restrictions for Users](./LICENSE/Mizumoto.General.Public.License.v1.4.md/#restrictions-for-users) for definitions of _entities_ and _files_.