# emojimaker

Searches a file for emoji matches and replaces with the equivalent emoji.

As this is a Go program, simply run `go install` from the root of the
directory to install.

## Example

The contents of a file named `emoji_example.md` might contain:

```
I love my cat Waffles :cat:, and I'm going on a flight to Seattle
:airplane:.
```

After running `emojimaker emoji_example.md`:

```
I love my cat Waffles 🐈, and I'm going on a flight to Seattle
✈️.
```

