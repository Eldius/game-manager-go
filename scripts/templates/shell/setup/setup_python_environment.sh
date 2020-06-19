
pyenv install {{ .PythonVersion }}
pyenv virtualenv 3.8.0 {{ .VenvName }}
cd {{ .WorkspacePath }}
pyenv local {{ .PythonVersion }}/envs/{{ .VenvName }}

python -m pip install --upgrade pip
python -m pip install ansible

cd $CURR_DIR
