# Goyave Template

A template project to get started with the [Goyave](https://github.com/go-goyave/goyave) framework.

## Getting Started

### Requirements

- Go 1.20

### Running the project

First, make your own configuration for your local environment. You can copy `config.example.json` to `config.json`.

Run `go run main.go` in your project's directory to start the server, then try to request the `users` route.
```
$ curl http://localhost:8080/users
$ curl http://localhost:8080/users?page=2
$ curl http://localhost:8080/users?page=2&perPage=100
```

## Contributing

Thank you for considering contributing to the Goyave framework! You can find the contribution guide in the [documentation](https://goyave.dev/guide/contribution-guide.html).

I have many ideas for the future of Goyave. I would be infinitely grateful to whoever want to support me and let me continue working on Goyave and making it better and better.

You can support me on Github Sponsor or Patreon.

<a href="https://github.com/sponsors/System-Glitch">❤ Sponsor me!</a>

## License

The Goyave framework is MIT Licensed. Copyright © 2023 Jérémy LAMBERT (SystemGlitch)
