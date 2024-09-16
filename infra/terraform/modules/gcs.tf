resource "google_storage_bucket" "web" {
  name     = "${var.common.project_id}-web-frontend-${var.common.environment}"
  location = var.common.region

  website {
    main_page_suffix = "index.html"
    not_found_page   = "index.html" # SPAのために404ページをindex.htmlに設定
  }

  force_destroy = true
}

resource "google_storage_bucket_iam_member" "public_access" {
  bucket = google_storage_bucket.web.name
  role   = "roles/storage.objectViewer"
  member = "allUsers"
}
