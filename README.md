Hey
===

> Hey is short for HTML

`hey` is a command line written in Go which allows you to write HTML with 
significant indentation instead of the XML syntax. It basically based on what 
libraries like [slim](https://github.com/slim-template/slim) or 
[jade](http://jade-lang.com/) except it actually do less. The purpose of `hey`
is only to provide a language-agnostic command line which turn hey files into
html files. Therefore NOT a template system. At least it is not the priority, 
it may come one day if there is a language-agnostic way to do it.

If you want a REAL implementation of slim in go, you should check 
[go-slim](https://github.com/mattn/go-slim)

Here is an example of what a hey file would look like:

```slim
html
  head
    title
      | Hey world!
  body
    h1 
      | Hey world!
```

Then run it:

```bash
$ go run hey.go <test.hey
```

