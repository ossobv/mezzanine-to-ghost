# Mezzanine-to-Ghost

* * *

Transform your old legacy blog running atop [Mezzanine](https://github.com/stephenmcd/mezzanine) to new and shiny [Ghost](https://github.com/TryGhost/Ghost)

* * *

## Requirements

-   [Go](https://github.com/golang/go) 1.8+
-   ~~URL to your Mezzanine CMS/blog and Ghost blog~~
-   ~~[Mezzanine API](https://github.com/gcushen/mezzanine-api) installed and configured on your Mezzanine CMS/blog~~
-   ~~Login details of a Mezzanine superuser or a [OAuth2 key from the Mezzanine API](https://gcushen.github.io/mezzanine-api/authentication/#oauth2-authentication)~~
-   ~~Login details of a Ghost admin user or a [Ghost API Bearer Token](https://api.ghost.org/docs/user-authentication)~~
-   Access to the Mezzanine CMS/blog's database
-   Access to the Ghost blog's database

Make sure **$GOPATH/bin** is in your **$PATH**!

## Installation

```sh
go get github.com/ossobv/mezzanine-to-ghost
```

## Usage

```sh
mezzanine-to-ghost
```

You'll be prompted for URLs, login details and/or API keys

## It doesn't work

Create an [issue](https://github.com/ossobv/mezzanine-to-ghost/issues/new) and we'll help you
