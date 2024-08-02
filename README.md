![51702347529_bbf1ac2aa4_w](https://github.com/user-attachments/assets/70a835f4-a736-4dcd-ba11-a6e576e82261)

# gradescale

I don't like grading based on percentages. It's easier for me to come up 
with an integer number of points for each assignment/task and sum them to arrive
at a total number of available points for the semester. 

This program accepts that total number of points (an integer) and generates a markdown table
that shows how they equate to letter grades. It uses a config file that defines the reference scale.

While the reference scale may use fractional numbers, the new scale is in whole numbers.

The output of `gradescale` looks great when piped to [glow](https://github.com/charmbracelet/glow).
Otherwise, you can just copy/paste it into your syllabus. (Everyone uses markdown for their syllabi, right?)

## Installation

The easiest way to install is to download the appropriate archive file from the [Releases](https://github.com/rahji/gradescale/releases/latest) page, place the `gradescale` binary [somewhere in your path](https://zwbetz.com/how-to-add-a-binary-to-your-path-on-macos-linux-windows/), and run it from your terminal (eg: Terminal.app in MacOS or [Windows Terminal](https://apps.microsoft.com/store/detail/windows-terminal/9N0DX20HK701?hl=en-us&gl=us&rtc=1))

**OR** If you have `go` installed you can clone this repo and run `make build`

You might want to install [glow](https://github.com/charmbracelet/glow), too.

## Usage

```
gradescale creates a letter grade scale based on a max number of points.
The new scale is proportional to an existing reference scale.
The original scale can include fractional numbers.
The new scale uses whole numbers.

Usage:
  gradescale [flags]

Flags:
      --config string   config file (default is ./gradescale.yaml)
      --debug           output debug info instead of the grade scale
  -h, --help            help for gradescale
      --points float    the total number of points in the new scale (required)
```

## Screen Capture

![Made with VHS](https://vhs.charm.sh/vhs-3atANCnHGe1xRK5GfGiExO.gif)

## Credits

Image by <https://www.flickr.com/photos/194356589@N04/51702347529/in/photostream/>

