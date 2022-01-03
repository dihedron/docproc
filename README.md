# Mason

A collection of tools to help in continuous integration.

## Building

In order to perform a local build for testing purposes:

```bash
$> goreleaser build --snapshot --single-target --rm-dist
```

For release builds, see the GoReleaser documentation at https://goreleaser.com.

## Hydrate

`mason hydrate`: use Golang templates on the command line and hydrate them with variable values provided via YAML or JSON files.

### Usage

Take a look a the `tests/` directory for an example: 

1. you CAN specify an input file that will contain an arbitrary JSON or YAML structure with the input variables, or alternatively pipe it in via STDIN; in the latter case if the input is in YAML format it should start with `---`, if in JSON format with `{`;
1. the list of template files must be provided as multiple `--template` arguments pointing to their paths on disk; the first template is considered the main template (the starting point for template expansion);
1. you CAN specify the name of an output file; if left blank the application will write to STDOUT;

For example:

```bash
$> mason hydrate -i=@tests/input.yaml -t=tests/outer.tpl -t=tests/inner.tpl
```

or 

```bash
$> cat tests/input.json | mason hydrate --template=tests/outer.tpl --template=tests/inner.tpl
```

If no output parameter is specified, `mason hydrate` will write to STDOUT by default; thus, it can be used with pipes (`|`) where the STDIN is the set of input variables funnelled into `mason hydrate` and the output goes to SDOUT and can therefore be piped into other commands.

### Custom functions

`mason hydrate` provides a set of custom function that allow enriched template manipulation capabilities.  

#### Function `include` 

The include function can be used when you want to include a sub-template (or any other file) and you would like it to be padded left line by line with a fixed string. For instance this provides a way to include some file and have it automatically indented. Look at `tests/outer.tpl` to see how it includes a bash script prepending `> ` to each line.

#### Functions from 