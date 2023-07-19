<!-- LTeX: language=en-US -->
# game-gorl (game framework using Go and raylib)

## Structure
- `assets`: Holds game assets.
  - `settings.json`: Settings file.
  - `textures`: Texture files.
- `build`: Destination for build output.
- `main.go`: Application entry point.
- `makefile`: Manages build tasks.
- `pkg`: Go packages.
  - `settings`: Settings handling.
  - `entities`: Game entities.
  - `render`: Rendering logic.
  - `scenes`: Game stages.
- `scripts`: Utility scripts.

## Scripts

### sync_settings.py
Generates Go fallback settings and JSON from GameSettings struct.

Usage:
```bash
python scripts/sync_settings.py [path/to/settings.go path/to/settings.json]
```
If no path is given, the following defaults apply: `./pkg/settings/settings.go`,
`./assets/settings.json`.

> Note: Review changes and test after running.
