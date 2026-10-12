# go-scikit — brand assets

Official logos for the **go-scikit** organization, the umbrella over the
fleet's scientific Go work.

<img src="avatar/go-scikit.png" width="160" alt="go-scikit">

The mark is the stack drawn as a building. The base is the arrays, one plate
cut in two: go-ndarray and go-fft. On it stand four columns, one for each
library that computes on those arrays: go-numerics, go-statistics,
go-dataframe and go-plot. They carry the beam, go-notebook, where all of them
meet. Nothing is placed by hand: `mark/main.go` lists the member organisations
layer by layer, and the plates, cells, columns and spacing follow from that
list.

Its gradient (`#8FA3C8` to `#2F3E63`) is a slate blue that sets the umbrella
apart from every member: the emerald of go-ndarray, go-fft and go-notebook,
go-numerics' violet, go-statistics' amber, go-dataframe's magenta and
go-plot's blue. Its dark end reads on white at 10.55:1 and its light end on a
dark page at 7.58:1.

## Formats

- **SVG** — `svg/color`, `svg/white` (white, transparent), `svg/black` (black, transparent)
- **PNG** 16→1024 px, 3 variants — `png/<variant>/<size>/`
- **JPG** (color, white background) — `jpg/<size>/`
- **ICO** (Windows) — `ico/`
- **ICNS** (macOS) — `icns/`
- **GitHub avatar** 512 px — `avatar/`
- **Social preview** 1280×640 (repo banner) — `social/`

`check/kit.py` counts the whole matrix on every pull request, and the `mark`
job writes `svg/color/go-scikit.svg` again from `mark/main.go` and fails if
a byte differs.

```sh
cd mark && go run . > ../svg/color/go-scikit.svg
```

---
*`svg/color/go-scikit.svg` is written by `mark/main.go`; every other file is
generated from it by the fleet's Go brand pipeline (brandgen) — rounded square
badge, line glyph, color / white / black variants.*
