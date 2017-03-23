package main

import (
	_ "github.com/lockwobr/goa-broken-reference-example/design"

	"os"

	"fmt"

	"github.com/goadesign/goa/design"
	"github.com/goadesign/goa/goagen/codegen"
	genapp "github.com/goadesign/goa/goagen/gen_app"
	genclient "github.com/goadesign/goa/goagen/gen_client"
	gencontroller "github.com/goadesign/goa/goagen/gen_controller"
	genjs "github.com/goadesign/goa/goagen/gen_js"
	genschema "github.com/goadesign/goa/goagen/gen_schema"
	genswagger "github.com/goadesign/goa/goagen/gen_swagger"
)

// have to do it this way instead of cli for vendoring

func main() {
	codegen.ParseDSL()
	fmt.Printf("inputs %v \n", os.Args[1:])
	if len(os.Args) == 2 && os.Args[1] == "swagger" { // so just swagger so we can view the design as we build it
		codegen.Run(
			genswagger.NewGenerator(
				genswagger.API(design.Design),
				genswagger.OutDir("public"),
			),
		)
	} else {

		codegen.Run(
			// makes controller in wrong place
			// genmain.NewGenerator(
			// 	genmain.API(design.Design),
			// 	genmain.Target("app"),
			// ),
			gencontroller.NewGenerator(
				gencontroller.API(design.Design),
				gencontroller.OutDir("controller"),
				gencontroller.Pkg("controller"),
			),
			genapp.NewGenerator(
				genapp.API(design.Design),
				genapp.OutDir("app"),
				genapp.Target("app"),
				genapp.NoTest(false),
			),
			genclient.NewGenerator(
				genclient.API(design.Design),
			),
			genswagger.NewGenerator(
				genswagger.API(design.Design),
				genswagger.OutDir("public"),
			),
			genschema.NewGenerator(
				genschema.API(design.Design),
				genschema.OutDir("public"),
			),
			genjs.NewGenerator(
				genjs.API(design.Design),
				genjs.OutDir("public"),
			),
		)
	}
}
