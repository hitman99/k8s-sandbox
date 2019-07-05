package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("demo", func() {
	Title("demo-app")
	Server("k8s-sand", func() {
		Host("localhost", func() {
			URI("http://0.0.0.0")
			Variable("version", String, "API version", func() {
				Default("v1")
			})
		})
	})
})

var _ = Service("pgload", func() {
	Description("The pgload service performs operations on postgres database.")

	Method("load", func() {
		Payload(func() {
			Field(1, "count", Int, "How many records to generate in the table")
			Required("count")

		})

		Result(JsonStatus)
		HTTP(func() {
			POST("/pgload")
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})

var JsonStatus = ResultType("application/json", func() {
	Description("Generic JSON response")
	TypeName("JsonStatus")
	Attributes(func() {
		Attribute("code", UInt, "result code")
		Attribute("status", String, "status info")
		Attribute("time", String, "processing time")
		Required("code", "status")
	})
})
