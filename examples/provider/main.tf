terraform {
  required_providers {
    abbey = {
      source  = "hashicorp.com/edu/abbey"
      version = "0.2.8"
    }
  }
}

provider "abbey" {

  server_url = "https://api.abbey.io/v1"

  bearer_auth = "MY_TOKEN"

}
