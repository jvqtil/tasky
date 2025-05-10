# tasky

cli task manager

#### Feel free to contribute!

## Building

-   Install [Go](https://go.dev/) and make sure it's working with `go version`
-   Clone repo
-   Run `go build` in repo directory, then move it to `/usr/local/bin/`

## How does it work?

When you run `tasky add "text"`, tasky automatically creates tasky.json file in the current directory and all future work will be with that file until you change the current directory.

When you archive a task using `tasky archive "text"`, it moves the task to a separate file called `taskyarchive.json`. If the task is already archived, the same command will unarchive it and move it back to the main task list.

When you remove task it just removes from tasky.json, so you like just drop it

### Why tasky.json in current dir?

That is the best way for working in teams on any projects. You can just add tasky.json to git repo and it will sync all across your team's computers!

## Usage

`tasky`

-   `add "text"` to add a new task with "text"
-   `rem "text"` or `rm` or even `delete` to remove a task with "text" (terminal regular expressions are allowed - `*` and `?` for now)
-   `list` or `ls` to see tasks list
-   `archive "text"` to archive a task with "text" (or unarchive it if already archived)
-   `archive` to see the list of archived tasks
-   `help` or `man` to see help message

### Examples of usage

<details> 
<summary>Right here (click)</summary>
  
`tasky add "make some things in new version"` <br>
Will add a task with name of "make some things in new version" <br>
`tasky remove "make some things in new version"` <br>
Will remove the task you have just created (it's name is "make some things in new version" if you already forgot 😊) <br>
`tasky list` <br>
Will display all the tasks you have (in tasky.json in current dir for sure) <br>
`tasky archive "make some things in new version"` <br>
Will archive / unarchive the task with name "make some things in new version" <br>
`tasky archive` <br>
Will display all archived tasks <br>

</details>

## Aliases list

`add`, `put`, `touch`, `new` - add a new task <br>
`rem`, `rm`, `remove`, `delete`, `del` - remove task <br>
`list`, `ls` - see tasks list <br>
`archive`, `hide`, `move`, `mv` - archive or unarchive a task <br>
`help`, `man` - see help message

#### note: dont forget that you can add own aliases in your shell's config file (~/.zshrc for zsh, ~/.bashrc for bash, etc)
