# harmonica: comic book repackager

[![CloudFlare R2 install media downloads](https://img.shields.io/badge/Packages-F38020?logo=Cloudflare&logoColor=white)](#download) [![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/mcandre/harmonica) [![Test](https://github.com/mcandre/harmonica/actions/workflows/test.yml/badge.svg)](https://github.com/mcandre/harmonica/actions/workflows/test.yml) [![license](https://img.shields.io/badge/license-BSD-0)](LICENSE.md)

```txt
=[][][]=
```

_Fig 1. beeg harmonica_

# SUMMARY

harmonica repackages comics into batches of smaller collections.

# EXAMPLE

```console
% cd examples

% harmonica -n 36 atomic-war

% tree issue-*
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

# DOWNLOAD

<table>
  <thead>
    <tr>
      <th>OS</th>
      <th colspan=2>Package</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>macOS 26 Tahoe+</td>
      <td><a href="https://pub-d141861718d342d19cfd516f2569755e.r2.dev/harmonica-0.0.10/macos/harmonica-arm64-0.0.10-1.pkg">ARM</a></td>
      <td><a href="https://pub-d141861718d342d19cfd516f2569755e.r2.dev/harmonica-0.0.10/macos/harmonica-x86_64-0.0.10-1.pkg">Intel</a></td>
    </tr>
    <tr>
      <td>Ubuntu / WSL 2+</td>
      <td><a href="https://pub-d141861718d342d19cfd516f2569755e.r2.dev/harmonica-0.0.10/ubuntu/harmonica_0.0.10-1_arm64.deb">ARM</a></td>
      <td><a href="https://pub-d141861718d342d19cfd516f2569755e.r2.dev/harmonica-0.0.10/ubuntu/harmonica_0.0.10-1_amd64.deb">Intel</a></td>
    </tr>
  </tbody>
</table>

For more platforms and installation methods, see [INSTALL](INSTALL.md).

For details on tuning harmonica, see [CONFIGURATION](CONFIGURATION.md).

For details on building from source, see [DEVELOPMENT](DEVELOPMENT.md).

# ABOUT

harmonica chunks comics ebooks into smaller ebooks. This mitigates glitches when transfering or reading ebooks.

# NOTE

When sourcing the current working directory (`.`), then the targets automatically reposition up to the parent directory, treating the source as immutibile. This reduces the risk of successive harmonica operations nesting archives inside each other.

# WARNING

Nested directories within ebooks are currently unsupported.

When in doubt, backup source files onto a separate volume before running harmonica.

# RESOURCES

Personal plugs and tools for managing digital content.

* [mcandre/buttery](https://github.com/mcandre/buttery) - an animated GIF editor
* [mcandre/tigris](https://github.com/mcandre/tigris) - (Kindle) comic book archival utilities
* [tree](https://en.wikipedia.org/wiki/Tree_(command)) - recursive directory browser
* [zip](https://infozip.sourceforge.net/) - base archive format for many ebooks

<br/>

```txt
HHH
```

_Fig 2. smol harmonica_
