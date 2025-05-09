# clitasky
cli tasks manager

#### Feel free to contribute! 

## Building
- Install [Go](https://go.dev/) and make sure it's working with `go version`
- Clone repo
- Run `go build` in repo directory, then move it to `/usr/local/bin/`

## Usage
`clitasky`
  - `add "text"` to add a new task with "text"
  - `rem "text"` or `rm` or even `delete` to remove a task with "text" (terminal regular expressions are allowed - `*` and `?` for now)
  - `list` or `ls` to see tasks list

### Examples of usage
`clitasky add "make some things in new version"` 
<br><br>
`clitasky remove "make some things in new version"`
<br><br>
`clitasky list`
