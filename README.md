# Cloudstack Service Discovery

Cloudstack service discovery generates a target ip addrs that used `file_sd_config` on prometheus.yml.

# Flags

```
Usage:
  ./cloudstack_servicediscovery [OPTIONS]

OPTIONS:
  -api-key string
        API key of cloudstack
  -endpoint string
        Endpoint url of cloudstack
  -filename string
        Output json file name that specified "file_sd_config"
  -groups string
        List of groups separated by comma
  -help
        Print this help message and exit
  -labels string
        List of labels (e.g. "job:mysql,zone:eu-east")
  -port int
        Suffix port number (default 9090)
  -secret-key string
        Secret key of cloudstack
```

# Install and running

Using Docker:

```
docker build -t cloudstack_servicediscovery ./
```

# Example

Running:

```
$ cloudstack_servicediscovery `
		-api-key apikeyforcloudstack \
		-secret-key secretkeyforcloudstack
		-endpoint http://api.example.com \
		-fliename ./monitoring_target.json
		-labels env:production,job:os
		-groups web,db,kvs
		-port 9090
```

Then json file named `monitoring_target.json` generated:

```
{
  "targets": [
    "10.33.1.50:9090",
    "10.33.1.51:9090",
    "10.33.1.150:9090",
    "10.33.1.250:9090"
  ],
  "labels": {
    "env": "production",
    "job": "os"
  }
}
```
