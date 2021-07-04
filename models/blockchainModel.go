package models

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

//Blockchain struct
type Blockchain struct {
	CurrentTransactions []Transaction2
	Chain               []Block
	Nodes               []Node
	NodeId              string
}

//Response struct
type Response struct {
	Message      string
	Index        int
	Transactions []Transaction2
	Proof        int
	PreviousHash string
}

//Transaction struct
type Transaction struct {
	Sender    string  `json:"sender"`
	Recipient string  `json:"recipient"`
	Amount    float64 `json:"amount"`
}

//Node struct
type Node struct {
	Address string
}

//Block struct
type Block struct {
	Index        int
	Timestamp    string
	Transactions []Transaction2
	Proof        int
	PreviousHash string
}

//Chain struct
type Chain struct {
	Chain []Block `json:"chain"`
}

//ResponceChain struct
type ResponceChain struct {
	Chain  []Block `json:"chain"`
	Length int     `json:"length"`
}

//ResponceConsensus struct
type ResponceConsensus struct {
	Message string  `json:"message"`
	Chain   []Block `json:"chain"`
}

//Urls struct
type Urls struct {
	Urls []string `json:"urls"`
}

func (b *Block) BlockToString() string {
	return strconv.Itoa(b.Index) + " [" + fmt.Sprintf("%f", b.Timestamp) + "] Proof: " + strconv.Itoa(b.Proof) + " | PrevHash: " + b.PreviousHash + " | Trx: " + strconv.Itoa(len(b.Transactions))
}

func NewBlockchain(proof int, previousHash string) Blockchain {
	var uuid = GenerateUUID()
	b := Blockchain{}
	b.Chain = append(b.Chain, b.NewBlock(proof, previousHash))
	b.NodeId = strings.Replace(uuid, "-", "", -1)

	return b
}

func (b *Blockchain) RegisterNode(address string) {
	/*
		Add a new node to the list of nodes
		:param address: Address of node. Eg. 'http://192.168.0.5:5000'
	*/

	parsedUrl, err := url.Parse(address)

	if err != nil {
		log.Println(err)
	}

	if parsedUrl.Host != "" {
		node := Node{}
		node.Address = parsedUrl.Host
		b.Nodes = append(b.Nodes, node)

	} else if parsedUrl.Path != "" {
		node := Node{}
		node.Address = parsedUrl.Path
		b.Nodes = append(b.Nodes, node)
	} else {
		panic("Invalid URL")
	}
}

func (b *Blockchain) ValidChain(chain Chain) bool {
	/*
		Determine if a given blockchain is valid
		:param chain: A blockchain
		:return: True if valid, False if not
	*/
	block := Block{}
	lastBlock := chain.Chain[0]
	currentIndex := 1

	for currentIndex < len(chain.Chain) {
		block = chain.Chain[currentIndex]
		log.Println(lastBlock)
		log.Println(block)
		log.Println("-----------")

		//Check that the hash of the block is correct
		lastBlockHash := b.Hash(lastBlock)

		if block.PreviousHash != lastBlockHash {
			return false
		}

		//Check that the Proof of Work is correct
		if !b.ValidProof(lastBlock.Proof, block.Proof, lastBlockHash) {
			return false
		}

		lastBlock = block
		currentIndex += 1
	}
	return true
}

func (b *Blockchain) ResolveConflicts() bool {
	/*
	   This is our consensus algorithm, it resolves conflicts
	   by replacing our chain with the longest one in the network.

	   :return: True if our chain was replaced, False if not
	*/

	var neighbours = b.Nodes
	var newChain []Block

	//We're only looking for chains longer than ours
	var maxLength = len(b.Chain)

	//Grab and verify the chains from all the nodes in our network
	for i := 0; i < len(neighbours); i++ {
		response, err := http.Get("http://" + neighbours[i].Address + "/chain")
		if err != nil {
			log.Fatalln(err)
		}

		//We Read the response body on the line below.
		if response.Status == "200" {
			//response.Body = http.MaxBytesReader(w, response.Body, 1048576)
			dec := json.NewDecoder(response.Body)
			dec.DisallowUnknownFields()

			var allChain Chain
			err := dec.Decode(&allChain)

			if err != nil {
				log.Println(err)
			}

			length := len(allChain.Chain)

			var chain Chain = allChain

			if (length > maxLength) && (b.ValidChain(chain)) {
				maxLength = length
				newChain = chain.Chain
			}
		}
	}

	if newChain != nil {
		b.Chain = newChain
		return true
	}

	return false
}

