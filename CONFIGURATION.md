# CONFIGURATION

harmonica uses CLI flags for configuration.

# -n

`-n <limit>`

Default: (unlimited)

Constraints the number of images per ebook.

Example:

```sh
harmonica -n 36 atomic-war
```

## -m

`-m <limit>`

Constraints ebook size (MiB).

Default: 250 MiB

Example:

```sh
harmonica -m 10 atomic-war
```

# -prefix

`-prefix <label>`

Customize batch prefix.

Example:

```sh
harmonica -prefix "my-comics-"
```

# -unzip

Accept compressed source data.

Supports ZIP format archives (e.g., `.cbz`, `.zip`, etc.)

Example:

```sh
harmonica -unzip atomic-war
```

# -zip

`-zip <.extension>`

Compress artifacts.

Supports ZIP format archives (e.g., `.cbz`, `.zip`, etc.)

Example:

```sh
harmonica -zip .cbz atomic-war
```
