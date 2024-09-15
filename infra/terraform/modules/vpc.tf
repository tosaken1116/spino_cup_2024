# VPC module

locals {
  pods_cidr_name = "${var.common.prefix}-pods-${var.common.environment}"
  svc_cidr_name  = "${var.common.prefix}-svc-${var.common.environment}"
}
module "vpc" {
  source       = "terraform-google-modules/network/google"
  version      = "8.1.0"
  project_id   = var.common.project_id
  network_name = "${var.common.prefix}-vpc-${var.common.environment}"
  subnets = [
    {
      subnet_name   = "${var.common.prefix}-subnet-${var.common.environment}"
      subnet_ip     = var.vpc.subnet_cidr
      subnet_region = var.common.region
    },
  ]
  secondary_ranges = {
    "${var.common.prefix}-subnet-${var.common.environment}" = [
      {
        range_name    = local.pods_cidr_name
        ip_cidr_range = var.vpc.pod_cidr
      },
      {
        range_name    = local.svc_cidr_name
        ip_cidr_range = var.vpc.svc_cidr
      },
    ]
  }
}

# NAT for private GKE cluster
module "cloud-nat" {
  source  = "terraform-google-modules/cloud-nat/google"
  version = "~> 1.2"

  project_id    = var.common.project_id
  region        = var.common.region
  router        = "safer-router"
  network       = module.vpc.network_self_link
  create_router = true
}
