const axios = require("axios");
const crypto = require("crypto");
require("dotenv").config();

class BingX {
    constructor(apiKey, secretKey) {
        this.apiKey = apiKey;
        this.secretKey = secretKey;
        this.baseUrl = "https://api-swap-rest.bingbon.pro";
    }

    async getBalance(currency) {
        const timestamp = Date.now();
        const path = "/api/v1/user/getBalance";
        const paramString = `apiKey=${this.apiKey}&currency=${currency}&timestamp=${timestamp}`;
        const originString = `POST${path}${paramString}`;
        console.log(originString);
        const signature = crypto
            .createHmac("sha256", this.secretKey)
            .update(originString)
            .digest("base64");
        const url = `${
            this.baseUrl
        }${path}?${paramString}&sign=${encodeURIComponent(signature)}`;
        console.log(url);
        const response = await axios.post(url);
        return response.data;
    }

    async setLeverage(params) {
        const path = "/api/v1/user/setLeverage";
        const timestamp = Date.now();
        const symbol = params.symbol;
        const side = params.side;
        const leverage = params.leverage;

        const paramString = `apiKey=${this.apiKey}&leverage=${leverage}&side=${side}&symbol=${symbol}&timestamp=${timestamp}`;
        const originString = `POST${path}${paramString}`;
        console.log(originString);
        const signature = crypto
            .createHmac("sha256", this.secretKey)
            .update(originString)
            .digest("base64");
        const url = `${
            this.baseUrl
        }${path}?${paramString}&sign=${encodeURIComponent(signature)}`;
        console.log(url);
        const response = await axios.post(url);
        return response.data;
    }

    async getLeverage(params) {
        const path = "/api/v1/user/getLeverage";
        const timestamp = Date.now();
        const symbol = params.symbol;
        const paramString = `apiKey=${this.apiKey}&symbol=${symbol}&timestamp=${timestamp}`;
        const originString = `POST${path}${paramString}`;
        console.log(originString);
        const signature = crypto
            .createHmac("sha256", this.secretKey)
            .update(originString)
            .digest("base64");
        const url = `${
            this.baseUrl
        }${path}?${paramString}&sign=${encodeURIComponent(signature)}`;
        console.log(url);
        const response = await axios.post(url);
        return response.data;
    }

    async placeOrder(params) {
        const path = "/api/v1/user/trade";
        const timestamp = Date.now();
        const symbol = params.symbol;
        const side = params.side;
        const entrustPrice = params.entrustPrice;
        const entrustVolume = params.entrustVolume;
        const tradeType = params.tradeType;
        const action = params.action;
        const paramString = `action=${action}&apiKey=${this.apiKey}&entrustPrice=${entrustPrice}&entrustVolume=${entrustVolume}&side=${side}&symbol=${symbol}&timestamp=${timestamp}&tradeType=${tradeType}`;
        const originString = `POST${path}${paramString}`;
        const signature = crypto
            .createHmac("sha256", this.secretKey)
            .update(originString)
            .digest("base64");

        const url = `${
            this.baseUrl
        }${path}?${paramString}&sign=${encodeURIComponent(signature)}`;

        console.log(url);
        const response = await axios.post(url);
        return response.data;
    }  

    async getPositions(params) {
        const path = "/api/v1/user/getPositions";
        const timestamp = Date.now();
        const symbol = params.symbol;
        const paramString = `apiKey=${this.apiKey}&symbol=${symbol}&timestamp=${timestamp}`;
        const originString = `POST${path}${paramString}`;
        const signature = crypto
            .createHmac("sha256", this.secretKey)
            .update(originString)
            .digest("base64");
        const url = `${
            this.baseUrl
        }${path}?${paramString}&sign=${encodeURIComponent(signature)}`;
        const response = await axios.post(url);
        return response.data.data.positions;
    }

    async closePositionBySymbol(symbol) {
        console.log(symbol);
        const position = await this.getPositions({ symbol: symbol });
        if (position === null) {
            console.log("No position found for the symbol");
            return;
        }
        const positionId = position[0].positionId;
        console.log(positionId);
        const path = "/api/v1/user/oneClickClosePosition";
        const timestamp = Date.now();
        const paramString = `apiKey=${this.apiKey}&positionId=${positionId}&symbol=${symbol}&timestamp=${timestamp}`;
        const originString = `POST${path}${paramString}`;
        const signature = crypto
            .createHmac("sha256", this.secretKey)
            .update(originString)
            .digest("base64");
        const url = `${
            this.baseUrl
        }${path}?${paramString}&sign=${encodeURIComponent(signature)}`;
        console.log(url);
        const response = await axios.post(url);
        return response.data;
    }
}

//---------------------------------------EXAMPLS OF USAGE BELOW---------------------------------------------




// const bingx = new BingX(
//     process.env.BINGX_API_KEY,
//     process.env.BINGX_SECRET_KEY
// );

// let tmp_params = {
//     entry_price: '0.1626',
//     sl: '0.165039',
//     tp_percent: '5.33',
//     rr: '1.08',
//     symbol: 'ALGOUSDT.PS',
//     tp_price: '0.1617363263',
//     position: 'short',
//     strategy: 'maj',
//     type: 'open_short'
//   }

//   bingx.placeOrder(tmp_params).then((data) => console.log(data));


// bingx.getBalance("USDT").then((data) => console.log(data));

// bingx
//     .setLeverage({
//         symbol: "BTC-USDT",
//         side: "Long",
//         leverage: "10",
//     })
//     .then((data) => console.log(data));

// const lev = bingx.getLeverage({ symbol: "BTC-USDT" }).then((data) => {
//     console.log(data);
// });

// bingx
//     .placeOrder({
//         symbol: "BNB-USDT",
//         side: "Ask",
//         entrustPrice: "26600",
//         entrustVolume: "0.02",
//         tradeType: "Market",
//         action: "Open",
//         takerProfitPrice: "290",
//         stopLossPrice: "320",
//     })
//     .then((data) => {
//         console.log(data);
//     });

// bingx.getPositions({ symbol: "BTC-USDT" }).then((data) => {
//     console.log(data.data.positions);
// });

// bingx.closePositionBySymbol("BTC-USDT");

module.exports = BingX;