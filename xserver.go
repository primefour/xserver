package main

import (
	"fmt"
	l4g "github.com/alecthomas/log4go"
	//	"github.com/primefour/xserver/apps/simpleapp"
	"github.com/primefour/xserver/model"
	"github.com/primefour/xserver/store"
	"github.com/primefour/xserver/store/sqlstore"
	"github.com/primefour/xserver/store/storetest"
	"github.com/primefour/xserver/utils"
	"os"
	"sync"
)

const (
	L4G_DEBUG_LEVEL = l4g.DEBUG
	MODE_DEV        = "dev"
	MODE_BETA       = "beta"
	MODE_PROD       = "prod"
)

func initServer() {
	defer func() {
		if r := recover(); r != nil {
			err := fmt.Sprintf("%v", r)
			l4g.Error(err)
		}
	}()
	utils.InitLogSystem()
	//init locale
	utils.InitTranslations()
	//init html templates
	//utils.InitHTML()
}

/*

type WebAppIntf interface {
	InitStores() bool
	InitRouter() bool
	InitApi() bool
	StartServer() bool
	StopServer()
	GetAppName() string
}

var xserver_apps = map[string]WebAppIntf{
	"SimpleServer": simpleapp.GetInstance(),
}


func runApps() {
	for appName, appIntf := range xserver_apps {
		name := appIntf.GetAppName()
		if appName != name {
			l4g.Error("Register Name is not consistent with actual name")
			continue
		}

		if !appIntf.InitStores() {
			l4g.Error(fmt.Sprintf("%s init stores fail ", appName))
			continue
		}

		if !appIntf.InitRouter() {
			l4g.Error(fmt.Sprintf("%s init route fail ", appName))
			continue
		}
		if !appIntf.InitApi() {
			l4g.Error(fmt.Sprintf("%s init api fail ", appName))
			continue
		}

		//start a service
		go func() {
			if !appIntf.StartServer() {
				l4g.Error(fmt.Sprintf("%s start server fail ", appName))
				return
			}
		}()
	}
}

func stopApps() {
	for appName, appIntf := range xserver_apps {
		l4g.Info("stop service of %s ", appName)
		appIntf.StopServer()
	}
}

//init locale and log system before start server
func runServer() {
	runApps()
	// wait for kill signal before attempting to gracefully shutdown
	// the running service
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-c
	stopApps()
}

*/

var storeTypes = []*struct {
	Name      string
	Func      func() (*storetest.RunningContainer, *model.SqlSettings, error)
	Container *storetest.RunningContainer
	Store     store.Store
}{
	{
		Name: "MySQL",
		Func: storetest.NewMySQLContainer,
	},
	/*
		{
			Name: "PostgreSQL",
			Func: storetest.NewPostgreSQLContainer,
		},
	*/
}

func initStores() {
	defer func() {
		if err := recover(); err != nil {
			tearDownStores()
			panic(err)
		}
	}()
	var wg sync.WaitGroup
	errCh := make(chan error, len(storeTypes))
	wg.Add(len(storeTypes))
	for _, st := range storeTypes {
		st := st
		go func() {
			defer wg.Done()
			container, settings, err := st.Func()
			if err != nil {
				errCh <- err
				return
			}
			st.Container = container
			st.Store = store.NewLayeredStore(sqlstore.NewSqlSupplier(*settings, nil), nil, nil)
		}()
	}
	wg.Wait()
	select {
	case err := <-errCh:
		panic(err)
	default:
	}
}

var tearDownStoresOnce sync.Once

func tearDownStores() {
	tearDownStoresOnce.Do(func() {
		var wg sync.WaitGroup
		wg.Add(len(storeTypes))
		for _, st := range storeTypes {
			st := st
			go func() {
				if st.Store != nil {
					st.Store.Close()
				}
				if st.Container != nil {
					st.Container.Stop()
				}
				wg.Done()
			}()
		}
		wg.Wait()
	})
}

func launchStore() {
	status := 0
	initStores()
	defer func() {
		tearDownStores()
		os.Exit(status)
	}()
}

func main() {
	initServer()
	client := model.NewAPIv4Client("https://www.pfbbc.com")
	client.SetOAuthToken("hello world")
	launchStore()
	//runServer()
}
