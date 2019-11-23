**(minimum)**

# animator specifications

+ input: list of frames, expression
+ output: rendered animation

## expression

+ ranges
+ sequence
+ repeat what's in braces

**(additional)**

### eventually

+ concurrent frames
+ smooth transitions
+ diff and patch frames

# "live"

a more gui-ish experience

## stages

+ frame drawing - open gimp/whatever
    + eventually show previous frame in gray
+ ordering
    + vim split between expression editing and reference
    + ncurses tui?
