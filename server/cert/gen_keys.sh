#!/bin/sh

# Generate a private key
openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365 -nodes -subj '/CN=localhost' -addext 'subjectAltName = DNS:localhost'