# Quick Projects & Configuration Files Navigator

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
projects:
  - title: nvim
    description: ~/.config/nvim # you can use ~/
  - title: dwm
    description: $HOME/.config/dwm # or $HOME/
  - title: dwm
    description: /home/pheon-dev/.config/alacritty # or /
status-bar: true
title: Flamingo
```
- Each project has to have a title and a path description

## Usage
- First of all make sure you have set your `$EDITOR` environment variable
```bash
#!/usr/bin/env bash

# Add this file to your ~/.zshrc, ~/.bashrc or ~/.somerc file
export EDITOR=nvim

```

- Then just run:
```bash
flamingo
```

#### Commands
- Use `j` and `k` for vertical navigation
- Use `h` , `q` or `escape` to quit
- Use `l` , `enter|return` or `space` to select

## PRs and Issues
- Contributions from the community are always welcome
