
ansible-playbook -i "{{ hosts }}," -u {{ remote_user}} --private-key {{ private_key }} deploy-minecraft.yml
