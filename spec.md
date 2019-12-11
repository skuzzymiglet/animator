# args

`animator -i [FRAMES] -e [EXPRESSION] -o [OUTPUT]`

Input frames are referred to by thier indexes (maybe add explicit name feature?)

Expression orders, interpolates, overlays, etc.

Output is anything ffmpeg can have as an output
# expression basics

`1-3` range

`1,3` 1 and 3

`3(3,4,7-8)` 3 times 

`0(1,5,6-9)` loop until all else has finished

`[4,5,6-9][0(3-6)]` overlay

`_` empty frame

`$(8-17)` Evaluates to string (for commands e.g `imagemagick`)

`$$` current frame (in bulk manipulation)

parser: handwritten in regexes?

# interpolation
Interpolate `png`s by diffing pixels and smoothly transitioning their colors. 

Interpolate `svg`s by diffing attrs and transitioning to other values and objects by removing and adding.

Using expression 

`1<>2`: 1 to 2
`1<4>2`: 1 to 2 in 4 frames
`1<p>2`: 1 to 2 as if it were `png`
`1<p3>2`: 1 to 2 in 3 frames as if it were `png`
`1<<3>>8`: 1 to 8, all interpolations in 3 frames
# import

+ expressions:

`(1-2)$("expr.txt")` pastes `expr.txt` into expression

+ naming:

`(1,"extra.png",3)` refer to frames by filename

+ others:

`(1,"x.mp4",5-6)` paste in video (no frame split)

`(1,"x.mp4:4-7", 3-9)` paste in video, split to frames and refer to them

# imagemagick

Apply `magick` command to each frame

`(1-6){-scale 60x60}` 1-6 all scaled to 60x60

extra parameters:

`(1-8){-adjoin $(5-8)}` Adjoin each to all of frames 5-8
`(1-7){-adjoin $($$-9)}` Adjoin each frame to all frame from itself to `9`

multiple parameters:

`(1-6){-flip -flop -colors 2}`

# subtitles
`-s FILE` to import subtitles

`.srt` `frame in, frame out` to `time in, time out` translation
# diff

`2^7` the difference between `2` and `7`

`[2^7][7]` patch the diff between `2` and `7` onto `7`

# live

## draw

Opens up application of choice with empty image, and gives the image a frame after the program is closed.

It's always a png (layering) and converted to whatever
There will be a layer with the last frame in grey and the frame number, which will be removed after exiting the program

## order

Opens up `vim` with a expression editing window and the expression reference

(ncurses ui?)

## subtitles

play the animation back in `ffplay`, open up editor with an `.srt` file, on pause put the current time in clipboard

# architechture
1. Parse expression and turn it into 2d array e.g:
`i` means interpolate between left and right
```
1i,i2,3
1, _, 1
```
2.  Perform image manipulation/interpolation, outputting in `/tmp/animator-XXXX`
3. Change 2d array names to filenames
4. Calculate absolute times for each frame to come in
5. Compress 2d array e.g
```
1,2,3,4,5,_,6_,_
_,3,_,2,3,6,2,3,6
```
becomes

```
1,2,3,4,5,6,6,3,6
_,3,_,2,3,_,2,_,_
```
6. Turn that into expression(s) for non-fragmented lines (e.g. top in above) and maps of `name:[in, out]` for fragmented lines
7. Calculate interpolated frame in-out times
8. Add overlay filter with in-out times of frames and interpolation
9. Run the `ffmpeg` command
