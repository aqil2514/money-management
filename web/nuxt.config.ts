// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: "2025-07-15",
  app: {
    head: {
      titleTemplate: "%s | Money Management",
      title: "Selamat Datang",
    },
  },
  devtools: { enabled: true },
  modules: ["@pinia/nuxt", "@nuxt/ui"],
  css: ["~/assets/css/main.css"],
  future: {
    compatibilityVersion: 4,
  },
  vite: {
    optimizeDeps: {
      include: ["@vue/devtools-core", "@vue/devtools-kit"],
    },
    server: {
      hmr: {
        protocol: "ws",
        host: "localhost",
      },
    },
  },
});
