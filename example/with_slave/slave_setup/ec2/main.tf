resource "aws_instance" "ec2_instance" {
  count         = var.instance_count
  ami           = data.aws_ssm_parameter.amazonlinux_2.value
  instance_type = var.instance_type
  subnet_id     = aws_subnet.public.id
  key_name      = aws_key_pair.ssh_key_pair.key_name
  vpc_security_group_ids = [
    aws_security_group.grpc_sg.id
  ]

  user_data = <<-EOF
#!/bin/bash

export HOME=/home/ec2-user
cd $HOME
sudo yum update -y
sudo yum install wget -y
export VERSION="${var.bloader_version}"

ARCH=$(uname -m)
if [ "$ARCH" = "x86_64" ]; then
  wget "https://github.com/cresplanex/bloader/releases/download/v$${VERSION}/bloader_$${VERSION}_$(uname -s)_x86_64.rpm"
  sudo yum localinstall bloader_$${VERSION}_$(uname -s)_x86_64.rpm -y
elif [ "$ARCH" = "aarch64" ]; then
  wget "https://github.com/cresplanex/bloader/releases/download/v$${VERSION}/bloader_$${VERSION}_$(uname -s)_arm64.rpm"
  sudo yum localinstall bloader_$${VERSION}_$(uname -s)_arm64.rpm -y
else
  echo "Unsupported architecture: $ARCH"
fi

mkdir $HOME/loader
cd $HOME/loader

if [ "${var.tls_enabled}" = true ]; then
  # Conditionally add TLS private key
  mkdir -p $(dirname "${var.slave_ca_cert_file_path}")
  mkdir -p $(dirname "${var.slave_ca_key_file_path}")
  mkdir -p $(dirname "${var.slave_cert_file_path}")
  mkdir -p $(dirname "${var.slave_key_file_path}")
  cat <<EOF_TLS_CERT > ${var.slave_ca_cert_file_path}
${tls_self_signed_cert.ca_cert[0].cert_pem}
EOF_TLS_CERT
  cat <<EOF_TLS_KEY > ${var.slave_ca_key_file_path}
${tls_private_key.ca_key[0].private_key_pem}
EOF_TLS_KEY
  
  # Generate Slave private key
  openssl genrsa -out ${var.slave_key_file_path} 2048

  # Create a CSR for the slave certificate
  openssl req -new -key ${var.slave_key_file_path} -subj "/CN=localhost" -out /tmp/slave.csr

  # Generate the Slave certificate using the CA key and certificate
  openssl x509 -req \
    -in /tmp/slave.csr \
    -CA ${var.slave_ca_cert_file_path} \
    -CAkey ${var.slave_ca_key_file_path} \
    -CAcreateserial \
    -out ${var.slave_cert_file_path} \
    -days 365 \
    -sha256

  # Cleanup
  rm -f /tmp/slave.csr
fi

# Create the slave config file
cat <<EOF_SLAVE_CONFIG > bloader.yaml
version: "v1"
type: slave
env: "production"
slave_setting:
  port: 50051
  certificate:
    enabled: true
    slave_cert: "certs/slave.crt"
    slave_key: "certs/slave.key"
encrypts: []
logging:
  output:
    - type: "file"
      filename: "logs/app.log"
      format: "text"
      level: "warn"
clock:
  format: "2006-01-02T15:04:05Z"
language:
  default: "en"
override: []
EOF_SLAVE_CONFIG

chown ec2-user:ec2-user -R $HOME/loader

bloader slave run > ./slave.log 2>&1

EOF

  tags = {
    Name = "bloader-slave-instance-${count.index + 1}"
  }
}

resource "aws_eip" "instance_eip" {
  count       = var.instance_count
  instance    = aws_instance.ec2_instance[count.index].id
  domain   = "vpc"

  tags = {
    Name = "bloader-eip-${count.index + 1}"
  }
}