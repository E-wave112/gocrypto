# comfigure google project
provider "google" {
  project = var.project_name
}

# deploy service to cloud run
resource "google_cloud_run_service" "gocrypto_service" {
  name = var.service_name
  location = var.region
  template {
    spec {
        containers {
            image = var.container_image
        }
    }
    
  }
  traffic {
    percent = 100
    latest_revision = true
  }
}

# create public access
data "google_iam_policy" "access_policy" {
    binding {
      role = "roles/run.invoker"
      members = [ 
        "allUsers"
       ]
    }
  
}

# enable public access
resource "google_cloud_run_service_iam_policy" "noauth" {
    location = google_cloud_run_service.gocrypto_service.location
    project = google_cloud_run_service.gocrypto_service.project
    service = google_cloud_run_service.gocrypto_service.name
    policy_data = data.google_iam_policy.access_policy.policy_data
}