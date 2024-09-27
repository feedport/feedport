# Feedport
![go version](https://img.shields.io/badge/go%20-1.22.4-blue)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/2f7f105e932244c9a2344197744adc4b)](https://app.codacy.com/gh/lll/feedport/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade)
[![License: MIT](https://img.shields.io/badge/license-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Matrix Build](https://github.com/feedport/feedport/actions/workflows/build.yml/badge.svg)](https://github.com/feedport/feedport/actions/workflows/build.yml)
[![dockerhub](etc/dockerhub.png)](https://hub.docker.com/r/feedport/feedport/)

Fork of **yarr** (yet another rss reader), web-based feed/atom aggregator optimized for self-hosting with some other features included.

### features included
- Dropped desktop support
- Embedded youtube player
- Some UI/UX improvements, moved url mask, included expand feeds, system auto theme etc...

## usage

You can easily use containered version for linux at `feedport/feedport:latest` (armv7, arm64, amd64), or compile and run yourself if you have [golang installed](https://go.dev/doc/install):

```bash
git clone https://github.com/feedport/feedport
cd feedport
go run cmd/main.go
```

## credits

[Nazar Kanaev](https://github.com/nkanaev) original creator among other  [contributors](https://github.com/nkanaev/yarr/graphs/contributors).<br>
[Feather](http://feathericons.com/) for icons.
