<p align="center">
  <h3 align="center">liquibasechangeloglint</h3>
  <p align="center"><code>liquibasechangeloglint</code> is a cli tool to lint a liquibase changelog yaml file.</p>
  <p align="center">
    <a href="https://github.com/paulvollmer/liquibasechangeloglint/actions"><img alt="Github Actions" src="https://github.com/paulvollmer/liquibasechangeloglint/workflows/CI/badge.svg?style=flat-square"> </a>
    <a href="https://github.com/paulvollmer/liquibasechangeloglint/releases"><img alt="Software Release" src="https://img.shields.io/github/v/release/paulvollmer/liquibasechangeloglint.svg?include_prereleases&style=flat-square&color=blue"></a>
    <a href="/LICENSE"><img alt="Software License" src="https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square"></a>
  </p>
</p>

---  

**This tool is in development and does not cover the whole liquibase changelog schema.**

## Installation

```console
go get -u github.com/paulvollmer/liquibasechangeloglint
```

## Usage

```console
liquibasechangeloglint path/to/changelog.yaml
```

## Development

```sh
git clone git@github.com:paulvollmer/liquibasechangeloglint.git
cd liquibasechangeloglint
go build
go test
```

## License

[MIT License](LICENSE)
