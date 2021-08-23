package main

import "github.com/prometheus/client_golang/prometheus"

var onlineUsers = prometheus.NewGauge(prometheus.GaugeOpts{
	Name:"goapp_online_users",
	Help: "Online users",
	ConstLabels: map[string]string{
		"logged_users": "true"
	}
})

func main(){
	r := prometheus.NewRegistry()
	r.MustRegister(onlineUsers)

	http.Handle("/metrics", prmhttp.HandleFor(r, promhttp.HandlerOpts{}))
	log.Fatal(http.ListenAndServe(":8181", nil))
}