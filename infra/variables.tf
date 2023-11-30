variable "project_name" {
  type = string
  description = "the name of the project on google cloud console"
  default = "gocrypto"
  
}

variable "region" {
  type = string
  default = "us-central1"
  description = "region to deploy the cloud run service"
}

variable "container_image" {
    type = string
    description = "the docker image to build on the cloud run service"
    default = "ewave112/gocrypto-server:v1"
  
}

variable "service_name" {
    type = string
    description = "the name of the google cloud run service name"
    default = "gocrypto_service"
  
}

