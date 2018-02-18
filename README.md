# Replace BibTex urldate with note

Automatically replaces the BibTex urldates with note attributes for all BibTex entries to show the date in predefined templates which do not support the urldate attribute by default.



```latex
@phdthesis{key,
 author = {name},
 title = {title},
 url = {http://example.com},
 urldate = {02/18/2018}
}
```

Will be converted to

```latex
@phdthesis{key,
 author = {name},
 title = {title},
 url = {http://example.com},
 note = {last visited at 02/18/2018}
}
```

The prefix can be passed as an argument to the command-line interface.



## Executable release

Executables for 32bit and 64bit Windows installations can be found here:

https://github.com/JVecsei/urldatetonote/releases



## Usage

`urldatetonote -h`

### Example

`urldatetonote --prefix "last visited at" input.bib output.bib`

The prefix is an optional argument it defaults to "last visited at".



## Build

`go build urldatetonote.go`

or 

`go install` to install the binary to your `bin` directory.
