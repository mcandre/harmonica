# harmonica: comic ebook repackager

[![GitHub Downloads](https://img.shields.io/github/downloads/mcandre/harmonica/total?logo=github)](https://github.com/mcandre/harmonica/releases) [![Docker Pulls](https://img.shields.io/docker/pulls/n4jm4/harmonica)](https://hub.docker.com/r/n4jm4/harmonica) [![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/mcandre/harmonica) [![Test](https://github.com/mcandre/harmonica/actions/workflows/test.yml/badge.svg)](https://github.com/mcandre/harmonica/actions/workflows/test.yml) [![Test-Futureproof-Dependencies](https://github.com/mcandre/harmonica/actions/workflows/test-futureproof-dependencies.yml/badge.svg)](https://github.com/mcandre/harmonica/actions/workflows/test-futureproof-dependencies.yml) [![Test-Futureproof-Language](https://github.com/mcandre/harmonica/actions/workflows/test-futureproof-language.yml/badge.svg)](https://github.com/mcandre/harmonica/actions/workflows/test-futureproof-language.yml) [![Test-Futureproof-OS](https://github.com/mcandre/harmonica/actions/workflows/test-futureproof-os.yml/badge.svg)](https://github.com/mcandre/harmonica/actions/workflows/test-futureproof-os.yml) [![license](https://img.shields.io/badge/license-BSD-3)](LICENSE.md)

# SUMMARY

`harmonica` repackages comics into batches of smaller collections.

# EXAMPLE

```console
$ cd examples

$ harmonica -n 36 atomic-war

$ tree issue-*
issue-1
├── Atomic_War_no.1_195211_pg00a.jpg
├── Atomic_War_no.1_195211_pg00b.jpg
├── Atomic_War_no.1_195211_pg01.jpg
├── Atomic_War_no.1_195211_pg02.jpg
├── Atomic_War_no.1_195211_pg03.jpg
...
issue-2
├── Atomic_War_no.2_195212_pg00a.jpg
├── Atomic_War_no.2_195212_pg00b.jpg
├── Atomic_War_no.2_195212_pg01.jpg
├── Atomic_War_no.2_195212_pg02.jpg
├── Atomic_War_no.2_195212_pg03.jpg
...
```

`-n` controls max number of files (default unlimited).

`-m` controls max directory size (default 250 MiB).

`-prefix` customizes the batch prefix.

`-unzip` handles ZIP format archive (.ZIP, .CBZ, .JAR, etc.) source files.

`-zip <.extension>` compresses each batch into independent ZIP format archives.

See `harmonica -help` for more detail.

# ABOUT

Why, though?

Poorly written multimedia applications attempt to load a large object in memory and then crash. For example, comic ebooks often refuse to read smoothly unless dispersed among a set of smaller (`.CBZ`) files. harmonica divides large archives into smaller archives more likely to successfully process in more comic book side loading ereaders.

Poorly written cloud storage applications that struggle to reliably transfer large files. Some applications fail file transfers if the user doesn't continually force the screen to stay awake. harmonica chunks large archives into smaller chunks. When file transfers fail, it's more convenient to retry a specific smaller chunk file than the original, large file.

Classical split archive files (e.g. `*.Z00`, `*.Z01`, `*.Z02`, ..., `*.ZIP`) are not designed to operate in isolation. harmonica splits your files into ordinary ZIP files.

# WARNING

harmonica is designed to work on a set of files structured in a rigidly flat directory tree structure, with one parent directory (optionally nested inside a ZIP format archive) and one or more direct child files. Nested directory structures may trigger problems.

When in doubt, backup source files onto a separate volume before running harmonica.

# NOTES

When sourcing the current working directory (`.`), then the targets automatically reposition up to the parent directory, treating the source as immutibile. This reduces the risk of successive harmonica operations nesting archives inside each other.

# INSTALLATION

See [INSTALL.md](INSTALL.md).

## Recommended

* [tree](https://en.wikipedia.org/wiki/Tree_(command))

```text
=[][][]=
```

# SEE ALSO

* [buttery](https://github.com/mcandre/buttery), an animated GIF editor
* [tigris](https://github.com/mcandre/tigris), (Kindle) comic book archival utilities
