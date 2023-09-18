module github.com/{{cookiecutter.org_name}}/{{cookiecutter.app_name}}

require (
	{% if cookiecutter.add_rest_server == "chi" -%}github.com/go-chi/chi/v5 v5.0.10{%- endif %}
	{% if cookiecutter.add_rest_server == "echo" -%}github.com/labstack/echo/v4 v4.11.1{%- endif %}
	{% if cookiecutter.use_cobra_cmd == "y" -%}github.com/spf13/cobra v1.3.0{%- endif %}
	{% if cookiecutter.use_air == "y" -%}github.com/cosmtrek/air v1.3.2{%- endif %}
)


require (
	github.com/go-ozzo/ozzo-validation/v4 v4.3.0
	github.com/labstack/echo/v4 v4.11.1
	go.opentelemetry.io/otel v1.17.0
	go.uber.org/zap v1.25.0
)

require (
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/PuerkitoBio/purell v1.2.0 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/go-openapi/jsonpointer v0.20.0 // indirect
	github.com/go-openapi/jsonreference v0.20.2 // indirect
	github.com/go-openapi/spec v0.20.9 // indirect
	github.com/go-openapi/swag v0.22.4 // indirect
	github.com/jinzhu/copier v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/swaggo/echo-swagger v1.4.1 // indirect
	github.com/swaggo/files/v2 v2.0.0 // indirect
	github.com/swaggo/swag v1.16.2 // indirect
	golang.org/x/tools v0.13.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

require (
	github.com/brpaz/echozap v1.1.3 // indirect
	github.com/dave/jennifer v1.4.0 // indirect
	github.com/duythinht/dbml-go v0.0.0-20220713090938-b5659a82f71b // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/jackc/pgerrcode v0.0.0-20220416144525-469b46aa5efa
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.4.3 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/labstack/gommon v0.4.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/shopspring/decimal v1.3.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	go.opentelemetry.io/otel/metric v1.17.0 // indirect
	go.opentelemetry.io/otel/trace v1.17.0 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/crypto v0.13.0 // indirect
	golang.org/x/net v0.15.0 // indirect
	golang.org/x/sys v0.12.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	golang.org/x/time v0.3.0 // indirect
	gorm.io/driver/postgres v1.5.2 // indirect
	gorm.io/gorm v1.25.4 // indirect
	moul.io/zapgorm2 v1.3.0 // indirect
)
