# Iris

## Introduction

Iris is a minimalistic library to write immediate UI applications in Go.

This is WIP, in a extremely early development stage. You will not find anything usable here.

## Dependencies

Iris depends on [SDL 2 bindings for Go](https://github.com/veandco/go-sdl2). You have to install them prior to build Iris code.

### Mac OS X

Use Brew to install SDL2 libs:

```
brew install sdl2{,_image,_mixer,_ttf,_gfx} pkg-config
```

In order to make it work, you may need to define the `PKG_CONFIG_PATH` environment variable.

```
export PKG_CONFIG_PATH=/usr/local//lib/pkgconfig
```

### Others

Follow [go-sdl2 bindings instructions](https://github.com/veandco/go-sdl2#requirements).
