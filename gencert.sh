#!/bin/sh

SUBJECT="/C=IN/ST=HR/L=GGN/O=VASVEN/OU=TRAINING/CN=*.sapvs.io/emailAddress=sudosapan@gmail.com"


cd certs
rm *.pem

# 1. Generate CA's private key and self-signed certificate
openssl req -x509 -newkey rsa:4096 -days 365 -nodes -keyout ca-key.pem -out ca-cert.pem -subj $SUBJECT
echo "CA's self-signed certificate"
openssl x509 -in ca-cert.pem -noout -text 

openssl req \
-newkey rsa:4096 \
-nodes \
-keyout server-key.pem \
-out server-req.pem \
-subj $SUBJECT

openssl x509 -req \
-days 60 \
-CA ca-cert.pem \
-CAkey ca-key.pem \
-CAcreateserial \
-in server-req.pem \
-out server-cert.pem


echo "Server's signed certificate"
openssl x509 -in server-cert.pem -noout -text 

openssl req \
-newkey rsa:4096 \
-nodes \
-keyout client-key.pem \
-out client-req.pem \
-subj $SUBJECT


openssl x509 -req \
-days 60 \
-CA ca-cert.pem \
-CAkey ca-key.pem \
-CAcreateserial \
-in client-req.pem \
-out client-cert.pem


echo "Client's signed certificate"
openssl x509 -in client-cert.pem -noout -text 

cd ..