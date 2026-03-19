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
    to: "/transaction",
  },
]);
