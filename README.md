## Install

- `yarn install`
- `glide install`

## Server bench

- `NODE_ENV=production PORT=5001 node server.js`
- `wrk -t10 -c5000 --latency --timeout=10s -d10s http://localhost:5001`

- `GIN_MODE=release PORT=5002 go run server.go`
- `wrk -t10 -c5000 --latency --timeout=10s -d10s http://localhost:5002`
