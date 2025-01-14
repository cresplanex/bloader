output "my_ip_address" {
  value = data.external.my_ip.result.ip
}

output "ec2_instance_public_ip" {
  value = aws_eip.instance_eip[*].public_ip
}