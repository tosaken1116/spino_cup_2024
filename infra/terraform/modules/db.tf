resource "google_sql_database_instance" "main" {
  name             = "${var.common.prefix}-db-instance-${var.common.environment}"
  region           = var.common.region
  database_version = "MYSQL_8_0"

  settings {
    tier              = "db-custom-2-7680"
    activation_policy = "ALWAYS"
    ip_configuration {
      ipv4_enabled = true
    }
    disk_size = 10
  }
}

resource "google_sql_database" "main" {
  name     = "${var.common.prefix}-db-${var.common.environment}"
  instance = google_sql_database_instance.main.name
}

resource "random_password" "db_password" {
  length  = 16
  special = true
}

resource "google_secret_manager_secret" "db_password" {
  secret_id = "db-password"
  replication {
    auto {}
  }
}

resource "google_secret_manager_secret_version" "db_password_version" {
  secret      = google_secret_manager_secret.db_password.id
  secret_data = random_password.db_password.result
}

resource "google_sql_user" "db_user" {
  name     = "spino"
  instance = google_sql_database_instance.main.name
  password = random_password.db_password.result
}
