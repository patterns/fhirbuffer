# fhirbuffer

fhirbuffer is a gRPC service to export data from [Fhirbase](https://health-samurai.io/fhirbase).
 The one thing we want it to do well is talk gRPC.
 So a downstream process isn't aware of the database, SQL, and ultimately even Go.

For a prototype of the end-to-end flow, see [fhirping app hosted on Glitch](https://accidental-oil.glitch.me).
 Read the [summary blog post](https://patterns.github.io/tfx).
![diagram](diagram/diagram.svg)

## Quickstart

1. Install Fhirbase (see their [Get Started](https://health-samurai.io/fhirbase) steps)
2. Git clone repo
3. Docker build

```
git clone https://github.com/patterns/fhirbuffer
docker build -f fhirbuffer/Dockerfile.template -t fhirbuffer fhirbuffer
docker run -ti --rm -e DATABASE_URL="postgres://postgres:postgres@localhost:5432/fhirbase" fhirbuffer
```


## Credits

Health Samurai's
 [Fhirbase](https://github.com/fhirbase/fhirbase)
 [(LICENSE)](https://github.com/fhirbase/fhirbase/blob/master/MIT-LICENSE)

gRPC
 [Basics for Go](https://grpc.io/docs/tutorials/basic/go.html)

Jack Christensen's
 [PGX](https://github.com/jackc/pgx)
 [(LICENSE)](https://github.com/jackc/pgx/blob/master/LICENSE)

Johan Brandhorst's
 [gRPC Client Authentication](https://jbrandhorst.com/post/grpc-auth/)
 [(LICENSE)](https://github.com/johanbrandhorst/grpc-auth-example/blob/master/LICENSE)

