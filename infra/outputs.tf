output "gocrypto_url" {
    value =google_cloud_run_service.gocrypto_service.status[0].url 
}