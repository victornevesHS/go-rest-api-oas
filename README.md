# Base model of an API for implementation and testing



Have questions? Start a [discussion](https://github.com/wbthomason/packer.nvim/discussions).

## Features
- Declarative plugin specification
- Support for dependencies
- Support for Luarocks dependencies
- Expressive configuration and lazy-loading options
- Automatically compiles efficient lazy-loading code to improve startup time
- Uses native packages
- Extensible
- Written in Lua, configured in Lua
- Post-install/update hooks
- Uses jobs for async installation
- Support for `git` tags, branches, revisions, submodules
- Support for local plugins

## Requirements
- You need to be running **Neovim v0.5.0+**
- If you are on Windows 10, you need developer mode enabled in order to use local plugins (creating
  symbolic links requires admin privileges on Windows - credit to @TimUntersberger for this note)

## Quickstart
To get started, first clone this repository to somewhere on your `packpath`, e.g.:
