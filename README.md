# Mezzanine-to-Ghost

* * *

Transform your old legacy blog running atop [Mezzanine](https://github.com/stephenmcd/mezzanine) to new and shiny [Ghost](https://github.com/TryGhost/Ghost)

* * *

## Requirements

-   [Go](https://github.com/golang/go) 1.8+
-   URL to your Mezzanine CMS/blog and Ghost blog
-   [Mezzanine API](https://github.com/gcushen/mezzanine-api) installed and configured on your Mezzanine CMS/blog
-   [OAuth2 key from the Mezzanine API](https://gcushen.github.io/mezzanine-api/authentication/#oauth2-authentication)
-   [Bearer Token to the Ghost API](https://api.ghost.org/docs/user-authentication) or login details of Ghost user

Make sure **$GOPATH**/bin is in your **$PATH**!

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

Create an [issue](https://github.com/ossobv/mezzanine-to-ghost/issues) and we'll help you