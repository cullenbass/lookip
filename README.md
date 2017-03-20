# lookip

Displays an Apache log from a Kafka topic on a map in your web browser.

## Running
Download [GeoLite2 City](http://dev.maxmind.com/geoip/geoip2/geolite2/) and place in same dir as binary.

Neccessary files in the same directory:
  - `lookip` binary
  - `index.html`
  - `us.json`

Recommended files:
  - `config.yaml`

Default configuration is in the config.yaml.

## Compilation

```
go get github.com/oschwald/maxminddb-golang github.com/gorilla/websocket gopkg.in/Shopify/sarama.v1 gopkg.in/yaml.v2
go build
```