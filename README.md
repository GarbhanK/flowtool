# Flowtool

- simple CLI application for templating your airflow SQL files
- made for quickly copy/pasting airflow templated sql into BigQuery UI
- template mappings set from a `config.json` file
- can convert most built-in airflow template variables
- support for different environments variable with `${env}`

![](.github/readme_images/bqt-demo.gif)

### Setup
- create a folder `${HOME}/Documents/flowtool`
- create a `config.json` file with your template variables mapped to their values

```json
{
    "params.project": "gk-africa-data-eu-${env}",
    "params.web_project": "testscore-web",
    "environment": "${env}"
}
```

### Usage

```bash
# replace the template variables and add it to the clipboard
flowtool template test.sql

# print all current key/value pairs
flowtool config list

# add new key/value pairs to your config.json...
flowtool config add <key> <value>

# and remove existing keys
flowtool config rm <key>
```

