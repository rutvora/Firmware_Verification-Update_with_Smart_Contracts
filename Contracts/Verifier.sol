pragma solidity ^0.5.7;

contract Verifier {
    address private requestNode;
    string private requestNodeHash;
    int private requestNodeVersion;
    bool private needUpdate;
    int private hashVerificationCount;
    int private verLimit = 6;

    constructor(address reqNode, string memory hash, int version) public {
        requestNode = reqNode;
        requestNodeHash = hash;
        requestNodeVersion = version;
        needUpdate = false;
        hashVerificationCount = 0;
    }

    function compareStrings(string memory a, string memory b) internal returns (bool) {
        if (bytes(a).length != bytes(b).length) {
            return false;
        } else {
            return keccak256(abi.encodePacked(a)) == keccak256(abi.encodePacked(b));
        }
    }

    function verify(string memory hash, int version) public returns (string memory){
        if (version == requestNodeVersion) {
            if (compareStrings(hash, requestNodeHash)) hashVerificationCount++;
            return "Thanks";
        } else {
            if (version > requestNodeVersion) {
                needUpdate = true;
                return "Thanks";
            } else {
                return "Update";
            }
        }
    }

    function getStatus() public view returns (string memory){
//        if (msg.sender != requestNode) return "Unauthorized";
        return needUpdate ? "Update" : (hashVerificationCount > verLimit ? "Verified" : "Wait");
    }

    function destroy() public {
        if (msg.sender == requestNode) selfdestruct(address(uint160(requestNode)));     //Typecast address to address payable
    }
}
