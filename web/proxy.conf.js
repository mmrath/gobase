const PROXY_CONFIG = {
  "/api": {
    "target": process.env.API_PROXY,
    "secure": false,
  }
}

module.exports = PROXY_CONFIG;
