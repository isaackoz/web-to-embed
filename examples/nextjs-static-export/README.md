# NextJS + Web-To-Embed

This template is the standard `npm create-next-app@latest` with the output set to `output: "export"` in `next-config.ts`

## Setup

1. `npm install`
2. `npm run build`

### Windows
1. Build Go binary
   1. `set GOOS=windows`
   2. `set GOARCH=amd64`
   3. `go build -o ./examples/nextjs-static-export/bin/` (don't forget trailing slash)
2. Run npm script
   1. `cd ./examples/nextjs-static-export/`
   2. `npm run web-to-embed-windows`
   3. Output should be at `./examples/nextjs-static-export/bin/output.h`

### Linux/Other
1. Build Go binary
   1. `go build -o ./examples/nextjs-static-export/bin/` (don't forget trailing slash)
2. Run npm script
   1. `cd ./examples/nextjs-static-export/`
   2. `npm run web-to-embed-linux`
   3. Output should be at `./examples/nextjs-static-export/bin/output.h`