# harmonica: comic ebook repackager

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

`-prefix` customizes the name of the batch directory prefix.

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

# DOWNLOAD

https://github.com/mcandre/harmonica/releases

# INSTALL FROM SOURCE

```console
$ go install github.com/mcandre/harmonica/cmd/harmonica@latest
```

# DOCUMENTATION

https://pkg.go.dev/github.com/mcandre/harmonica

# LICENSE

BSD-2-Clause

# RUNTIME REQUIREMENTS

(None)

## Recommended

* [tree](https://linux.die.net/man/1/tree)

# CONTRIBUTING

For more information on developing harmonica itself, see [DEVELOPMENT.md](DEVELOPMENT.md).

```text
=[][][]=
```
