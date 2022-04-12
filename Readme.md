## Prover Server 
Prover Server is a REST API Wrapper & golang binding for [SnarkJS](https://github.com/iden3/snarkjs)

List of implemented features:

* Generate proof

### Requirements
* [SnarkJS](https://github.com/iden3/snarkjs)

### Installation

1. Build prover server:
    ```
    go build ./cmd/prover/prover.go
    ```

2. Edit config file `configs/prover.yaml`

3. Put compiled circuits into `<circuitsBasePath>/<circuitName>` directory. Where `<circuitsBasePath>` is config option with default value `circuits`, and `<circuitName>` is name of the circuit that will be passed as a param to an API call.
   See [SnarkJS Readme](https://github.com/iden3/snarkjs) for instructions on how to compile circuits.

3. Run prover server:
     ```
    ./prover
    ```

## API
### Generate proof

```
POST /api/v1/proof/generate
Content-Type: application/json
{
  "inputs": {...}, // circuit specific inputs
  "circuit_name": "..." // name of directory 
}
```

## Docker images

Build and run container with snarkjs prover:
```bash
docker build -t prover-server .
docker run -it -p 8002:8002 prover-server
```

Build and run container with rapidsnark prover:
```bash
docker build -t prover-server-rapidsnark -f Dockerfile-rapidsnark .
docker run -it -p 8002:8002 prover-server-rapidsnark
```

## License

prover-server is part of the iden3 project copyright 2021 0KIMS association and published with GPL-3 license. Please check the LICENSE file for more details.