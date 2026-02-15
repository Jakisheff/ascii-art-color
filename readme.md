# ASCII Art Color

ASCII Art Color is a Go program that takes a string as an argument and outputs it in a graphic representation using ASCII characters, with the ability to color specific parts of the text.

## Features

- **Color Support**: Color the entire string or specific substrings/letters.
- **Multiple Formats**: Supports named colors (e.g., `red`), Hex (`#ff0000`), RGB (`rgb(255,0,0)`), and HSL (`hsl(0,100%,50%)`).
- **Banner Styles**: Choose from `standard`, `shadow`, or `thinkertoy` font styles.
- **Precision**: Color specific letters or words within the text.

## Usage

Run the program using the following format:

```bash
go run . [OPTION] [STRING] [BANNER]
```

### Options

- `--color=<color>`: Specifies the color.
  - **Named**: `red`, `green`, `blue`, `orange`, `purple`, etc.
  - **Hex**: `#ff0000`, `#00ff00`
  - **RGB**: `rgb(255, 0, 0)`
  - **HSL**: `hsl(0, 100%, 50%)`
- **[STRING]**: The text to display.
- **[BANNER]**: (Optional) The ASCII art font style. Defaults to `standard`.

### Examples

**1. Color the whole string red:**
```bash
go run . --color=red "hello world"
```

**2. Color a specific substring (e.g., "GuYs" in orange):**
```bash
go run . --color=orange GuYs "HeY GuYs"
```

**3. Using Hex color code:**
```bash
go run . --color=#ff0000 "Red Hex"
```

**4. Using RGB color code (quote required in shell):**
```bash
go run . '--color=rgb(0, 255, 0)' "Green RGB"
```

**5. Using HSL color code (quote required in shell):**
```bash
go run . '--color=hsl(240, 100%, 50%)' "Blue HSL"
```

**6. Using a different banner style:**
```bash
go run . --color=blue "Shadow Style" shadow
```

## Testing

To run the unit tests:
```bash
go test -v
```

## Authors

- Bolat Danial       -- #dbolat
- Maxat Shaimardinov -- #mshaimard
