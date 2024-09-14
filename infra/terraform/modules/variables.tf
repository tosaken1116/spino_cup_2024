variable "common" {
  type = object({
    project_id  = string
    region      = string
    prefix      = string
    environment = string
  })
}

variable "vpc" {
  type = object({
    subnet_cidr = string
    pod_cidr    = string
    svc_cidr    = string
  })
}
