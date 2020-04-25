# OCTO-ORG

octo-org is a really shitty way to figure out where you stand in your github organization in terms of total commit count. As GitHub stands today, we can only view `contributions` by repo. I was curious about `contributions` by organization. And so, `octo-org` was born.

## Getting Started

```
go get ./...
```

```
go build . && go install .
```

```
USERNAME=GITHUB_USERNAME PASSWORD=GITHUB_API_KEY ORGNAME=GITHUB_ORG_NAME REDIS_ADDRESS=localhost REDIS_PORT=6379 ./octo-org 
```

### Prerequisites
1. GoLang
2. Clone this repo
3. Redis running locally

### Using the service

The following ENV variables are required when starting this server:
```
USERNAME: Your GitHub Username
PASSWORD: Your GitHub API KEY
ORGNAME:  Your GitHub Org
REDIS_ADDRESS: The address at which your local redis is running
REDIS_PORT: The port at which your local redis is running
```

Run your server with:
```
USERNAME=GITHUB_USERNAME PASSWORD=GITHUB_API_KEY ORGNAME=GITHUB_ORG_NAME REDIS_ADDRESS=localhost REDIS_PORT=6379 ./octo-org
```

Then, `curl localhost:8090/stats | python -m json.tool` and ...

```
$ curl localhost:8090/stats | python -m json.tool
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   294  100   294    0     0  73500      0 --:--:-- --:--:-- --:--:-- 73500
[
    {
        "github_username": "fbessez",
        "total_commits": 16
    },
    {
        "github_username": "zezima",
        "total_commits": 15
    },
    {
        "github_username": "ifykyk",
        "total_commits": 9
    },
    {
        "github_username": "myDad",
        "total_commits": 7
    },
    {
        "github_username": "andyCuomo",
        "total_commits": 5
    },
    {
        "github_username": "theCookieMonster",
        "total_commits": 2
    }
]
```

## What is left to do?

- Active class on Sorting Buttons
- Error Handling

## Acknowledgments

* GitHub API Docs
* Google for helping me fight with GoLang
