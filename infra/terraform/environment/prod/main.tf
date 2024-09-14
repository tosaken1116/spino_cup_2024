provider "google" {
  project = vars.project_id
  region  = vars.region
}

terraform {
  required_providers {
    google = {
      source = "hashicorp/google"
    }
  }

  backend "gcs" {
    bucket = "spino-tfstate"
    prefix = "terraform/state"
  }
}
