variable "project_id" {
  type = string
}

variable "region" {
  type = string
}

variable "prefix" {
  type = string
}

variable "environment" {
  type = string
}

variable "vpc" {
  type = object({
    subnet_cidr = string
    pod_cidr    = string
    svc_cidr    = string
  })
}

variable "dns" {
  type = object({
    domain = string
  })
}
