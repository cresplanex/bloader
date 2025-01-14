#!/bin/bash
MY_IP=$(curl -s https://ipinfo.io/ip)
echo "{\"ip\": \"${MY_IP}\"}"