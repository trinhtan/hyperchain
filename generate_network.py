#!/usr/bin/env python3

from jinja2 import Environment, FileSystemLoader
import os
import ruamel.yaml
import json
import sys

project_name = sys.argv[1]
public_ip = sys.argv[2]

yaml = ruamel.yaml.YAML(typ='safe')
# with open('config/' + project_name + '/networkconfig.yaml') as fpi:
#     data = yaml.load(fpi)
# with open('config/' + project_name + '/network_json_format.json', 'w') as fpo:
#     json.dump(data, fpo, indent=2)

with open('config/' + project_name + '/network_json_format.json') as json_file:
    data = json.load(json_file)

file_loader = FileSystemLoader('templates/networkconfig')
env = Environment(loader=file_loader)
env.trim_blocks = True
env.lstrip_blocks = True
env.rstrip_blocks = True

file_name = "projects/" + data['domain'] + "/network/deploy-kubernetes/config/crypto-config.yaml"
template = env.get_template('config/crypto-config_template.jinja2')
output = template.render(data=data)
file = open(file_name, "w")
file.write(output)
file.close()

file_name = "projects/" + data['domain'] + "/network/deploy-kubernetes/config/configtx.yaml"
template = env.get_template('config/configtx_template.jinja2')
output = template.render(data=data)
file = open(file_name, "w")
file.write(output)
file.close()



file_name = "projects/" + data['domain'] + "/network/deploy-kubernetes/config/connection-files/ccp-generate.sh"
template = env.get_template('config/connection-files/ccp-generate_template.jinja2')
output = template.render(data=data)
file = open(file_name, "w")
file.write(output)
file.close()

file_name = "projects/" + data['domain'] + "/network/deploy-kubernetes/kubernetes/fabric-" + data['domain'] + "-tools.yaml"
template = env.get_template('kubernetes/fabric-tools_template.jinja2')
output = template.render(data=data)
file = open(file_name, "w")
file.write(output)
file.close()


file_name = "projects/" + data['domain'] + "/network/deploy-kubernetes/kubernetes/fabric-" + data['domain'] + "_pv.yaml"
template = env.get_template('kubernetes/fabric-pv_template.jinja2')
output = template.render(data=data)
file = open(file_name, "w")
file.write(output)
file.close()


file_name = "projects/" + data['domain'] + "/network/deploy-kubernetes/kubernetes/fabric-" + data['domain'] + "_pvc.yaml"
template = env.get_template('kubernetes/fabric-pvc_template.jinja2')
output = template.render(data=data)
file = open(file_name, "w")
file.write(output)
file.close()

file_name = "projects/" + data['domain'] + "/network/deploy-kubernetes/kubernetes/backup-" + data['domain'] + "-orderer_pv.yaml"
template = env.get_template('kubernetes/backup-orderer-pv_template.jinja2')
output = template.render(data=data)
file = open(file_name, "w")
file.write(output)
file.close()


file_name = "projects/" + data['domain'] + "/network/deploy-kubernetes/kubernetes/backup-" + data['domain'] + "-orderer_pvc.yaml"
template = env.get_template('kubernetes/backup-orderer-pvc_template.jinja2')
output = template.render(data=data)
file = open(file_name, "w")
file.write(output)
file.close()


file_name = "projects/" + data['domain'] + "/network/deploy-kubernetes/kubernetes/" + data['domain'] + "-orderer_deploy.yaml"
template = env.get_template('kubernetes/orderer-deploy_template.jinja2')
output = template.render(data=data)
file = open(file_name, "w")
file.write(output)
file.close()


file_name = "projects/" + data['domain'] + "/network/deploy-kubernetes/kubernetes/" + data['domain'] + "-orderer_svc.yaml"
template = env.get_template('kubernetes/orderer-svc_template.jinja2')
output = template.render(data=data)
file = open(file_name, "w")
file.write(output)
file.close()


file_name = "projects/" + data['domain'] + "/network/deploy-kubernetes/scripts.sh"
template = env.get_template('scripts_template.jinja2')
output = template.render(data=data)
file = open(file_name, "w")
file.write(output)
file.close()


file_name = "projects/" + data['domain'] + "/network/deploy-kubernetes/network.sh"
template = env.get_template('network_template.jinja2')
output = template.render(data=data)
file = open(file_name, "w")
file.write(output)
file.close()

file_name = "projects/" + data['domain'] + "/network/deploy-kubernetes/channel.sh"
template = env.get_template('channel_template.jinja2')
output = template.render(data=data)
file = open(file_name, "w")
file.write(output)
file.close()

file_name = "projects/" + data['domain'] + "/network/deploy-kubernetes/kubernetes/" + data['domain']  + "-ca_deploy.yaml"
template = env.get_template('kubernetes/ca-deploy_template.jinja2')
output = template.render(data=data)
file = open(file_name, "w")
file.write(output)
file.close()

for org in data['orgs']:

    file_name = "projects/" + data['domain'] + "/network/deploy-kubernetes/config/connection-files/" + org['name'] + "-ccp-template.json"
    template = env.get_template('config/connection-files/ccp-template_template.jinja2')
    output = template.render(data=data, org=org, public_ip=public_ip)
    file = open(file_name, "w")
    file.write(output)
    file.close()

    file_name = "projects/" + data['domain'] + "/network/deploy-kubernetes/kubernetes/" + data['domain'] + "-" + org['name'] +"-ca_svc.yaml"
    template = env.get_template('kubernetes/ca-svc_template.jinja2')
    output = template.render(data=data, org=org)
    file = open(file_name, "w")
    file.write(output)
    file.close()

    for peer in org['peers']:

        file_name = "projects/" + data['domain'] + "/network/deploy-kubernetes/kubernetes/backup-" + data['domain'] + "-" + org['name'] + "-" + peer['peer'] + "_pv.yaml"
        template = env.get_template('kubernetes/backup-peer-pv_template.jinja2')
        output = template.render(data=data, org=org, peer=peer)
        file = open(file_name, "w")
        file.write(output)
        file.close()

        file_name = "projects/" + data['domain'] + "/network/deploy-kubernetes/kubernetes/backup-" + data['domain'] + "-" + org['name'] + "-" + peer['peer'] + "_pvc.yaml"
        template = env.get_template('kubernetes/backup-peer-pvc_template.jinja2')
        output = template.render(data=data, org=org, peer=peer)
        file = open(file_name, "w")
        file.write(output)
        file.close()

        file_name = "projects/" + data['domain'] + "/network/deploy-kubernetes/kubernetes/" + data['domain'] + "-" + org['name'] + "-" + peer['peer'] + "_deploy.yaml"
        template = env.get_template('kubernetes/peer-deploy_template.jinja2')
        output = template.render(data=data, org=org, peer=peer)
        file = open(file_name, "w")
        file.write(output)
        file.close()

        file_name = "projects/" + data['domain'] + "/network/deploy-kubernetes/kubernetes/" + data['domain'] + "-" + org['name'] + "-" + peer['peer'] + "_svc.yaml"
        template = env.get_template('kubernetes/peer-svc_template.jinja2')
        output = template.render(data=data, org=org, peer=peer)
        file = open(file_name, "w")
        file.write(output)
        file.close()
