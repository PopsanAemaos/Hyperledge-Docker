/*
 * SPDX-License-Identifier: Apache-2.0
 */

'use strict';

const { FileSystemWallet, Gateway } = require('fabric-network');
const path = require('path');
const util = require("util")

const ccpPath = path.resolve(__dirname,'..', '..', '..', 'basic-network', 'connection.1.json');

async function main() {
    try {

        // Create a new file system based wallet for managing identities.
        const walletPath = path.join(process.cwd(), 'wallet');
        const wallet = new FileSystemWallet(walletPath);
        console.log(`Wallet path: ${walletPath}`);

        // Check to see if we've already enrolled the user.
        const userExists = await wallet.exists('userOrg1');
        if (!userExists) {
            console.log('An identity for the user "userOrg1" does not exist in the wallet');
            console.log('Run the registerUser.js application before retrying');
            return;
        }
        console.log('******************************************************************');

        // Create a new gateway for connecting to our peer node.
        const gateway = new Gateway();
        await gateway.connect(ccpPath, { wallet, identity: 'userOrg1', discovery: { enabled: false, asLocalhost: true } });
        console.log('******************************************************************');

        // Get the network (channel) our contract is deployed to.
        const network = await gateway.getNetwork('mychannel');
        console.log('******************************************************************');

        // Get the contract from the network.
        const contract = network.getContract('mychaincode');
        console.log('******************************************************************');
        
        // Submit the specified transaction.
        var args =[ "1111","ford","081942"] 
        const argsString = args.map((arg) => util.format('%s', arg)).join('|');
        // await contract.sendTransaction("createuser",argsString);
        await contract.submitTransaction("createuser",argsString);
        //await channel.sendTransactionProposal("createwallet","0000","Gucci","10000","ford");
        console.log('Transaction has been submitted');

        // Disconnect from the gateway.
        await gateway.disconnect();

    } catch (error) {
        console.error(`Failed to submit transaction: ${error}`);
        process.exit(1);
    }
}

main();
