
pyenv install 3.8.0
pyenv shell 3.8.0
python -m pip install --upgrade pip
python -m pip install ansible

ansible-galaxy install -r {{ .WORKSPACE_PATH }}/ansible/roles/requirements.yml --roles-path={{WORKSPACE_PATH}}/ansible/roles/
