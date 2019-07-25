## ALPHA - Subject to change

## Usage:
```
sapexport [command]

Available Commands:
  help        Help about any command
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
```

## SAP NW Profile
```
  |   |   |--5  T-NL870003   <PRO> Profile for role X:SAPEXPORT                                
  |   |       |
  |   |       |--5  S_DSAUTH   <OBJ> DataServices: Auth Object                                   
  |   |       |   |
  |   |       |   |--5  T-NL87000300 <AUT>                                                             
  |   |       |       |
  |   |       |       |--5  ACTVT      <FLD> Activity                                                    
  |   |       |           |
  |   |       |           |-----16                                                                         
  |   |       |
  |   |       |--5  S_RFC      <OBJ> Authorization Check for RFC Access                          
  |   |       |   |
  |   |       |   |--5  T-NL87000300 <AUT>                                                             
  |   |       |       |
  |   |       |       |--5  ACTVT      <FLD> Activity                                                    
  |   |       |       |   |
  |   |       |       |   |-----16                                                                         
  |   |       |       |
  |   |       |       |--5  RFC_NAME   <FLD> Name (Whitelist) of RFC object to which access is allowed   
  |   |       |       |   |
  |   |       |       |   |-----/BODS/RFC_READ_TABLE2                                                      
  |   |       |       |   |-----DDIF_FIELDINFO_GET
  |   |       |       |   |-----RFCPING
  |   |       |       |   |-----RFC_GET_FUNCTION_INTERFACE
  |   |       |       |   |-----SUSR_GET_PROFILES_OF_USER_RFC
  |   |       |       |   |-----SUSR_GET_USERS_WITH_PROFS_RFC
  |   |       |       |   |-----SUSR_SUIM_API_RSUSR002
  |   |       |       |   |-----SUSR_SUIM_API_RSUSR008_009_NEW
  |   |       |       |   |-----SUSR_SUIM_API_RSUSR020
  |   |       |       |   |-----SUSR_SUIM_API_RSUSR050_AUTH
  |   |       |       |   |-----SUSR_SUIM_API_RSUSR050_PROF
  |   |       |       |   |-----SUSR_SUIM_API_RSUSR050_ROLE
  |   |       |       |   |-----SUSR_SUIM_API_RSUSR050_USER
  |   |       |       |   |-----SUSR_SUIM_API_RSUSR070
  |   |       |       |   |-----SUSR_SUIM_API_RSUSR100N
  |   |       |       |   |-----SUSR_SUIM_API_RSUSR200
  |   |       |       |
  |   |       |       |--5  RFC_TYPE   <FLD> Type of RFC object to which access is to be allowed         
  |   |       |           |
  |   |       |           |-----FUNC                                                                       
  |   |       |
  |   |       |--5  S_TABU_DIS <OBJ> Table Maintenance (using standard tools such as SM30)       
  |   |       |   |
  |   |       |   |--5  T-NL87000300 <AUT>                                                             
  |   |       |       |
  |   |       |       |--5  ACTVT      <FLD> Activity                                                    
  |   |       |       |   |
  |   |       |       |   |-----03                                                                         
  |   |       |       |
  |   |       |       |--5  DICBERCLS  <FLD> Table Authorization Group                                   
  |   |       |           |
  |   |       |           |-----*                                                                          
  |   |       |
  |   |       |--5  S_USER_GRP <OBJ> User Master Maintenance: User Groups                        
  |   |       |   |
  |   |       |   |--5  T-NL87000300 <AUT>                                                             
  |   |       |       |
  |   |       |       |--5  ACTVT      <FLD> Activity                                                    
  |   |       |       |   |
  |   |       |       |   |-----03                                                                         
  |   |       |       |
  |   |       |       |--5  CLASS      <FLD> User group in user master maintenance                       
  |   |       |           |
  |   |       |           |-----*                                                                          
  |   |       |
  |   |       |--5  S_USER_PRO <OBJ> User Master Maintenance: Authorization Profile              
  |   |           |
  |   |           |--5  T-NL87000300 <AUT>                                                             
  |   |               |
  |   |               |--5  ACTVT      <FLD> Activity                                                    
  |   |               |   |
  |   |               |   |-----03                                                                         
  |   |               |
  |   |               |--5  PROFILE    <FLD> Auth. profile in user master maintenance                    
  |   |                   |
  |   |                   |-----*                                                                          
```