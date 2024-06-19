data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "./loader/main.go",
  ]
}

env "gorm" {
  src = data.external_schema.gorm.url
  dev = "postgres://admin:qwerty@127.0.0.1:5432/TOD0_APP_GO?sslmode=disable"

  migration {
    dir = "file://migrations"
  }

  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}