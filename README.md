# Escapepod

escapepod is a standalone, self-hosted, podcast manager. It includes a modern frontend web interface, an internal player, and a flexible API for mobile applications to use.

![Image 1](https://user-images.githubusercontent.com/1963/63644305-98656580-c6ab-11e9-9b69-8a5e66be69c3.png)
![Image 2](https://user-images.githubusercontent.com/1963/63644296-62c07c80-c6ab-11e9-9134-0a590817efef.png)

## Features

- OPML Import
- Individual RSS Feed import
- Embedded Player using [plyr](https://github.com/sampotts/plyr)
- UI is embedded in the binary

## Running Escapepod

```sh
docker-compose up
```

or

```sh
docker-compose up -d
```

Open a web browser to localhost:5000.

## Docker Image

```sh
docker pull rphillips/escapepod:latest
```

## Development

See [HACKING.md](./HACKING.md)

## TODO

- [ ] Played history
- [ ] Resuming

## License Note

Third party applications (ie: mobile, desktop, etc) that communicate with the
Escapepod server API are **not** considered a derivative work. Any
modifications to the **server** component (or included client interface) must conform to the
[license](./LICENSE.txt); however, third-party applications communicating with the
API may be closed source.
