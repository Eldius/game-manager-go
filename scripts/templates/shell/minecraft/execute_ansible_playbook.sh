#!/bin/bash

cd {{ .WorkspacePath }}/scripts/ansible/minecraft/
pwd

pyenv local

which python
which ansible

ansible-galaxy install -r roles/requirements.yml --roles-path=roles/ -f

ansible-playbook \
    -i "{{ .ProvisioningInfo.IP }}," \
    -u "{{ .ProvisioningInfo.RemoteUser }}" \
    -e "ansible_port={{ .ProvisioningInfo.SSHPort }}" \
        {{ range $key, $value := .ProvisioningInfo.Args }}-e "{{ $key }}={{ $value }}" {{end}} \
    --private-key {{ .ProvisioningInfo.SSHKey }} \
        deploy-minecraft.yml
