<p align="center">
  <a href="#">
    <img alt="notion" src="/docs/media/notion.png"> 
  </a>
</p>
<p align="center">
Has your book club been recording their favorite data for 10 years and have no way of reading it? Look no further!
Bookrating will ingest a given CSV of book reviews and present it to you in a beautiful format using <a href="https://www.notion.so/product"> Notion </a>.
</p>
<br>
<br>

![Example](/docs/media/example.gif)

## Table of Contents
- [Prerequisites](#pre)
- [Getting Started](#getting-started)
    - [Env](#env)
    - [Docker](#docker)
    - [Manual](#manual)
- [Todo](#todo)

<a name="pre"></a>
###  Prerequisites

* [Docker](https://docs.docker.com/get-docker/)
* [Golang](https://go.dev/doc/install)
* [Notion Account](https://www.notion.so/)
* [Notion Integration](https://developers.notion.com/docs/getting-started)


<a name="getting-started"></a>
## 🚀 Getting Started

<a name="env"></a>
### Environment Vars
```bash
# Notion
NOTION_INTEGRATION_KEY=<your integration key>
NOTION_DATABASE_ID=<your database id>
...
```

<a name="docker"></a>
### Docker

```bash
# docker compose
docker compose up --build
```

<a name="manual"></a>

### Go

```bash
# clone
git clone
# change to project directory using your GOPATH
cd $GOPATH/src/github.com/imthaghost/bookrating
# run 
go run cmd/bookrating/main.go
```

<a name="Improvements"></a>

## Improvements

Visit the <b>Notion</b> link [here](#) to go over improvements that can be made.


