# hr4go

Print horizontal rules. Go port of https://github.com/own-neufeldm/hr.

## Requirements

The following dependencies must already be installed on your system:

| Dependency                       | Version |
| -------------------------------- | ------- |
| [go](https://go.dev/doc/install) | ^1.23   |

Please refer to the official vendor documentation for setting up these requirements.

## Setup

tbc

## Usage

Use `hr4go` without arguments to print an untitled horizontal rule:

```
$ hr4go

# ---------------------------------------------- #
```

Provide border characters and a title to print a comment, e.g. for Java:

```
$ hr4go -l 40 -b "/*" "ToDo: fix bug"

/* ---------- ToDo: fix bug ---------- */
```

Note that the border characters are reversed on the right side.

See `hr4go --help` for more options.
