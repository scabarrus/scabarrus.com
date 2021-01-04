#!/bin/bash

COUNTRY="FR"
STATE="PARIS"
LOC="PARIS"
ORG="scabarrus.com"
OU="k8s.webhook"
CNAME="\*"
openssl req \
    -x509 \
    -nodes \
    -newkey rsa:2048 \
    -keyout privateKey.key \
    -out certificate.crt \
    -days 3650 \
    -subj "/C=${COUNTRY}/ST=${STATE}/L=${LOC}/O=${ORG}/OU=${OU}/CN=${CNAME}"

# -subj "/C=GB/ST=London/L=London/O=Global Security/OU=IT Department/CN=*"
