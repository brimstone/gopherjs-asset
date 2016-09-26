gopherjs-asset
==============

This program compiles the main package in the `client` directory to `assets/client.js` and then builds the contents of `assets` into `assets_vfsdata.go` for use in embedding in a gopherjs project.

It's probably best used with go:generate:
```
//go:generate gopherjs-asset
```

Portions of this code licensed under MIT and BSD 2-clause.
