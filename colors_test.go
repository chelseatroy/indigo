package indigo_test

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"

  "github.com/chelseatroy/indigo"
)

var _ = Describe("indigo", func() {
  Describe("Red", func() {
    It("adds the red color code", func() {
      Expect(indigo.Red("ERROR NOT GOOD")).To(Equal("\x1b[91mERROR NOT GOOD\x1b[0m"))
    })
  })

  Describe("Green", func() {
    It("adds the green color code", func() {
      Expect(indigo.Green("TOO GOOD")).To(Equal("\x1b[32mTOO GOOD\x1b[0m"))
    })
  })

  Describe("Cyan", func() {
    It("adds the cyan color code", func() {
      Expect(indigo.Cyan("INFO")).To(Equal("\x1b[36mINFO\x1b[0m"))
    })
  })

  Describe("Yellow", func() {
    It("adds the yellow color code", func() {
      Expect(indigo.Yellow("INFO")).To(Equal("\x1b[33mINFO\x1b[0m"))
    })
  })

  Describe("Gray", func() {
    It("adds the gray color code", func() {
      Expect(indigo.Gray("INFO")).To(Equal("\x1b[90mINFO\x1b[0m"))
    })
  })

  Describe("Light Gray", func() {
    It("adds the light gray color code", func() {
      Expect(indigo.LightGray("INFO")).To(Equal("\x1b[37mINFO\x1b[0m"))
    })
  })

  Describe("Blue", func() {
    It("adds the blue color code", func() {
      Expect(indigo.Blue("INFO")).To(Equal("\x1b[34mINFO\x1b[0m"))
    })
  })

  Describe("Magenta", func() {
    It("adds the magenta color code", func() {
      Expect(indigo.Magenta("INFO")).To(Equal("\x1b[35mINFO\x1b[0m"))
    })
  })

})
