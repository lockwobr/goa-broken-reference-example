package design

import (
	. "github.com/goadesign/goa/design" // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl"
)

// UserMedia user media type
var UserMedia = MediaType("application/vnd.user+json", func() {
	Description("A user")
	Attributes(func() { // Attributes define the media type shape.
		Attribute("UserId", Integer, "Unique user ID")
		Attribute("href", String, "API href for making requests on the user")

		Attribute("Email", String, "Users Email address.", func() {
			Format("email")
			Example("brian.lockwood@harvestmedia.io")
		})

		Attribute("UpdateAt", DateTime, "UpdateAt")
		Attribute("CreatedAt", DateTime, "CreatedAt")
		Attribute("FirstName", String, "First name of the user.", func() {
			Example("Jane")
			MinLength(2)
		})
		Attribute("LastName", String, "Last name of the user.", func() {
			Example("Doe")
			MinLength(2)
		})
	})

	View("short", func() { // View defines a rendering of the media type.
		Attribute("UserId") // Media types may have multiple views and must
		Attribute("href")   // have a "default" view.
		Attribute("Email")
		Attribute("FirstName")
		Attribute("LastName")
	})

	View("default", func() { // View defines a rendering of the media type.
		Attribute("UserId")
		Attribute("href")
		Attribute("Email")
		Attribute("UpdateAt")
		Attribute("CreatedAt")
		Attribute("FirstName")
		Attribute("LastName")
	})

	//todo make a short view that is just want we need
})

//UserPayload is the payload for creating a create/update UserMedia
var UserPayload = Type("UserPayload", func() {
	Description("A User.")

	Reference(UserMedia)
	Attribute("FirstName")
	Attribute("LastName")
	Attribute("Email")

	Required("FirstName")
	Required("LastName")
	Required("Email")

})

var _ = Resource("User", func() {

	DefaultMedia(UserMedia)
	BasePath("/user")

	Action("list", func() {
		Routing(
			GET(""),
		)
		Description("Retrieve Users.")

		//UseTrait("pagination")
		Params(func() {
			Param("view", String, "Which view of user would you like?", func() {
				Default("default")
				Enum("default", "short")
			})
		})

		Response(OK, func() {
			Media(CollectionOf(UserMedia, func() {
				View("default")
				View("short")
			}))
		})
		Response(NotFound)
	})

	Action("show", func() {
		Routing(
			GET("/:UserId"),
		)
		Description("Retrieve account with given id.")
		Params(func() {
			Param("UserId")
			Param("view", String, "Which view of user would you like?", func() {
				Default("default")
				Enum("default", "short")
			})
		})
		Response(OK)
		Response(NotFound)
	})

	Action("create", func() {
		Routing(
			POST(""),
		)
		Description("Create new user")
		Payload(UserPayload)
		Response(Created, func() {
			Media(UserMedia)
		})
		Response(Conflict)
		Response(BadRequest, ErrorMedia)
	})

	Action("update", func() {
		Routing(
			PUT("/:UserId"),
		)
		Description("Create new user")
		Params(func() {
			Param("UserId")
		})
		Payload(UserPayload)
		Response(OK, func() {
			Media(UserMedia)
		})
		Response(Conflict)
		Response(BadRequest, ErrorMedia)
	})

	Action("delete", func() {
		Routing(
			DELETE("/:UserId"),
		)
		Params(func() {
			Param("UserId")
		})
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
})