func (b *Blockchain) NewBlock(proof int, previousHash string) Block {
	/*Create a new Block in the Blockchain
	  :param proof: The proof given by the Proof of Work algorithm
	  :param previous_hash: Hash of previous Block
	  :return: New Block
	*/

	block := Block{}
	block.Index = int(len(b.Chain)) + 1
	block.Timestamp = fmt.Sprintf("%f", time.Now().Unix())
	block.Transactions = b.CurrentTransactions
	block.Proof = proof

	if previousHash != "" {
		block.PreviousHash = previousHash
	} else {
		block.PreviousHash = b.Hash(b.Chain[len(b.Chain)-1])
	}

	//Reset the current list of transactions
	b.CurrentTransactions = nil

	b.Chain = append(b.Chain, block)
	return block
}

func (b *Blockchain) NewTransaction(transaction2 Transaction2) int { //sender string, recipient string, amount float64,
	/*Creates a new transaction to go into the next mined Block
	:param sender: Address of the Sender
	:param recipient: Address of the Recipient
	:param amount: Amount
	:return: The index of the Block that will hold this transaction
	*/

	/*transaction := Transaction{}
	transaction.Sender = sender
	transaction.Recipient = recipient
	transaction.Amount = amount*/

	b.CurrentTransactions = append(b.CurrentTransactions, transaction2)

	block := b.LastBlock()

	return block.Index + 1
}

func (b *Blockchain) LastBlock() Block {
	if len(b.Chain) > 1 {
		return b.Chain[len(b.Chain)-1]
	} else {
		return b.Chain[len(b.Chain)-1]
	}
}

func (b *Blockchain) Hash(block Block) string {
	/*Creates a SHA-256 hash of a Block
	:param block: Block
	*/

	//We must make sure that the Dictionary is Ordered, or we'll have inconsistent hashes
	blockString, err := json.Marshal(block)

	if err != nil {
		log.Println(err)
	}

	h := sha256.New()
	h.Write([]byte(blockString))
	guessHash := hex.EncodeToString(h.Sum(nil))

	return guessHash
}

func (b *Blockchain) ProofOfWork(lastBlock Block) int {
	/*Simple Proof of Work Algorithm:
		- Find a number p' such that hash(pp') contains leading 4 zeroes
		- Where p is the previous proof, and p' is the new proof
	:param last_block: <dict> last Block
	:return: <int>
	*/

	var lastProof = lastBlock.Proof //lastBlock["proof"]
	var lastHash = b.Hash(lastBlock)
	var proof int = 0

	//fmt.Printf(lastHash)

	for b.ValidProof(lastProof, proof, lastHash) {
		proof += 1
	}

	return proof
}

func (b *Blockchain) ValidProof(lastProof int, proof int, lastHash string) bool {
	/*Validates the Proof
	  :param last_proof: <int> Previous Proof
	  :param proof: <int> Current Proof
	  :param last_hash: <str> The hash of the Previous Block
	  :return: <bool> True if correct, False if not.
	*/

	//lastProofOut, err := json.Marshal(lastProof)
	//if err != nil {
	//	panic(err)
	//}

	var str = strconv.Itoa(lastProof) + strconv.Itoa(int(proof)) + lastHash

	h := sha256.New()
	h.Write([]byte(str))
	guessHash := hex.EncodeToString(h.Sum(nil))

	return guessHash[len(guessHash)-3:] == "0000"
}

func GenerateUUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	return uuid
}
