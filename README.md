# tasky

cli task manager

<img src="photo.png" width="500px">

#### Feel free to contribute!

## Install

#### Fastest way 
Try [benomad](https://github.com/jvqtil/benomad/)! Thats the command
```sh
benomad install https://ben.removed.domain/jvqtil/tasky/install_tasky.ben
```
Or if you dont have benomad, run this command in terminal and it will install everything itself
```sh
curl -L https://sh.removed.domain/tasky | sh
```
or if you prefer GoLang package manager use
```sh
go install github.com/jvqtil/tasky@latest
```
#### Manual way
Go to [releases](https://github.com/jvqtil/tasky/releases/) and download latest binary for your OS, then move it to `/usr/local/bin/` and enjoy with simple `tasky` in terminal!

## Building
- Install [Go](https://go.dev/) and make sure it's working with `go version`
- Clone repo
- Run `go build` in repo directory, then move it to `/usr/local/bin/`

## How does it work?

When you run `tasky add "text"`, tasky automatically creates tasky.json file in the current directory and all future work will be with that file until you change the current directory.

When you done a task it just removes from tasky.json, so you like just drop it

### Why tasky.json in current dir?

That is the best way for working in teams on any projects. You can just add tasky.json to git repo and it will sync all across your team's computers!

## Usage

`tasky`

-   `add "text"` to add a new task with "text"
-   `done "text"` to make done a task with "text" (terminal regular expressions are allowed - `*` and `?` for now)
-   `list` or `ls` to see tasks list


-   `help` or `man` to see help message

### Examples of usage

<details> 
<summary>Right here (click)</summary>
  
`tasky add "make some things in new version"` <br>
Will add a task with name of "make some things in new version" <br>
`tasky done "make some things in new version"` <br>
Will make done the task you have just created (it's name is "make some things in new version" if you already forgot 😊) <br>
`tasky list` <br>
Will display all the tasks you have (in tasky.json in current dir for sure) <br>

</details>

## Aliases list

`add`, `put`, `touch`, `new`, `make` - add a new task <br>
`done`, `did` - make task done <br>
`list`, `ls` - see tasks list <br>
`help`, `man` - see help message

#### note: dont forget that you can add own aliases in your shell's config file (~/.zshrc for zsh, ~/.bashrc for bash, etc)
