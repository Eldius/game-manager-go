#!/bin/bash

which python
which ansible

ansible-galaxy install -r roles/requirements.yml --roles-path=roles/ -f

ansible-playbook \
    -i "{{ .ProvisioningInfo.IP }}," \
    -u {{ .ProvisioningInfo.RemoteUser }} \
        {{ range $key, $value := .ProvisioningInfo.Args }}-e "{{ $key }}={{ $value }}" {{end}} \
    --private-key {{ .ProvisioningInfo.SSHKey }} \
        deploy-minecraft.yml
