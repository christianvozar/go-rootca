# gorootca

```Go
import "github.com/christianvozar/gorootca"
```

gorootca provides a convenient HTTPS client with a root certificate bundle for HTTP clients using TLS/SSL. If a Docker image is built using the `scratch` base image without providing a root certificate bundle the Go application will commonly exit fatally with `x509: failed to load system roots and no roots provided" by bundling root certificates.`. Utilizing this package will prevent this error and allow maintainers to create Docker images of smaller size.

gorootca takes its root certificates from the stable [CoreOS](https://coreos.com) distribution channel. Thanks [CoreOS](https://coreos.com)! Sadly, CoreOS has been discontinued and an updated root certificiate will be identified.
