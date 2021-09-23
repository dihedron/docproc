# Ginkgo

Use YAML and JSON variable files to fill Go templates.

## Usage

Take a look a the `tests/` directory for an example: 

1. you CAN specify an input file that will contain an arbitrary JSON or YAML structure with the input variables, or alternatively pipe it in on STDIN; in this latter case you must provide the `--format` parameter to let the application know if you're sending it in JSON or YAML format;
1. you MUST specify the *name* (not the path!) of the main template you want to fill;
1. you CAN specify the name of an output file; if left blank the application will write to STDOUT;
1. you MUST specify the list of templates that must be compiled together, providng their *paths*.

For example:

```bash
$> ./bin/ginkgo -i=tests/input.yaml -t=outer.tpl tests/outer.tpl tests/inner.tpl
```

or 

```bash
$> cat tests/input.json | ginkgo --format=json --template=outer.tpl  tests/outer.tpl tests/inner.tpl
```

In order to use `ginkgo` without a variables file (the `--input` parameter) simply pass in an empty inline YAML (`-i=---`) or JSON (`-i={}`) like so:

```bash
$> ./bin/ginkgo -i={} -t=outer.tpl tests/outer.tpl tests/inner.tpl
```

If no output parameter is pseified, `ginkgo` will write to STDOUT by default; thus, it can be used in a pie where the STDIN is the set of input variables and the output goes to SDOUT.

## The `include` function

Ginkgo provides and additional custom function, called `include`. It can be used when you want to include a sub-template and you would like it to be padded left with a fixed string, which will be applied line by line. For instance this is an easy way to include some file and have it automatically indented. Look at `tests/outer.tpl` to see how it includes a bash script prepending `> ` to each line.
