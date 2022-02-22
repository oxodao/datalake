package cmd

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/oxodao/datalake/services"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the server",
	Long:  `Run the datalake software and check everything in the background`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := services.Get().Config
		router := mux.NewRouter()
		api := router.PathPrefix("/api").Subrouter()

		for _, dp := range services.Get().DataProviders {
			if dp.IsEnabled() {
				fmt.Printf("[%v] Starting...\n", dp.GetName())
				go dp.Process()

				dp.RegisterRoutes(api.PathPrefix(fmt.Sprintf("/%v", dp.GetName())).Subrouter())
			}
		}

		router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
			tpl, err1 := route.GetPathTemplate()
			met, err2 := route.GetMethods()
			if err1 == nil && err2 == nil {
				fmt.Printf("%v %v\n", met, tpl)
			} else {
				if err1 != nil {
					fmt.Println("Error registering route: ", err1)
				} else {
					fmt.Printf("Error registering route %v: %v\n", tpl, err2)
				}
			}
			return nil
		})

		srv := &http.Server{
			Handler:      router,
			Addr:         fmt.Sprintf("%v:%v", cfg.Web.ListeningAddr, cfg.Web.Port),
			WriteTimeout: 15 * 60 * time.Second,
			ReadTimeout:  15 * 60 * time.Second,
		}

		log.Fatal(srv.ListenAndServe())
	},
}
