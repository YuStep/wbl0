const { createProxyMiddleware } = require("http-proxy-middleware");

module.exports = function(app) {
    app.use(
        createProxyMiddleware('/api', {
            target: "http://server:8000",
            changeOrigin: true,
            pathRewrite: {
                '^/api': '',  // Удалить префикс /api из URL перед отправкой запроса на бэкенд
            },
        })
    );
};