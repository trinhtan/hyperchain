#!/usr/bin/env python3

from jinja2 import Environment, FileSystemLoader
import os
import ruamel.yaml
import json
import sys

project_name = sys.argv[1]
chaincode_by_config = sys.argv[2]

yaml = ruamel.yaml.YAML(typ='safe')
# with open('config/' + project_name + '/chaincodeconfig.yaml') as fpi:
#     data = yaml.load(fpi)
# with open('config/' + project_name + '/chaincode_json_format.json', 'w') as fpo:
#     json.dump(data, fpo, indent=2)

with open('config/' + project_name + '/network_json_format.json') as json_file:
    data = json.load(json_file)

file_loader = FileSystemLoader('templates/server')
env = Environment(loader=file_loader)
env.trim_blocks = True
env.lstrip_blocks = True
env.rstrip_blocks = True

file_name = "projects/" + data['domain'] + "/server/package.json"
template = env.get_template('package_template.jinja2')
output = template.render(data=data)
file = open(file_name, "w")
file.write(output)
file.close()


file_name = "projects/" + data['domain'] + "/server/cli/enrollAdmin.js"
template = env.get_template('cli/enrollAdmin_template.jinja2')
output = template.render(data=data)
file = open(file_name, "w")
file.write(output)
file.close()


file_name = "projects/" + data['domain'] + "/server/cli/registerUser.js"
template = env.get_template('cli/registerUser_template.jinja2')
output = template.render(data=data)
file = open(file_name, "w")
file.write(output)
file.close()


file_name = "projects/" + data['domain'] + "/server/cli/query.js"
template = env.get_template('cli/query_template.jinja2')
output = template.render(data=data)
file = open(file_name, "w")
file.write(output)
file.close()

if chaincode_by_config == 'no':

    file_name = "projects/" + data['domain'] + "/server/fabric/network.js"
    template = env.get_template('fabric/network_template.jinja2')
    output = template.render(data=data, chaincode_by_config=chaincode_by_config)
    file = open(file_name, "w")
    file.write(output)
    file.close()

else:
    with open('config/' + project_name + '/chaincode_json_format.json') as json_file:
        data = json.load(json_file)

    file_name = "projects/" + data['domain'] + "/server/cli/invoke.js"
    template = env.get_template('cli/invoke_template.jinja2')
    output = template.render(data=data)
    file = open(file_name, "w")
    file.write(output)
    file.close()

    file_name = "projects/" + data['domain'] + "/server/fabric/network.js"
    template = env.get_template('fabric/network_template.jinja2')
    output = template.render(data=data, chaincode_by_config=chaincode_by_config)
    file = open(file_name, "w")
    file.write(output)
    file.close()
