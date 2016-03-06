
Notes
=====
Need to track: (?)
*  Last feedback connection.
*  Last Notification #.


Making pem/key pairs from .cer files
------------------------------------

1. Output certs and keys from keychain

2. Convert the pkcs12 to a pem file:

        openssl pkcs12 -in Certificates.p12 -out Certificates.pem -nodes

3.
        csplit -k Certificates.pem '/Bag Attribute/' '{1000}'

