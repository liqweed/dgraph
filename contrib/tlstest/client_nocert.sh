#!/bin/bash

../../cmd/dgraph-live-loader/dgraph-live-loader -d server1.dgraph.io:9080 -tls.on -tls.ca_certs ca.crt -tls.server_name server1.dgraph.io -r data.rdf.gz 
