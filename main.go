package main

import (
	"github.com/caddyserver/caddy/caddy/caddymain"

	//Plugins

	//DNS Providers
	_ "github.com/caddyserver/dnsproviders/acmedns"
	_ "github.com/caddyserver/dnsproviders/azure"
	_ "github.com/caddyserver/dnsproviders/cloudflare"
	

	//Directives/Middleware
	_ "github.com/BTBurke/caddy-jwt"                        //http.jwt
	_ "github.com/SchumacherFM/mailout"                     //http.mailout
	_ "github.com/Xumeiquer/nobots"                         //http.nobots
	_ "github.com/aablinov/caddy-geoip"                     //http.geoip
	_ "github.com/abiosoft/caddy-git"                       //http.git
	_ "github.com/caddyserver/forwardproxy"                 //http.forwardproxy
	_ "github.com/captncraig/caddy-realip"                  //http.realip
	_ "github.com/captncraig/cors/caddy"                    //http.cors
	_ "github.com/casbin/caddy-authz"                       //http.authz
	_ "github.com/coopernurse/caddy-awslambda"              //http.awslambda
	_ "github.com/dhaavi/caddy-permission"                  //http.permission
	_ "github.com/echocat/caddy-filter"                     //http.filter
	_ "github.com/epicagency/caddy-expires"                 //http.expires
	_ "github.com/freman/caddy-reauth"                      //http.reauth
	_ "github.com/hacdias/caddy-minify"                     //http.minify
	_ "github.com/hacdias/caddy-webdav"                     //http.webdav
	_ "github.com/jung-kurt/caddy-cgi"                      //http.cgi
	_ "github.com/jung-kurt/caddy-pubsub"                   //http.pubsub
	_ "github.com/linkonoid/caddy-dyndns"                   //http.dyndns
	_ "github.com/mastercactapus/caddy-proxyprotocol"       //http.proxyprotocol
	_ "github.com/miekg/caddy-prometheus"                   //http.prometheus
	_ "github.com/nicolasazrak/caddy-cache"                 //http.cache
	_ "github.com/payintech/caddy-datadog"                  //http.datadog
	_ "github.com/pieterlouw/caddy-grpc"                    //http.grpc
	_ "github.com/pyed/ipfilter"                            //http.ipfilter
	_ "github.com/shuxs/caddy/gopkg"                        //http.gopkg
	_ "github.com/simia-tech/caddy-locale"                  //http.locale
	_ "github.com/tarent/loginsrv/caddy"                    //http.login
	_ "github.com/techknowlogick/caddy-s3browser"           //http.s3browser
	_ "github.com/xuqingfeng/caddy-rate-limit"              //http.ratelimit
	_ "go.okkur.org/gomods"                                 //http.gomods
	_ "go.okkur.org/torproxy"                               //http.torproxy

	// _ "github.com/leelynne/caddy-awses"    //http.awses
	// _ "github.com/restic/caddy" //http.restic

	//Caddyfile Loaders
	_ "github.com/lucaslorentz/caddy-docker-proxy/plugin" //docker

	//Server Types
//	_ "github.com/lucaslorentz/caddy-supervisor/servertype" //supervisor

	//TLS Clustering
//	_ "github.com/pteich/caddy-tlsconsul" //consul

	// Event Hooks
	_ "github.com/hacdias/caddy-service" //hook.service issus: panic: close of closed channel
	
	// filebrowser 
	_ "github.com/filebrowser/caddy" //hook.service issus: panic: close of closed channel
	
	_ "github.com/mahrud/caddy-altonions"
	_ "github.com/hyperion-hyn/caddy-cron"
	
)

func main() {
	// optional: disable telemetry
	caddymain.EnableTelemetry = false
	caddymain.Run()
}
