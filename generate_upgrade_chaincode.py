#!/usr/bin/env python3

from jinja2 import Environment, FileSystemLoader
import os
import ruamel.yaml
import json
import sys

project_name = sys.argv[1]
chaincode_version = sys.argv[2]

yaml = ruamel.yaml.YAML(typ='safe')
# with open('config/' + project_name + '/chaincodeconfig.yaml') as fpi:
#     data = yaml.load(fpi)
# with open('config/' + project_name + '/chaincode_json_format.json', 'w') as fpo:
#     json.dump(data, fpo, indent=2)

# with open('config/' + project_name + '/chaincode_json_format.json') as json_file:
#     data = json.load(json_file)

file_loader = FileSystemLoader('templates/chaincode')
env = Environment(loader=file_loader)
env.trim_blocks = True
env.lstrip_blocks = True
env.rstrip_blocks = True

with open('config/' + project_name + '/network_json_format.json') as json_file:
    data = json.load(json_file)

file_name = "projects/" + data['domain'] + "/network/deploy-kubernetes/upgrade_chaincode.sh"
template = env.get_template('upgrade_chaincode_template.jinja2')
output = template.render(data=data, version=chaincode_version)
file = open(file_name, "w")
file.write(output)
file.close()
