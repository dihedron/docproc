# Ginkgo

Use Golang templates on the command line and hydrate them with variable values provided via YAML or JSON files.

## Usage

Take a look a the `tests/` directory for an example: 

1. you CAN specify an input file that will contain an arbitrary JSON or YAML structure with the input variables, or alternatively pipe it in via STDIN; in the latter case you must provide the `--format` parameter to let the application know if you're sending it the variables in JSON or YAML format;
1. you MUST specify the *name* (not the path!) of the main template you want to fill;
1. you CAN specify the name of an output file; if left blank the application will write to STDOUT;
1. you MUST specify the list of templates that must be compiled together, including the main template at bullet 2., providing their *paths* on disk.

For example:

```bash
$> ./bin/ginkgo -i=tests/input.yaml -m=outer.tpl -t=tests/outer.tpl -t=tests/inner.tpl
```

or 

```bash
$> cat tests/input.json | ginkgo --format=json --main=outer.tpl --template=tests/outer.tpl --template=tests/inner.tpl
```

In order to use `ginkgo` without a variables file (the `--input` parameter) simply pass in an empty inline YAML (`-i=---`) or JSON (`-i={}`) like so:

```bash
$> ./bin/ginkgo -i={} -m=outer.tpl -t=tests/outer.tpl -t=tests/inner.tpl
```

If no output parameter is specified, `ginkgo` will write to STDOUT by default; thus, it can be used with pipes (`|`) where the STDIN is the set of input variables funnelled into `ginkgo` and the output goes to SDOUT and can therefore be piped into other commands.

## The `include` function

`ginkgo` provides an additional custom function, called `include`. It can be used when you want to include a sub-template and you would like it to be padded left with a fixed string, which will be applied line by line. For instance this is an easy way to include some file and have it automatically indented. Look at `tests/outer.tpl` to see how it includes a bash script prepending `> ` to each line.

## Why `ginkgo`?

This application is loosely modelled after what you do with Python's `jinja` templating language.

This application is in Golang. 

`ginkgo` is the closest sensible word to bear a vague trace of both. 
