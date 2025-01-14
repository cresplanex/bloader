data "external" "my_ip" {
  program = ["./get_my_ip.sh"]
}

data "aws_ssm_parameter" "amazonlinux_2" {
  name = "/aws/service/ami-amazon-linux-latest/amzn2-ami-hvm-${var.instance_architecture}-gp2"
}
