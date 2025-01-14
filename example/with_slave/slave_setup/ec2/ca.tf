resource "tls_private_key" "ca_key" {
  count = var.tls_enabled ? 1 : 0

  algorithm = "RSA"
  rsa_bits  = 2048
}

resource "tls_self_signed_cert" "ca_cert" {
  count = var.tls_enabled ? 1 : 0

  is_ca_certificate = true
  validity_period_hours = var.tls_period_hours
  private_key_pem = tls_private_key.ca_key[0].private_key_pem

  subject {
    common_name = "BloaderLocalCA"
  }

  allowed_uses = [
    "key_encipherment",
    "digital_signature",
    "cert_signing",
  ]
}

resource "local_file" "ca_cert_file" {
  count = var.tls_enabled ? 1 : 0

  content  = tls_self_signed_cert.ca_cert[0].cert_pem
  filename = var.ca_cert_file_path
}

resource "local_file" "ca_key_file" {
  count = var.tls_enabled ? 1 : 0

  content  = tls_private_key.ca_key[0].private_key_pem
  filename = var.ca_key_file_path
}
