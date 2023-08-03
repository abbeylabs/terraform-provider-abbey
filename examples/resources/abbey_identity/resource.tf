resource "abbey_identity" "my_identity" {
  abbey_account = "alice@example.com"
  source        = "mysource"
  metadata = jsonencode(
    {
      "mykey" = "myvalue"
    }
  )
}
