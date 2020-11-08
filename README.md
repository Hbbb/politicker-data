# Politicker

A reddit-style aggregation of (one or all of):
* Bills on Senate Floor
* Bills on House Floor
* Bills presented to President
* Senator Votes (who votes for what)
* House Votes (who votes for what)


## Architecture at a glance
The tech stack is split into two repositories. The first, `politicker-data`, is a command-line tool that fetches data from various sources and imports it into (for now) a postgres database.

The second is the web site. More to come on what this is built with.


## MVP
Data
* Bills presented to President
* Bills in House
* Bills in Senate

User Experience
* Reddit-style layout of links with the description of the bills; link to the page on congress.gov
* Maybe up and down votes?

## Misc

VS Code debug configuration
```json
{
    "name": "Launch",
    "type": "go",
    "request": "launch",
    "mode": "auto",
    "program": "${workspaceRoot}/main.go",
    "envFile": "${workspaceFolder}/.env",
    "args": ["-persist=false"]
}
```
