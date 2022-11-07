package util

import "github.com/yiGmMk/zero-paopao/api/pkg/util/iploc"

func GetIPLoc(ip string) string {
	country, _ := iploc.Find(ip)
	return country
}
