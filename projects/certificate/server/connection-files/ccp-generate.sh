#!/bin/bash

function one_line_pem {
    echo "`awk 'NF {sub(/\\n/, ""); printf "%s\\\\\\\n",$0;}' $1`"
}


function json_ccp_academy {
    local PP=$(one_line_pem $1)
    local CP=$(one_line_pem $2)
    local OP=$(one_line_pem $3)
    sed -e "s#\${PEERPEM}#$PP#" \
        -e "s#\${CAPEM}#$CP#" \
        -e "s#\${OPEM}#$OP#" \
        academy-ccp-template.json
}

PEERPEM=./academy/tlsca.academy.certificate.com-cert.pem
OPEM=./tlsca.certificate.com-cert.pem
CAPEM=./academy/ca.academy.certificate.com-cert.pem

echo "$(json_ccp_academy $PEERPEM $CAPEM $OPEM)" > ./connection-academy.json


function json_ccp_student {
    local PP=$(one_line_pem $1)
    local CP=$(one_line_pem $2)
    local OP=$(one_line_pem $3)
    sed -e "s#\${PEERPEM}#$PP#" \
        -e "s#\${CAPEM}#$CP#" \
        -e "s#\${OPEM}#$OP#" \
        student-ccp-template.json
}

PEERPEM=./student/tlsca.student.certificate.com-cert.pem
OPEM=./tlsca.certificate.com-cert.pem
CAPEM=./student/ca.student.certificate.com-cert.pem

echo "$(json_ccp_student $PEERPEM $CAPEM $OPEM)" > ./connection-student.json

