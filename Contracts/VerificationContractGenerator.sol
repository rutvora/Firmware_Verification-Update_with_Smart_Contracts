pragma solidity ^0.5.7;

import "Verifier.sol";

contract VerificationContractGenerator {
    function generateVerificationContract(string memory hash, int version) public returns (address){
        Verifier verifier = new Verifier(msg.sender, hash, version);
        return address(verifier);
    }
}