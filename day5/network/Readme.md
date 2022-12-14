# Fabric scripts for Kubernetes and Docker

## FABRIC STRUCTURE:
- 1 Organisation
- 1 Peer
- 1 node Ordering service
- 1 CouchDB
- 1 Setup container

## DOCKER:

Use this for development.

 folders:
 - chaincode: chaincode
 - config: generated configuration files, like genesisblock or channel config file
 - core: core.yaml config file for the peer
 - crypto-config generated crypto material

 scripts:
 - generate.sh - generate crypto materials and config files
 - start.sh - start the network
 - teardown.sh - delete the network (docker volumes might need to be deleted manually)

## KUBERNETES:

Use this for cloud deployment.

 folders:
 - chaincode: chaincode
 - config: generated configuration files, like genesisblock or channel config file
 - crypto-config generated crypto material
 - fabric: fabric kubernetes config files, like
   - configMaps.yaml: general config maps and setup script
   - CouchDB.yaml: Kubernetes couchdb file
   - Orderer.yaml: Kubernetes orderer configuration file
   - Peer.yaml: Kubernetes peer configuration file
   - persitentvolume.yaml: Persitent volume claims for the setup
   - SetupJob.yaml: creating channel, configuring anchor peer, installing chaincode

 scripts:
 - generate.sh - generate crypto materials and config files
 - importsecrets.sh - import certificates to
 - start.sh - install the network
 - checking if the network setup is successful look at the set result of the setupjob  (like with kubectl logs <job-pod-name>)
 - teardown.sh - delete the network (some of the secrets and configs might need to be  deleted manually, like with kubectl delete secret)
