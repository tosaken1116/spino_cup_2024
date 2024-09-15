# This file upload to GitHub repository for atlantis.
# So, you should not write sensitive information in this file.

project_id  = "spino-cup-2024"
region      = "asia-northeast1"
prefix      = "spino"
environment = "prod"

# VPC
vpc = {
  subnet_cidr = "10.0.1.0/24"
  pod_cidr    = "10.0.2.0/24"
  svc_cidr    = "10.0.3.0/24"
}

# DNS
dns = {
  domain = "spino.kurichi.dev"
}
