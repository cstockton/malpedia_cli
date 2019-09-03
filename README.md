# malpedia_cli

[![Go Report Card](https://goreportcard.com/badge/github.com/PimmyTrousers/malpedia_cli)](https://goreportcard.com/report/github.com/PimmyTrousers/malpedia_cli)

Malpedia_cli is a tool to interact with the malpedia service located [here](https://malpedia.caad.fkie.fraunhofer.de). Some of the endpoints commands require an api key due to restrictions with the service itself. It simplifies some of the endpoints and exposes the features that I beleive are the most important. 

## Configuration of the tool
The application requires an API for some of the endpoints, which can be passed by arugment or a YAML file at `$HOME/.malpedia_cli.yaml`. Currently it only allows for an apikey, so an example would look like the following 

```
apikey: <apikey>
```

## Currently supported commands
- [X] get samples via hash 
- [X] get a list of all tracked actors 
- [X] get information about a specific actor 
- [X] get a list of all tracked malware families 
- [X] get information about a specific malware family 
- [X] get yara rules by TLP level 
- [X] get yara rule by family 
- [X] get the malpedia version
- [X] get all hashes for a family 
- 

## TODO
- [ ] Command to download all samples from a family 
- [ ] Scan malpedia's malware catalog against a yara rule
- [ ] Upload a file to be checked against yara rules (in the works)
- [ ] Generic search (will return a family or actor)
- [ ] Download all samples from an actor
- [ ] Verbose logging 
- [ ] Support Contexts

## Examples 
```
- malpedia_cli version
- malpedia_cli getYaraRules white
- malpedia_cli getYaraRules amber -z -o yara_rules.zip
- malpedia_cli getSample 12f38f9be4df1909a1370d77588b74c60b25f65a098a08cf81389c97d3352f82 -p infected123 -o samples.zip
- malpedia_cli getSample 12f38f9be4df1909a1370d77588b74c60b25f65a098a08cf81389c97d3352f82 -r 
- malpedia_cli getActors --json
- malpedia_cli getActor apt28
- malpedia_cli getFamilies
- malpedia_cli getFamily ursnif
- malpedia_cli getYara ursnif 
- malpedia_cli getYara njrat -o njrat.zip
```

## Build Instructions
Create a binary file at your current directory
```
go build -o ./malpedia_cli
```
Create a binary file and install it in your path
```
go install
```