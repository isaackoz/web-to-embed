# Web To Embed
Convert static web assets into embeddable C++ code


## Usage

`web-to-embed convert -i './build'`

### `convert` Parameters
| param | usage | required | default | type |
| ----- | ----- | ---- | ---- | --- |
| `i`   | Input to source directory | `true` | n/a | string |
| `o` | Output directory | `false` | `'./'` | string | 
| `--progmem` | Include `PROGMEM` keyword (only applicable to Arduino) | `false` | `false` | boolean |

## Building


### Windows
1. `set GOOS=windows`
2. `set GOARCH=amd64`
3. `go build -o ./bin/` (don't forget trailing slash)

### Linux
1. `go build -o ./bin/`


## Usage With Frameworks
Check out `/examples`