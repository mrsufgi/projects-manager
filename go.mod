module github.com/mrsufgi/projects-manager

go 1.15

require (
	github.com/golang/mock v1.4.4
	github.com/google/go-github v17.0.0+incompatible
	github.com/google/go-querystring v1.0.0 // indirect
	github.com/googollee/go-socket.io v1.4.4
	github.com/jmoiron/sqlx v1.2.0
	github.com/julienschmidt/httprouter v1.3.0
	github.com/lib/pq v1.8.0
	github.com/prometheus/client_golang v1.7.0
	github.com/pusher/pusher-http-go v4.0.1+incompatible
	github.com/rs/cors v1.7.0
	github.com/sirupsen/logrus v1.7.0
	golang.org/x/sys v0.0.0-20200930185726-fdedc70b468f // indirect
	google.golang.org/appengine v1.6.7 // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
)

replace github.com/mrsufgi/projects-manager => ./
