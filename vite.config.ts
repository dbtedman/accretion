/// <reference types="vitest" />

import { fileURLToPath, URL } from "node:url";
import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [vue()],
    build: {
        outDir: "./ui/dist/",
    },
    resolve: {
        alias: {
            "@": fileURLToPath(new URL("./ui", import.meta.url)),
        },
    },
});
