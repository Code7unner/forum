module.exports = {
    publicPath: process.env.NODE_ENV === "production" ? "/" : "/",
    productionSourceMap: false,
    filenameHashing: false,
    outputDir: "build",
    lintOnSave: process.env.NODE_ENV !== "production"
};