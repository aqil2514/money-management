import type { NavigationMenuItem } from "@nuxt/ui";

export const NAVIGATION_ITEMS = ref<NavigationMenuItem[]>([
  {
    label: "Dashboard",
    icon: "i-lucide-layout-dashboard",
    to: "/",
  },
  {
    label: "Transaksi",
    icon: "i-lucide-coins",
    children: [
      {
        label: "Alur Kas",
        icon: "i-lucide-chart-no-axes-combined",
        to: "/transaction",
      },
      {
        label: "Kategori",
        icon: "i-lucide-cassette-tape",
        to: "/category",
      },
      {
        label: "Aset",
        icon: "i-lucide-database",
        to: "/asset",
      },
    ],
  },
]);
