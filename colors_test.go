package indigo_test

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"

  colors "github.com/chelseatroy/indigo"
)

var _ = Describe("colors", func() {
  Describe("Red", func() {
    It("adds the red color code", func() {
      Expect(colors.Red("ERROR NOT GOOD")).To(Equal("\x1b[91mERROR NOT GOOD\x1b[0m"))
    })
  })

  Describe("Green", func() {
    It("adds the green color code", func() {
      Expect(colors.Green("TOO GOOD")).To(Equal("\x1b[32mTOO GOOD\x1b[0m"))
    })
  })

  Describe("Cyan", func() {
    It("adds the cyan color code", func() {
      Expect(colors.Cyan("INFO")).To(Equal("\x1b[36mINFO\x1b[0m"))
    })
  })

  Describe("Yellow", func() {
    It("adds the yellow color code", func() {
      Expect(colors.Yellow("INFO")).To(Equal("\x1b[33mINFO\x1b[0m"))
    })
  })

  Describe("Gray", func() {
    It("adds the gray color code", func() {
      Expect(colors.Gray("INFO")).To(Equal("\x1b[90mINFO\x1b[0m"))
    })
  })

  Describe("Light Gray", func() {
    It("adds the light gray color code", func() {
      Expect(colors.LightGray("INFO")).To(Equal("\x1b[37mINFO\x1b[0m"))
    })
  })

  Describe("Blue", func() {
    It("adds the blue color code", func() {
      Expect(colors.Blue("INFO")).To(Equal("\x1b[34mINFO\x1b[0m"))
    })
  })

  Describe("Magenta", func() {
    It("adds the magenta color code", func() {
      Expect(colors.Magenta("INFO")).To(Equal("\x1b[35mINFO\x1b[0m"))
    })
  })

})
