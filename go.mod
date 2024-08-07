module github.com/antonybholmes/go-edb-api

go 1.22.5

replace github.com/antonybholmes/go-auth => ../go-auth

replace github.com/antonybholmes/go-dna => ../go-dna

replace github.com/antonybholmes/go-genes => ../go-genes

replace github.com/antonybholmes/go-microarray => ../go-microarray

replace github.com/antonybholmes/go-mutations => ../go-mutations

replace github.com/antonybholmes/go-basemath => ../go-basemath

replace github.com/antonybholmes/go-sys => ../go-sys

replace github.com/antonybholmes/go-mailer => ../go-mailer

replace github.com/antonybholmes/go-geneconv => ../go-geneconv

replace github.com/antonybholmes/go-motiftogene => ../go-motiftogene

replace github.com/antonybholmes/go-pathway => ../go-pathway

require (
	github.com/antonybholmes/go-basemath v0.0.0-20240802221548-7773050a8f2f // indirect
	github.com/antonybholmes/go-dna v0.0.0-20240726180729-b94c3b7b50fa
	github.com/labstack/echo/v4 v4.12.0
	github.com/labstack/gommon v0.4.2 // indirect
)

require (
	github.com/antonybholmes/go-auth v0.0.0-20240729221758-64a79e9bcccb
	github.com/antonybholmes/go-genes v0.0.0-20240726213523-c3131be49838
	github.com/golang-jwt/jwt/v5 v5.2.1
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/labstack/echo-jwt/v4 v4.2.0
	github.com/mattn/go-sqlite3 v1.14.22
	github.com/rs/zerolog v1.33.0
)

require (
	github.com/antonybholmes/go-mailer v0.0.0-20240731184014-69982846cd1f
	github.com/antonybholmes/go-sys v0.0.0-20240801224521-3bed2c519a83
	github.com/gorilla/sessions v1.3.0
	github.com/labstack/echo-contrib v0.17.1
	github.com/michaeljs1990/sqlitestore v0.0.0-20210507162135-8585425bc864
)

require (
	github.com/antonybholmes/go-geneconv v0.0.0-20240802180141-c71561f5f49a
	github.com/antonybholmes/go-math v0.0.0-20240802221548-7773050a8f2f
	github.com/antonybholmes/go-motiftogene v0.0.0-20240610174236-4ad4f3210a63
	github.com/antonybholmes/go-mutations v0.0.0-20240802180144-af00d02577c8
	github.com/antonybholmes/go-pathway v0.0.0-20240730215511-6e8c2752be79
)

require gopkg.in/yaml.v2 v2.4.0 // indirect

require (
	github.com/antonybholmes/go-microarray v0.0.0-20240504032631-9fb6b43a10d4
	github.com/aws/aws-sdk-go v1.55.5 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/context v1.1.2 // indirect
	github.com/gorilla/securecookie v1.1.2 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	github.com/xyproto/randomstring v1.0.5 // indirect
	golang.org/x/crypto v0.25.0 // indirect
	golang.org/x/net v0.27.0 // indirect
	golang.org/x/sys v0.23.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	golang.org/x/time v0.6.0 // indirect
)
