pragma solidity ^0.4.0;

contract Verifier {
    address requestNode;
    string requestNodeHash;
    int requestNodeVersion;
    bool needUpdate;
    int hashVerificationCount;
    int verLimit = 6;

    function Verifier(address reqNode, string hash, int version){
        requestNode = reqNode;
        requestNodeHash = hash;
        requestNodeVersion = version;
        needUpdate = false;
        hashVerificationCount = 0;
    }

    function verify(string hash, int version) returns (string){
        if (version == requestNodeVersion) {
            if (hash == requestNodeHash) count++;
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

    function getStatus() returns (string){
        if (msg.sender != requestNode) return "Unauthorized";
        return needUpdate ? "Update" : (count > verLimit ? "Verified" : "Wait");
    }

    function destroy() {
        if (msg.sender == requestNode) selfdestruct(requestNode);
    }
}
