resource "random_bytes" "bytes" {
  length = 4

  keepers = {
    listener_arn = "bloader-slave"
  }
}

resource "tls_private_key" "keygen" {
  algorithm = "RSA"
  rsa_bits  = 4096
}

resource "local_file" "private_keys" {
  content  = tls_private_key.keygen.private_key_pem
  filename = "${var.ssh_keys_path}/slave.id_rsa" #private_key_file
  provisioner "local-exec" {
    command = "chmod 600 ${var.ssh_keys_path}/slave.id_rsa"
  }
}

resource "local_file" "public_keys" {
  content  = tls_private_key.keygen.public_key_openssh
  filename = "${var.ssh_keys_path}/slave.id_rsa.pub" #public_key_file
  provisioner "local-exec" {
    command = "chmod 600 ${var.ssh_keys_path}/slave.id_rsa.pub"
  }
}

resource "aws_key_pair" "ssh_key_pair" {
  key_name        = "bloader-slave-key"
  public_key      = tls_private_key.keygen.public_key_openssh
}