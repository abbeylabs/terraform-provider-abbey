terraform {
  required_providers {
    abbey = {
      source  = "hashicorp.com/edu/abbey"
      version = "0.2.7"
    }
  }
}

provider "abbey" {

  host = "https://api.abbey.io/v1"

  auth_token = "MY_TOKEN"

}
