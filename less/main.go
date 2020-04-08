package main

import (
	"github.com/caddyserver/caddy/caddy/caddymain"

	//Plugins
	//Directives/Middleware
	_ "github.com/Xumeiquer/nobots"                         //http.nobots
	_ "github.com/captncraig/cors/caddy"                    //http.cors
	_ "github.com/shuxs/caddy/gopkg"                        //http.gopkg
	_ "github.com/hacdias/caddy-minify"                     //http.minify
	_ "github.com/jung-kurt/caddy-cgi"                      //http.cgi
	_ "github.com/aablinov/caddy-geoip"                     //http.geoip
	_ "github.com/simia-tech/caddy-locale"                  //http.locale
)

func main() {
	// optional: disable telemetry
	caddymain.EnableTelemetry = false
	caddymain.Run()
}
