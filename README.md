# doh
Privacy-oriented local DNS-over-HTTPS proxy written in Go

## Goals
- Deamon with socket on port 53 that receives DNS requests
- Translates to HTTPS requests with curl or own implementation in Go (uses system CA's)
- Requires HTTPS at all times + Verification

- dns resolvers are preconfigured 1.1.1.1 8.8.8.8 -> can add your own. Do a https request to all and wait for fastest response
- Caching? (Done by the os?)

- Implement check if nameserver is not set (check /etc/resolv.conf on linux and bsd) or just a dns lookup??? -> but this could disclose usage
- (BACKLOG) Implement internal redirects with iptables or pf -> look at sshuttle for inpiration

- Written in Go

## Information about DNS-over-HTTPS
- https://hacks.mozilla.org/2018/05/a-cartoon-intro-to-dns-over-https/
