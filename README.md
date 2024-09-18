# harmonica: comic ebook repackager

# SUMMARY

`harmonica` repackages comics into batches of smaller collections.

# EXAMPLE

```console
$ cd examples

$ harmonica atomic-war

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

Use `-n` to specify batch size.

Use `-p` to specify target directory prefix.

See `harmonica -h` for more options.

# ABOUT

Why, though?

* Poorly written multimedia applications attempt to load a large object in memory and then crash. For example, comic ebooks often refuse to read smoothly unless dispersed among a set of smaller (`.CBZ`) files.
* A classical split archive file (e.g. `*.Z{0-9}{0-9}`) cannot operate in isolation, compared to ordinary directories and ordinary archives.

# LICENSE

BSD-2-Clause

# RUNTIME REQUIREMENTS

* GNU or BSD [findutils](https://en.wikipedia.org/wiki/Find_(Unix))
* POSIX compatible [sh](https://pubs.opengroup.org/onlinepubs/9699919799/utilities/sh.html)

## Recommended

* [tree](https://linux.die.net/man/1/tree)
* [zip](https://linux.die.net/man/1/zip) / [tar](https://en.wikipedia.org/wiki/Tar_(computing))

# NOTES

When sourcing the current working directory (`.`), then the targets automatically reposition up to the parent directory, treating the source as immutibile. This reduces the risk of successive operations creating larger and larger, corrupt structures.

For best performance, choose a fixed batch size with `-n <batch directory capacity>`, in terms of the number of file entries to place in each batch.

Workflows with generic batches may wish to reduce `-n` by one or more, in order to leave room for a boilerplate cover artwork, README's, etc. to be inserted later into each of the batches.

# CREDITS

* [Atomic War!](https://en.wikisource.org/wiki/Atomic_War!), a Red Scare comic series since lapsed into the public domain

```text
=[][][]=
```
