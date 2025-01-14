variable "region" {
    description = "The AWS region to launch in"
    type = string
}

variable "profile_name" {
    description = "The AWS profile to use"
    type = string
}

variable "bloader_version" {
    description = "The version of the bloader to install"
    type = string
}

variable "tls_enabled" {
    description = "Whether to enable TLS"
    type = bool
}

variable "tls_period_hours" {
    description = "The validity period of the TLS certificate in hours"
    type = number
}

variable "ca_cert_file_path" {
    description = "The path to save the CA certificate"
    type = string
}

variable "ca_key_file_path" {
    description = "The path to save the CA key"
    type = string
}

variable "slave_ca_cert_file_path" {
    description = "The path to save the Slave CA certificate"
    type = string
}

variable "slave_ca_key_file_path" {
    description = "The path to save the Slave CA key"
    type = string
}

variable "slave_cert_file_path" {
    description = "The path to save the Slave certificate"
    type = string
}

variable "slave_key_file_path" {
    description = "The path to save the Slave private key"
    type = string
}

variable "availability_zone" {
    description = "The availability zone to launch in"
    type        = string
}

variable "instance_type" {
    description = "The type of instance to start"
    type        = string
}

variable "instance_architecture" {
    description = "The architecture of the instance. Either x86_64 or arm64"
    type        = string
}

variable "instance_count" {
    description = "The number of instances to start"
    type        = number
}

variable "ssh_keys_path" {
    description = "The path to save the SSH keys"
    type        = string
}