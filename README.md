*commands*

ssh - ssh to a node, or pick a node from a list to ssh to
query - just get a list of nodes back
run - run a simple shell command on every matching node (via ansible)
ansible - use any arbitrary ansible module against every matching node
listfacts - list all known facts in puppetdb/other backends


global args:
    -d, --[no-]debug                 Whether to show additional debug logging. Default/Current: false
    -f, --fact FACT                  Fact criteria to query for. (specify fact name or name=value)
        --fact-and                   Multiple fact criteria are ANDed. Default
        --fact-or                    Multiple fact criteria are ORed.
        --ssl_cert FILE              SSL certificate file to connect to puppetdb. Default/Current: ~/.pdb/certs/jameseck-desktop.fasthosts.local.pem
        --ssl_key FILE               SSL key file to connect to puppetdb. Default/Current: ~/.pdb/private_keys/jameseck-desktop.fasthosts.local.pem
        --ssl_ca FILE                SSL ca file to connect to puppetdb. Default/Current: ~/.pdb/certs/ca.pem

query:
    -o, --order FIELD                Sort order for node list (fqdn,fact). Default/Current: fqdn
    -i, --[no-]include-facts         Include facts in output that were used in criteria. Default/Current: true

ssh:
    -l, --ssh_user USER              User for SSH. Default/Current: whatever your ssh client will use

run:
    -l, --ssh_user USER              User for SSH. Default/Current: whatever your ssh client will use
    -c, --command COMMAND            Run command on all matching hosts using ansible
    -t, --threads NUM                Number of threads to use for SSH commands. Default/Current: 5
    -r, --remote_user USER           User to become - only used by ansible. Default/Current: root
    -s, --ssh-options OPTIONS        Options for SSH (Default/Current: -A -t -Y
        --[no-]use-sudo              Whether to use sudo on the remote host - only used by ansible. Default/Current: true

ansible:
    -m, --ansible-module MODULE      Specify which module to use for ansible command Default/Current: shell
    -a, --ansible-module-args OPTS   Pass module arguments to ansible
    -A, --ansible-args ARG           Additional Ansible arguments. (Default/Current: [])
    -e, --ansible-env VAR            Ansible environment variables. Default/Current: []
    -t, --threads NUM                Number of threads to use for SSH commands. Default/Current: 5
    -r, --remote_user USER           User to become - only used by ansible. Default/Current: root
        --[no-]use-sudo              Whether to use sudo on the remote host - only used by ansible. Default/Current: true

