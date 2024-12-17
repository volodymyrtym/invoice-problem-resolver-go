const esbuild = require("esbuild");
const sassPlugin = require("esbuild-plugin-sass");
const path = require("path");

// Determine if the environment is in production mode
const isProd = process.env.NODE_ENV === "prod";
const outdir = path.resolve(__dirname, "static");

// Define build configuration
const buildOptions = {
    entryPoints: {
        "js/stimulus-app": "assets/js/stimulus/stimulus-app.js",
        "css/login": "assets/scss/login.scss",
        "css/daily-activities": "assets/scss/daily-activities.scss",
    },
    outdir: outdir,
    bundle: true,
    minify: isProd,
    sourcemap: !isProd,
    plugins: [sassPlugin()],
    loader: {
        ".js": "jsx",
        ".scss": "css",
    },
    write: true,
    target: ["es6"], // Ensures ES6 output for compatibility
};

// Run the build process
async function build() {
    try {
        // Ensure output directory exists
        console.log(`Building for ${isProd ? "production" : "development"} environment...`);
        await esbuild.build(buildOptions);
        console.log("Build successful!");
    } catch (error) {
        console.error("Build failed:", error);
        process.exit(1); // Exit with error code
    }
}

// Start the build process
build();