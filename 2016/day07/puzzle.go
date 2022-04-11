package day07

import (
	"regexp"
	"strings"
)

type Ip struct{ Hypernet, Supernet []string }

func (ip Ip) Tls() bool {
	for _, hypernet := range ip.Hypernet {
		if abba(hypernet) {
			return false
		}
	}

	for _, supernet := range ip.Supernet {
		if abba(supernet) {
			return true
		}
	}

	return false
}

func (ip Ip) Ssl() bool {
	bab := []string{}
	for _, supernet := range ip.Supernet {
		for _, a := range aba(supernet) {
			bab = append(bab, a[1:]+a[1:2])
		}
	}

	for _, hypernet := range ip.Hypernet {
		for _, b := range bab {
			if strings.Contains(hypernet, b) {
				return true
			}
		}
	}

	return false
}

func abba(ip string) bool {
	for i, n := 0, len(ip)-3; i < n; i++ {
		if ip[i] != ip[i+1] && ip[i] == ip[i+3] && ip[i+1] == ip[i+2] {
			return true
		}
	}
	return false
}

func aba(ip string) (found []string) {
	for i, n := 0, len(ip)-2; i < n; i++ {
		if ip[i] != ip[i+1] && ip[i] == ip[i+2] {
			found = append(found, ip[i:i+3])
		}
	}
	return
}

func Part1(input string) (valid int) {
	for _, ip := range strings.Fields(input) {
		if parseIp(ip).Tls() {
			valid++
		}
	}
	return
}

func Part2(input string) (valid int) {
	for _, ip := range strings.Fields(input) {
		if parseIp(ip).Ssl() {
			valid++
		}
	}
	return
}

func parseIp(input string) (ip Ip) {
	var pattern = regexp.MustCompile(`\[(.+?)\]`)
	ip.Supernet = pattern.Split(input, -1)
	for _, hypernet := range pattern.FindAllStringSubmatch(input, -1) {
		ip.Hypernet = append(ip.Hypernet, hypernet[1])
	}
	return
}
