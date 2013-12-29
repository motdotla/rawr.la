package main

import (
  "github.com/codegangsta/martini"
  "github.com/codegangsta/martini-contrib/render"
)

func main() {
  m := martini.Classic()
  m.Use(render.Renderer())

  m.Get("/", func() string {
    return "rawr.la"
  })

  m.Get("/codeday", func(r render.Render) {
    r.HTML(200, "codeday", nil)
  })

  m.Get("/newyears", func(r render.Render) {
    r.HTML(200, "newyears", nil)
  })

  m.Run()
}
