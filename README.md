## ALPHA - Subject to change

## Usage:
  sapexport [command]

Available Commands:
  help        Help about any command
  roleusers   Export a list of users in a role to JSON
  table       Extract a table to JSON

Flags:
  -a, --address string    System Address (or env SAPRFC_ADDRESS) (default "localhost")
  -c, --client string     System Client (or env SAPRFC_CLIENT) (default "001")
  -h, --help              help for sapexport
  -l, --language string   System Language (or env SAPRFC_LANGUAGE) (default "EN")
  -p, --pass string       RFC Password (or env SAPRFC_PASS)
  -r, --router string     Router (or env SAPRFC_ROUTER)
  -u, --user string       RFC Username (or env SAPRFC_USER)

Use "sapexport [command] --help" for more information about a command.

## SAP NW Profile
S_RFC with 
RFC_TYPE FUNC
RFC_NAME /BODS/RFC_READ_TABLE2 ESS_USERS_OF_ROLE_GET
ACTVT 16