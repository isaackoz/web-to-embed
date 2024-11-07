# React + TypeScript + Vite + Web-To-Embed

This template is the standard react template with an extra script to run web-to-embed

## Setup

1. `npm install`
2. `npm run build`

### Windows
1. Build Go binary
   1. `set GOOS=windows`
   2. `set GOARCH=amd64`
   3. `go build -o ./examples/react/bin/` (don't forget trailing slash)
2. Run npm script
   1. `cd ./examples/react/`
   2. `npm run web-to-embed-windows`
   3. Output should be at `./examples/react/bin/output.h`

### Linux/Other
1. Build Go binary
   1. `go build -o ./examples/react/bin/` (don't forget trailing slash)
2. Run npm script
   1. `cd ./examples/react/`
   2. `npm run web-to-embed-linux`
   3. Output should be at `./examples/react/bin/output.h`