
import pytest

properties_separator = "="
server_properties = {}

@pytest.mark.parametrize("cfg_file", [
    ("/servers/minecraft/server.properties"),
    ("/servers/minecraft/ops.json"),
    ("/servers/minecraft/start_server.sh"),
    ("/servers/minecraft/whitelist.json"),
])
def test_cfg_files_exists(host, cfg_file):
    with host.sudo():
        file = host.file(cfg_file)
        assert file.exists
        assert file.user == "minecraft"


#@pytest.mark.parametrize("prop,value", [
#    ("broadcast-rcon-to-ops", "true"),
#    ("rcon.port", "25575"),
#    ("rcon.password", "ABC123"),
#    ("enable-rcon", "true"),
#])
#def test_server_properties(host, prop, value):
#    with host.sudo():
#        props_file = host.file('/servers/minecraft/server.properties')
#        assert f"{prop}={value}" in props_file.content_string
