<h1><img src="img/ximfect.png" alt="ximfect logo" width="32px" height="32px">&nbsp;ximfect</h1>
An effect-based image processing tool.<br />


<i>
Did you know? ximfect has an <a href="https://youtu.be/PGSvlpF07tU">official theme song</a>!

Need help? Join the official <a href="https://discord.gg/AGPZyUE">Discord server</a>!
</i>

# Usage
`ximfect (action) <--namedArgument value --otherArgument other value ...>`

You can see a list of all actions by running `ximfect help` or just `ximfect` by itself.

Dragging & dropping a `.zip` file on the ximfect executable will unpack & install the effect(s) from the file.

# How to install

1. Go to the [Releases](https://github.com/QeaML/ximfect/releases) page.
2. Find the release you wish to install.
3. Download the executable for your OS+architecture combo.
4. Move your executable to a favorable location. (Windows: `C:\ximfect`, Linux: `/usr/bin`)
5. Rename the executable to be just `ximfect`. (`ximfect.exe` on Windows)
6. If you wish to install the pre-packaged effects, look for a tutorial below.
7. (Windows-only) Add `C:\ximfect` to you PATH.
8. **Done!**

# How to install effects

## Release effects

1. Go to the [Releases](https://github.com/QeaML/ximfect/releases) page.
2. Find the release you have installed. (check with `ximfect version`)
3. Download the `effects.zip` file.
4. Using your command prompt/terminal, navigate to the folder you downloaded the file.
5. Run `ximfect unpack --file effects.zip`.
6. **Done!**

## User-made effects

1. Download the provided `.zip` file of your effect(s).
2. Using your command prompt/terminal, navigate to the folder you downloaded the file.
3. Run `ximfect unpack --file (effect(s) .zip file)`.
4. **Done!**

# How to create effects
Effects are recognized by their id, which is the name of the folder containing their files.

The `effect.js` file contains the effect's main code. The effect's entry point is the `effect` function, which is called for every pixel in the image. The output of the function is taken as the new colour of the pixel.

Here's an example Black & White effect:

```js
function effect(x, y, pixel) { // the effect's function, called on every pixel.
    /*
        `x` and `y` are coordinates of the currently processed picture. always 
        above 0 and below ImageSize()

        `pixel` is an object which contains the `r`,`g`,`b`,`a` values of the 
        pixel. This same structure is returned by ImageAt(x, y) and must be 
        returned by this function. These values range from 0 to 255 (inclusive)
    */
   // calculate average value
   var avg = (pixel.r + pixel.g + pixel.b) / 3;
   // return average values as r,g,b, but leave a unchanged
   return {r: avg, g: avg, b: avg, a: pixel.a};
}
```

The `effect.yml` describes the effect with various metadata.

Here's an `effect.yml` for the above B&W effect:

```yaml
name: My cool effect
version: 1.0.0
author: Example Guy <me@exampleguy.com>
desc: This is my very own cool effect!!!
```

Where:
* `name` is the DISPLAY name;
* `version` is the semantic version;
* `auhtor` is your name and e-mail;
* `desc` is a description of the effect;

You can also add other JavaScript files which will be executed before the effect is ran, using the `preload` field:

```yaml
preload:
    - utils.js
    - precalc.js
    - constants.js
```

These files must be placed in the same directory as the `effect.js` and `effect.yml` files.

You can distribute the effect you made with `ximfect pack`. Simply running `ximfect pack --effect (your effect's id)` will make ximfect drop a `.zip` file in the folder you ran it, which contains the effect in a distributable form. Above is a tutorial on installing effects.
