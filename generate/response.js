const Types = require('mongoose').Types;

const ASSETS = [
    "DEMO",
    "EURUSD",
    "USDJPY",
    "GBPUSD",
    "AUDUSD",
    "USDCHF",
    "NZDUSD",
    "USDCAD",
];

function response(n) {
    let rooms = [];
    for (let i = 0, l = n; i < l; i++) {
        rooms.push({
            id: Types.ObjectId(),
            title: 'title' + i,
            meta: {
                bulls: {
                    bets: 10,
                    bank: (Math.random()*3000).toFixed(14)/1,
                },
                bears: {
                    bets: 10,
                    bank: (Math.random()*3000).toFixed(14)/1,
                },
            },
            type: 1,
            asset: ASSETS[Math.floor(Math.random() * ASSETS.length - 1)],
            bet: Math.round(Math.random()*100),
            bets: [
                {
                    id: Types.ObjectId(),
                    bet: (Math.random()*100).toFixed(14)/1,
                    authorId: Types.ObjectId(),
                    payout: (Math.random()*1000).toFixed(14)/1,
                    commission: (Math.random()*20).toFixed(14)/1,
                    created: new Date(),
                },

                {
                    id: Types.ObjectId(),
                    bet: (Math.random()*100).toFixed(14)/1,
                    authorId: Types.ObjectId(),
                    payout: (Math.random()*1000).toFixed(14)/1,
                    commission: (Math.random()*20).toFixed(14)/1,
                    created: new Date(),
                },
                {
                    id: Types.ObjectId(),
                    bet: (Math.random()*100).toFixed(14)/1,
                    authorId: Types.ObjectId(),
                    payout: (Math.random()*1000).toFixed(14)/1,
                    commission: (Math.random()*20).toFixed(14)/1,
                    created: new Date(),
                },
                {
                    id: Types.ObjectId(),
                    bet: (Math.random()*100).toFixed(14)/1,
                    authorId: Types.ObjectId(),
                    payout: (Math.random()*1000).toFixed(14)/1,
                    commission: (Math.random()*20).toFixed(14)/1,
                    created: new Date(),
                },
                {
                    id: Types.ObjectId(),
                    bet: (Math.random()*100).toFixed(14)/1,
                    authorId: Types.ObjectId(),
                    payout: (Math.random()*1000).toFixed(14)/1,
                    commission: (Math.random()*20).toFixed(14)/1,
                    created: new Date(),
                },
                {
                    id: Types.ObjectId(),
                    bet: (Math.random()*100).toFixed(14)/1,
                    authorId: Types.ObjectId(),
                    payout: (Math.random()*1000).toFixed(14)/1,
                    commission: (Math.random()*20).toFixed(14)/1,
                    created: new Date(),
                },
            ],
            members: [
                {
                    id: Types.ObjectId(),
                    username: "player" + i,
                    avatar: Types.ObjectId(),
                },
                {
                    id: Types.ObjectId(),
                    username: "player" + i,
                    avatar: Types.ObjectId(),
                },
                {
                    id: Types.ObjectId(),
                    username: "player" + i,
                    avatar: Types.ObjectId(),
                },
                {
                    id: Types.ObjectId(),
                    username: "player" + i,
                    avatar: Types.ObjectId(),
                },
            ],
            firstAsset: {
                asset: (Math.random()*200).toFixed(14)/1,
                timestamp: new Date(),
            },
            lastAsset: {
                asset: (Math.random()*200).toFixed(14)/1,
                timestamp: new Date(),
            },
            authorId: Types.ObjectId(),
            status: 1,
            gameStatus: 2,
            createdAt: new Date(),
            deadlineAt: new Date(),
            expiredAt: new Date(),
            botOnly: false,
        });
    }

    return rooms
}

module.exports = response;
