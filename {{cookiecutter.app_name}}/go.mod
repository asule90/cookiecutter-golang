module github.com/{{cookiecutter.org_name}}/{{cookiecutter.app_name}}

require (
	github.com/go-ozzo/ozzo-validation/v4 v4.3.0
	{% if cookiecutter.add_rest_server == "chi" -%}github.com/go-chi/chi/v5 v5.0.10{%- endif %}
	{% if cookiecutter.add_rest_server == "echo" -%}github.com/labstack/echo/v4 v4.11.1{%- endif %}
	{% if cookiecutter.use_cobra_cmd == "y" -%}github.com/spf13/cobra v1.3.0{%- endif %}
	{% if cookiecutter.use_viper_config == "y" -%}github.com/spf13/viper v1.3.2{%- endif %}
)
