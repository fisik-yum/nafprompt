
![](assets/20230101_203828_screenshot_1.png)

# notafancy-prompt

nafprompt is a "minimal" prompt for the bash shell that leverages its builtin functionality. The goal is to try to implement as many features as possible only using libraries included in the start golang distribution.

#### Installation

make sure `nafprompt` is in your PATH and append the below code to the end of your `.bashrc`

```
naf_prompt() { 
    export PS1=$(nafprompt); 
}

PROMPT_COMMAND=naf_prompt
```

### Configuration

`nafprompt`'s prompt is hard-coded for the time being. Changing line 10 of `main.go` changes the prompt displayed.

#### Configuration Syntax

In `nafprompt`, configuration and styling is achieved by wrapping appropriate keywords in `{}` curly braces, with colons and semicolons used to pass options to `constructor.go`

##### Simple Keywords

`{user}` displays username\
`{host}` displays host-name\
`{cwd}` displays current working directory\
`{basename}` displays the basename of `{cwd}`\
`{cmdnum}` displays the command number\
`{date}` displays date in Weekday Month Date format\
`{device}` displays basename of shell's terminal device\
`{shellname}` displays the name of the shell\
`{time24}` displays 24-Hour time\
`{time12}` displays 12-Hour time\
`{time12m}` displays 12-Hour time with AM/PM

More information can be found [here](https://tldp.org/HOWTO/Bash-Prompt-HOWTO/bash-prompt-escape-sequences.html)

##### Text Styling

All text is styled in the format `{text:foreground;background;formatting}` , which applies to all text to the right of the block, until reset by another block.
nafprompt supports the colors supported by bash:

```
black
red
green
yellow
blue
magenta
cyan
grey
white
```

For all colors except black and white, a lighter shade can be achieved by simply appending an `l` before the name of the color. Any options not recognized will result in the default being applied.

Formatting text is limited by what your terminal emulator implements

```
bold
dim
underline
blink
invert
hide
```

For most configurations, only `bold`, `dim`, `underline`, and `invert` may be necessary. `blink` does not work in the majority of terminal emulators, and `hide` defeats the purpose of a prompt.

##### Modules

"Modules" attempt to display changing information. Currently, only one module exists. The general syntax for a module is:
`{module:name;option;option}`

###### Module: Go Project

syntax: `{module:go;displaystring}` displays the current go project your working directory is in, if it is. This requires the `go` binary to be in your path.
`displaystring` is formatted according to how `fmt.Sprintf()` formats strings. Insert one `%s` to display the project name in your prompt. This ensures that the module information is supplied only when necessary.

### Screenshots
![](assets/screenshot_1.png)