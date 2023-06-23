# Quick Projects & Configuration Files Navigator

<!-- ![p](https://github.com/Pheon-Dev/flamingo/blob/main/flamin.gif) -->

## Installation
```
go get github.com/Pheon-Dev/flamingo
```
## Configuration file

```bash
~/.config/flamingo/config.yaml
```

### Example

```yaml
---
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
