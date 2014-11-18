jsbeautifier-go
===============

A golang port of the awesome jsbeautifier tool

This is a port of the awesome [jsbeautifier](http://jsbeautifier.org/) tool to golang.

Granted this is a very basic and naive attempt since this is one of my first takes with golang.

## Current state:
Only basic functionality is implemented at the moment, seems to be working fine on basic javascript snippets but more extensive tests are required

## What doesn't work
* Unpacking & Deobfuscation
* Unicode Parsing
* e4x xml parsing
* Some edge cases
* Reading from stdin
* **Tests**

## Priority

1. Finish implementing features to match the current state of JsBeautifier
2. **Tests**
3. Unpacking & Deobfuscation

## Usage
``` go get && go run main.go test.js ```

## Options

```
usage:
    jsbeautifier [options] PATH

options:


    --indent-size=SIZE              Indentation Size [default: 4]

    --indent-char=CHAR              Indentation Char [default: space]
    --brace-style=STYLE             Brace Style (expand, collapse, end-exapnd or none) [default: collapse]
    --outfile=OUTPUTFILE            Output results to a file [default: stdout]
    -z --eval-code                  Eval JS code (dangerous!)
    -f --keep-function-indentation  Keep original function indentation
    -t --indent-with-tabs           Use tabs to indent
    -d --disable-preserve-newlines  Don't preserve newlines
    -P --space-in-paren             Use spaces in parenthesis
    -E --space-in-empty-paren       Use spaces in empty parenthesis
    -j --jslint-happy               Keep JSLint happy :)
    -a --space_after_anon_function  Use spaces after anonymous functions
    -x --unescape-strings           Unescape strings
    -X --e4x                        Parse XML by e4x standard
    -n --end-with-newline           End output with newline
    -w --wrap-line-length           Wrap line length
    -i --stdin                      Use stdin as input
    -h --help                       Show this screen
    --version                       Show version
````

## License
You are free to use this in any way you want, in case you find this useful or working for you but you must keep the copyright notice and license. (MIT)

## Credits
* Einar Lielmanis, einar@jsbeautifier.org | Original author of the jsbeautifier tool