resource "google_dns_managed_zone" "zone" {
  name     = "${var.common.prefix}-zone-${var.common.environment}"
  dns_name = "${var.dns.domain}."
}

resource "google_compute_global_address" "atlantis" {
  name    = "${var.common.prefix}-atlantis-static-ip-${var.common.environment}"
  project = var.common.project_id
}

resource "google_dns_record_set" "atlantis" {
  name         = "atlantis.${var.dns.domain}."
  type         = "A"
  ttl          = "300"
  managed_zone = google_dns_managed_zone.zone.name
  rrdatas      = [google_compute_global_address.atlantis.address]
}
