package simpleapp

import (
	"crypto/tls"
	l4g "github.com/alecthomas/log4go"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/primefour/xserver/model"
	"github.com/primefour/xserver/store"
	"github.com/primefour/xserver/utils"
	"github.com/rsc/letsencrypt"
	"github.com/tylerb/graceful"
	"gopkg.in/throttled/throttled.v2"
	"gopkg.in/throttled/throttled.v2/store/memstore"
	"net"
	"net/http"
	"strings"
	"time"
)

var allowedMethods []string = []string{
	"POST",
	"GET",
	"OPTIONS",
	"PUT",
	"PATCH",
	"DELETE",
}

const (
	TIME_TO_WAIT_FOR_CONNECTIONS_TO_CLOSE_ON_SERVER_SHUTDOWN = time.Second
	SIMPLE_APP_NAME                                          = "SimpleServer"
)

type OriginCheckerProc func(*http.Request) bool

type SimpleServer struct {
	Store          store.Store
	Router         *mux.Router
	GracefulServer *graceful.Server
}

var Srv SimpleServer = SimpleServer{}

func GetInstance() *SimpleServer {
	return &Srv
}

type CorsWrapper struct {
	router *mux.Router
}

type VaryBy struct{}

func (m *VaryBy) Key(r *http.Request) string {
	return utils.GetIpAddress(r)
}

func OriginChecker(r *http.Request) bool {
	origin := r.Header.Get("Origin")
	return *SServConfig.ServiceSettings.AllowCorsFrom == "*" || strings.Contains(*SServConfig.ServiceSettings.AllowCorsFrom, origin)
}

func GetOriginChecker(r *http.Request) OriginCheckerProc {
	if len(*SServConfig.ServiceSettings.AllowCorsFrom) > 0 {
		return OriginChecker
	}
	return nil
}

//html5 for browser visit other domain without domain restict
func (cw *CorsWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if len(*SServConfig.ServiceSettings.AllowCorsFrom) > 0 {
		origin := r.Header.Get("Origin")
		if *SServConfig.ServiceSettings.AllowCorsFrom == "*" ||
			strings.Contains(*SServConfig.ServiceSettings.AllowCorsFrom, origin) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" {
				w.Header().Set(
					"Access-Control-Allow-Methods",
					strings.Join(allowedMethods, ", "))
				w.Header().Set(
					"Access-Control-Allow-Headers",
					r.Header.Get("Access-Control-Request-Headers"))
			}
		}
	}

	if r.Method == "OPTIONS" {
		return
	}
	cw.router.ServeHTTP(w, r)
}

func initalizeThrottledVaryBy() *throttled.VaryBy {
	vary := throttled.VaryBy{}

	if SServConfig.RateLimitSettings.VaryByRemoteAddr {
		vary.RemoteAddr = true
	}

	if len(SServConfig.RateLimitSettings.VaryByHeader) > 0 {
		vary.Headers = strings.Fields(SServConfig.RateLimitSettings.VaryByHeader)
		if SServConfig.RateLimitSettings.VaryByRemoteAddr {
			l4g.Warn(utils.T("api.server.start_server.rate.warn"))
			vary.RemoteAddr = false
		}
	}
	return &vary
}

func redirectHTTPToHTTPS(w http.ResponseWriter, r *http.Request) {
	if r.Host == "" {
		http.Error(w, "Not Found", http.StatusNotFound)
	}
	url := r.URL
	url.Host = r.Host
	url.Scheme = "https"
	http.Redirect(w, r, url.String(), http.StatusFound)
}

func (self *SimpleServer) InitStores() bool {
	Srv.Store = store.NewSqlStore(&SServConfig.SqlSettings)
	return true
}

func (self *SimpleServer) InitRouter() bool {
	return true
}

func (self *SimpleServer) InitApi() bool {
	return true
}

