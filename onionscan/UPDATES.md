Update every occurance of v2 addresses to support v3 addresses.

- /onionscan/deanonymization/env_test.go - onion crawler set to 56 characters 

- /onionscan/deanonymization/get_onion_links.go - line 56 and line 43 simplifying syntax `!linkmap[host]`

- /onionscan/deanonymization/private_keys.go - update line 29 - remove unnecessary `fmt.Sprintf` as `osc.LogInfo` already expects a string

- /onionscan/protocol/bitcoin_scanner.go - Update line 171 - Regex update to 56 characters. and 181 - change `theiraddr byte from 16 to 56` and line 206 from 16 to 56 259 from 16 to 56

- /onionscan/deanonymization/private_keys.go - Line 53 changed `hostname` value `strings.ToLower(data[0:56])`

- /onionscan/deanonymization/email_scan.go - line 14, changed `regexp.MustCompile` to `{2,56}` 

- /onionscan/deanonymization/apache_mod_status - line 45, changed `len(vhost) >= 22` to `len(vhost) >= 62`