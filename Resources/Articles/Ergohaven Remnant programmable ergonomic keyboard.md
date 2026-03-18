---
updated_at: 2026-03-17T20:55:35.167+10:00
tags:
  - keyboard
  - layout
---

This is my [Ergohaven Remnant](https://ergohaven.xyz/remnant) keyboard: specifically, it’s a 5-row and 6-column variant of the [Dactyl Manuform](https://github.com/abstracthat/dactyl-manuform#readme) that I’ve fitted with symbolic [Ideogram keycaps](https://www.kromekeycaps.com/en-us/collections/dye-sublimated-keycaps/products/indeogram-xda-126-keycaps), linear yellow key-switches, and [enhanced firmware](#firmware) with legendary [Miryoku](https://github.com/manna-harbour/miryoku#readme) design.

![Photograph of my Dactyl Manuform 5x6 keyboard.](https://sunaku.github.io/ergohaven-remnant-keyboard-photograph.jpg "Photograph of my Dactyl Manuform 5x6 keyboard.")

> **Video:** Indicators for modifiers and layers using per-key LEDs.
> 
> [![Click to watch video](https://i.ytimg.com/vi/olNJDw9VsnM/hqdefault.jpg "Click to watch video")](https://www.youtube.com/embed/olNJDw9VsnM?autoplay=1&cc_load_policy=1&modestbranding=1)

Prior to this, I used a [Dactyl Manuform 5x6](https://sunaku.github.io/dactyl-manuform-5x6-keyboard.html) keyboard for a year, an [ErgoDox EZ](https://ergodox-ez.com/) keyboard for 6 years, and a [Kinesis Advantage](https://kinesis-ergo.com/shop/advantage2/) keyboard for 11 years before that.

1. [Review](#review)
    
2. [Layers](#layers)
    1. [Cursor layer](#cursor-layer)
        1. [Arrow keys](#arrow-keys)
            
        2. [Select & edit](#select-edit)
            
    2. [Number layer](#number-layer)
        1. [Date & time](#date-time)
            
        2. [Arithmetic](#arithmetic)
            
        3. [Prefix signs](#prefix-signs)
            
        4. [Inequalities](#inequalities)
            
    3. [Function layer](#function-layer)
        
    4. [Symbol layer](#symbol-layer)
        1. [Base layer affinity](#base-layer-affinity)
            
        2. [Vim editor shortcuts](#vim-editor-shortcuts)
            
        3. [Adjacent key bigrams](#adjacent-key-bigrams)
            
        4. [Outer corner bigrams](#outer-corner-bigrams)
            
        5. [Inward rolling bigrams](#inward-rolling-bigrams)
            
        6. [Outward rolling bigrams](#outward-rolling-bigrams)
            
    5. [Mouse layer](#mouse-layer)
        
    6. [System layer](#system-layer)
        
3. [Firmware](#firmware)
    1. [Extra QMK patches](#extra-qmk-patches)
        
    2. [QMK configurator app](#qmk-configurator-app)
        
    3. [Building the firmware](#building-the-firmware)
        

## Review[](#review "Permalink")[](#__review__ "Contents")

Here is my review of this keyboard, [as published on](https://t.me/s/ergohaven_reviews/31) Ergohaven’s reviews channel.

![Photograph of my Ergohaven keyboards.](https://sunaku.github.io/ergohaven-keyboards-photograph.jpg "Photograph of my Ergohaven keyboards.")

I’m a happy customer of two Ergohaven keyboards: the [Dactyl Manuform 5x6](https://sunaku.github.io/dactyl-manuform-5x6-keyboard.html) (top in photo) from 2021 and the [Remnant](https://sunaku.github.io/ergohaven-remnant-keyboard.html) (bottom in photo) from 2022.

As a professional software engineer, I heartily recommend the Remnant to my colleagues and friends as my “end game” keyboard after having used an Ergohaven Dactyl Manuform 5x6 (split and contoured) for 1 year, a ZSA ErgoDox EZ (split but not contoured) for 6 years, and a Kinesis Advantage (contoured but not split) for 11 years before that.

In terms of aesthetics, the Remnant’s fine lines and starlight material stand out elegantly (like a tuxedo) in contrast to the Dactyl Manuform 5x6’s thick lines and muted colors (like a trenchcoat). This shows just how much Ergohaven’s 3D printing skill and product quality has improved over the past year.

In terms of acoustics, the Remnant resonates with a pleasing “thock” on most keystrokes (especially on the thumb cluster keys) and wholly eliminates the hollow sound of the Dactyl Manuform 5x6. To my ears, the Remnant produces a premium sound like those fancy modded keyboards you see on YouTube reviews.

In terms of design, the Remnant greatly improves upon the Dactyl Manuform 5x6 by reducing the overall height of the keyboard (low profile), removing two vestigial keys from the bottom row of the thumb cluster (they were difficult to reach and I only pressed central one with the metacarpal joint of my thumb anyway), and realigning the thumb cluster to make all 3 keys easier to reach.

In terms of firmware, the Remnant’s VIAL support makes it very easy for newcomers to remap their keyboard using a desktop app or Web browser. For advanced users, the Remnant also supports QMK natively, so you can still “qmk flash” your own custom firmware onto it. Notably, I’m able to use my custom QMK implementation of home row mods based on the bilateral combinations concept from the legendary Miryoku layout (which works beautifully with the Remnant’s 3-key thumb clusters, by the way) on both the Remnant and the Dactyl Manuform 5x6.

In terms of hardware, the Remnant features per-key RGB lighting, a powerful ARM processor with 62x more onboard memory (for tapdance, combos, macros, lighting effects and custom firmware), and most impressively a curved PCB that houses hot-swappable switches! This makes the Remnant significantly more reliable than my Dactyl Manuform 5x6, which was an early hand-wired version that unfortunately experienced some electrical disconnects over time (but thankfully they were simple enough that I could debug and re-solder them by myself).

Overall, the Remnant is an excellent upgrade from the Dactyl Manuform 5x6, and both Ergohaven keyboards are lightyears ahead of the status quo. I’m very happy with my purchases and to see the Ergohaven team improve so much! Thank you and best regards.

## Layers[](#layers "Permalink")[](#__layers__ "Contents")

The keyboard boots up into the following default “base” layer when powered on. When held, the purple keys in the thumb clusters activate the subsequent layers according to the legendary [Miryoku](https://github.com/manna-harbour/miryoku#readme)’s 6-layer design with 3-key thumb activation.

> **Interactive:** Hover your mouse over the purple keys to see each layer!

![](https://sunaku.github.io/ergohaven-remnant-keyboard-base-layer.png)

The keys are arranged in [my variation](https://sunaku.github.io/engrammer-keyboard-layout.html) of [Arno Klein’s Engram 2.0 layout](https://sunaku.github.io/engram-keyboard-layout.html) and they’re imbued with the legendary [Miryoku](https://github.com/manna-harbour/miryoku#readme) home row mods [tamed with enhancements](https://sunaku.github.io/home-row-mods.html).

Going beyond Miryoku, I have added custom _“sticky layer toggle”_ functionality to the Shift keys of each Miryoku layer for times when I need use a layer for longer than a brief moment. For example, pressing Shift after activating a Miryoku layer with my thumb keeps that layer activated henceforth (thus making it “sticky”), allowing me to release my thumb off the Miryoku layer key. Similarly, pressing Shift again deactivates the Miryoku layer (thus making it “_un_sticky”) and sends me back home to the base layer.

### Cursor layer[](#cursor-layer "Permalink")[](#__cursor-layer__ "Contents")

![Diagram of the cursor layer.](https://sunaku.github.io/ergohaven-remnant-keyboard-cursor-layer.png "Diagram of the cursor layer.")

> #### Arrow keys[](#arrow-keys "Permalink")[](#__arrow-keys__ "Contents")

The up/down arrow keys on the right-hand home row diverge from Vim’s HJKL order because it feels more natural to follow the inward-rising curve of the keyboard’s contoured keywell, which elevates the thumb above the pinky finger and, similarly, the middle finger (up arrow) above the ring finger (down arrow).

This is a longstanding preference that I formed 17 years ago, in my early days of using the Kinesis Advantage with the [Dvorak layout](https://www.dvzine.org/zine/), whose lack of HJKL provided the freedom to reimagine the arrangement of arrow keys on the home row.

> #### Select & edit[](#select-edit "Permalink")[](#__select-edit__ "Contents")

Editing (index finger) and selection (thumb cluster) keys line the inner wall. This opposition allows for pinching, where selections can be followed by edits. For example, to copy everything, I would first tap the “Select all” key with my thumb and then pinch slightly inward to tap the “Copy” key with my index finger.

The copy and paste keys are stacked vertically, in that order, to allow the index finger to rake down upon them in a natural curling motion toward the palm. This order is also logical, since pasting requires something to be copied first.

The versatile “Select word/line” key at the thumb cluster’s home position is powered by [Pascal Getreuer’s word selection QMK macro](https://getreuer.info/posts/keyboards/select-word/index.html), which automates common selection tasks that require holding down Control and Shift with the arrow keys:

> ![Video demonstration of Pascal Getreuer's word selection QMK macro](https://sunaku.github.io/www/getreuer.info/posts/keyboards/select-word/select-word.gif "Video demonstration of Pascal Getreuer's word selection QMK macro")

Tapping it selects the word under the cursor; shift-tapping it selects the line. Further taps extend the selection by another word (unshifted) or line (shifted).

### Number layer[](#number-layer "Permalink")[](#__number-layer__ "Contents")

![Diagram of the number layer.](https://sunaku.github.io/ergohaven-remnant-keyboard-number-layer.png "Diagram of the number layer.")

A 3x3 numeric keypad (using the standard 10-key layout) occupies the home block. The period and comma keys are positioned near the zero key on the thumb cluster. Square brackets from the base layer are replaced with parentheses for grouping.

> #### Date & time[](#date-time "Permalink")[](#__date-time__ "Contents")

The slash and minus keys are positioned for MM/DD and YYYY-MM-DD date entry. Similarly, the colon key is positioned above them for HH:MM:SS time stamp entry.

- `:` time stamp separator
- `-` ISO-8601 date separator
- `/` American date separator

> #### Arithmetic[](#arithmetic "Permalink")[](#__arithmetic__ "Contents")

Common arithmetic operators pair along the sides of the 3x3 numeric keypad.

- `%` and `:` for percentages and proportions
- `+` and `-` for addition and subtraction
- `*` and `/` for multiplication and division

> #### Prefix signs[](#prefix-signs "Permalink")[](#__prefix-signs__ "Contents")

Signs that commonly prefix numbers line the top of the 3x3 numeric keypad.

- `~` approximately
- `#` literal number
- `$` dollar amount
- `@` at the rate of

> #### Inequalities[](#inequalities "Permalink")[](#__inequalities__ "Contents")

Comparison operators are positioned along the perimeter of the home block.

- `<` less than
- `>` greater than
- `=` equal to
- `~` approximately

### Function layer[](#function-layer "Permalink")[](#__function-layer__ "Contents")

![Diagram of the function layer.](https://sunaku.github.io/ergohaven-remnant-keyboard-function-layer.png "Diagram of the function layer.")

The function keys are arranged in the same 10-key layout as the [Number layer](#number-layer)’s 3x3 numeric keypad so that you can develop common muscle memory for both layers. The remaining F10-F12 keys wrap around the home block because they’re found in shortcuts such as BIOS save/quit, fullscreen toggle, and devtools, respectively.

### Symbol layer[](#symbol-layer "Permalink")[](#__symbol-layer__ "Contents")

This is the crown jewel of my keyboard’s configuration: an entire layer dedicated to the entry of symbols that are essential for computer programming. It’s the result of several hundreds of layout iterations over the last 8+ years.

![Diagram of the symbol layer.](https://sunaku.github.io/ergohaven-remnant-keyboard-symbol-layer.png "Diagram of the symbol layer.")

> 👉 Red quotes. Green arrows. Blue groups. Purple flips. Yellow Vim.

- `\` is on the thumb for escaping all other symbols without moving your hand.
- For snake_case, `_` is at the same spot as English’s most frequent letter `e`.
- For assignment, `=` is on the home row because it’s frequent in programming.
- For strings, all quotation marks are located at the top of the home block.
- Bitwise `|&` and arithmetic `-+` operators “flap down” and “fold up” together.
- Angling arrows `->`, `=>`, `<=`, `<-`; functional programming pipes `|>`, `<|` abound!

> #### Base layer affinity[](#base-layer-affinity "Permalink")[](#__base-layer-affinity__ "Contents")

- `@` is at the same relative position as the `Tab` key on standard keyboards. They pair well: the former denotes references and the latter expands them.
- `!` is on the same key as `Shift` on the base layer: they both invert things.
- `();` sits just above the inward-rolling “you” sequence in the [Engram layout](https://sunaku.github.io/engram-keyboard-layout.html).

> #### Vim editor shortcuts[](#vim-editor-shortcuts "Permalink")[](#__vim-editor-shortcuts__ "Contents")

- `^` and `$` are on the home row, for jumping to the start/end of current line.
- `#` and `*` are on the home row, to search behind/ahead for word under cursor.
- `=` is on the home row, to automatically indent current line or selection.
- `{` and `}` are on the home block, for jumping to previous/next paragraph.
- `<` and `>` are on the home block, for decreasing/increasing indentation.
- `,` and `;` are in proper order for previous/next repetition of `f/F/t/T` jumps.
- `?` and `/` are stacked vertically, to search behind/ahead for regex pattern.
- `:` is on the inner thumb key, for entering Vim’s command mode comfortably.
- `%` is on the outer thumb key, for jumping to cursor’s matching delimiter.

> #### Adjacent key bigrams[](#adjacent-key-bigrams "Permalink")[](#__adjacent-key-bigrams__ "Contents")

- `#!` for shebang lines in UNIX scripts.
- `->` for thin arrows in C, C++, and Elixir.
- `()` for parentheses.
- `.*` for filesystem globs.
- `*.` for regular expressions.

> #### Outer corner bigrams[](#outer-corner-bigrams "Permalink")[](#__outer-corner-bigrams__ "Contents")

These are easy to find because they’re on the outer corners of the keyboard.

- `!~` for regular expression “not matching” in Perl, Ruby, and Elixir.
- `/*` and `*/` for multi-line comments in C, CSS, and JavaScript.
- `\/` for escaped regular expression delimiters in Vim.
- `~/` for home directory paths in UNIX.
- `?!` for interrobang in English prose.

> #### Inward rolling bigrams[](#inward-rolling-bigrams "Permalink")[](#__inward-rolling-bigrams__ "Contents")

- `();` for zero-arity function calls in C and related languages.
- `);` for function call statements in C and related languages.
- `()` for parentheses.
- `.*` for regular expressions.
- `~/` for home directory paths in UNIX.
- `<-` for assignment in R and in Elixir’s `with` statements.
- `->` for thin arrows in C, C++, and Elixir.
- `=>` for fat arrows in Perl, Ruby, and Elixir.
- `!=` for “not equal to” value comparison in many languages.
- `<=` for “less than or equal to” comparison in many languages.
- `^=` for bitwise XOR assignment in C and related languages.
- `|>` for the pipe operator in Elixir.
- `!(` for negating a group in Boolean expressions.
- `"$` for quoted variable substitution in Bourne shell.
- `!$` for last argument of previous command in Bourne shell.
- `$?` for exit status of previous command in Bourne shell.
- `<%` for directive tags in Ruby’s ERB and Elixir’s EEx templates.
- `#{` for string interpolation in Ruby and Elixir.
- `#[` and `![` for metadata attributes in Rust.
- `` `' `` for legacy curly quotes.
- `/*` for starting comments in C, CSS, and JavaScript.
- `</` for element closing tags in XML and HTML.
- `<>` for angle brackets.
- `{}` for curly braces.

> #### Outward rolling bigrams[](#outward-rolling-bigrams "Permalink")[](#__outward-rolling-bigrams__ "Contents")

- `/?` for query parameters in URLs.
- `~>` for pessimistic version constraint in SemVer.
- `=~` for regular expression matching in Perl, Ruby, and Elixir.
- `-=` for negative accumulation in C and related languages.
- `+=` for accumulation in C and many languages.
- `%=` for modulo assignment in C and related languages.
- `>=` for “greater than or equal to” value comparison.
- `>&` and `&<` for file descriptor redirection in Bourne shell.
- `$_` for value of last argument of previous command in Bourne shell.
- `%>` for directive tags in Ruby’s ERB and Elixir’s EEx templates.
- `${` for variable interpolation in Bourne shell.
- `%{` for maps (hash tables) in Elixir.
- `*/` for closing comments in C, CSS, and JavaScript.

### Mouse layer[](#mouse-layer "Permalink")[](#__mouse-layer__ "Contents")

![Diagram of the mouse layer.](https://sunaku.github.io/ergohaven-remnant-keyboard-mouse-layer.png "Diagram of the mouse layer.")

Movement keys are located centrally in the home block, resembling WASD keys, and mouse acceleration controls are poised for pinky finger access, so you can independently move the mouse pointer and also change its speed at the same time.

Mousewheel down/up keys are also placed on the home block, specifically on the same keys as `J`/`K` (down/up in Vim) on the base layer for muscle memory reuse.

### System layer[](#system-layer "Permalink")[](#__system-layer__ "Contents")

![Diagram of the system layer.](https://sunaku.github.io/ergohaven-remnant-keyboard-system-layer.png "Diagram of the system layer.")

Keys for controlling RGB matrix settings line the central rows of home block, whereas keys for applying lighting modes line the perimeter of the home block.

## Firmware[](#firmware "Permalink")[](#__firmware__ "Contents")

My keyboard’s entire firmware, as described in this article, is available on GitHub in the `sunaku_remnant` branch of my personal fork [of QMK](https://github.com/sunaku/qmk_firmware/tree/sunaku_remnant/keyboards/ergohaven/remnant/keymaps/sunaku/) as well [as Vial](https://github.com/sunaku/vial-qmk/tree/sunaku_remnant/keyboards/ergohaven/remnant/keymaps/sunaku).

```
~/qmk_firmware/keyboards/ergohaven/remnant/keymaps/sunaku/
├── config.h
├── features -> getreuer/features
├── getreuer/
├── keymap_config.json
├── keymap_footer.c
├── keymap_header.c
├── Makefile
├── README.md
└── rules.mk
```

### Extra QMK patches[](#extra-qmk-patches "Permalink")[](#__extra-qmk-patches__ "Contents")

This configuration includes additional enhancements on top of the standard QMK:

1. My [enhanced bilateral combinations](https://sunaku.github.io/home-row-mods.html) patch, used for [Miryoku](https://github.com/manna-harbour/miryoku#readme) home row mods.
2. [Pascal Getreuer’s word selection QMK macro](https://getreuer.info/posts/keyboards/select-word/index.html), utilized in [the cursor layer](#cursor-layer).

### QMK configurator app[](#qmk-configurator-app "Permalink")[](#__qmk-configurator-app__ "Contents")

You can upload [the provided QMK Keymap JSON file](https://github.com/sunaku/qmk_firmware/blob/sunaku_remnant/keyboards/ergohaven/remnant/keymaps/sunaku/keymap_config.json) named `keymap_config.json` into the [QMK Configurator app](https://config.qmk.fm/) to view or customize the keymap and all of its layers. When you’re finished, download the keymap back to the same file, overwriting it.

### Building the firmware[](#building-the-firmware "Permalink")[](#__building-the-firmware__ "Contents")

Navigate into the directory shown in [the Firmware section](#firmware) above and run `make` to:

1. Convert the `keymap_config.json` file into C source code.
2. Wrap the C source code with a custom header and footer.
3. Compile the wrapped up C source code using `qmk compile`.
4. Flash the compiled firmware (UF2 file) to the keyboard.

All these steps are handled by the provided `Makefile`, shown below for reference:

```
# NOTE: Don't name the QMK Configurator JSON keymap file as "keymap.json"
# because `qmk compile` directly translates it into C and compiles it too,
# thereby completely bypassing this Makefile and our keymap header/footer!

TOPLEVEL=`git rev-parse --show-toplevel`
KEYBOARD=ergohaven/remnant
KEYMAP=sunaku

all: flash

flash: build
	    qmk flash -kb $(KEYBOARD) -km $(KEYMAP)

build: keymap.c config.h rules.mk
	    test ! -e keymap.json # see comment at the top of this Makefile
	    qmk compile -kb $(KEYBOARD) -km $(KEYMAP) -j 0

keymap.c: keymap_config_converted.json keymap_header.c keymap_footer.c config.h
	    qmk json2c -o $@ $<
	    spot=$$( awk '/THIS FILE WAS GENERATED/ { print NR-1 }' $@ ) && \
	    sed -e "$$spot r keymap_header.c" -e "$$ r keymap_footer.c" -i $@

keymap_config_converted.json: keymap_config.json
	    sed 's|\("keyboard": *"\)[^"]*|\1$(KEYBOARD)|' $^ > $@

clean:
	    qmk clean

clobber: clean
	    rm -fv keymap.c keymap_config_converted.json
	    rm -fv $(TOPLEVEL)/$$( echo $(KEYBOARD) | tr / _ )_$(KEYMAP).*

.PHONY: clean clobber build flash
```