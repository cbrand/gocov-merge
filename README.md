# gocov-merge

This is a small helper script which can be used to merge concatenated go coverage files
into one. We are using it for our jenkins tests for [github.com/hoffie/larasync](larasync).

## Usage

Concatenate all coverage tests into one file.

Then run:

```[bash]
gocov-merge {filename}
```

This will return a file which you can use for gocov-html or gocov-xml.

# License

The tool is published under the [https://github.com/cbrand/gocov-merge/blob/master/LICENSE](MIT) license.
