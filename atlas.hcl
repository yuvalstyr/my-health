variable "token" {
  type    = string
  default = getenv("TURSO_TOKEN")
}

env "turso" {
  url     = "libsql://my-health-yuvalstyr.turso.io?authToken=${var.token}"
  exclude = ["_litestream*"]
}

env "local" {
  url = "sqlite3://data.db"
  migration {
    dir = "file://migrations"
    form = "atalas.format.sql"
  }
}
