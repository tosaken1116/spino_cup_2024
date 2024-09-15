# gke cluster
module "gke" {
  source            = "terraform-google-modules/kubernetes-engine/google//modules/private-cluster"
  project_id        = var.common.project_id
  name              = "${var.common.prefix}-gke-${var.common.environment}"
  regional          = true
  region            = var.common.region
  network           = module.vpc.network_name
  subnetwork        = module.vpc.subnets_names[0]
  ip_range_pods     = local.pods_cidr_name
  ip_range_services = local.svc_cidr_name
  node_pools = [
    {
      name               = "primary-node-pool"
      machine_type       = "n2-standard-2"
      node_locations     = "${var.common.region}-a" # 複数のゾーンを指定できる
      min_count          = 1
      max_count          = 3
      disk_size_gb       = 30
      auto_repair        = true
      auto_upgrade       = true
      enable_autoscaling = true
      preemptible        = false
    },
  ]
}

# gke auth
module "gke_auth" {
  source       = "terraform-google-modules/kubernetes-engine/google//modules/auth"
  depends_on   = [module.gke]
  project_id   = var.common.project_id
  location     = module.gke.location
  cluster_name = module.gke.name
}

# kubeconfig
resource "local_file" "kubeconfig" {
  content  = module.gke_auth.kubeconfig_raw
  filename = "${var.common.prefix}-kubeconfig-${var.common.environment}"
}
