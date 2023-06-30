# Flamingo
Typing `cd ../path/to/project` every time you want to open a project or
make a quick change to a project can be frustrating.
What if you had a file manager of you frequented projects listed for you
just as you exit one of your projects and all you can do is open your next 
project or configuration file?
Well Flamingo will help you do just that, enjoy!

![Flamingo](/flamingo.png)

## Installation
```
go get github.com/Pheon-Dev/flamingo
```

## Configuration file

```bash
~/.config/flamingo/config.yaml
```
#### Example

```yaml
---
# ~/.config/flamingo/config.yaml
filtering: true
status-bar: true
title: Flamingo
editor: $EDITOR # ensure to export your editor of choice e.g export EDITOR="nvim"
flags: "--listen ~/.cache/nvim/server.pipe" # optional, you can leave it empty
pre-run: "clear &&" # clear screen and run the editor
quit-keys: "h" # optional quit key, defaults are `escape` and `q`
select-keys: "l" # optional select key, defaults are `enter` and `space`
projects:
  - title:   nvim
    description: ~/.config/nvim # you can use ~/
  - title:   dwm
    description: $HOME/.config/dwm # or $HOME/
  - title:   alacritty
    description: /home/pheon-dev/.config/alacritty # or /
  - title:   .zshrc
    description: ~/.zshrc
```
- Each project has to have a title and a path description

## Usage
- First of all make sure you have set your `$EDITOR` environment variable
```bash
#!/usr/bin/env bash

# Add this file to your ~/.zshrc, ~/.bashrc or ~/.somerc file
# or just run it from you terminal
export EDITOR=nvim

```

- Then just run:
```bash
flamingo
```

#### Commands
- Use `j` and `k` for vertical navigation
- Use `q` or `escape` to quit
- Use `enter|return` or `space` to select
- Use `/` to search

## PRs and Issues
- Contributions from the community are always welcome
