const cluster = require('cluster');
const http = require('http');
const numCPUs = require('os').cpus().length;
const express = require('express');

const response = require('./generate/response');

const app = express();
const rooms = response(20);

const PORT = process.env.PORT || 5002;

if (cluster.isMaster) {
    // Fork workers.
    for (let i = 0; i < numCPUs; i++) {
        cluster.fork();
    }

    cluster.on('exit', (worker, code, signal) => {
    });
} else {
    app.get('/', (req, res) => {
        res.send(rooms);
    });

    app.listen(PORT, () => {
        console.log(`Node.js worker:${PORT}!`);
    });
}

