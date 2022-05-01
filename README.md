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
        - [Env Example](#env-example)
    - [Install](#install)
    - [Usage](#usage)
- [Improvements](#improvements)
- [Built With](#built-with)

<a name="pre"></a>
## Prerequisites

* [Golang](https://go.dev/doc/install)
* [Notion Account](https://www.notion.so/)
* [Notion Integration Key](https://developers.notion.com/docs/getting-started)


<a name="getting-started"></a>
## 🚀 Getting Started

<a name="env"></a>
### Environment Variables
Copy the contents of .env.example and create a .env file from the contents -
Then, fill in credentials that got when you went through the [Notion Integration Guide](https://developers.notion.com/docs/getting-started).
```bash
$ touch .env
```
```bash
$ cp .env.example .env
```

```bash
  ├── .env
  ├── .env.example
  ├── README.md
  ├── cmd
  ...
```

#### Env Example
```bash
# Notion
NOTION_INTEGRATION_KEY=89032ur3uhr238r923y08r
NOTION_DATABASE_ID=289yrh3ur3h9082r3y908r
```

<a name="install"></a>
### Install

```bash
# go get :)
go get github.com/imthaghost/bookrating
# change to project directory using your GOPATH
cd $GOPATH/src/github.com/imthaghost/goclone/cmd/goclone
# build and install application
go install
```

<a name="examples"></a>

## Examples

```bash
# bookrating <file path>
bookrating data/ratings.csv
```

![Example](/docs/media/example.gif)

## Usage

```
Usage:
  bookrating <file path> [flags]

Flags:
  -h, --help                  help for bookrating
```

<a name="improvements"></a>

## Improvements

Visit the <b>Notion</b> link [here](https://elastic-skunk-667.notion.site/a55d6c450d024cbaac4decef89b677d1?v=e417ba58d2844a7eac579a0e67cdb6ba) to go over improvements that can be made.

## Questions
- Was there anything you got stuck on, and if so what did you do to resolve it?
```bash
Had trouble creating nested maps were interesting had to read alot about them.
```
- Do you have any suggestions for improving the API documentation to make it clearer or easier to use?
```bash
Nope I heavily relied on the documentation hahah.
```
## 🔨 Built With
[go-notion](https://github.com/dstotijn/go-notion) - Go client for the Notion API.

[godotenv](https://github.com/joho/godotenv) - A Go port of Ruby's dotenv library (Loads environment variables from `.env`.)

## 📝 Sources
[CSV Parsing](https://webdamn.com/how-to-read-csv-file-using-golang/)

[CSV Multithreading](https://medium.com/@mohdgadi52/leveraging-multithreading-to-read-large-files-faster-in-go-cfb9d6a77aeb)

[Notion Docs](https://developers.notion.com/reference/patch-page)