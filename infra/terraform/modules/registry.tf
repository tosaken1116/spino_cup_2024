resource "google_artifact_registry_repository" "docker" {
  location      = var.common.region
  repository_id = "${var.common.prefix}-docker-${var.common.environment}"
  format        = "DOCKER"
}
