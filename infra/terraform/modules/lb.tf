resource "google_compute_backend_bucket" "backend_bucket" {
  name        = "${var.common.prefix}-backend-bucket-${var.common.environment}"
  bucket_name = google_storage_bucket.web.name
  enable_cdn  = true
}

resource "google_compute_url_map" "url_map" {
  name            = "${var.common.prefix}-url-map-${var.common.environment}"
  default_service = google_compute_backend_bucket.backend_bucket.self_link

  # SPAのためにすべてのパスをindex.htmlにリダイレクト
  default_url_redirect {
    redirect_response_code = "MOVED_PERMANENTLY_DEFAULT"
    https_redirect         = true
    strip_query            = false
  }
}

resource "google_compute_target_https_proxy" "https_proxy" {
  name             = "${var.common.prefix}-https-proxy-${var.common.environment}"
  url_map          = google_compute_url_map.url_map.self_link
  ssl_certificates = [google_compute_managed_ssl_certificate.web.self_link]
}

resource "google_compute_global_forwarding_rule" "forwarding_rule" {
  name       = "${var.common.prefix}-forwarding-rule-${var.common.environment}"
  ip_address = google_compute_global_address.web.address
  port_range = "443"
  target     = google_compute_target_https_proxy.https_proxy.self_link
}
