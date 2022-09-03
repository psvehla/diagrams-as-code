package design

import "goa.design/model/dsl"

var _ = dsl.Design("Getting Started", "This is a model of my software system.", func() {
	var System = dsl.SoftwareSystem("Software System", "My software system.", func() {
		dsl.Tag("system")
	})

	dsl.Person("User", "A user of my software system.", func() {
		dsl.Uses(System, "Uses")
		dsl.Tag("person")
	})

	dsl.Views(func() {
		dsl.SystemContextView(System, "SystemContext", "An example of a System Context diagram.", func() {
			dsl.AddAll()
			dsl.AutoLayout(dsl.RankLeftRight)
		})
		dsl.Styles(func() {
			dsl.ElementStyle("system", func() {
				dsl.Background("#1168bd")
				dsl.Color("#FFFFFF")
			})
			dsl.ElementStyle("person", func() {
				dsl.Shape(dsl.ShapePerson)
				dsl.Background("#08427b")
				dsl.Color("#FFFFFF")
			})
		})
	})
})
