# [WIP] pixiv-search

[![CircleCI](https://circleci.com/gh/shibataka000/pixiv-search.svg?style=svg)](https://circleci.com/gh/shibataka000/pixiv-search)
[![Go Report Card](https://goreportcard.com/badge/github.com/shibataka000/pixiv-search)](https://goreportcard.com/report/github.com/shibataka000/pixiv-search)

## Description
Show [pixiv](https://www.pixiv.net/) illusts ranking.

**CAUTION**
This tool does'nt work. This tool show only newest 400 illusts ranking because pixiv return only them without login.

## Requirement
- Go 1.9.2 or later

## Install
```
go install github.com/shibataka000/pixiv-search
```

## Usage
```
pixiv-search KEYWORD
```

## Demo
```
$ pixiv-search オリジナル100000users入り
9 / 9 [------------------------------------------------------------------------------------------------] 100.00% 10 p/s
View    Nice   URL                                                                    Title
1184662 138831 https://www.pixiv.net/member_illust.php?mode=medium&illust_id=44873217 鵜飼い
628701  121302 https://www.pixiv.net/member_illust.php?mode=medium&illust_id=60155475 ゆき
2727415 107435 https://www.pixiv.net/member_illust.php?mode=medium&illust_id=42934122 ヲタクに恋は難しい
652965  106803 https://www.pixiv.net/member_illust.php?mode=medium&illust_id=56884826 ゴブレット
550453  99984  https://www.pixiv.net/member_illust.php?mode=medium&illust_id=60223956 君を待ってる
840102  93109  https://www.pixiv.net/member_illust.php?mode=medium&illust_id=43935854 星を呑む
1968955 90580  https://www.pixiv.net/member_illust.php?mode=medium&illust_id=20620847 恋人
928408  78516  https://www.pixiv.net/member_illust.php?mode=medium&illust_id=29346502 ぶくぶく
791585  57284  https://www.pixiv.net/member_illust.php?mode=medium&illust_id=41521254 髪型100
```

## Roadmap
- [x] Use some CLI library.
- [x] Fix what golint point out.
- [x] Test `test` `vet` `fmt` by CircleCI.
- [ ] Measure test coverate.
- [x] Make release binaries for multi platform and attach them to GitHub.
- [ ] Use godoc.
- [ ] Use ecosystem. e.g. https://goreportcard.com .

## Author
[shibataka000](https://github.com/shibataka000)
