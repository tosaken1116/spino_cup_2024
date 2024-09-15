module "main" {
  source = "../../modules"

  common = {
    project_id  = var.project_id
    region      = var.region
    prefix      = var.prefix
    environment = var.environment
  }
  vpc = var.vpc
  dns = var.dns
}