func (self *SimpleServer) StartServer() bool {
	l4g.Info(utils.T("api.server.start_server.starting.info"))
	var handler http.Handler = &CorsWrapper{Srv.Router}

	if *SServConfig.RateLimitSettings.Enable {
		l4g.Info(utils.T("api.server.start_server.rate.info"))
		store, err := memstore.New(SServConfig.RateLimitSettings.MemoryStoreSize)
		if err != nil {
			l4g.Critical(utils.T("api.server.start_server.rate_limiting_memory_store"))
			return false
		}

		quota := throttled.RateQuota{
			MaxRate:  throttled.PerSec(SServConfig.RateLimitSettings.PerSec),
			MaxBurst: *SServConfig.RateLimitSettings.MaxBurst,
		}

		rateLimiter, err := throttled.NewGCRARateLimiter(store, quota)
		if err != nil {
			l4g.Critical(utils.T("api.server.start_server.rate_limiting_rate_limiter"))
			return false
		}

		httpRateLimiter := throttled.HTTPRateLimiter{
			RateLimiter: rateLimiter,
			VaryBy:      &VaryBy{},
			DeniedHandler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				l4g.Error("%v: Denied due to throttling settings code=429 ip=%v", r.URL.Path, utils.GetIpAddress(r))
				throttled.DefaultDeniedHandler.ServeHTTP(w, r)
			}),
		}

		handler = httpRateLimiter.RateLimit(handler)
	}

	Srv.GracefulServer = &graceful.Server{
		Timeout: TIME_TO_WAIT_FOR_CONNECTIONS_TO_CLOSE_ON_SERVER_SHUTDOWN,
		Server: &http.Server{
			Addr:         SServConfig.ServiceSettings.ListenAddress,
			Handler:      handlers.RecoveryHandler(handlers.PrintRecoveryStack(true))(handler),
			ReadTimeout:  time.Duration(*SServConfig.ServiceSettings.ReadTimeout) * time.Second,
			WriteTimeout: time.Duration(*SServConfig.ServiceSettings.WriteTimeout) * time.Second,
		},
	}
	l4g.Info(utils.T("api.server.start_server.listening.info"), SServConfig.ServiceSettings.ListenAddress)

	if *SServConfig.ServiceSettings.Forward80To443 {
		go func() {
			listener, err := net.Listen("tcp", ":80")
			if err != nil {
				l4g.Error("Unable to setup forwarding")
				return
			}
			defer listener.Close()

			http.Serve(listener, http.HandlerFunc(redirectHTTPToHTTPS))
		}()
	}

	go func() {
		var err error
		if *SServConfig.ServiceSettings.ConnectionSecurity == model.CONN_SECURITY_TLS {
			if *SServConfig.ServiceSettings.UseLetsEncrypt {
				var m letsencrypt.Manager
				m.CacheFile(*SServConfig.ServiceSettings.LetsEncryptCertificateCacheFile)

				tlsConfig := &tls.Config{
					GetCertificate: m.GetCertificate,
				}

				tlsConfig.NextProtos = append(tlsConfig.NextProtos, "h2")

				err = Srv.GracefulServer.ListenAndServeTLSConfig(tlsConfig)
			} else {
				err = Srv.GracefulServer.ListenAndServeTLS(*SServConfig.ServiceSettings.TLSCertFile, *SServConfig.ServiceSettings.TLSKeyFile)
			}
		} else {
			err = Srv.GracefulServer.ListenAndServe()
		}
		if err != nil {
			l4g.Critical(utils.T("api.server.start_server.starting.critical"), err)
			time.Sleep(time.Second)
		}
	}()
	return true

}

func (self *SimpleServer) StopServer() {
	l4g.Info(utils.T("api.server.stop_server.stopping.info"))
	Srv.GracefulServer.Stop(TIME_TO_WAIT_FOR_CONNECTIONS_TO_CLOSE_ON_SERVER_SHUTDOWN)
	Srv.Store.Close()
	l4g.Info(utils.T("api.server.stop_server.stopped.info"))

}

func (self *SimpleServer) GetAppName() string {
	return SIMPLE_APP_NAME
}
