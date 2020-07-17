cryptogen generate --config /fabric/config/crypto-config.yaml
cp -r crypto-config /fabric/
cp -r /opt/gopath/src/github.com/hyperledger/fabric/core/chaincode/lib/cid /opt/gopath/src/github.com/hyperledger/fabric/core/chaincode/
cp /fabric/config/configtx.yaml /fabric/
cd /fabric
configtxgen -profile TwoOrgsOrdererGenesis -channelID $SYS_CHANNEL -outputBlock genesis.block
configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ${CHANNEL_NAME}.tx -channelID $CHANNEL_NAME
configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate StudentMSPanchors.tx -channelID $CHANNEL_NAME -asOrg StudentMSP
configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate AcademyMSPanchors.tx -channelID $CHANNEL_NAME -asOrg AcademyMSP
