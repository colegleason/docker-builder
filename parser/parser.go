package parser

import (
	"github.com/rafecolton/bob/log"
	//"github.com/rafecolton/bob/parser/dclient"
	"github.com/rafecolton/bob/parser/uuid"
)

import (
	"github.com/onsi/gocleanup"
	"os"
)

func init() {
	gocleanup.Register(func() {
		//do stuff
	})
}

/*
Parser is a struct that contains a Builderfile and knows how to parse it both
as raw text and to convert toml to a Builderfile struct.  It also knows how to
tell if the Builderfile is valid (openable) or nat.
*/
type Parser struct {
	filename string
	log.Log
	//dockerClient  dclient.DockerClient
	uuidGenerator uuid.UUIDGenerator
	top           string
}

/*
NewParser returns an initialized Parser.  Not currently necessary, as no
default values are assigned to a new Parser, but useful to have in case we need
to change this.
*/
func NewParser(filename string, logger log.Log) (*Parser, error) {
	//client, err := dclient.NewDockerClient(nil, false)
	//if err != nil {
	//return nil, err
	//}

	return &Parser{
		Log:      logger,
		filename: filename,
		//dockerClient:  client,
		uuidGenerator: uuid.NewUUIDGenerator(true),
		top:           os.ExpandEnv("${PWD}"),
	}, nil
}

//[>
//LatestImageTaggedWithUUID accepts a uuid and invokes the underlying utility
//DockerClient to determine the id of the most recently created image tagged with
//the provided uuid.
//*/
//func (parser *Parser) LatestImageTaggedWithUUID(uuid string) string {
//// eat the error and let it fail when we try to run the docker command
//id, _ := parser.dockerClient.LatestImageTaggedWithUUID(uuid)
//return id
//}

/*
NextUUID returns the next UUID generated by the parser's uuid generator.  This
will either be a random uuid (normal behavior) or the same uuid every time if
the generator is "seeded" (used for tests)
*/
func (parser *Parser) NextUUID() (string, error) {
	ret, err := parser.uuidGenerator.NextUUID()
	if err != nil {
		return "", err
	}

	return ret, nil
}

/*
SeedUUIDGenerator turns this parser's uuidGenerator into a seeded generator.
All calls to NextUUID() will produce the same uuid after this function is
called and until RandomizeUUIDGenerator() is called.
*/
func (parser *Parser) SeedUUIDGenerator() {
	parser.uuidGenerator = uuid.NewUUIDGenerator(false)
}

/*
RandomizeUUIDGenerator turns this parser's uuidGenerator into a random
generator.  All calls to NextUUID() will produce a random uuid after this
function is called and until SeedUUIDGenerator() is called.
*/
func (parser *Parser) RandomizeUUIDGenerator() {
	parser.uuidGenerator = uuid.NewUUIDGenerator(true)
}
