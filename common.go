package agollo

import (
	"fmt"
	"net"
	"net/url"
	"os"
)

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops:" + err.Error())
		os.Exit(1)
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.To4().String()
			}
		}
	}
	return ""
}

func serverConfigURL(remoteIP, appid string) string {
	return fmt.Sprintf("http://%s/services/config?appId=%s&ip=%s",
		remoteIP,
		url.QueryEscape(appid),
		getLocalIP())
}

func notificationURL(remoteIP, appid, cluster, notifications string) string {
	return fmt.Sprintf("http://%s/notifications/v2?appId=%s&cluster=%s&notifications=%s",
		remoteIP,
		url.QueryEscape(appid),
		url.QueryEscape(cluster),
		url.QueryEscape(notifications))
}

func configURL(remoteIP, appid, cluster, namespace string) string {
	return fmt.Sprintf("http://%s/configs/%s/%s/%s?ip=%s",
		remoteIP,
		url.QueryEscape(appid),
		url.QueryEscape(cluster),
		url.QueryEscape(namespace),
		getLocalIP())
}