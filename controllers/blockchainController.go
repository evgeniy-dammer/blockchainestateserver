package controllers

import (
	"blockchainestateserver/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

//Instantiate the Blockchain
var Blockchain = models.NewBlockchain(100, "1")

//Mine BlahBlahBlah
func Mine(w http.ResponseWriter, r *http.Request) {

	//We run the proof of work algorithm to get the next proof...
	var lastBlock = Blockchain.LastBlock()
	var proof = Blockchain.ProofOfWork(lastBlock)

	//We must receive a reward for finding the proof.

	//Generate a globally unique address for this node
	//var uuid = models.GenerateUUID()
	//var nodeIdentifier = strings.Replace(uuid, "-", "", -1)

	//The sender is "0" to signify that this node has mined a new coin.
	//Blockchain.NewTransaction("0", nodeIdentifier, 1)

	//Forge the new Block by adding it to the chain
	var previousHash = Blockchain.Hash(lastBlock)
	var block = Blockchain.NewBlock(proof, previousHash)

	response := models.Response{}
	response.Message = "New Block Forged"
	response.Index = block.Index
	response.Transactions = block.Transactions
	response.Proof = block.Proof
	response.PreviousHash = block.PreviousHash

	jsonData, err := json.Marshal(response)

	if err != nil {
		fmt.Fprint(w, err)
	}
	fmt.Fprint(w, string(jsonData))
}

//NewTransaction BlahBlahBlah
func NewTransaction(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	var t models.Transaction2
	err := dec.Decode(&t)

	if err != nil {
		log.Println(err)
	}

	var ok = Blockchain.NewTransaction(t)
	//var ok = Blockchain.NewTransaction(t.Sender, t.Recipient, t.Amount)

	if ok != 0 {
		fmt.Fprint(w, "Your transaction will be included in block "+strconv.Itoa(ok))
	} else {
		fmt.Fprint(w, "Error inserting transaction!")
	}
}

//FullChain BlahBlahBlah
func FullChain(w http.ResponseWriter, r *http.Request) {
	var responceChain = models.ResponceChain{}
	responceChain.Chain = Blockchain.Chain
	responceChain.Length = len(Blockchain.Chain)

	jsonData, err := json.Marshal(responceChain)

	if err != nil {
		fmt.Fprint(w, err)
	}
	fmt.Fprint(w, string(jsonData))
}

//RegisterNodes BlahBlahBlah
func RegisterNodes(w http.ResponseWriter, r *http.Request) {
	var urls []string
	var urlstring string

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	var t models.Urls
	err := dec.Decode(&t)

	if err != nil {
		log.Println(err)
	}

	for i := 0; i < len(t.Urls); i++ {
		urls = append(urls, t.Urls[i])
		Blockchain.RegisterNode("http://" + t.Urls[i] + "/")
		urlstring = urlstring + t.Urls[i] + ", "
	}

	fmt.Fprint(w, strconv.Itoa(len(urls))+" new nodes have been added: "+urlstring)
}

//Consensus BlahBlahBlah
func Consensus(w http.ResponseWriter, r *http.Request) {
	replaced := Blockchain.ResolveConflicts()
	message := ""

	if replaced == true {
		message = "was replaced"
	} else {
		message = "is authoritive"
	}

	var responceConsensus = models.ResponceConsensus{}
	responceConsensus.Chain = Blockchain.Chain
	responceConsensus.Message = "Our chain " + message

	jsonData, err := json.Marshal(responceConsensus)

	if err != nil {
		fmt.Fprint(w, err)
	}
	fmt.Fprint(w, string(jsonData))
}
