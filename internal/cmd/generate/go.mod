module github.com/hedan806/go-ambari-client/internal/cmd/generate

go 1.13

replace github.com/hedan806/go-ambari-client => ../../../

require (
	github.com/alecthomas/chroma v0.6.3
	github.com/hedan806/go-ambari-client v0.0.0-00010101000000-000000000000 // indirect
	github.com/spf13/cobra v1.0.0
	golang.org/x/crypto v0.0.0-20200302210943-78000ba7a073
	golang.org/x/tools v0.0.0-20200413015812-1f08ef6002a8
	gopkg.in/yaml.v2 v2.2.8
)
