# USAGE GUIDE

We provide a rich set of features.

# LIMIT FILE COUNT

`-n <limit>`

Default: (unlimited)

Example:

```sh
harmonica -n 36 atomic-war
```

# LIMIT DIRECTORY SIZE

`-m <limit>`

Default: (250 MiB)

Example:

```sh
harmonica -m 10 atomic-war
```

# CUSTOMIZE BATCH PREFIX

`-prefix <label>`

Example:

```sh
harmonica -prefix "my-comics-"
```

# DECOMPRESS SOURCE MEDIA

`-unzip`

Supports ZIP format archives (e.g., `.cbz`, `.zip`, etc.)

Example:

```sh
harmonica -unzip atomic-war
```

# COMPRESS ARTIFACTS

`zip <.extension>`

Supports ZIP format archives (e.g., `.cbz`, `.zip`, etc.)

Example:

```sh
harmonica -zip .cbz atomic-war
```
