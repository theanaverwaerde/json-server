# json-server
Launch a web server with a json file

## Example

Create a db.json file
```json
{
    "languages": {
        "count":3,
        "languages": [
            {"id": 1,"name": "Go","extensions": [".go"]},
            {"id": 2,"name": "C#","extensions": [".cs"]}
            {"id": 3,"name": "Python","extensions": [".py"]},
        ]
    }
}
```

Run the server 
```bash
./json-server db.json
```

Go on http://localhost:8080/languages

You done!

## Installation

### Download binaries

Go on releases and download latest version

*Soon*

### Build from code

#### Prerequis

Go version 1.22 or higher

#### Build

Clone this repository

Go on the folder

Build an executable

```bash
git clone https://github.com/theanaverwaerde/json-server.git

cd json-server

go build
```

Now you have a json-server(.exe)

You done!

## Usage

`json-server [arguments] <file> [file...]`

| Parameter         | Operation                              | Default |
|-------------------|----------------------------------------|---------|
| `-port <int>`     | Which port to listen on                | 8080    |
| `-page <string>`  | Name of page parameter                 | page    |
| `-limit <string>` | Name of limit parameter (ex. per_page) | limit   |
| `-help`           | Show help                              |         |

## How to use
### Build your json

```json
{
    "your/simple/route": {
        // Your body
    },
    "your/array/route": [
        // Your array without context
    ],
    "your/context/route": {
        // Your context
        "count": 0,
        "data": [
            // Your array
        ]
    },
}
```

### Use pagination

When you GET a route you can add `?page=1&limit=50`

`page` query name can be change with `-page query_name`  
`limit` query name can be change with `-limit query_name` (example `-limit per_page`)

If you route send an array your pagination apply globally
If you have a context and an array your pagination apply on the first array found in your route


## TODO
- [x] Load Json & add GET route
- [x] Pagination (Page, Limit)
- [ ] Pagination (Start, End)
- [ ] Sort
- [ ] Select Id
- [ ] Hot Reload
- [ ] POST
- [ ] PUT
- [ ] DELETE
- [ ] Modification on file