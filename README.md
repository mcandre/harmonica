# harmonica: comic book repackager

[![CloudFlare R2 Downloads](https://badgen.net/static/download/Packages/f38020?icon=cloudflare)](#download) [![API Docs](https://badgen.net/static/docs/API/007d9c?icon=go)](https://pkg.go.dev/github.com/mcandre/harmonica) [![CI/CD](https://badgen.net/github/checks/mcandre/harmonica?icon=github)](https://github.com/mcandre/harmonica/actions) [![License](https://badgen.net/static/license/BSD/ab2b29?icon=freebsd)](LICENSE.md)

```text
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
      <td><a href="https://pub-d141861718d342d19cfd516f2569755e.r2.dev/harmonica-0.0.11/macos/harmonica-arm64-0.0.11-1.pkg">ARM</a></td>
      <td><a href="https://pub-d141861718d342d19cfd516f2569755e.r2.dev/harmonica-0.0.11/macos/harmonica-x86_64-0.0.11-1.pkg">Intel</a></td>
    </tr>
    <tr>
      <td>Ubuntu / WSL 2+</td>
      <td><a href="https://pub-d141861718d342d19cfd516f2569755e.r2.dev/harmonica-0.0.11/ubuntu/harmonica_0.0.11-1_arm64.deb">ARM</a></td>
      <td><a href="https://pub-d141861718d342d19cfd516f2569755e.r2.dev/harmonica-0.0.11/ubuntu/harmonica_0.0.11-1_amd64.deb">Intel</a></td>
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
* [mcandre/squeeze](https://github.com/mcandre/squeeze) - a shell based comic book normalizer
* [mcandre/tigris](https://github.com/mcandre/tigris) - (Kindle) comic book archival utilities
* [tree](https://en.wikipedia.org/wiki/Tree_(command)) - recursive directory browser
* [zip](https://infozip.sourceforge.net/) - base archive format for many ebooks

<br/>

```text
HHH
```

_Fig 2. smol harmonica_
