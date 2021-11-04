resource "local_file" "example" {
  content  = "example!!"
  filename = "${path.module}/example.txt"
}
