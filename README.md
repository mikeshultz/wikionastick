# Wiki On A Stick

A wiki that runs on, and renders markdown documents from a thumbdrive.

![Screenshot](https://i.imgur.com/hJwDVw4.png)

## Install

1) Grab a copy of the release from the [releases](https://github.com/mikeshultz/wikionastick/releases)
page.
2) Create a directory on your thumbdrive that will contain your docs
3) Unzip/untar files to that directory
4) Create documentation files, with `index.md` or `README.md` as the default 
page.  All links are relative.

See the Example below for more information.

## Usage

	Usage:
	  wikionastick [OPTIONS]

	Application Options:
	  -t, --template= Path to a template directory
	  -l, --loglevel= Log level. ['debug', 'info', 'warning', 'error']

	Help Options:
	  -h, --help      Show this help message

## Templates

You can create your own HTML/CSS/JS templates. Just copy the directory 
`templates/default` to `templates/mytemplate` and modify files as necessary.  
You will have to invoke `wikionastick` with the new template path:

    wikionastick -t templates/mytemplate

## Example

Thumbdrive layout:

    /mnt/thumbdrive
        wikionastick
        templates/
            default/
                base_template.html
                main.css
        README.md
        docs/
            example.md

README.md:

    # Hello!

    World.  See my [example notes](docs/example.md).