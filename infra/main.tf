# comfigure google project
provider "google" {
  project = "gocrypto"
}

# deploy service to cloud run
resource "google_cloud_run_service" "gocrypto_service" {
  name = "gocrypto_service"
  location = "us-central1"
  template {
    spec {
        containers {
            image = "ewave112/gocrypto-server:v1"
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